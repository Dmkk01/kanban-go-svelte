package emoji

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/rodaine/table"
)

type Emoji struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Category string `json:"category"`
	ImageURL string `json:"image_url"`
}

type EmojiURL struct {
	Category string
	URL      string
}

var jsonPath = "../frontend/src/utils/emojis.json"

var emojipediaURL = "https://emojipedia.org"
var emojisMap map[string]Emoji = make(map[string]Emoji)

var wg sync.WaitGroup
var mutex sync.Mutex

func GetEmojis() {
	allEmojis := readEmojis()

	categories := []string{"people", "nature", "food-drink", "activity", "travel-places", "objects", "symbols", "flags"}
	emojiURLs := []EmojiURL{}

	wg.Add(len(categories))

	for _, category := range categories {
		go func(category string) {
			defer wg.Done()
			newEmojiUrls := scrapeEmojiList(category)
			mutex.Lock()
			emojiURLs = append(emojiURLs, newEmojiUrls...)
			mutex.Unlock()
		}(category)
	}

	wg.Wait()

	wg.Add(len(emojiURLs))

	for _, item := range emojiURLs {
		go func(emoji EmojiURL) {
			defer wg.Done()

			newEmojis := scrapeEmojiType(emoji)

			mutex.Lock()
			allEmojis = append(allEmojis, newEmojis...)
			mutex.Unlock()

		}(item)
	}

	wg.Wait()

	fmt.Println("Saving emojis data...")
	fmt.Println("Total emojis:", len(allEmojis))

	emojisJSON, _ := json.Marshal(allEmojis)
	absPath, _ := filepath.Abs(jsonPath)

	os.WriteFile(absPath, emojisJSON, 0644)

	fmt.Println("Successfully wrote to file:", absPath)
}

func scrapeEmojiList(category string) []EmojiURL {
	url := emojipediaURL + "/" + category

	c := colly.NewCollector()

	emojis := []EmojiURL{}

	c.OnHTML(".emoji-list > li > a", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		var emoji EmojiURL
		emoji.Category = category
		emoji.URL = href

		emojis = append(emojis, emoji)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.StatusCode, err)
	})

	c.OnRequest(func(r *colly.Request) {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
		c.IgnoreRobotsTxt = true

		fmt.Println("Company Visiting", r.URL.String())
	})

	c.Visit(url)

	return emojis
}

func scrapeEmojiType(emojiURL EmojiURL) []Emoji {
	url := emojipediaURL + emojiURL.URL
	c := colly.NewCollector()
	name := strings.ReplaceAll(emojiURL.URL, "/", "")

	vendorOptions := []string{"Apple", "Microsoft Teams"}

	emojis := []Emoji{}

	c.OnHTML("section.vendor-list > ul > li", func(e *colly.HTMLElement) {
		el := e.DOM.Find("h2").FilterFunction(func(i int, s *goquery.Selection) bool {
			for _, vendor := range vendorOptions {
				if strings.Contains(s.Text(), vendor) {
					return true
				}
			}

			return false
		})

		vendor := el.Text()

		if vendor == "" {
			return
		}

		imageEl := e.DOM.Find("img")
		src := imageEl.AttrOr("src", "not found")

		if src == "not found" {
			return
		}

		vendor = strings.ToLower(vendor)
		slug := strings.ReplaceAll(vendor+"-"+name, " ", "-")

		if _, ok := emojisMap[slug]; ok {
			return
		}

		var data Emoji
		data.Name = name
		data.Slug = slug
		data.Category = emojiURL.Category
		data.ImageURL = src

		emojis = append(emojis, data)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.StatusCode, err)
	})

	c.OnRequest(func(r *colly.Request) {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
		c.IgnoreRobotsTxt = true

		fmt.Println("Company Visiting", r.URL.String())
	})

	c.Visit(url)

	return emojis
}

func readEmojis() []Emoji {
	absPath, _ := filepath.Abs(jsonPath)

	file, _ := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	emojis := []Emoji{}
	err := json.NewDecoder(file).Decode(&emojis)
	if err != nil {
		fmt.Println("Error Decoding:", err)
		return []Emoji{}
	}

	for _, emoji := range emojis {
		emojisMap[emoji.Slug] = emoji
	}

	return emojis
}

func ListEmojis() {
	emojis := readEmojis()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Index", "Name", "Category")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i, emoji := range emojis {
		category := strings.Split(emoji.Slug, "-")[0]
		tbl.AddRow(i, emoji.Name, category)
	}

	tbl.Print()
}

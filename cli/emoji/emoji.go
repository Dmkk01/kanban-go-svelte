package emoji

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/rodaine/table"
)

type Emoji struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

var jsonPath = "../frontend/src/utils/emojis.json"

// var imagesPath = "../frontend/public/images/emojis/"

var emojisMap map[string]Emoji = make(map[string]Emoji)

func GetEmojis() {
	emojis := readEmojis()

	for _, emoji := range emojis {
		emojisMap[emoji.Slug] = emoji
	}

	categories := []string{"apple", "microsoft-teams"}
	newEmojis := []Emoji{}

	for _, category := range categories {
		emojis := scrapeEmojis(category)
		newEmojis = append(newEmojis, emojis...)
	}

	totalEmojis := append(newEmojis, emojis...)

	fmt.Println("Saving emojis data...")
	fmt.Println("Total emojis:", len(totalEmojis))

	emojisJSON, _ := json.Marshal(totalEmojis)
	absPath, _ := filepath.Abs(jsonPath)

	os.WriteFile(absPath, emojisJSON, 0644)

	fmt.Println("Successfully wrote to file:", absPath)

	// fmt.Println("Saving emojis images...")
	// for _, emoji := range newEmojis {
	// 	absPath, _ = filepath.Abs(imagesPath + emoji.Slug + ".png")

	// 	img, _ := os.Create(absPath)
	// 	defer img.Close()

	// 	resp, err := http.Get(emoji.ImageURL)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}
	// 	defer resp.Body.Close()

	// 	io.Copy(img, resp.Body)
	// }

	// fmt.Println("Successfully saved emojis images")
}

func scrapeEmojis(category string) []Emoji {
	url := "https://emojipedia.org/" + category + "/"

	c := colly.NewCollector()

	emojis := []Emoji{}

	c.OnHTML(".emoji-grid > li > a", func(e *colly.HTMLElement) {
		var data Emoji
		href := e.Attr("href")
		src := e.ChildAttr("img", "data-src")

		if href == "" || src == "" {
			return
		}

		name := strings.ReplaceAll(href, "/", "")
		slug := category + "-" + name

		if _, ok := emojisMap[slug]; ok {
			return
		}

		data.Name = name
		data.Slug = slug
		data.URL = "https://emojipedia.org" + href
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

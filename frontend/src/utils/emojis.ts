import emojiJSON from './emojis.json' assert { type: 'json' }
import store from '../store/index'

export const emojis = emojiJSON as Emoji[]

export const getInitEmoji = () => {
  return emojis[0]
}

export const getEmojisByCategory = (category: EmojiCategory): Emoji[] => {
  return emojis.filter((emoji) => emoji.category === category)
}

export const getEmojiByCategoryVendor = (category: EmojiCategory, vendor: EmojiVendor): Emoji[] => {
  return emojis.filter((emoji) => emoji.category === category && emoji.slug.includes(vendor))
}

export const getEmojiBySlug = (slug: string): Emoji => {
  return emojis.find((emoji) => emoji.slug === slug) || emojis[0]
}

export const getEmojiURLBySlug = (slug: string): string => {
  const emoji = getEmojiBySlug(slug)

  return emoji.image_url
}

export const getEmojiByName = (name: string): Emoji => {
  return emojis.find((emoji) => emoji.name === name) || emojis[0]
}

export const searchEmojisByName = (name: string): Emoji[] => {
  return emojis.filter((emoji) => emoji.name.includes(name))
}

export const openEmojiSelector = (key: string, initEmoji: string) => {
  store.update((s) => {
    s.emojis.activeKey = key
    s.emojis.isOpen = true
    if (!s.emojis.keys.find((k) => k.key === key)) {
      s.emojis.keys.push({ key, emoji: initEmoji })
    }
    return s
  })
}

export const closeEmojiSelector = () => {
  store.update((s) => {
    s.emojis.isOpen = false
    s.emojis.activeKey = null
    return s
  })
}

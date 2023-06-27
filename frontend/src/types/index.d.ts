declare module 'prettier-plugin-tailwindcss'

type TokenResponse = {
  token: string
  expirationTime: string
}

type MessageResponse = {
  message: string
}

type StatusResponse = {
  status: string
}

type EmojiVendor = 'apple' | 'microsoft-teams'

type EmojiCategory = 'people' | 'nature' | 'food-drink' | 'activity' | 'travel-places' | 'objects' | 'symbols' | 'flags'

type Emoji = {
  name: string
  slug: string
  image_url: string
  category: EmojiCategory
}

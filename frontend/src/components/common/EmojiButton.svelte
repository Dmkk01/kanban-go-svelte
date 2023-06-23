<script lang="ts">
  import store from '@/store'
  import { getEmojiURLBySlug, getInitEmoji, openEmojiSelector } from '@/utils/emojis'

  export let emojiKey: string
  export let emojiSlug: string
  export let extraStyles: string = ''
  export let imageStyles: string = 'w-11/12 h-full object-contain'

  $: {
    if ($store.emojis.isOpen) {
      const possibleEmoji = $store.emojis.keys.find((item) => item.key === emojiKey)
      if (possibleEmoji) {
        emojiSlug = possibleEmoji.emoji
      }
    }
  }
</script>

<button
  type="button"
  on:click={() => openEmojiSelector(emojiKey, emojiSlug)}
  class={`px-1 flex items-center justify-center  rounded-lg ${extraStyles}`}
>
  <img
    src={getEmojiURLBySlug(emojiSlug)}
    alt={emojiSlug}
    class={imageStyles}
  />
</button>

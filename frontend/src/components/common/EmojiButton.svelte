<script lang="ts">
  import store from '../../store'
  import { getEmojiURLBySlug, openEmojiSelector } from '../../utils/emojis'
  export let emojiKey: string
  export let emojiSlug: string
  export let extraStyles: string = ''

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
    class="w-11/12 h-full object-contain"
  />
</button>

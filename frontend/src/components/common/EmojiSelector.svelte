<script lang="ts">
  import { closeEmojiSelector, getEmojiByCategoryVendor, searchEmojisByName } from '@/utils/emojis'
  import store from '@/store'
  import { clickOutside } from '@/directives/click-outside'

  let currentEmojis: Emoji[] = []
  let currentVendor: EmojiVendor = 'apple'
  let currentCategory: EmojiCategory = 'people'
  let selectedEmoji: string = ''
  let emojiSearch: string = ''

  $: {
    if (emojiSearch !== '') {
      currentEmojis = searchEmojisByName(emojiSearch)
    } else {
      currentEmojis = getEmojiByCategoryVendor(currentCategory, currentVendor)
    }
  }

  let dialog: HTMLDialogElement

  $: if (dialog && $store.emojis.isOpen) {
    if (dialog.open === false) {
      dialog.showModal()
    }
  }

  const categoryOptions: EmojiCategory[] = ['people', 'nature', 'food-drink', 'activity', 'travel-places', 'objects', 'symbols', 'flags']

  const onSelectEmoji = (emoji: string) => {
    selectedEmoji = emoji
    const newKeys = $store.emojis.keys.map((item) => {
      if (item.key === $store.emojis.activeKey) {
        item.emoji = emoji
      }

      return item
    })

    console.log(newKeys)

    $store.emojis.keys = newKeys
  }
</script>

<dialog
  bind:this={dialog}
  on:click|self={closeEmojiSelector}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeEmojiSelector()
    }
  }}
  on:close={closeEmojiSelector}
  class="absolute backdrop:bg-transparent p-0 rounded-md inset-0 border-none bg-white/30 flex justify-center items-center z-10"
>
  <div class="min-w-[30rem] static flex p-2 flex-col gap-2 bg-white/90 z-[10000] border border-tgray-200 rounded-lg">
    <div class={`flex flex-row gap-2 w-full justify-around`}>
      <button
        type="button"
        class={`w-full border rounded-md ${currentVendor === 'apple' ? 'border-tgray-200' : 'border-transparent'}`}
        on:click={() => (currentVendor = 'apple')}
      >
        <p class="font-medium text-lg text-center">Apple</p>
      </button>
      <button
        type="button"
        class={`w-full border rounded-md ${currentVendor === 'microsoft-teams' ? 'border-tgray-200' : 'border-transparent'}`}
        on:click={() => (currentVendor = 'microsoft-teams')}
      >
        <p class="font-medium text-lg text-center">Microsoft</p>
      </button>
    </div>
    <div class="w-auto mx-4 border border-tgray-200 rounded-md flex flex-row gap-2 px-2 py-1">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
        />
      </svg>
      <input
        type="text"
        class="w-full outline-none bg-transparent"
        placeholder="Search emoji"
        bind:value={emojiSearch}
      />
    </div>
    <div class="grid grid-cols-8 max-h-60 overflow-y-auto">
      {#each currentEmojis as emoji}
        <button
          type="button"
          on:click={() => onSelectEmoji(emoji.slug)}
          class={`w-full aspect-square rounded-md max-w-[2.5rem] border ${selectedEmoji === emoji.slug ? 'border-tgray-200' : 'border-transparent'}`}
        >
          <img
            src={emoji.image_url}
            alt={emoji.slug}
            class={`w-full h-full p-1`}
          />
        </button>
      {/each}
    </div>
    <hr class="border-tgray-200" />
    <div class="w-full flex flex-row gap-0 justify-around">
      {#each categoryOptions as category}
        <button
          type="button"
          on:click={() => (currentCategory = category)}
          class={`w-full aspect-square rounded-md max-w-[2.5rem] border ${currentCategory === category ? 'border-tgray-200' : 'border-transparent'}`}
        >
          <img
            src={`/images/emoji/category/${category}.png`}
            alt={category}
            class={`w-full h-full  p-1 `}
          />
        </button>
      {/each}
    </div>
    <!-- <hr class="border-tgray-200" />
    <div class="mx-auto">
      <button
        type="button"
        on:click={() => ($store.isEmojiOpen = false)}
        class="w-fit px-16 border border-tgray-200 text-xl shadow-lg text-black font-semibold bg-white/40 rounded-2xl py-1"
      >
        Close
      </button>
    </div> -->
  </div>
</dialog>

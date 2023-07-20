<script lang="ts">
  import { getEmojiURLBySlug } from '@/utils/emojis'

  export let item: string = 'board'
  export let onClickHandler: () => void
</script>

<div class="mx-auto flex w-3/4 max-w-md flex-col items-center justify-center gap-8 rounded-lg bg-white/30 px-4 py-8 drop-shadow-lg sm:w-full">
  <img
    src={getEmojiURLBySlug('microsoft-teams-loudly-crying-face')}
    alt="emoji"
    class="h-20 w-20 lg:h-24 lg:w-24"
  />
  <p class="text-center text-xl font-semibold lg:text-2xl">
    You don't have any {item}s yet.
  </p>
  <button
    type="button"
    class="flex flex-row items-center gap-0"
    on:click={onClickHandler}
  >
    <div class="flex aspect-square h-12 w-auto items-center justify-center rounded-lg bg-white/20 shadow-lg">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-11 w-11 text-white"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M12 4.5v15m7.5-7.5h-15"
        />
      </svg>
    </div>
    <div class="hidden h-9 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg md:flex">
      <p class="px-3 text-sm font-semibold text-tgray-600">Add new {item}</p>
    </div>
  </button>
</div>

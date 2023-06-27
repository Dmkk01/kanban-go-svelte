<script lang="ts">
  import store from '@/store'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { link } from 'svelte-routing'
  import { slide, fade } from 'svelte/transition'

  export let board: Board

  let isHovering = false
</script>

<a
  href={`/home/board/${board.id}`}
  use:link
  class="flex w-full flex-row items-center gap-0"
  on:mouseover={() => (isHovering = true)}
  on:mouseleave={() => (isHovering = false)}
  on:focus={() => (isHovering = true)}
  on:blur={() => (isHovering = false)}
>
  <div class="relative flex aspect-square h-14 w-auto items-center justify-center rounded-lg bg-white/50">
    <img
      src={getEmojiURLBySlug(board.emoji || '')}
      alt="sidebar-emoji"
      class="h-8 w-8"
    />

    {#if !$store.isSidebarOpen && isHovering}
      <div
        transition:fade
        class="absolute left-12 top-1/2 flex -translate-y-1/2 flex-row items-center"
      >
        <svg
          class="h-5 w-5"
          viewBox="0 0 9 10"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M8.33013 0.5L8.33013 9.16025L0.830127 4.83013L8.33013 0.5Z"
            fill="white"
          />
        </svg>

        <div class="-ml-[3px] w-fit rounded-md bg-white px-3 py-1.5">
          <p class="whitespace-nowrap text-[12px] font-semibold text-tgray-600">{board.name}</p>
        </div>
      </div>
    {/if}
  </div>

  <div class={`h-11 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg ${$store.isSidebarOpen ? 'flex' : 'hidden'}`}>
    <p class="pl-4 text-base font-semibold text-tgray-600">
      {board.name}
    </p>
    <div>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-6 w-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
        />
      </svg>
    </div>
  </div>
</a>

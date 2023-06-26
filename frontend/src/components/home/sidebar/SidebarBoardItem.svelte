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
  class="flex flex-row gap-0 w-full items-center"
  on:mouseover={() => (isHovering = true)}
  on:mouseleave={() => (isHovering = false)}
  on:focus={() => (isHovering = true)}
  on:blur={() => (isHovering = false)}
>
  <div class="flex items-center relative justify-center bg-white/50 h-14 w-auto aspect-square rounded-lg">
    <img
      src={getEmojiURLBySlug(board.emoji || '')}
      alt="sidebar-emoji"
      class="w-8 h-8"
    />

    {#if !$store.isSidebarOpen && isHovering}
      <div
        transition:fade
        class="absolute top-1/2 -translate-y-1/2 left-12 flex flex-row items-center"
      >
        <svg
          class="w-5 h-5"
          viewBox="0 0 9 10"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M8.33013 0.5L8.33013 9.16025L0.830127 4.83013L8.33013 0.5Z"
            fill="white"
          />
        </svg>

        <div class="py-1.5 px-3 bg-white rounded-md w-fit -ml-[3px]">
          <p class="whitespace-nowrap text-[12px] text-tgray-600 font-semibold">{board.name}</p>
        </div>
      </div>
    {/if}
  </div>

  {#if $store.isSidebarOpen}
    <div class="flex flex-row bg-white/20 items-center justify-between w-full h-11 shadow-lg rounded-r-lg">
      <p class="font-semibold text-base text-tgray-600 pl-4">
        {board.name}
      </p>
      <div>
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
            d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
          />
        </svg>
      </div>
    </div>
  {/if}
</a>

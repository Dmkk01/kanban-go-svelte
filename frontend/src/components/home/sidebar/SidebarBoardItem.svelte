<script lang="ts">
  import store from '../../../store'
  export let board: Board
</script>

<div class="flex flex-row gap-0 w-full items-center">
  <div class="flex items-center justify-center bg-white/50 h-14 w-auto aspect-square rounded-lg">
    <p class="text-2xl">
      {board.emoji}
    </p>
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
</div>

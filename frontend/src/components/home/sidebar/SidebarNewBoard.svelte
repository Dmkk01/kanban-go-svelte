<script lang="ts">
  import store from '@/store'
</script>

<button
  type="button"
  on:click={() => {
    $store.boardDrawer.isOpen = true
  }}
  class="flex flex-row gap-0 w-full items-center"
>
  <div class="flex items-center justify-center bg-white/20 h-14 w-auto aspect-square rounded-lg shadow-lg">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="w-12 h-12 text-white"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 4.5v15m7.5-7.5h-15"
      />
    </svg>
  </div>
  {#if $store.isSidebarOpen}
    <div class="flex flex-row bg-white/20 items-center justify-between w-full h-11 shadow-lg rounded-r-lg">
      <p class="font-semibold text-base text-tgray-600 pl-4">Add new board</p>
    </div>
  {/if}
</button>

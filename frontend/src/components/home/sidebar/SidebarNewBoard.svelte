<script lang="ts">
  import store from '@/store'
</script>

<button
  type="button"
  on:click={() => {
    $store.isSidebarOpen = false
    $store.boardDrawer.isOpen = true
  }}
  class="flex w-full flex-row items-center gap-0"
>
  <div class="flex aspect-square h-14 w-auto items-center justify-center rounded-lg bg-white/20 shadow-lg">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="h-12 w-12 text-white"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 4.5v15m7.5-7.5h-15"
      />
    </svg>
  </div>
  {#if $store.isSidebarOpen}
    <div class="flex h-11 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg">
      <p class="pl-4 text-base font-semibold text-tgray-600">Add new board</p>
    </div>
  {/if}
</button>

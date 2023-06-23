<script lang="ts">
  import { Link } from 'svelte-routing'
  import store from '@/store'
</script>

<div class="flex flex-col gap-6 w-full">
  <div class="flex flex-row gap-0 w-full items-center">
    <Link
      to="/home/settings"
      class="flex items-center shadow-lg justify-center bg-white/50 h-14 w-auto aspect-square rounded-lg"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-9 h-9"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    </Link>
    {#if $store.isSidebarOpen}
      <div class="flex flex-row bg-white/20 items-center justify-between w-full h-11 shadow-lg rounded-r-lg">
        <p class="font-semibold text-base text-tgray-600 pl-4">Others</p>
      </div>
    {/if}
  </div>
  <button
    type="button"
    class="flex flex-row gap-0 w-full items-center"
    on:click={() => {
      $store.isSidebarOpen = !$store.isSidebarOpen
    }}
  >
    <div class="flex items-center shadow-lg justify-center bg-white/50 h-14 w-auto aspect-square rounded-lg">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-9 h-9"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M18.75 19.5l-7.5-7.5 7.5-7.5m-6 15L5.25 12l7.5-7.5"
        />
      </svg>
    </div>
    {#if $store.isSidebarOpen}
      <div class="flex flex-row bg-white/20 items-center justify-between w-full h-11 shadow-lg rounded-r-lg">
        <p class="font-semibold text-base text-tgray-600 pl-4">Hide sidebar</p>
      </div>
    {/if}
  </button>
</div>

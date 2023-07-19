<script lang="ts">
  import store from '@/store'
</script>

<div class="flex w-full flex-col gap-6">
  <div class="flex w-full flex-row items-center gap-0">
    <a
      href="/home/settings"
      class="flex aspect-square h-14 w-auto items-center justify-center rounded-lg bg-white/50 shadow-lg"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-9 w-9"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M8.625 12a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H8.25m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0H12m4.125 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm0 0h-.375M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    </a>
    {#if $store.isSidebarOpen}
      <div class="flex h-11 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg">
        <p class="pl-4 text-base font-semibold text-tgray-600">Others</p>
      </div>
    {/if}
  </div>
  <button
    type="button"
    class="flex w-full flex-row items-center gap-0"
    on:click={() => {
      $store.isSidebarOpen = !$store.isSidebarOpen
    }}
  >
  {#if !$store.isSidebarOpen}
    <div class="flex aspect-square h-14 w-auto items-center justify-center rounded-lg bg-white/50 shadow-lg">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-9 w-9"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M11.25 4.5l7.5 7.5-7.5 7.5m-6-15l7.5 7.5-7.5 7.5"
        />
      </svg>
    </div>
    {/if}
    {#if $store.isSidebarOpen}
      <div class="flex aspect-square h-14 w-auto items-center justify-center rounded-lg bg-white/50 shadow-lg">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="h-9 w-9"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M18.75 19.5l-7.5-7.5 7.5-7.5m-6 15L5.25 12l7.5-7.5"
          />
        </svg>
      </div>
      <div class="flex h-11 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg">
        <p class="pl-4 text-base font-semibold text-tgray-600">Hide sidebar</p>
      </div>
    {/if}
  </button>
</div>

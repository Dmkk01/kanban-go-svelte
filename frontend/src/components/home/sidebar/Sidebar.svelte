<script lang="ts">
  import { useQuery } from '@sveltestack/svelte-query'
  import UserAPI from '@/api/user'
  import { navigate } from 'svelte-routing/src/history'
  import BoardsAPI from '@/api/board'
  import SidebarBoardItem from './SidebarBoardItem.svelte'
  import SidebarNewBoard from './SidebarNewBoard.svelte'
  import SidebarOthers from './SidebarOthers.svelte'
  import store from '@/store'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { fade, fly, slide } from 'svelte/transition'

  export let boardID: number = 0

  const settings = useQuery('settings', UserAPI.getUserSettings, {
    onSuccess: async (data) => {
      if (data.app_name === '') {
        navigate('/getting-started')
      }

      if (data.app_emoji !== '' && data.app_name !== '') {
        localStorage.setItem(
          'default-join-data',
          JSON.stringify({
            name: data.app_name,
            emoji: data.app_emoji,
          })
        )
      }
    },
  })

  const boards = useQuery('boards', BoardsAPI.getBoards)
</script>

<div class="flex flex-row items-center gap-3 px-3 py-3 lg:hidden">
  <button
    type="button"
    on:click={() => {
      $store.isSidebarOpen = !$store.isSidebarOpen
    }}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2.5"
      stroke="currentColor"
      class="h-10 w-10"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
      />
    </svg>
  </button>
  <div class={`aspect-square h-full w-14 p-1`}>
    <img
      src={getEmojiURLBySlug($settings.data?.app_emoji || '')}
      alt="sidebar-emoji"
    />
  </div>
  <h1 class="text-4xl font-bold">
    {$settings.data?.app_name}
  </h1>
</div>

{#if $store.isSidebarOpen}
  <div
    transition:fly={{ x: -200 }}
    class={`fixed z-50 flex  min-h-screen w-[90%] max-w-md flex-col justify-between gap-6 bg-tgray-200 px-5 py-8 lg:hidden`}
  >
    {#if !$settings.isLoading && !$boards.isLoading}
      <div class="flex w-full flex-col gap-6">
        <div class="flex flex-row items-center gap-3">
          <div class={`aspect-square h-full w-14 p-1`}>
            <img
              src={getEmojiURLBySlug($settings.data?.app_emoji || '')}
              alt="sidebar-emoji"
            />
          </div>
          <h1 class="text-4xl font-bold">
            {$settings.data?.app_name}
          </h1>
        </div>
        <div class="flex w-full flex-col gap-4">
          <p class="text-sm font-bold text-tgray-600">
            ALL BOARDS ({$boards.data?.length || 0})
          </p>
          {#each $boards.data || [] as board (board.id)}
            <div class="relative">
              {#if boardID === board.id}
                <div
                  transition:fly={{ x: -200 }}
                  class="absolute -left-3 top-1/2 h-2 w-2 -translate-y-1/2 rounded-full bg-white"
                />
              {/if}
              <SidebarBoardItem {board} />
            </div>
          {/each}
          <SidebarNewBoard />
        </div>
      </div>
      <SidebarOthers />
    {/if}
  </div>
{/if}

<div
  class={`hidden h-screen w-full flex-col justify-between gap-6 bg-white/20 py-8 shadow-xl lg:flex ${
    !$store.isSidebarOpen ? 'max-w-fit px-4' : 'max-w-xs px-6 '
  }`}
>
  {#if !$settings.isLoading && !$boards.isLoading}
    <div class="flex w-full flex-col gap-6">
      <div class="mb-10 flex flex-row items-center gap-3">
        <div class={`aspect-square h-full w-14 p-1 ${!$store.isSidebarOpen ? 'h-auto rounded-lg bg-white/50 ' : ''}`}>
          <img
            src={getEmojiURLBySlug($settings.data?.app_emoji || '')}
            alt="sidebar-emoji"
          />
        </div>
        {#if $store.isSidebarOpen}
          <h1 class="text-4xl font-bold">
            {$settings.data?.app_name}
          </h1>
        {/if}
      </div>
      <div class="flex w-full flex-col gap-4">
        {#if $store.isSidebarOpen}
          <p class="text-sm font-bold text-tgray-600">
            ALL BOARDS ({$boards.data?.length || 0})
          </p>
        {/if}
        {#each $boards.data || [] as board (board.id)}
          <div class="relative">
            {#if boardID === board.id}
              <div
                transition:fly={{ x: -200 }}
                class="absolute -left-3 top-1/2 h-2 w-2 -translate-y-1/2 rounded-full bg-white"
              />
            {/if}
            <SidebarBoardItem {board} />
          </div>
        {/each}
        <SidebarNewBoard />
      </div>
    </div>
    <SidebarOthers />
  {/if}
</div>

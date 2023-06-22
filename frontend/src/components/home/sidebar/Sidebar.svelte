<script lang="ts">
  import { useQuery } from '@sveltestack/svelte-query'
  import UserAPI from '../../../api/user'
  import { navigate } from 'svelte-routing/src/history'
  import BoardsAPI from '../../../api/board'
  import SidebarBoardItem from './SidebarBoardItem.svelte'
  import SidebarNewBoard from './SidebarNewBoard.svelte'
  import SidebarOthers from './SidebarOthers.svelte'
  import store from '../../../store'

  const settings = useQuery('settings', UserAPI.getUserSettings, {
    onSuccess: async (data) => {
      if (data.app_name === '') {
        navigate('/getting-started')
      }
    },
  })

  const boards = useQuery('boards', BoardsAPI.getBoards)
</script>

<div
  class={`w-full h-screen bg-white/20 shadow-xl py-8 flex flex-col justify-between gap-6 ${
    !$store.isSidebarOpen ? 'max-w-fit px-4' : 'max-w-xs px-6 '
  }`}
>
  {#if !$settings.isLoading && !$boards.isLoading}
    <div class="flex flex-col gap-6 w-full">
      <div class="flex flex-row gap-3 items-center mb-10">
        <p
          class={`text-3xl text-center ${
            !$store.isSidebarOpen ? 'bg-white/50 h-14 w-auto aspect-square rounded-lg flex items-center justify-center' : ''
          }`}
        >
          {$settings.data.app_emoji}
        </p>
        {#if $store.isSidebarOpen}
          <h1 class="text-4xl font-bold">
            {$settings.data.app_name}
          </h1>
        {/if}
      </div>
      <div class="flex flex-col w-full gap-4">
        {#if $store.isSidebarOpen}
          <p class="font-bold text-sm text-tgray-600">
            ALL BOARDS ({$boards.data.length || 0})
          </p>
        {/if}
        {#each $boards.data as board}
          <SidebarBoardItem {board} />
        {/each}
        <SidebarNewBoard />
      </div>
    </div>
    <SidebarOthers />
  {/if}
</div>

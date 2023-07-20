<script lang="ts">
  import BoardsAPI from '@/api/board'
  import UserAPI from '@/api/user'
  import MainLayout from '@/layout/MainLayout.svelte'
  import NoItems from '@/components/common/NoItems.svelte'
  import { useQuery } from '@sveltestack/svelte-query'
  import { navigate } from 'svelte-routing/src/history'
  import store from '@/store'

  const settings = useQuery('settings', UserAPI.getUserSettings, {
    refetchOnMount: true,
    onSuccess: async (data) => {
      if (data.primary_board_id) {
        navigate(`/home/board/${data.primary_board_id}`, { replace: true, state: {} })
      }
    },
  })

  const boards = useQuery('boards', BoardsAPI.getBoards)
</script>

<MainLayout>
  <div>
    {#if $settings.isLoading || $boards.isLoading}
      <p>Loading...</p>
    {:else if $boards.data && $boards.data.length < 1}
      <div class="h-full w-full pt-10">
        <NoItems
          item="board"
          onClickHandler={() => {
            $store.isSidebarOpen = false
            $store.boardDrawer.isOpen = true
          }}
        />
      </div>
    {/if}
  </div>
</MainLayout>

<script lang="ts">
  import AuthAPI from '@/api/auth'
  import UserAPI from '@/api/user'
  import BoardDrawer from '@/components/drawers/board/BoardDrawer.svelte'
  import ColumnDrawer from '@/components/drawers/column/ColumnDrawer.svelte'
  import TaskDrawer from '@/components/drawers/task/add-edit/TaskEditDrawer.svelte'
  import Sidebar from '@/components/home/sidebar/Sidebar.svelte'
  import Loading from '@/components/common/Loading.svelte'
  import store from '@/store'
  import { useQuery } from '@sveltestack/svelte-query'
  import { navigate } from 'svelte-routing'
  import TaskViewDrawer from '@/components/drawers/task/view/TaskViewDrawer.svelte'

  export let boardID: number = 0

  const user = useQuery('user-check', async () => await UserAPI.getUser(), {
    // onError: (err) => {
    //   navigate('/')
    // },
  })

  const refreshMutation = useQuery('refresh', () => AuthAPI.refresh(), {
    refetchOnWindowFocus: false,
    onSuccess: async () => {
      console.log('refreshed')
    },
  })

  setInterval(() => {
    $refreshMutation.refetch()
  }, 1000 * 60 * 2)
</script>

<div class="relative flex h-screen w-full flex-col overflow-hidden bg-[#C0C2CC] lg:flex-row">
  <Sidebar {boardID} />
  {#if !$user.isLoading}
    <div class="h-screen w-full overflow-hidden px-0 py-2 sm:px-4">
      <slot />
    </div>
  {:else}
    <div class="flex w-full items-center justify-center">
      <Loading />
    </div>
  {/if}
  {#if $store.boardDrawer.isOpen}
    <BoardDrawer />
  {/if}
  {#if $store.columnDrawer.isOpen}
    <ColumnDrawer />
  {/if}
  {#if $store.taskDrawer.isOpen}
    <TaskDrawer />
  {/if}
  {#if $store.taskDrawerView.isOpen}
    <TaskViewDrawer />
  {/if}
</div>

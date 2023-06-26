<script lang="ts">
  import UserAPI from '@/api/user'
  import BoardDrawer from '@/components/drawers/board/BoardDrawer.svelte'
  import ColumnDrawer from '@/components/drawers/column/ColumnDrawer.svelte'
  import TaskDrawer from '@/components/drawers/task/TaskDrawer.svelte'
  import Sidebar from '@/components/home/sidebar/Sidebar.svelte'
  import store from '@/store'
  import { useQuery } from '@sveltestack/svelte-query'
  import { navigate } from 'svelte-routing'

  export let boardID: number = 0

  const user = useQuery('user-check', async () => await UserAPI.getUser(), {
    onError: (err) => {
      console.log(err)
      console.log('BIG ERRROR')
      navigate('/')
    },
  })
</script>

<div class="w-full h-screen flex flex-row bg-[#C0C2CC] relative overflow-hidden">
  {#if !$user.isLoading}
    <Sidebar {boardID} />

    <div class="px-4 py-2 h-screen w-full overflow-hidden">
      <slot />
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
</div>

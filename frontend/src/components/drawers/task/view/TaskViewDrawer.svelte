<script lang="ts">
  import ColumnAPI from '@/api/column'
  import TaskAPI from '@/api/task'
  import store from '@/store'
  import { useQuery } from '@sveltestack/svelte-query'
  import { fly } from 'svelte/transition'
  import TaskViewSubTasks from './TaskViewSubTasks.svelte'
  import TaskViewLinks from './TaskViewLinks.svelte'
  import TaskViewInfos from './TaskViewInfos.svelte'
  import TaskViewTags from './TaskViewTags.svelte'
  import TaskViewDescription from './TaskViewDescription.svelte'
  import TaskViewHeader from './TaskViewHeader.svelte'

  const task = useQuery(
    ['task', $store.taskDrawerView.taskID],
    async () => {
      const task = await TaskAPI.getTask($store.taskDrawerView.taskID as number)
      const column = await ColumnAPI.getColumn(task.column_id)

      return { task, column }
    },
    {
      enabled: $store.taskDrawerView.isOpen,
      onSuccess: (data) => {
        console.log(data)
      },
    }
  )

  const closeDrawer = () => {
    $store.taskDrawerView.isOpen = false
  }
</script>

<div
  on:click|self={closeDrawer}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeDrawer()
    }
  }}
  class="absolute inset-0 overflow-hidden backdrop-blur-sm"
  transition:fly={{ x: 200, duration: 1000 }}
>
  <div class="absolute bottom-0 right-0 top-0 flex min-h-screen w-full max-w-md flex-col gap-2 bg-white/90 px-6 py-3 drop-shadow-lg">
    {#if $task.data && $task.data.task && !$task.isLoading}
      <div class="flex w-full flex-col gap-6">
        <TaskViewHeader
          data={$task.data}
          {closeDrawer}
        />
        <TaskViewTags tags={$task.data.task.tags} />
        <TaskViewDescription description={$task.data.task.description} />
        <TaskViewInfos
          created_at={$task.data.task.created_at}
          due_date={$task.data.task.due_date}
          time_needed={$task.data.task.time_needed}
        />
        <TaskViewSubTasks
          subtasks={$task.data.task.sub_tasks}
          boardID={$task.data.task.board_id}
        />
        <TaskViewLinks links={$task.data.task.links} />
      </div>
    {/if}
  </div>
</div>

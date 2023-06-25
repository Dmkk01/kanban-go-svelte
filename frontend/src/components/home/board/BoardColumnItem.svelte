<script lang="ts">
  import ColumnAPI from '@/api/column'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { flip } from 'svelte/animate'
  import { dndzone, TRIGGERS } from 'svelte-dnd-action'
  import BoardTaskItem from './BoardTaskItem.svelte'
  import TaskAPI from '@/api/task'
  import store from '@/store'
  import { createEventDispatcher } from 'svelte'

  export let column: ColumnBoardFull | undefined

  const dispatch = createEventDispatcher<{ 'move-task-to-column': number }>()

  $: {
    if (column) {
      taskItems = column.tasks
    }
  }

  // export let columnID: number = 0
  // let currColumnID = columnID
  // const column = useQuery(['column', currColumnID], async () => await ColumnAPI.getColumn(currColumnID), {})
  // const tasks = useQuery(['column', currColumnID, 'tasks'], async () => await ColumnAPI.getTasks(currColumnID), {
  //   onSuccess: (data) => {
  //     taskItems = data
  //   },
  // })
  // $: {
  //   if (currColumnID !== columnID) {
  //     currColumnID = columnID
  //     void $column.refetch()
  //     void $tasks.refetch()
  //   }
  // }

  let taskItems: Task[] = []

  const handleSort = (e: CustomEvent<DndEvent<Task>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })

    taskItems = newItems
  }

  const handleSortFinalized = (e: CustomEvent<DndEvent<Task>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })
    const toUpdate: { task_id: number; position: number; column_id: number }[] = []

    taskItems = newItems

    const columnID = column?.column.id as number

    newItems.forEach((item) => {
      if (typeof item.id === 'number') {
        toUpdate.push({ task_id: item.id, position: item.position, column_id: columnID })
      }
    })

    if (toUpdate.length > 0) {
      void $positionMutation.mutate(toUpdate)
    }
  }

  const positionMutation = useMutation(async (data: { task_id: number; position: number; column_id: number }[]) => {
    return await TaskAPI.taskUpdatePosition(data)
  })
</script>

<div class="bg-white/30 shadow-lg rounded-lg py-4 px-2 w-[17rem]">
  {#if column}
    <div class="flex flex-row justify-between items-center pb-6">
      <div class="flex flex-row gap-2 items-center">
        <img
          src={getEmojiURLBySlug(column.column.emoji)}
          alt="column-emoji"
          class="w-5 h-5"
        />
        <h3 class="font-bold text-sm">
          {column.column.name} ({taskItems.length ?? 0})
        </h3>
      </div>
      <button type="button">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          class="w-6 h-6 text-tgray-600"
          ><path
            fill="currentColor"
            d="M8.5 17c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7-10c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2zm-7 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm0 7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm-7-14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2z"
          /></svg
        >
      </button>
    </div>

    <div
      use:dndzone={{ items: taskItems, flipDurationMs: 200, type: `board-column-tasks` }}
      on:consider={handleSort}
      on:finalize={handleSortFinalized}
      class="flex flex-col gap-3 w-full pb-4 min-h-[7rem]"
    >
      {#each taskItems as task (task.id)}
        <div
          animate:flip={{ duration: 200 }}
          draggable="true"
        >
          <BoardTaskItem {task} />
        </div>
      {/each}
    </div>
    <div class="w-full flex justify-center items-center">
      <button
        type="button"
        on:click={() => {
          $store.taskDrawer.isOpen = true
          $store.taskDrawer.ids.board = column?.column.board_id ?? 0
          $store.taskDrawer.ids.column = column?.column.id ?? 0
        }}
        class="p-0.5 rounded-lg shadow-lg bg-white/20"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-8 h-8 text-white"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M12 4.5v15m7.5-7.5h-15"
          />
        </svg>
      </button>
    </div>
  {/if}
</div>

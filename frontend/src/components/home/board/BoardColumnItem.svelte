<script lang="ts">
  import ColumnAPI from '@/api/column'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { flip } from 'svelte/animate'
  import { dndzone, TRIGGERS, SOURCES } from 'svelte-dnd-action'
  import BoardTaskItem from './BoardTaskItem.svelte'
  import TaskAPI from '@/api/task'
  import store from '@/store'
  import { createEventDispatcher } from 'svelte'
  import { longpress } from '@/directives/longpress'

  export let column: ColumnBoardFull | undefined

  $: {
    if (column) {
      taskItems = column.tasks
    }
  }

  let taskItems: Task[] = []

  let dragDisabled = true

  const handleSort = (e: CustomEvent<DndEvent<Task>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })

    taskItems = newItems

    if (e.detail.info.source === SOURCES.KEYBOARD && e.detail.info.trigger === TRIGGERS.DRAG_STOPPED) {
      dragDisabled = true
    }
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

    if (e.detail.info.source === SOURCES.POINTER) {
      dragDisabled = true
    }
  }

  const positionMutation = useMutation(async (data: { task_id: number; position: number; column_id: number }[]) => {
    return await TaskAPI.taskUpdatePosition(data)
  })

  const openEditColumnDrawer = () => {
    if (!column) return
    $store.columnDrawer = {
      isOpen: true,
      columnID: column?.column.id,
      boardID: column?.column.board_id,
    }
  }
</script>

<div class="w-[90vw] md:w-[20rem] xl:w-[21rem] rounded-lg bg-white/30 px-2 py-4 shadow-lg">
  {#if column}
    <div class="flex flex-row items-center justify-between pb-6">
      <div class="flex flex-row items-center gap-2">
        <img
          src={getEmojiURLBySlug(column.column.emoji)}
          alt="column-emoji"
          class="h-5 w-5"
        />
        <h3 class="text-md font-bold">
          {column.column.name} ({taskItems.length ?? 0})
        </h3>
      </div>
      <button
        type="button"
        on:click={openEditColumnDrawer}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          class="h-7 w-7 text-tgray-600"
          ><path
            fill="currentColor"
            d="M8.5 17c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7-10c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2zm-7 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm0 7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm-7-14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2z"
          /></svg
        >
      </button>
    </div>

    <div
      use:dndzone={{ items: taskItems, flipDurationMs: 200, type: `board-column-tasks`, dragDisabled }}
      on:consider={handleSort}
      on:finalize={handleSortFinalized}
      class="flex min-h-[7rem] w-full flex-col gap-3 pb-4"
    >
      {#each taskItems as task (task.id)}
        <div
          animate:flip={{ duration: 200 }}
          draggable={!dragDisabled}
          use:longpress={{ category: `board-column-tasks` }}
          on:long={(e) => {
            if (e.detail.category === 'board-column-tasks') {
              dragDisabled = false
            }
          }}
          on:mouseup={(e) => {
            if (dragDisabled) {
              $store.taskDrawerView.taskID = task.id
              $store.taskDrawerView.isOpen = true
            }
          }}
          on:touchend={(e) => {
            if (dragDisabled) {
              $store.taskDrawerView.taskID = task.id
              $store.taskDrawerView.isOpen = true
            }
          }}
        >
          <BoardTaskItem {task} />
        </div>
      {/each}
    </div>
    <div class="flex w-full items-center justify-center">
      <button
        type="button"
        on:click={() => {
          $store.taskDrawer.isOpen = true
          $store.taskDrawer.ids.task = null
          $store.taskDrawer.ids.board = column?.column.board_id ?? 0
          $store.taskDrawer.ids.column = column?.column.id ?? 0
        }}
        class="rounded-lg bg-white/20 p-0.5 shadow-lg"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="h-10 w-10 text-white"
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

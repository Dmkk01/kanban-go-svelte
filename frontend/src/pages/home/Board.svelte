<script lang="ts">
  import BoardsAPI from '@/api/board'
  import BoardColumnItem from '@/components/home/board/BoardColumnItem.svelte'
  import MainLayout from '@/layout/MainLayout.svelte'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { onMount } from 'svelte'
  import store from '@/store'
  import { flip } from 'svelte/animate'
  import { dndzone } from 'svelte-dnd-action'

  export let id: string = '0'
  let boardID = parseInt(id)

  let columnItems: ColumnBoard[] = []

  const board = useQuery(`board-${boardID}`, async () => await BoardsAPI.getBoardById(boardID), {
    onSuccess: () => {
      void $columns.refetch()
    },
  })
  const columns = useQuery(`board-${boardID}-columns`, async () => await BoardsAPI.getColumns(boardID), {
    onSuccess: (data) => {
      columnItems = data
    },
  })

  $: {
    if (boardID !== parseInt(id)) {
      boardID = parseInt(id)
      void $board.refetch()
      void $columns.refetch()
    }
  }

  const handleSort = (e: CustomEvent<DndEvent<ColumnBoard>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })

    const toUpdate: { id: number; position: number }[] = []

    newItems.forEach((item, index) => {
      if (item.position !== columnItems[index].position) {
        toUpdate.push({ id: item.id, position: item.position })
      }
    })

    columnItems = newItems

    if (toUpdate.length > 0) {
      void $positionMutation.mutate(toUpdate)
    }
  }

  const positionMutation = useMutation(async (data: { id: number; position: number }[]) => {
    return await BoardsAPI.updateColumnPosition(boardID, data)
  })
</script>

<MainLayout {boardID}>
  <div class="flex flex-col gap-2">
    <div class="flex flex-row justify-between items-center px-6 pt-4">
      {#if !$board.isLoading && $board.data}
        <div class="flex flex-row gap-2 items-center">
          <img
            src={getEmojiURLBySlug($board.data.emoji)}
            alt="emoji"
            class="w-8 h-8 rounded-full"
          />
          <h1 class="font-bold text-xl">
            {$board.data.name}
          </h1>
        </div>
        <button
          type="button"
          class="flex flex-row gap-0 items-center"
          on:click={() => {
            $store.columnDrawer.boardID = boardID
            $store.columnDrawer.isOpen = true
            $store.columnDrawer.columnID = null
          }}
        >
          <div class="flex items-center justify-center bg-white/20 h-12 w-auto aspect-square rounded-lg shadow-lg">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-11 h-11 text-white"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
          </div>
          <div class="flex flex-row bg-white/20 items-center justify-between w-full h-9 shadow-lg rounded-r-lg">
            <p class="font-semibold text-sm text-tgray-600 px-3">Add new column</p>
          </div>
        </button>
      {/if}
    </div>
    {#if !$columns.isLoading && columnItems}
      <div
        use:dndzone={{ items: columnItems, flipDurationMs: 200, type: 'board-column' }}
        on:consider={handleSort}
        on:finalize={handleSort}
        class="flex flex-row gap-8 w-full px-6 pb-4 overflow-x-auto"
      >
        {#each columnItems as column (column.id)}
          <div
            animate:flip={{ duration: 200 }}
            draggable="true"
          >
            <BoardColumnItem columnID={column.id} />
          </div>
        {/each}
      </div>
    {/if}
    <p>
      {JSON.stringify(columnItems)}
    </p>
  </div>
</MainLayout>

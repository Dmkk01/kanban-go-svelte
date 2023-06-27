<script lang="ts">
  import BoardsAPI from '@/api/board'
  import BoardColumnItem from '@/components/home/board/BoardColumnItem.svelte'
  import MainLayout from '@/layout/MainLayout.svelte'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { longpress } from '@/directives/longpress'
  import store from '@/store'
  import { flip } from 'svelte/animate'
  import { dndzone, SOURCES, TRIGGERS } from 'svelte-dnd-action'

  export let id: string = '0'
  let boardID = parseInt(id)

  let columnItems: ColumnBoard[] = []

  const board = useQuery(`board-${boardID}`, async () => await BoardsAPI.getBoardFull(boardID), {
    refetchOnWindowFocus: false,
    onSuccess: (data) => {
      console.log(data)
      columnItems = data.columns.map((item) => item.column)
    },
  })
  // const columns = useQuery(`board-${boardID}-columns`, async () => await BoardsAPI.getColumns(boardID), {
  //   onSuccess: (data) => {
  //     columnItems = data
  //   },
  // })

  $: {
    if (boardID !== parseInt(id)) {
      boardID = parseInt(id)
      void $board.refetch()
      // void $columns.refetch()
    }
  }

  let dragDisabled = true

  const handleSort = (e: CustomEvent<DndEvent<ColumnBoard>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })

    columnItems = newItems

    if (e.detail.info.source === SOURCES.KEYBOARD && e.detail.info.trigger === TRIGGERS.DRAG_STOPPED) {
      dragDisabled = true
    }
  }

  const handleSortFinalized = (e: CustomEvent<DndEvent<ColumnBoard>>) => {
    const newItems = e.detail.items.map((item, index) => {
      item.position = index + 1
      return item
    })

    const toUpdate: { id: number; position: number }[] = []

    newItems.forEach((item, index) => {
      if (typeof item.id === 'number') {
        toUpdate.push({ id: item.id, position: item.position })
      }
    })

    columnItems = newItems

    if (toUpdate.length > 0) {
      void $positionMutation.mutate(toUpdate)
    }

    if (e.detail.info.source === SOURCES.POINTER) {
      dragDisabled = true
    }
  }

  const positionMutation = useMutation(async (data: { id: number; position: number }[]) => {
    return await BoardsAPI.updateColumnPosition(boardID, data)
  })
</script>

<MainLayout {boardID}>
  <div class="flex flex-col gap-2">
    <div class="flex flex-row items-center justify-between px-6 pt-4">
      {#if !$board.isLoading && $board.data}
        <div class="flex flex-row items-center gap-2">
          <img
            src={getEmojiURLBySlug($board.data.board.emoji)}
            alt="emoji"
            class="h-8 w-8 rounded-full"
          />
          <h1 class="text-xl font-bold">
            {$board.data.board.name}
          </h1>
        </div>
        <button
          type="button"
          class="flex flex-row items-center gap-0"
          on:click={() => {
            $store.columnDrawer.boardID = boardID
            $store.columnDrawer.isOpen = true
            $store.columnDrawer.columnID = null
          }}
        >
          <div class="flex aspect-square h-12 w-auto items-center justify-center rounded-lg bg-white/20 shadow-lg">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-11 w-11 text-white"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
          </div>
          <div class="flex h-9 w-full flex-row items-center justify-between rounded-r-lg bg-white/20 shadow-lg">
            <p class="px-3 text-sm font-semibold text-tgray-600">Add new column</p>
          </div>
        </button>
      {/if}
    </div>
    {#if columnItems && $board.data}
      <div
        use:dndzone={{ items: columnItems, flipDurationMs: 200, type: 'board-column', dragDisabled }}
        on:consider={handleSort}
        on:finalize={handleSortFinalized}
        class="flex w-full flex-row gap-8 overflow-x-auto px-6 pb-4"
      >
        {#each columnItems as column (column.id)}
          <div
            animate:flip={{ duration: 200 }}
            draggable={!dragDisabled}
            use:longpress={{ category: 'board-column' }}
            on:long={(e) => {
              console.log(e)
              if (e.detail.category === 'board-column') {
                dragDisabled = false
              }
            }}
          >
            <BoardColumnItem column={$board.data.columns.find((item) => item.column.id === column.id)} />
          </div>
        {/each}
      </div>
    {/if}
  </div>
</MainLayout>

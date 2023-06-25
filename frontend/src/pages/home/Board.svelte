<script lang="ts">
  import BoardsAPI from '@/api/board'
  import BoardColumnItem from '@/components/home/board/BoardColumnItem.svelte'
  import MainLayout from '@/layout/MainLayout.svelte'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useQuery } from '@sveltestack/svelte-query'
  import { onMount } from 'svelte'
  import store from '@/store'

  export let id: string = '0'
  let boardID = parseInt(id)

  const board = useQuery(`board-${boardID}`, async () => await BoardsAPI.getBoardById(boardID), {
    onSuccess: () => {
      $columns.refetch()
    },
  })
  const columns = useQuery(`board-${boardID}-columns`, async () => await BoardsAPI.getColumns(boardID), {})

  $: {
    if (boardID !== parseInt(id)) {
      boardID = parseInt(id)
      $board.refetch()
      $columns.refetch()
    }
  }
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
    {#if !$columns.isLoading && $columns.data}
      <div class="flex flex-row gap-8 w-full overflow-x-auto px-6 pb-4">
        {#each $columns.data as column}
          <div class="relative">
            <BoardColumnItem columnID={column.id} />
          </div>
        {/each}
      </div>
    {/if}
  </div>
</MainLayout>

<script lang="ts">
  import TagAPI from '@/api/tag'
  import { useMutation, useQueryClient } from '@sveltestack/svelte-query'
  import { createEventDispatcher } from 'svelte'

  export let tag: Tag
  export let boardID: number = 0

  const colorOptions = ['#A2E8DD', '#DFFFD6', '#ABE4FF', '#FFDAAB', '#FFDDFF', '#D4DCFF', '#E8D7FF', '#E8F4C8']
  const queryClient = useQueryClient()
  const dispatch = createEventDispatcher<{ 'delete-tag-from-board': void }>()

  const updateMutation = useMutation(
    ({ title, color, id }: { title: string; color: string; id: number }) => TagAPI.updateTag(boardID, id, color, title),
    {
      onSuccess: async () => {
        await queryClient.invalidateQueries(`board-tags-${boardID}`)
      },
    }
  )

  const deleteMutation = useMutation(() => TagAPI.deleteTag(tag.id), {
    onSuccess: async () => {
      await queryClient.invalidateQueries(`board-tags-${boardID}`)
    },
  })

  const saveTagHandler = () => {
    $updateMutation.mutateAsync({
      id: tag.id,
      title: tag.title,
      color: tag.color,
    })
  }

  const handleDeleteTagFromBoard = () => {
    $deleteMutation.mutateAsync()
    dispatch('delete-tag-from-board')
  }
</script>

<div class="my-2 flex flex-col items-center gap-1 rounded-md border border-tgray-200 px-2 py-2">
  <div class="flex w-full flex-row items-center gap-2">
    <input
      type="text"
      class="w-full rounded-md border border-tgray-600 px-1 py-0.5 text-sm"
      bind:value={tag.title}
      placeholder="Tag name"
      required
    />
    <button
      type="button"
      on:click={saveTagHandler}
      class="flex aspect-square h-6 w-6 items-center justify-center rounded-md border border-tgray-600"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2.5"
        stroke="currentColor"
        class="h-4 w-4"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M4.5 12.75l6 6 9-13.5"
        />
      </svg>
    </button>
  </div>
  <div class="flex w-full flex-row items-center justify-around gap-2">
    <div class="flex flex-row gap-1.5 px-2 py-1">
      {#each colorOptions as colorOption}
        <button
          type="button"
          class="flex h-6 w-6 items-center justify-center"
          style="background-color: {colorOption}"
          on:click={() => (tag.color = colorOption)}
        >
          {#if colorOption === tag.color}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2.5"
              stroke="currentColor"
              class="h-6 w-6 text-black"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4.5 12.75l6 6 9-13.5"
              />
            </svg>
          {/if}
        </button>
      {/each}
    </div>
    <button
      type="button"
      on:click={handleDeleteTagFromBoard}
      class="flex flex-row items-center gap-1 px-1"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="currentColor"
        class="h-4 w-4"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
        />
      </svg>

      <p class="text-sm">Delete</p>
    </button>
  </div>
</div>

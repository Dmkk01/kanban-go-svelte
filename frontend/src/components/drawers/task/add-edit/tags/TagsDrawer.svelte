<script lang="ts">
  import BoardsAPI from '@/api/board'
  import TagAPI from '@/api/tag'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { createEventDispatcher } from 'svelte'
  import { fly } from 'svelte/transition'

  const dispatch = createEventDispatcher<{ 'close-drawer': void }>()

  export let selectedTagIds: number[] = []
  export let boardID: number = 0

  const tags = useQuery(['board-tags', boardID], () => BoardsAPI.getTagsBoard(boardID), {
    refetchOnWindowFocus: false,
    onSuccess: (data) => {},
  })

  let selectedTags: Tag[] = []

  $: {
    selectedTags = $tags.data?.filter((tag) => selectedTagIds.includes(tag.id)) ?? []
  }

  const colorOptions = ['#A2E8DD', '#DFFFD6', '#ABE4FF', '#FFDAAB', '#FFDDFF', '#D4DCFF', '#E8D7FF', '#E8F4C8']

  const getRandomColor = () => {
    return colorOptions[Math.floor(Math.random() * colorOptions.length)]
  }

  const defaultNewTag = () => {
    return {
      title: '',
      color: getRandomColor(),
    }
  }

  let newTag: { title: string; color: string } = defaultNewTag()
  let tagOptionsOpen: number = 12

  const createMutation = useMutation(({ title, color }: { title: string; color: string }) => TagAPI.createTag(boardID, color, title), {
    onSuccess: () => {
      $tags.refetch()
      newTag = defaultNewTag()
    },
  })

  const updateMutation = useMutation(
    ({ title, color, id }: { title: string; color: string; id: number }) => TagAPI.updateTag(boardID, id, color, title),
    {
      onSuccess: () => {
        $tags.refetch()
      },
    }
  )

  const deleteMutation = useMutation((tag_id: number) => TagAPI.deleteTag(tag_id), {
    onSuccess: () => {
      $tags.refetch()
    },
  })

  const deleteFromSelected = (id: number) => {
    selectedTagIds = selectedTagIds.filter((tagId) => tagId !== id)
  }

  const addTagToSelected = (id: number) => {
    selectedTagIds = [...selectedTagIds, id]
  }

  const handleSubmit = (e: Event) => {
    console.log('handleSubmit')
    $createMutation.mutateAsync({
      title: newTag.title,
      color: newTag.color,
    })
  }

  const openTagsOption = (id: number) => {
    if (tagOptionsOpen === id) {
      tagOptionsOpen = 0
      return
    }
    tagOptionsOpen = id
  }

  const handleChangeColor = (id: number, color: string) => {
    const tag = $tags.data?.find((tag) => tag.id === id)
    $updateMutation.mutateAsync({
      id,
      title: tag?.title ?? '',
      color,
    })
  }

  const handleTagInputChange = (e: Event, id: number) => {
    const value = (e.target as HTMLInputElement).value
    const tag = $tags.data?.find((tag) => tag.id === id)

    $updateMutation.mutateAsync({
      id,
      title: value,
      color: tag?.color ?? '',
    })
  }

  const handleDeleteTagFromBoard = (id: number) => {
    selectedTagIds = selectedTagIds.filter((tagId) => tagId !== id)
    tagOptionsOpen = 0
    $deleteMutation.mutateAsync(id)
  }
</script>

<div
  class="flex h-full w-full flex-col gap-6 overflow-y-auto bg-white px-6 py-3"
  transition:fly={{ x: 200, duration: 1000 }}
>
  {#if !$tags.isLoading && $tags.data}
    <div class="flex flex-col gap-3">
      <p class="text-sm font-bold text-tgray-600">Tags for this task</p>
      <div class="flex min-h-[1.5rem] flex-row flex-wrap gap-3">
        {#if selectedTags.length > 0}
          {#each selectedTags as tag (tag.id)}
            <button
              type="button"
              on:click={() => deleteFromSelected(tag.id)}
              class="flex flex-row items-center gap-1 rounded-md px-2 py-0.5"
              style="background-color: {tag.color};"
            >
              <p class="text-sm font-semibold">
                {tag.title}
              </p>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="h-4 w-4"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          {/each}
        {:else}
          <p class="text-sm font-medium text-tgray-600">No tags selected</p>
        {/if}
      </div>
    </div>
    <div class="flex flex-col gap-3">
      <p class="text-sm font-bold text-tgray-600">Add tags</p>
      <div class="flex flex-col gap-1">
        <p class="text-sm font-medium text-tgray-600">Create a new tag</p>
        <form
          on:submit|preventDefault={handleSubmit}
          class="flex flex-row items-center gap-3"
        >
          <input
            type="text"
            class="text-md w-full rounded-md border border-tgray-200 px-2 py-1.5"
            bind:value={newTag.title}
            placeholder="Tag name"
            required
          />
          <button
            type="submit"
            class="flex aspect-square h-7 w-7 items-center justify-center rounded-md border border-tgray-600"
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
        </form>
      </div>
      <div class="flex flex-col gap-1 pt-6">
        <p class="text-sm font-medium text-tgray-600">Existing tags</p>
        <div class="flex flex-col gap-2">
          {#each $tags.data as tag (tag.id)}
            <button
              type="button"
              on:click={() => addTagToSelected(tag.id)}
              class="flex w-fit flex-row items-center gap-1 rounded-md px-2 py-0.5"
              style="background-color: {tag.color};"
            >
              <button
                type="button"
                on:click|stopPropagation={() => openTagsOption(tag.id)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="2"
                  stroke="currentColor"
                  class="-ml-1 h-4 w-4"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
                  />
                </svg>
              </button>
              <p class="text-sm font-semibold">
                {tag.title}
              </p>
            </button>
            {#if tagOptionsOpen === tag.id}
              <div class="my-2 flex flex-col items-center gap-1 rounded-md border border-tgray-200 px-2 py-2">
                <div class="w-full">
                  <input
                    type="text"
                    class="w-full rounded-md border border-tgray-600 px-1 py-0.5 text-sm"
                    bind:value={tag.title}
                    placeholder="Tag name"
                    required
                    on:input={(e) => handleTagInputChange(e, tag.id)}
                  />
                </div>
                <div class="flex w-full flex-row items-center justify-around gap-2">
                  <div class="flex flex-row gap-1.5 px-2 py-1">
                    {#each colorOptions as colorOption}
                      <button
                        type="button"
                        class="flex h-6 w-6 items-center justify-center"
                        style="background-color: {colorOption}"
                        on:click={() => handleChangeColor(tag.id, colorOption)}
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
                    on:click={() => handleDeleteTagFromBoard(tag.id)}
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
            {/if}
          {/each}
        </div>
      </div>
    </div>
  {/if}
  <button
    type="submit"
    on:click={() => dispatch('close-drawer')}
    class="mx-auto flex items-center justify-center gap-2 rounded-lg border border-tgray-600 bg-white/40 px-10 py-1.5"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="h-5 w-5"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
      />
    </svg>

    <p class="text-sm font-semibold">Done</p>
  </button>
</div>

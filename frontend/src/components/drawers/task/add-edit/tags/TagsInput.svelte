<script lang="ts">
  import BoardsAPI from '@/api/board'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { fade } from 'svelte/transition'
  import TagItem from './TagItem.svelte'
  import { createEventDispatcher, onMount } from 'svelte'
  import TagAPI from '@/api/tag'

  let tagItemOptionsOpen: number

  export let boardID: number = 0
  export let selectedTagIds: number[] = []

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
  let newTagName = ''
  let newTagColor = getRandomColor()
  let isFocus = false

  const createNew = async () => {
    if (newTagName !== '') {
      await $createMutation.mutateAsync({
        board_id: boardID,
        title: newTagName,
        color: newTagColor,
      })
    }
  }

  onMount(() => {
    $tags.refetch()
  })

  const createMutation = useMutation(
    ({ board_id, title, color }: { board_id: number; title: string; color: string }) => TagAPI.createTag(board_id, color, title),
    {
      onSuccess: () => {
        $tags.refetch()
        newTagName = ''
        newTagColor = getRandomColor()
      },
    }
  )
</script>

<div class="w-full relative">
  <div
    on:mouseover={() => {
      isFocus = true
      tagItemOptionsOpen = 0
    }}
    on:focus={() => {
      isFocus = true
      tagItemOptionsOpen = 0
    }}
    class="flex flex-row flex-wrap gap-4 w-full border border-tgray-600 rounded-md h-auto p-1"
  >
    {#each selectedTags as tag (`current-tag-${tag.id}`)}
      <button
        type="button"
        on:click={() => {
          selectedTagIds = selectedTagIds.filter((id) => id !== tag.id)
        }}
        class="px-2 py-0.5 w-fit relative rounded-md flex flex-row gap-0 items-center z-30"
        style="background-color: {tag.color}"
      >
        <p class="text-sm font-semibold select-none">
          {tag.title}
        </p>
        <button
          type="button"
          class="ml-1"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="currentColor"
            class="w-4 h-4"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </button>
    {/each}

    <input
      type="text"
      class="text-sm font-bold text-tgray-600 w-fit bg-transparent"
      bind:value={newTagName}
      on:focus={() => (isFocus = true)}
    />
  </div>
  <div class="relative mt-2">
    {#if !$tags.isLoading && $tags.data && isFocus}
      <div
        transition:fade
        class="bg-white/50 top-[110%] absolute border border-tgray-600 rounded-md backdrop-blur-lg p-3 flex flex-col w-full gap-2"
      >
        <p class="text-sm font-medium">Select an option or type to create one</p>
        {#each $tags.data as tag (`option-tag-${tag.id}`)}
          <TagItem
            {tag}
            bind:optionsOpen={tagItemOptionsOpen}
            refetch={$tags.refetch}
            on:select-new={() => {
              if (!selectedTagIds.includes(tag.id)) {
                selectedTagIds = [...selectedTagIds, tag.id]
              }
            }}
          />
        {/each}

        {#if newTagName !== ''}
          <div class="flex flex-row gap-2 items-center">
            <p class="text-sm font-medium">Create</p>
            <button
              type="button"
              on:click={createNew}
              class="px-2 py-0.5 w-fit relative rounded-md flex flex-row gap-0 items-center z-30"
              style="background-color: {newTagColor}"
            >
              <p class="text-sm font-semibold select-none">
                {newTagName}
              </p>
            </button>
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>

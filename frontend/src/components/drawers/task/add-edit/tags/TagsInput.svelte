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
  export let isFocus = false

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

{#if !$tags.isLoading && $tags.data && isFocus}
  <div
    on:mouseleave={() => {
      isFocus = false
    }}
    class="absolute -bottom-48 -left-0 -right-0 h-44 bg-transparent"
  >
    <div
      transition:fade
      class="absolute top-0 flex w-full flex-col gap-2 rounded-md border border-tgray-600 bg-white/50 p-3 backdrop-blur-lg"
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
        <div class="flex flex-row items-center gap-2">
          <p class="text-sm font-medium">Create</p>
          <button
            type="button"
            on:click={createNew}
            class="relative z-30 flex w-fit flex-row items-center gap-0 rounded-md px-2 py-0.5"
            style="background-color: {newTagColor}"
          >
            <p class="select-none text-sm font-semibold">
              {newTagName}
            </p>
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
<div class="relative w-full">
  <div
    on:mouseover={() => {
      isFocus = true
      tagItemOptionsOpen = 0
    }}
    on:focus={() => {
      isFocus = true
      tagItemOptionsOpen = 0
    }}
    class="flex h-auto w-full flex-row flex-wrap gap-4 rounded-md border border-tgray-600 p-1"
  >
    {#each selectedTags as tag (`current-tag-${tag.id}`)}
      <button
        type="button"
        on:click={() => {
          selectedTagIds = selectedTagIds.filter((id) => id !== tag.id)
        }}
        class="relative z-30 flex w-fit flex-row items-center gap-0 rounded-md px-2 py-0.5"
        style="background-color: {tag.color}"
      >
        <p class="select-none text-sm font-semibold">
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
            class="h-4 w-4"
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
      class="w-fit bg-transparent text-sm font-bold text-tgray-600"
      bind:value={newTagName}
      on:focus={() => (isFocus = true)}
    />
  </div>
</div>

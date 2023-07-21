<script lang="ts">
  import BoardsAPI from '@/api/board'
  import TagAPI from '@/api/tag'
  import TagsDrawerCreateTag from './TagsDrawerCreateTag.svelte'
  import TagsDrawerSelected from './TagsDrawerSelected.svelte'
  import TagDrawerBoardTag from './board-tag/TagDrawerBoardTag.svelte'
  import { useMutation, useQuery } from '@sveltestack/svelte-query'
  import { createEventDispatcher } from 'svelte'
  import { fly } from 'svelte/transition'

  const dispatch = createEventDispatcher<{ 'close-drawer': void }>()

  export let selectedTagIds: number[] = []
  export let boardID: number = 0

  const tags = useQuery(`board-tags-${boardID}`, () => BoardsAPI.getTagsBoard(boardID), {
    refetchOnWindowFocus: false,
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

  const deleteFromSelected = (id: number) => {
    selectedTagIds = selectedTagIds.filter((tagId) => tagId !== id)
  }

  const addTagToSelected = (id: number) => {
    selectedTagIds = [...selectedTagIds, id]
  }

  const handleSubmit = (e: Event) => {
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

  const handleDeleteTagFromBoard = (e: CustomEvent<number>) => {
    const id = e.detail
    selectedTagIds = selectedTagIds.filter((tagId) => tagId !== id)
    tagOptionsOpen = 0
    $tags.refetch()
  }
</script>

<div
  class="flex h-full w-full flex-col gap-6 overflow-y-auto bg-white px-6 py-3"
  transition:fly={{ x: 200, duration: 1000 }}
>
  {#if !$tags.isLoading && $tags.data}
    <TagsDrawerSelected
      {selectedTags}
      {deleteFromSelected}
    />
    <div class="flex flex-col gap-3">
      <p class="text-sm font-bold text-tgray-600">Add tags</p>
      <TagsDrawerCreateTag
        {handleSubmit}
        bind:tag={newTag.title}
      />
      <div class="flex flex-col gap-1 pt-6">
        <p class="text-sm font-medium text-tgray-600">Existing tags</p>
        <div class="flex flex-col gap-2">
          {#each $tags.data as tag (tag.id)}
            <TagDrawerBoardTag
              {addTagToSelected}
              {openTagsOption}
              {tagOptionsOpen}
              {tag}
              {boardID}
              on:delete-tag-from-board={handleDeleteTagFromBoard}
            />
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

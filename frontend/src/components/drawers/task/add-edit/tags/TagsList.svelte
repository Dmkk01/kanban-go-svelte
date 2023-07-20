<script lang="ts">
  import BoardsAPI from '@/api/board'
  import { useQuery } from '@sveltestack/svelte-query'

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
</script>

{#if !$tags.isLoading && $tags.data}
  <div class="flex flex-row flex-wrap items-center gap-2 rounded-md border border-tgray-200 px-2 py-1.5">
    {#if selectedTags.length > 0}
      {#each selectedTags as tag (tag.id)}
        <div
          class="flex w-fit flex-row items-center gap-1 rounded-md px-2 py-0.5"
          style="background-color: {tag.color};"
        >
          <p class="text-sm font-semibold">
            {tag.title}
          </p>
        </div>
      {/each}
    {:else}
      <p class="text-tgray-400 text-sm">No tags selected</p>
    {/if}
  </div>
{/if}

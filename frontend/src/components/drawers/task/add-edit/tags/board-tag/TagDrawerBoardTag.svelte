<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import TagDrawerTagOption from './TagDrawerTagOption.svelte'

  export let addTagToSelected: (id: number) => void
  export let openTagsOption: (id: number) => void
  export let tagOptionsOpen: number = 0
  export let tag: Tag
  export let boardID: number

  const dispatch = createEventDispatcher<{ 'delete-tag-from-board': number }>()
</script>

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
  <TagDrawerTagOption
    {tag}
    {boardID}
    on:delete-tag-from-board={() => dispatch('delete-tag-from-board', tag.id)}
  />
{/if}

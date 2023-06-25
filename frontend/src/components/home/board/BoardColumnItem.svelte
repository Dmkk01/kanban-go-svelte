<script lang="ts">
  import ColumnAPI from '@/api/column'
  import { getEmojiBySlug, getEmojiURLBySlug } from '@/utils/emojis'
  import { useQuery } from '@sveltestack/svelte-query'
  import { fade } from 'svelte/transition'

  export let columnID: number = 0
  let currColumnID = columnID
  const column = useQuery(['column', currColumnID], async () => await ColumnAPI.getColumn(currColumnID), {})

  $: {
    if (currColumnID !== columnID) {
      currColumnID = columnID
      void $column.refetch()
    }
  }
</script>

<div
  class="bg-white/30 shadow-lg rounded-lg py-4 px-2 w-[17rem]"
  transition:fade
>
  {#if !$column.isLoading && $column.data}
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row gap-2 items-center">
        <img
          src={getEmojiURLBySlug($column.data.emoji)}
          alt="column-emoji"
          class="w-5 h-5"
        />
        <h3 class="font-bold text-sm">
          {$column.data.name} (4)
        </h3>
      </div>
      <button type="button">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          class="w-6 h-6 text-tgray-600"
          ><path
            fill="currentColor"
            d="M8.5 17c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7-10c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2zm-7 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm7 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm0 7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2zm-7-14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2z"
          /></svg
        >
      </button>
    </div>
  {/if}
</div>

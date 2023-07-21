<script lang="ts">
  import BoardsAPI from '@/api/board'
  import store from '@/store'
  import { useQuery } from '@sveltestack/svelte-query'

  export let columnID: number = 0

  const columns = useQuery(`tasks-drawer-columns`, async () => await BoardsAPI.getColumns($store.taskDrawer.ids.board || 0), {})
</script>

{#if $columns.data}
  <select
    class="rounded-md border border-tgray-200 px-1 py-1 text-base font-medium text-tgray-600"
    bind:value={columnID}
  >
    {#each $columns.data as column}
      <option value={column.id}>{column.name}</option>
    {/each}
  </select>
{/if}

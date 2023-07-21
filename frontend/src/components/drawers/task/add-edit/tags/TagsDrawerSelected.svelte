<script lang="ts">
  export let selectedTags: Tag[] = []
  export let deleteFromSelected: (id: number) => void
</script>

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

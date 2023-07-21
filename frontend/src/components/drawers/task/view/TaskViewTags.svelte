<script lang="ts">
  export let tags: Tag[] = []
</script>

{#if tags.length > 0}
  <div class="flex flex-col gap-1">
    <h3 class="text-base font-bold text-tgray-600"># Tags</h3>
    <div class="ml-6 flex flex-row flex-wrap gap-3">
      {#each tags as tag (tag.id)}
        <div
          class="relative z-30 flex w-fit flex-row items-center gap-0 rounded-md px-2 py-0.5"
          style="background-color: {tag.color}"
        >
          <p class="select-none text-sm font-medium">
            {tag.title}
          </p>
        </div>
      {/each}
    </div>
  </div>
{/if}

<script lang="ts">
  export let task: Task | null = null
</script>

<div class="flex w-full flex-col gap-2 rounded-xl bg-white/50 px-4 py-3 drop-shadow-md">
  {#if task}
    <p class="text-md font-semibold text-tgray-600">
      {task.title}
    </p>

    <p class="text-sm font-semibold text-tgray-200">
      {task.sub_tasks.filter((item) => item.completed).length} of {task.sub_tasks.length} subtasks
    </p>

    <div class="flex flex-row flex-wrap gap-1">
      {#each task.tags as tag (tag.id)}
        <div
          class="relative z-30 flex w-fit flex-row items-center gap-0 rounded-md px-2 py-0.5"
          style="background-color: {tag.color}"
        >
          <p class="select-none text-sm font-bold">
            {tag.title}
          </p>
        </div>
      {/each}
    </div>
  {/if}
</div>

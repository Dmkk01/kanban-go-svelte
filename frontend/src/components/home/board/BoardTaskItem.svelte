<script lang="ts">
  export let task: Task | null = null
</script>

<div class="w-full flex flex-col gap-2 px-4 py-3 bg-white/50 drop-shadow-md rounded-xl">
  {#if task}
    <p class="text-sm font-semibold text-tgray-600">
      {task.title}
    </p>

    <p class="text-[12px] font-semibold text-tgray-200">
      {task.sub_tasks.filter((item) => item.completed).length} of {task.sub_tasks.length} subtasks
    </p>

    <div class="flex flex-row gap-1 flex-wrap">
      {#each task.tags as tag (tag.id)}
        <div
          class="px-2 py-0.5 w-fit relative rounded-md flex flex-row gap-0 items-center z-30"
          style="background-color: {tag.color}"
        >
          <p class="text-[12px] font-medium select-none">
            {tag.title}
          </p>
        </div>
      {/each}
    </div>
  {/if}
</div>

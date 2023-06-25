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
  {/if}
</div>

<script lang="ts">
  import TaskAPI from '@/api/task'
  import { useMutation, useQueryClient } from '@sveltestack/svelte-query'

  export let subtasks: SubTask[] = []
  export let boardID: number

  const queryClient = useQueryClient()

  const updateSubtask = useMutation(
    async ({ subtask_id, value }: { subtask_id: number; value: boolean }) => TaskAPI.updateSubtaskCompleted(subtask_id, value),
    {
      onSuccess: (data) => {
        console.log(data)
        queryClient.invalidateQueries(`board-${boardID}`)
      },
    }
  )

  const handleSubtaskChange = async (e: Event, subtask_id: number) => {
    const target = e.target as HTMLInputElement
    const checked = target.checked

    $updateSubtask.mutateAsync({ subtask_id, value: checked })
  }
</script>

{#if subtasks.length > 0}
  <div class="flex flex-col gap-1">
    <h3 class="text-base font-bold text-tgray-600">Subtasks</h3>
    <div class="my-2 ml-6 flex flex-col gap-2">
      {#each subtasks as subtask (subtask.id)}
        <div class="flex flex-row gap-4">
          <input
            type="checkbox"
            class="h-5 w-5 rounded-lg border border-tgray-600"
            checked={subtask.completed}
            on:input={(e) => handleSubtaskChange(e, subtask.id)}
          />
          <p class="text-base font-medium">
            {subtask.title}
          </p>
        </div>
      {/each}
    </div>
  </div>
{/if}

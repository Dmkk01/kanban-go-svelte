<script lang="ts">
  import ColumnAPI from '@/api/column'
  import TaskAPI from '@/api/task'
  import UserAPI from '@/api/user'
  import store from '@/store'
  import { getEmojiBySlug, getEmojiURLBySlug } from '@/utils/emojis'
  import { useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query'
  import { fly } from 'svelte/transition'

  const userSettings = useQuery('user-settings', async () => await UserAPI.getUserSettings())

  const task = useQuery(
    ['task', $store.taskDrawerView.taskID],
    async () => {
      const task = await TaskAPI.getTask($store.taskDrawerView.taskID as number)
      const column = await ColumnAPI.getColumn(task.column_id)

      return { task, column }
    },
    {
      enabled: $store.taskDrawerView.isOpen,
      onSuccess: (data) => {
        console.log(data)
      },
    }
  )

  const updateSubtask = useMutation(
    async ({ subtask_id, value }: { subtask_id: number; value: boolean }) => TaskAPI.updateSubtaskCompleted(subtask_id, value),
    {
      onSuccess: (data) => {
        console.log(data)
        if (!$task.data) return
        queryClient.invalidateQueries(`board-${$task.data?.task.board_id}`)
      },
    }
  )

  const handleSubtaskChange = async (e: Event, subtask_id: number) => {
    const target = e.target as HTMLInputElement
    const checked = target.checked

    $updateSubtask.mutateAsync({ subtask_id, value: checked })
  }
  const queryClient = useQueryClient()

  const closeDrawer = () => {
    $store.taskDrawerView.isOpen = false
  }

  const formatDate = (dateString: string) => {
    if (!$userSettings.data) return ''
    const format = $userSettings.data.date_format
    const date = new Date(dateString)

    const day = date.getDate()
    const month = date.getMonth() + 1
    const year = date.getFullYear()

    switch (format) {
      case 'DD/MM/YYYY':
        return `${day}/${month}/${year}`
      case 'MM/DD/YYYY':
        return `${month}/${day}/${year}`
      case 'YYYY/MM/DD':
        return `${year}/${month}/${day}`
      default:
        return `${day}/${month}/${year}`
    }
  }
</script>

<div
  on:click|self={closeDrawer}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeDrawer()
    }
  }}
  class="absolute inset-0 overflow-hidden backdrop-blur-sm"
  transition:fly={{ x: 200, duration: 1000 }}
>
  <div class="absolute bottom-0 right-0 top-0 flex min-h-screen w-full max-w-md flex-col gap-2 bg-white/90 px-6 py-3 drop-shadow-lg">
    {#if $task.data && $task.data.task && !$task.isLoading}
      <div class="flex w-full flex-col gap-6">
        <div class="flex flex-row items-center justify-between gap-2">
          <div class="flex flex-row items-center gap-4">
            <h2 class="text-xl font-bold text-tgray-600">
              {$task.data.task.title ?? ''}
            </h2>
            <div class="flex flex-row items-center gap-2 rounded-md border border-tgray-600 px-3 py-1">
              <img
                src={getEmojiURLBySlug($task.data.column.emoji)}
                alt={$task.data.column.name}
                class="h-5 w-5"
              />
              <p class="text-base font-medium">
                {$task.data.column.name}
              </p>
            </div>
          </div>
          <button
            type="button"
            class="text-tgray-600"
            on:click={() => {
              closeDrawer()
              $store.taskDrawer.ids.board = $task.data?.task.board_id ?? 0
              $store.taskDrawer.ids.column = $task.data?.task.column_id ?? 0
              $store.taskDrawer.ids.task = $task.data?.task.id ?? 0
              $store.taskDrawer.isOpen = true
            }}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10"
              />
            </svg>
          </button>
          <button
            on:click={closeDrawer}
            type="button"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        {#if $task.data.task.tags.length > 0}
          <div class="flex flex-col gap-1">
            <h3 class="text-base font-bold text-tgray-600"># Tags</h3>
            <div class="ml-6 flex flex-row flex-wrap gap-3">
              {#each $task.data.task.tags as tag (tag.id)}
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

        {#if $task.data.task.description.length > 0}
          <div class="flex flex-col gap-1">
            <h3 class="text-base font-bold text-tgray-600">Description</h3>
            <p class="ml-6 text-base font-medium text-tgray-600">
              {$task.data.task.description}
            </p>
          </div>
        {/if}

        <div class="flex w-full flex-row justify-between">
          <div class="flex flex-col gap-1">
            <h3 class="text-base font-bold text-tgray-600">Created date</h3>
            <div class="ml-6 flex flex-row gap-3">
              <img
                src={getEmojiBySlug('apple-calendar').image_url}
                alt={$task.data.column.name}
                class="h-5 w-5"
              />
              <p class="text-base font-medium">
                {formatDate($task.data.task.created_at)}
              </p>
            </div>
          </div>

          {#if !$task.data.task.due_date.includes('0001-01-01')}
            <div class="flex flex-col gap-1">
              <h3 class="text-base font-bold text-tgray-600">Due date</h3>
              <div class="ml-6 flex flex-row gap-3">
                <img
                  src={getEmojiBySlug('apple-calendar').image_url}
                  alt={$task.data.column.name}
                  class="h-5 w-5"
                />
                <p class="text-base font-medium">
                  {formatDate($task.data.task.due_date)}
                </p>
              </div>
            </div>
          {/if}

          {#if $task.data.task.time_needed > 0}
            <div class="flex flex-col gap-1">
              <h3 class="text-base font-bold text-tgray-600">Approx. time</h3>
              <div class="ml-6 flex flex-row gap-3">
                <img
                  src={getEmojiBySlug('apple-alarm-clock').image_url}
                  alt={$task.data.column.name}
                  class="h-5 w-5"
                />
                <p class="text-base font-medium">
                  {$task.data.task.time_needed} min
                </p>
              </div>
            </div>
          {/if}
        </div>

        {#if $task.data.task.sub_tasks.length > 0}
          <div class="flex flex-col gap-1">
            <h3 class="text-base font-bold text-tgray-600">Subtasks</h3>
            <div class="my-2 ml-6 flex flex-col gap-2">
              {#each $task.data.task.sub_tasks as subtask (subtask.id)}
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

        {#if $task.data.task.links.length > 0}
          <div class="flex flex-col gap-1">
            <h3 class="text-base font-bold text-tgray-600">Links</h3>
            <div class="my-2 ml-6 flex flex-col gap-2">
              {#each $task.data.task.links as link (link.id)}
                <div class="flex flex-row gap-4">
                  <img
                    src={getEmojiURLBySlug(link.emoji)}
                    alt={link.url}
                    class="h-5 w-5"
                  />
                  <a
                    href={link.url}
                    target="_blank"
                    class="text-base font-medium underline"
                  >
                    {link.url}
                  </a>
                </div>
              {/each}
            </div>
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>

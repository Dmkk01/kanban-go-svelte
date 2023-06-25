<script lang="ts">
  import { getInitEmoji } from '@/utils/emojis'
  import { z } from 'zod'
  import store from '@/store'
  import EmojiButton from '../../common/EmojiButton.svelte'
  import { useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query'
  import BoardsAPI from '@/api/board'
  import SubmitButton from '@/components/common/SubmitButton.svelte'
  import { fly } from 'svelte/transition'
  import ColumnAPI from '@/api/column'
  import TaskAPI from '@/api/task'
  import TaskDrawerSubtaskItem from './TaskDrawerSubtaskItem.svelte'
  import TaskDrawerLinkItem from './TaskDrawerLinkItem.svelte'

  const schema = z.object({
    board_id: z.number().min(1),
    column_id: z.number().min(1),
    description: z.string(),
    due_date: z.string(),
    no_due_date: z.boolean(),
    position: z.number(),
    time_needed: z.number(),
    title: z.string(),
    sub_tasks: z.array(z.object({ id: z.number(), title: z.string(), completed: z.boolean() })),
    links: z.array(z.object({ id: z.number(), title: z.string(), url: z.string(), emoji: z.string() })),
  })

  type Schema = z.infer<typeof schema>

  const queryClient = useQueryClient()

  const columns = useQuery(`tasks-drawer-columns`, async () => await BoardsAPI.getColumns($store.taskDrawer.ids.board || 0), {})
  const tasks = useQuery(`tasks-drawer-tasks`, async () => await ColumnAPI.getTasks($store.taskDrawer.ids.column || 0), {})

  const data: Schema = {
    board_id: $store.taskDrawer.ids.board || 0,
    column_id: $store.taskDrawer.ids.column || 0,
    description: '',
    due_date: '',
    no_due_date: false,
    position: $tasks.data?.length || 0,
    time_needed: 0,
    title: '',
    sub_tasks: [],
    links: [],
  }

  let message = ''
  let isSaved = false

  const updateIsSaved = () => {
    isSaved = true

    setTimeout(() => {
      isSaved = false
    }, 3000)
  }

  const updateMessage = (msg: string) => {
    message = msg

    setTimeout(() => {
      message = ''
    }, 3000)
  }

  let drawerType: 'edit' | 'create' = $store.taskDrawer.ids.task ? 'edit' : 'create'

  const mutationCreate = useMutation(({ columnID, data }: { data: TaskCreate; columnID: number }) => TaskAPI.createTask(columnID, data), {
    onSuccess: () => {
      updateIsSaved()
      queryClient.invalidateQueries(`board-${$store.taskDrawer.ids.board}`)
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const closeDrawer = () => {
    $store.taskDrawer.isOpen = false
  }

  const onSubmit = async (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(data)

    if (!result.success) {
      updateMessage(result.error.issues[0].message)
      return
    }

    if (data.no_due_date) {
      data.due_date = ''
    }

    if (drawerType === 'create') {
      const columnID = $store.taskDrawer.ids.column
      if (!columnID) {
        updateMessage('Board ID is not defined')
        return
      }
      $mutationCreate.mutate({ data, columnID })
    }

    closeDrawer()
  }

  const addNewSubtask = () => {
    const newSubtask = { id: Math.floor(Math.random() * 1000), title: '', completed: false }

    data.sub_tasks = [...data.sub_tasks, newSubtask]
  }

  const deleteSubtask = (e: CustomEvent<number>) => {
    data.sub_tasks = data.sub_tasks.filter((item) => item.id !== e.detail)
  }

  const addNewLink = () => {
    const newLink = { id: Math.floor(Math.random() * 1000), title: '', url: '', emoji: getInitEmoji().slug }

    data.links = [...data.links, newLink]
  }

  const deleteLink = (e: CustomEvent<number>) => {
    data.links = data.links.filter((item) => item.id !== e.detail)
  }
</script>

<div
  on:click|self={closeDrawer}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeDrawer()
    }
  }}
  class="absolute inset-0 backdrop-blur-sm overflow-hidden"
  transition:fly={{ x: 200, duration: 1000 }}
>
  <div class="absolute right-0 top-0 bottom-0 w-full max-w-md bg-white/90 drop-shadow-lg px-6 py-3 flex flex-col gap-2">
    <div class="flex flex-row gap-4 items-center">
      <h2 class="font-bold text-lg">{drawerType === 'edit' ? 'Edit' : 'Add New'} Task</h2>
      {#if $columns.data}
        <select
          class="text-base font-medium py-1 px-1 border border-tgray-200 rounded-md text-tgray-600"
          bind:value={data.column_id}
        >
          {#each $columns.data as column}
            <option value={column.id}>{column.name}</option>
          {/each}
        </select>
      {/if}
    </div>

    <form
      class="py-3 flex flex-col gap-5 w-full"
      on:submit={onSubmit}
    >
      <div class="flex flex-col gap-1 w-full">
        <label
          for="#task_name"
          class="text-sm font-bold text-tgray-600"
        >
          Title
        </label>
        <input
          type="text"
          id="task_name"
          bind:value={data.title}
          class="text-base w-full font-medium px-2 py-1 border border-tgray-200 rounded-md"
        />
      </div>
      <div class="flex flex-col gap-1 w-full">
        <label
          for="#task_description"
          class="text-sm font-bold text-tgray-600"
        >
          Description
        </label>
        <textarea
          id="task_description"
          bind:value={data.description}
          class="text-base w-full min-h-[7rem] font-medium px-2 py-1 border border-tgray-200 rounded-md"
        />
      </div>
      <div class="flex flex-row gap-8">
        <div class="flex flex-col gap-2 w-full">
          <label
            for="#task_description"
            class="text-sm font-bold text-tgray-600"
          >
            Due Date
          </label>
          <input
            type="date"
            id="task_due_date"
            bind:value={data.due_date}
            class="text-base w-full font-medium px-2 py-1 border border-tgray-200 rounded-md"
          />
          <div class="flex flex-row gap-2 items-center">
            <input
              type="checkbox"
              id="task_no_due_date"
              bind:checked={data.no_due_date}
              class="h-4 w-4 rounded-md"
            />
            <label
              for="#task_no_due_date"
              class="text-sm font-bold text-tgray-600"
            >
              No Due Date
            </label>
          </div>
        </div>
        <div class="flex flex-col gap-1 w-full">
          <label
            for="#task_time_needed"
            class="text-sm font-bold text-tgray-600"
          >
            Approximate time
          </label>
          <div class="flex flex-row gap-2 items-center">
            <input
              type="number"
              min="0"
              id="task_time_needed"
              bind:value={data.due_date}
              class="text-base w-full font-medium px-2 py-1 border border-tgray-200 rounded-md"
            />
            <span class="text-sm font-bold text-tgray-600">minutes</span>
          </div>
        </div>
      </div>

      <div class="flex flex-col gap-3 w-full">
        <p class="text-sm font-bold text-tgray-600">Subtasks</p>
        {#each data.sub_tasks as subtask (subtask.id)}
          <TaskDrawerSubtaskItem
            bind:completed={subtask.completed}
            bind:title={subtask.title}
            id={subtask.id}
            on:delete-subtask={deleteSubtask}
          />
        {/each}
        <button
          type="button"
          on:click={addNewSubtask}
          class="px-2 py-1 border border-tgray-200 flex flex-row gap-2 items-center w-fit rounded-md"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="2.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 4.5v15m7.5-7.5h-15"
            />
          </svg>
          <p class="font-medium text-sm">Add New Subtask</p>
        </button>
      </div>

      <div class="flex flex-col gap-3 w-full">
        <p class="text-sm font-bold text-tgray-600">Links</p>
        {#each data.links as link (link.id)}
          <TaskDrawerLinkItem
            bind:url={link.url}
            bind:emoji={link.emoji}
            id={link.id}
            on:delete-link={deleteLink}
          />
        {/each}
        <button
          type="button"
          on:click={addNewLink}
          class="px-2 py-1 border border-tgray-200 flex flex-row gap-2 items-center w-fit rounded-md"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="2.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 4.5v15m7.5-7.5h-15"
            />
          </svg>
          <p class="font-medium text-sm">Add New Link</p>
        </button>
      </div>

      <SubmitButton
        text={drawerType === 'edit' ? 'Save Task' : 'Create Task'}
        extraStyles="bg-transparent py-0.5"
        textStyles="text-base"
        {message}
        {isSaved}
      />
    </form>
  </div>
</div>

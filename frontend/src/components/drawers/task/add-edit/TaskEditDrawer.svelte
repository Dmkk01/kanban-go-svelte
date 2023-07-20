<script lang="ts">
  import { getInitEmoji } from '@/utils/emojis'
  import { Schema, z } from 'zod'
  import store from '@/store'
  import { useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query'
  import BoardsAPI from '@/api/board'
  import SubmitButton from '@/components/common/SubmitButton.svelte'
  import { fly } from 'svelte/transition'
  import ColumnAPI from '@/api/column'
  import TaskAPI from '@/api/task'
  import TaskDrawerSubtaskItem from './TaskEditDrawerSubtaskItem.svelte'
  import TaskDrawerLinkItem from './TaskEditDrawerLinkItem.svelte'
  import TagsList from './tags/TagsList.svelte'
  import TagsDrawer from './tags/TagsDrawer.svelte'

  const schema = z.object({
    board_id: z.number().min(1),
    column_id: z.number().min(1),
    description: z.string(),
    due_date: z.string(),
    no_due_date: z.boolean(),
    position: z.number(),
    time_needed: z.number(),
    title: z.string(),
    sub_tasks: z.array(z.object({ id: z.number(), title: z.string(), completed: z.boolean(), isNew: z.boolean().optional() })),
    links: z.array(z.object({ id: z.number(), title: z.string(), url: z.string(), emoji: z.string(), isNew: z.boolean().optional() })),
    tags: z.array(z.number().min(1)),
  })

  type Schema = z.infer<typeof schema>

  const queryClient = useQueryClient()

  const columns = useQuery(`tasks-drawer-columns`, async () => await BoardsAPI.getColumns($store.taskDrawer.ids.board || 0), {})
  const tasks = useQuery(`tasks-drawer-tasks`, async () => await ColumnAPI.getTasks($store.taskDrawer.ids.column || 0), {})
  const task = useQuery(
    ['task', $store.taskDrawer.ids.task],
    async () => {
      if (!$store.taskDrawer.ids.task) return
      return await TaskAPI.getTask($store.taskDrawer.ids.task)
    },
    {
      onSuccess: (res) => {
        if (res) {
          data.title = res.title
          data.description = res.description
          data.time_needed = res.time_needed
          data.sub_tasks = res.sub_tasks.map((item) => ({ id: item.id, title: item.title, completed: item.completed }))
          data.links = res.links.map((item) => ({ id: item.id, title: item.title, url: item.url, emoji: item.emoji }))
          data.tags = res.tags.map((item) => item.id)
          data.position = res.position

          if (res.due_date.includes('0001-01-01')) {
            data.no_due_date = true
          } else {
            data.due_date = res.due_date.slice(0, 10)
          }
        }
      },
    }
  )

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
    tags: [],
  }

  // $: console.log('DATA', data)

  let message = ''
  let isSaved = false
  let tagsDrawerOpen = false

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
      refreshCloseBoard()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const mutationDelete = useMutation(async (taskId: number) => TaskAPI.deleteTask(taskId), {
    onSuccess: () => {
      updateIsSaved()
      refreshCloseBoard()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const mutationUpdate = useMutation(async ({ taskId, data }: { taskId: number; data: TaskUpdate }) => TaskAPI.updateTask(taskId, data), {
    onSuccess: () => {
      updateIsSaved()
      refreshCloseBoard()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const closeDrawer = () => {
    $store.taskDrawer.isOpen = false
    $store.taskDrawer.ids = {
      board: null,
      column: null,
      task: null,
    }
  }

  const refreshCloseBoard = async () => {
    const boardKey = `board-${$store.taskDrawer.ids.board}`
    await queryClient.invalidateQueries(boardKey)
    closeDrawer()
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
    } else {
      const taskId = $store.taskDrawer.ids.task
      if (!taskId || !$task.data) {
        updateMessage('Task ID is not defined')
        return
      }
      const finalData: TaskUpdate = {
        description: data.description,
        due_date: data.no_due_date ? $task.data.due_date : data.due_date,
        position: data.position,
        title: data.title,
        time_needed: data.time_needed,
        sub_tasks: data.sub_tasks.map((item) => ({ id: item.isNew ? 0 : item.id, title: item.title, completed: item.completed })),
        links: data.links.map((item) => ({ id: item.isNew ? 0 : item.id, title: item.title, url: item.url, emoji: item.emoji })),
        tags: data.tags,
      }

      console.log('FINAL DATA', finalData)

      $mutationUpdate.mutate({ taskId, data: finalData })
    }
  }

  const addNewSubtask = () => {
    const newSubtask = { id: Math.floor(Math.random() * 1000), title: '', completed: false, isNew: true }

    data.sub_tasks = [...data.sub_tasks, newSubtask]
  }

  const deleteSubtask = (e: CustomEvent<number>) => {
    data.sub_tasks = data.sub_tasks.filter((item) => item.id !== e.detail)
  }

  const addNewLink = () => {
    const newLink = { id: Math.floor(Math.random() * 1000), title: '', url: '', emoji: getInitEmoji().slug, isNew: true }

    data.links = [...data.links, newLink]
  }

  const deleteLink = (e: CustomEvent<number>) => {
    data.links = data.links.filter((item) => item.id !== e.detail)
  }

  const handleDeleteTask = () => {
    if (!$store.taskDrawer.ids.task) {
      updateMessage('Task ID is not defined')
      return
    }
    $mutationDelete.mutate($store.taskDrawer.ids.task)
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
  <div class="absolute bottom-0 right-0 top-0 flex min-h-screen w-full max-w-md flex-col gap-2 bg-white/90 drop-shadow-lg">
    {#if tagsDrawerOpen && $store.taskDrawer.ids.board}
      <div class="absolute inset-0 z-10">
        <TagsDrawer
          bind:selectedTagIds={data.tags}
          boardID={$store.taskDrawer.ids.board}
          on:close-drawer={() => (tagsDrawerOpen = false)}
        />
      </div>
    {/if}
    <div class="flex flex-row justify-between px-6 py-3">
      <div class="flex flex-row items-center gap-4">
        <h2 class="text-lg font-bold">{drawerType === 'edit' ? 'Edit' : 'Add New'} Task</h2>
        {#if $columns.data}
          <select
            class="rounded-md border border-tgray-200 px-1 py-1 text-base font-medium text-tgray-600"
            bind:value={data.column_id}
          >
            {#each $columns.data as column}
              <option value={column.id}>{column.name}</option>
            {/each}
          </select>
        {/if}
      </div>
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
    {#if !$task.isLoading}
      <form
        class="flex w-full flex-col gap-5 overflow-y-auto px-6 py-3"
        on:submit={onSubmit}
      >
        <div class="flex w-full flex-col gap-1">
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
            class="w-full rounded-md border border-tgray-200 px-2 py-1 text-base font-medium"
          />
        </div>
        <div class="flex w-full flex-col gap-1">
          <label
            for="#task_description"
            class="text-sm font-bold text-tgray-600"
          >
            Description
          </label>
          <textarea
            id="task_description"
            bind:value={data.description}
            class="min-h-[7rem] w-full rounded-md border border-tgray-200 px-2 py-1 text-base font-medium"
          />
        </div>
        <div class="flex flex-row gap-8">
          <div class="flex w-full flex-col gap-2">
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
              class="w-full rounded-md border border-tgray-200 px-2 py-1 text-base font-medium"
            />
            <div class="flex flex-row items-center gap-2">
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
          <div class="flex w-full flex-col gap-1">
            <label
              for="#task_time_needed"
              class="text-sm font-bold text-tgray-600"
            >
              Approximate time
            </label>
            <div class="flex flex-row items-center gap-2">
              <input
                type="number"
                min="0"
                id="task_time_needed"
                bind:value={data.time_needed}
                class="w-full rounded-md border border-tgray-200 px-2 py-1 text-base font-medium"
              />
              <span class="text-sm font-bold text-tgray-600">minutes</span>
            </div>
          </div>
        </div>

        <div class="flex w-full flex-col gap-3">
          <p class="text-sm font-bold text-tgray-600"># tags</p>
          {#if $store.taskDrawer.ids.board}
            <TagsList
              selectedTagIds={data.tags}
              boardID={$store.taskDrawer.ids.board}
            />
          {/if}
          <div>
            <button
              type="button"
              on:click={() => (tagsDrawerOpen = true)}
              class="flex w-fit flex-row items-center gap-2 rounded-md border border-tgray-200 px-2 py-1"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="2.5"
                stroke="currentColor"
                class="h-5 w-5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 4.5v15m7.5-7.5h-15"
                />
              </svg>
              <p class="text-sm font-medium">Add Tags</p>
            </button>
          </div>
        </div>

        <div class="flex w-full flex-col gap-3">
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
            class="flex w-fit flex-row items-center gap-2 rounded-md border border-tgray-200 px-2 py-1"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2.5"
              stroke="currentColor"
              class="h-5 w-5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
            <p class="text-sm font-medium">Add New Subtask</p>
          </button>
        </div>

        <div class="flex w-full flex-col gap-3">
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
            class="flex w-fit flex-row items-center gap-2 rounded-md border border-tgray-200 px-2 py-1"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2.5"
              stroke="currentColor"
              class="h-5 w-5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
            <p class="text-sm font-medium">Add New Link</p>
          </button>
        </div>

        <SubmitButton
          text={drawerType === 'edit' ? 'Save Task' : 'Create Task'}
          extraStyles="bg-transparent py-0.5"
          textStyles="text-base"
          {message}
          {isSaved}
        />
        <div class="w-full">
          <button
            type="button"
            on:click={handleDeleteTask}
            class="mx-auto flex flex-row items-center gap-2 rounded-lg border border-[#F07575] px-8 py-2 text-[#F07575]"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-5 w-5"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
              />
            </svg>
            <p class="text-sm">Delete Task</p>
          </button>
        </div>
      </form>
    {/if}
  </div>
</div>

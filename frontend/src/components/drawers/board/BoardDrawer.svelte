<script lang="ts">
  import { getInitEmoji } from '@/utils/emojis'
  import { z } from 'zod'
  import store from '@/store'
  import EmojiButton from '../../common/EmojiButton.svelte'
  import { useMutation, useQueryClient } from '@sveltestack/svelte-query'
  import BoardsAPI from '@/api/board'
  import SubmitButton from '@/components/common/SubmitButton.svelte'
  import BoardDrawerColumn from './BoardDrawerColumn.svelte'
  import { fly } from 'svelte/transition'
  import { flip } from 'svelte/animate'
  import { dndzone } from 'svelte-dnd-action'

  const schema = z.object({
    emoji: z.string().min(1),
    name: z.string().min(1),
    columns: z.array(
      z.object({
        id: z.number(),
        emoji: z.string().min(1),
        name: z.string().min(1),
        position: z.number().min(1),
      })
    ),
  })

  type Schema = z.infer<typeof schema>

  type ColumnSchema = Schema['columns'][number]

  const queryClient = useQueryClient()

  const data: Schema = {
    name: 'Board',
    emoji: getInitEmoji().slug,
    columns: [],
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

  let drawerType: 'edit' | 'create' = $store.boardDrawer.boardID ? 'edit' : 'create'

  const mutationCreate = useMutation((data: Schema) => BoardsAPI.createNewBoard(data.name, data.emoji, data.columns), {
    onSuccess: () => {
      updateIsSaved()
      queryClient.invalidateQueries('boards')
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const closeDrawer = () => {
    $store.boardDrawer.isOpen = false
  }

  const createNewColumn = () => {
    const newColumn = {
      emoji: getInitEmoji().slug,
      name: 'Column',
      position: data.columns.length + 1,
      id: Math.floor(Math.random() * 100000),
    }

    data.columns = [...data.columns, newColumn]
  }

  const deleteColumn = (e: CustomEvent<number>) => {
    data.columns = data.columns.filter((item) => item.position !== e.detail)
  }

  const onSubmit = async (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(data)

    if (!result.success) {
      updateMessage(result.error.issues[0].message)
      return
    }

    if (drawerType === 'create') {
      $mutationCreate.mutate(result.data)
    }

    closeDrawer()
  }

  const handleSort = (e: CustomEvent<DndEvent<ColumnSchema>>) => {
    const newItems = e.detail.items as ColumnSchema[]

    data.columns = newItems.map((item, index) => {
      item.position = index + 1
      return item
    })
  }
</script>

<div
  on:click|self={closeDrawer}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeDrawer()
    }
  }}
  class="absolute inset-0 backdrop-blur-sm"
  transition:fly={{ x: 200, duration: 1000 }}
>
  <div class="absolute right-0 top-0 bottom-0 w-full max-w-md bg-white/90 drop-shadow-lg px-6 py-10 flex flex-col gap-2">
    <h2 class="font-bold text-lg">{drawerType === 'edit' ? 'Edit' : 'Add'} Board</h2>

    <form
      class="py-3 flex flex-col gap-5 w-full"
      on:submit={onSubmit}
    >
      <div class="flex flex-row gap-3 items-center w-full">
        <div class="flex flex-col gap-2">
          <p class="text-sm font-bold text-tgray-600">Emoji</p>

          <EmojiButton
            bind:emojiSlug={data.emoji}
            emojiKey="board-customization"
            extraStyles="w-12 h-12 border border-tgray-200"
            imageStyles="w-3/4 h-auto"
          />
        </div>
        <div class="flex flex-col gap-2 w-full">
          <label
            for="#board_name"
            class="text-sm font-bold text-tgray-600"
          >
            Name
          </label>
          <input
            type="text"
            id="board_name"
            bind:value={data.name}
            class="text-xl w-full font-bold h-12 px-2 border border-tgray-200 rounded-lg"
          />
        </div>
      </div>
      <div class="flex flex-col gap-3">
        <p class="text-base font-bold text-tgray-600">Columns</p>
        <section
          use:dndzone={{ items: data.columns, flipDurationMs: 200, type: 'board-create-columns', dropTargetStyle: { outline: 'none' } }}
          on:consider={handleSort}
          on:finalize={handleSort}
          class="flex flex-col gap-3"
        >
          {#each data.columns as item (item.id)}
            <div
              draggable="true"
              animate:flip={{ duration: 200 }}
            >
              <BoardDrawerColumn
                bind:emoji={item.emoji}
                bind:name={item.name}
                emojiKey={item.position}
                on:delete-column={deleteColumn}
              />
            </div>
          {/each}
        </section>
        <button
          type="button"
          on:click={createNewColumn}
          class="px-4 py-1 border border-tgray-200 flex flex-row gap-2 items-center w-fit rounded-md"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="2.5"
            stroke="currentColor"
            class="w-6 h-6"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 4.5v15m7.5-7.5h-15"
            />
          </svg>
          <p class="font-medium text-sm">Add New Column</p>
        </button>
      </div>

      <SubmitButton
        text={drawerType === 'edit' ? 'Save Board' : 'Create Board'}
        extraStyles="bg-transparent py-0.5"
        textStyles="text-base"
        {message}
        {isSaved}
      />
    </form>
  </div>
</div>

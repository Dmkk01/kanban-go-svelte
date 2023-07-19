<script lang="ts">
  import { getInitEmoji } from '@/utils/emojis'
  import { z } from 'zod'
  import store from '@/store'
  import EmojiButton from '../../common/EmojiButton.svelte'
  import { useMutation, useQueryClient, useQuery } from '@sveltestack/svelte-query'
  import BoardsAPI from '@/api/board'
  import SubmitButton from '@/components/common/SubmitButton.svelte'
  import BoardDrawerColumn from './BoardDrawerColumn.svelte'
  import { fly } from 'svelte/transition'
  import { flip } from 'svelte/animate'
  import { dndzone } from 'svelte-dnd-action'
  import ColumnAPI from '@/api/column'

  const board = useQuery(`board-${$store.boardDrawer.boardID}`, async () => await BoardsAPI.getBoardFull($store.boardDrawer.boardID || 0), {
    enabled: $store.boardDrawer.boardID !== null,
    onSuccess: (res) => {
      if (res) {
        data.name = res.board.name
        data.emoji = res.board.emoji
        data.columns = res.columns.map((item) => {
          return {
            id: item.column.id,
            emoji: item.column.emoji,
            name: item.column.name,
            position: item.column.position,
          }
        })
      }
    },
  })

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

  const mutationUpdate = useMutation(
    async (data: Schema) => {
      if (!$store.boardDrawer.boardID) return

      for (const item of data.columns) {
        await ColumnAPI.editColumn(item.id, item.emoji, item.name)
      }

      await BoardsAPI.updateColumnPosition($store.boardDrawer.boardID, data.columns)
      return await BoardsAPI.editBoard($store.boardDrawer.boardID, data.name, data.emoji)
    },
    {
      onSuccess: async () => {
        updateIsSaved()
        await queryClient.invalidateQueries('boards')
        await queryClient.invalidateQueries(`board-${$store.boardDrawer.boardID}`)
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

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
    } else {
      $mutationUpdate.mutate(result.data)
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
  <div class="absolute bottom-0 right-0 top-0 flex w-full max-w-md flex-col gap-2 bg-white/90 px-6 py-10 drop-shadow-lg">
    <button
      type="button"
      on:click={closeDrawer}
      class="absolute right-2 top-2"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="h-10 w-10"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M6 18L18 6M6 6l12 12"
        />
      </svg>
    </button>
    <h2 class="text-lg font-bold">{drawerType === 'edit' ? 'Edit' : 'Add'} Board</h2>

    {#if !$board.isLoading}
      <form
        class="flex w-full flex-col gap-5 py-3"
        on:submit={onSubmit}
      >
        <div class="flex w-full flex-row items-center gap-3">
          <div class="flex flex-col gap-2">
            <p class="text-sm font-bold text-tgray-600">Emoji</p>

            <EmojiButton
              bind:emojiSlug={data.emoji}
              emojiKey="board-customization"
              extraStyles="w-12 h-12 border border-tgray-200"
              imageStyles="w-3/4 h-auto"
            />
          </div>
          <div class="flex w-full flex-col gap-2">
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
              class="h-12 w-full rounded-lg border border-tgray-200 px-2 text-xl font-bold"
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
            class="flex w-fit flex-row items-center gap-2 rounded-md border border-tgray-200 px-4 py-1"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2.5"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
            <p class="text-sm font-medium">Add New Column</p>
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
    {/if}
  </div>
</div>

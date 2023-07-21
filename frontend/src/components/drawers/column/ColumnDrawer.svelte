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

  const schema = z.object({
    emoji: z.string().min(1),
    name: z.string().min(1),
    position: z.number().min(1),
  })

  type Schema = z.infer<typeof schema>

  const queryClient = useQueryClient()

  const columns = useQuery(`column-drawer`, async () => await BoardsAPI.getColumns($store.columnDrawer.boardID || 0), {})

  const defaultData = () => {
    const data = {
      name: 'Column',
      emoji: getInitEmoji().slug,
      position: $columns.data?.length || 1,
    }
    return data
  }

  const data: Schema = defaultData()

  const currentColumn = useQuery(
    `edit-column-${$store.columnDrawer.columnID}`,
    async () => await ColumnAPI.getColumn($store.columnDrawer.columnID || 0),
    {
      enabled: $store.columnDrawer.columnID !== null,
      onSuccess: (res) => {
        if (res) {
          data.name = res.name
          data.emoji = res.emoji
          data.position = res.position
        }
      },
    }
  )

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

  let drawerType: 'edit' | 'create' = $store.columnDrawer.columnID ? 'edit' : 'create'

  const mutationCreate = useMutation(
    ({ data, boardID }: { data: Schema; boardID: number }) => BoardsAPI.createColumn(boardID, data.name, data.emoji, data.position),
    {
      onSuccess: async () => {
        updateIsSaved()
        await queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}`)
        closeDrawer()
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

  const mutationDelete = useMutation(({ columnID }: { columnID: number }) => ColumnAPI.deleteColumn(columnID), {
    onSuccess: () => {
      updateIsSaved()
      queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}`)
      closeDrawer()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const mutationUpdate = useMutation(
    ({ data, boardID, columnID }: { data: Schema; boardID: number; columnID: number }) => {
      if (!$store.columnDrawer.columnID) {
        throw new Error('Column ID is not defined')
      }

      if ($currentColumn.data?.position !== data.position) {
        return BoardsAPI.updateColumnPosition(boardID, [
          {
            id: columnID,
            position: data.position,
          },
        ])
      }

      return ColumnAPI.editColumn(columnID, data.emoji, data.name)
    },
    {
      onSuccess: async () => {
        updateIsSaved()
        await queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}`)
        closeDrawer()
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

  const closeDrawer = () => {
    $store.columnDrawer.isOpen = false
  }

  const onSubmit = async (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(data)

    if (!result.success) {
      updateMessage(result.error.issues[0].message)
      return
    }

    if (drawerType === 'create') {
      const boardID = $store.columnDrawer.boardID
      if (!boardID) {
        updateMessage('Board ID is not defined')
        return
      }
      $mutationCreate.mutate({ data, boardID })
    } else {
      const columnID = $store.columnDrawer.columnID
      const boardID = $store.columnDrawer.boardID
      if (!columnID || !boardID) {
        updateMessage('Column ID or Board ID is not defined')
        return
      }
      $mutationUpdate.mutate({ data, boardID, columnID })
    }
  }

  const handleDelete = () => {
    const columnID = $store.columnDrawer.columnID
    if (!columnID) {
      updateMessage('Column ID is not defined')
      return
    }
    $mutationDelete.mutate({ columnID })
  }
</script>

<div
  on:click|self={closeDrawer}
  on:keydown={(e) => {
    if (e.key === 'Escape') {
      closeDrawer()
    }
  }}
  class="absolute inset-0 h-full overflow-hidden backdrop-blur-sm"
  transition:fly={{ x: 200, duration: 1000 }}
>
  <div class="absolute bottom-0 right-0 top-0 flex h-full w-full max-w-md flex-col justify-between gap-2 bg-white/90 px-6 py-10 drop-shadow-lg">
    <div class="flex w-full flex-col gap-2">
      <div class="flex flex-row justify-between">
        <h2 class="text-lg font-bold">{drawerType === 'edit' ? 'Edit' : 'Add'} Column</h2>
        <button
          type="button"
          on:click={closeDrawer}
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

      {#if !$currentColumn.isLoading}
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
          {#if !$columns.isLoading && $columns.data}
            <div class="flex flex-row items-center gap-2">
              <label
                for="#column_position"
                class="text-sm font-bold text-tgray-600"
              >
                Position
              </label>
              <select
                class="w-16 rounded-md border border-tgray-200 px-1 py-1 text-base font-medium text-tgray-600"
                bind:value={data.position}
                id="column_position"
              >
                {#each Array.from({ length: $columns.data.length + 1 }, (_, i) => i + 1) as position}
                  <option value={position}>{position}</option>
                {/each}
              </select>
            </div>
          {/if}

          <SubmitButton
            text={drawerType === 'edit' ? 'Save Column' : 'Create Column'}
            extraStyles="bg-transparent py-0.5"
            textStyles="text-base"
            {message}
            {isSaved}
          />
        </form>
      {/if}
    </div>
    {#if drawerType === 'edit'}
      <div class="w-full">
        <button
          type="button"
          on:click={handleDelete}
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
          <p class="text-sm">Delete Column</p>
        </button>
      </div>
    {/if}
  </div>
</div>

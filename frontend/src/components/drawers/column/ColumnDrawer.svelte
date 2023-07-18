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
      onSuccess: () => {
        updateIsSaved()
        queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}`)
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

  const mutationUpdate = useMutation(
    ({ data, boardID, columnID }: { data: Schema; boardID: number, columnID: number }) => {
      if (!$store.columnDrawer.columnID) {
        throw new Error('Column ID is not defined')
      }

      if ($currentColumn.data?.position !== data.position) {
        return BoardsAPI.updateColumnPosition(boardID, [{
          id: columnID,
          position: data.position,
        }])
      }

      return ColumnAPI.editColumn(columnID, data.emoji, data.name)
    },
    {
      onSuccess: () => {
        updateIsSaved()
        queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}`)
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

    closeDrawer()
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

    <h2 class="text-lg font-bold">{drawerType === 'edit' ? 'Edit' : 'Add'} Column</h2>

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
</div>

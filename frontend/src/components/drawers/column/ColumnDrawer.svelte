<script lang="ts">
  import { getInitEmoji } from '@/utils/emojis'
  import { z } from 'zod'
  import store from '@/store'
  import EmojiButton from '../../common/EmojiButton.svelte'
  import { useMutation, useQuery, useQueryClient } from '@sveltestack/svelte-query'
  import BoardsAPI from '@/api/board'
  import SubmitButton from '@/components/common/SubmitButton.svelte'
  import { onMount } from 'svelte'

  const schema = z.object({
    emoji: z.string().min(1),
    name: z.string().min(1),
    position: z.number().min(1),
  })

  type Schema = z.infer<typeof schema>

  const queryClient = useQueryClient()

  const columns = useQuery(`column-drawer`, async () => await BoardsAPI.getColumns($store.columnDrawer.boardID || 0), {})

  const data: Schema = {
    name: 'Column',
    emoji: getInitEmoji().slug,
    position: $columns.data?.length || 1,
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

  let drawerType: 'edit' | 'create' = $store.columnDrawer.columnID ? 'edit' : 'create'

  const mutationCreate = useMutation(
    ({ data, boardID }: { data: Schema; boardID: number }) => BoardsAPI.createColumn(boardID, data.name, data.emoji, data.position),
    {
      onSuccess: () => {
        updateIsSaved()
        queryClient.invalidateQueries(`board-${$store.columnDrawer.boardID}-columns`)
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
  class="absolute inset-0 backdrop-blur-sm"
>
  <div class="absolute right-0 top-0 bottom-0 w-full max-w-md bg-white/90 drop-shadow-lg px-6 py-10 flex flex-col gap-2">
    <h2 class="font-bold text-lg">{drawerType === 'edit' ? 'Edit' : 'Add'} Column</h2>

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
      {#if !$columns.isLoading && $columns.data}
        <div class="flex flex-row gap-2 items-center">
          <label
            for="#column_position"
            class="text-sm font-bold text-tgray-600"
          >
            Position
          </label>
          <select
            class="text-base font-medium w-16 py-1 px-1 border border-tgray-200 rounded-md text-tgray-600"
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
  </div>
</div>

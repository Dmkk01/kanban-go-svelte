<script lang="ts">
  import { useMutation, useQueryClient } from '@sveltestack/svelte-query'
  import { string, z } from 'zod'
  import UserAPI from '@/api/user'
  import SettingsSubmit from '../../common/SubmitButton.svelte'
  import EmojiButton from '@/components/common/EmojiButton.svelte'
  import store from '@/store'

  import { getEmojiURLBySlug, getInitEmoji, openEmojiSelector } from '@/utils/emojis'

  export let emoji: string
  export let name: string
  export let primary_board: number
  export let date_format: string
  export let boards: Board[]

  const schema = z.object({
    emoji: z.string().nonempty(),
    name: z.string().nonempty(),
    primary_board: z.number(),
    date_format: z.string(),
  })

  type Schema = z.infer<typeof schema>

  const data: Schema = {
    emoji: emoji,
    name: name,
    primary_board: primary_board,
    date_format: date_format,
  }

  const formatOptions = ['DD/MM/YYYY', 'MM/DD/YYYY', 'YYYY/MM/DD']
  let message = ''
  let isSaved = false

  const queryClient = useQueryClient()

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

  const mutationPrimaryBoard = useMutation((board_id: number) => UserAPI.updatePrimaryBoard(board_id), {
    onSuccess: () => {
      updateIsSaved()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const mutationSettings = useMutation(
    ({ name, emoji, format }: { name: string; emoji: string; format: string }) =>
      UserAPI.updateUserSettings({
        app_emoji: emoji,
        app_name: name,
        date_format: format,
      }),
    {
      onSuccess: () => {
        updateIsSaved()
        queryClient.invalidateQueries('settings')
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

  const onSubmit = (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(data)

    if (!result.success) {
      updateMessage('Invalid input')
      return
    }

    if (data.primary_board !== primary_board) {
      $mutationPrimaryBoard.mutate(data.primary_board)
    }

    $mutationSettings.mutate({
      name: data.name,
      emoji: data.emoji,
      format: data.date_format,
    })
  }
</script>

<div class="flex flex-col gap-2 w-full">
  <div class="flex flex-row gap-2 items-center text-black">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="currentColor"
      class="w-6 h-6"
    >
      <path
        d="M11.645 20.91l-.007-.003-.022-.012a15.247 15.247 0 01-.383-.218 25.18 25.18 0 01-4.244-3.17C4.688 15.36 2.25 12.174 2.25 8.25 2.25 5.322 4.714 3 7.688 3A5.5 5.5 0 0112 5.052 5.5 5.5 0 0116.313 3c2.973 0 5.437 2.322 5.437 5.25 0 3.925-2.438 7.111-4.739 9.256a25.175 25.175 0 01-4.244 3.17 15.247 15.247 0 01-.383.219l-.022.012-.007.004-.003.001a.752.752 0 01-.704 0l-.003-.001z"
      />
    </svg>
    <h2 class="font-bold text-lg">Customization</h2>
  </div>
  <form
    on:submit={onSubmit}
    class="flex flex-col gap-10 pt-10 px-4"
  >
    <div class="flex flex-row gap-3 items-center w-full">
      <div class="flex flex-col gap-0">
        <p class="text-[10px] text-tgray-200">Emoji</p>
        <EmojiButton
          bind:emojiSlug={data.emoji}
          emojiKey="settings-customization"
          extraStyles="w-12 h-12 bg-white/20 shadow-lg"
        />
      </div>
      <div class="flex flex-col gap-0 w-full">
        <label
          for="#app_name"
          class="text-[10px] text-tgray-200"
        >
          Name
        </label>
        <input
          type="text"
          id="app_name"
          bind:value={data.name}
          class="text-2xl w-full h-12 px-2 bg-white/20 rounded-lg shadow-lg"
        />
      </div>
    </div>
    <div class="flex flex-row gap-2">
      <label
        for="#date_format"
        class="text-tgray-400 text-sm font-semibold px-2 py-2 w-40"
      >
        Date format
      </label>
      <select
        class="w-full bg-transparent border max-w-[10rem] pl-3 py-1 border-tgray-200 text-tgray-600 rounded-lg"
        bind:value={data.date_format}
        id="date_format"
      >
        {#each formatOptions as format}
          <option
            value={format}
            class="text-tgray-200"
          >
            {format}
          </option>
        {/each}
      </select>
    </div>
    <div class="flex flex-row gap-2">
      <label
        for="#default_board"
        class="text-tgray-400 text-sm font-semibold px-2 py-2 w-40"
      >
        Default Board
      </label>
      <select
        class="w-full bg-transparent border max-w-[10rem] pl-3 py-1 border-tgray-200 text-tgray-600 rounded-lg"
        bind:value={data.primary_board}
        id="default_board"
      >
        {#each boards as board}
          <option
            value={board.id}
            class="text-tgray-200"
          >
            {board.name}
          </option>
        {/each}
      </select>
    </div>
    <SettingsSubmit
      {message}
      {isSaved}
    />
  </form>
</div>

<script lang="ts">
  import { useMutation } from '@sveltestack/svelte-query'
  import { z } from 'zod'
  import UserAPI from '@/api/user'
  import { navigate } from 'svelte-routing'
  import { fade } from 'svelte/transition'

  import { getEmojiURLBySlug, getInitEmoji } from '@/utils/emojis'
  import EmojiButton from '../components/common/EmojiButton.svelte'

  const schema = z.object({
    app_name: z.string().min(3, { message: 'App name must be at least 3 characters long' }),
    app_emoji: z.string().min(1, { message: 'App emoji must be at least 1 character long' }),
  })

  type GetStartedType = z.infer<typeof schema>

  const data: GetStartedType = {
    app_name: 'BanBan',
    app_emoji: getInitEmoji().slug,
  }

  let message = ''

  const gsMutation = useMutation((data: GetStartedType) => UserAPI.gettingStarted(data.app_name, data.app_emoji), {
    onSuccess: async (data) => {
      navigate('/home/new', { replace: true, state: {} })
    },
    onError: (err) => {
      message = err as string

      setTimeout(() => {
        message = ''
      }, 3000)
    },
  })

  const handleSubmit = async (e: Event) => {
    e.preventDefault()

    try {
      schema.parse(data)
    } catch (err) {
      if (err instanceof z.ZodError) {
        message = err.errors[0].message

        setTimeout(() => {
          message = ''
        }, 3000)
      }

      return
    }

    $gsMutation.mutate(data)
  }
</script>

<div class="flex min-h-screen w-full items-center justify-center bg-[#C0C2CC]">
  <div
    transition:fade
    class="mx-3 flex w-full max-w-xl flex-col items-start gap-4 rounded-3xl bg-white/30 px-4 py-6 shadow-lg md:gap-6 md:px-6 md:py-10"
  >
    <div class="flex w-full flex-row items-center justify-center gap-2">
      <img
        src={getEmojiURLBySlug(data.app_emoji)}
        alt="getting-started-emoji"
        class="h-full w-8 object-contain md:w-10"
      />
      <h1 class="text-2xl font-extrabold text-tgray-600 md:text-3xl">{data.app_name}</h1>
    </div>

    <div>
      <p class="text-xl font-bold text-tgray-600 md:text-2xl">
        Welcome! <br />
        Personalize your KanBan.
      </p>
      <p class="text-base font-normal text-tgray-600 sm:text-lg md:text-xl">Don't worry, you can change these in the settings later.</p>
    </div>

    <form
      class="w-full"
      on:submit={handleSubmit}
    >
      <div class="flex w-full flex-row items-center gap-3">
        <div class="relative flex flex-col gap-0">
          <p class="text-[10px] text-tgray-200">Emoji</p>
          <EmojiButton
            bind:emojiSlug={data.app_emoji}
            emojiKey="getting-started"
            extraStyles="w-12 h-12 md:w-14 md:h-14 bg-white/50 shadow-lg"
          />
        </div>
        <div class="flex w-full flex-col gap-0">
          <label
            for="#app_name"
            class="text-[10px] text-tgray-200"
          >
            Name
          </label>
          <input
            type="text"
            id="app_name"
            bind:value={data.app_name}
            class="h-12 w-full rounded-lg bg-white/50 px-2 text-2xl shadow-lg md:h-14 md:text-3xl"
          />
        </div>
      </div>

      <div class="relative mx-auto flex w-full items-center justify-center">
        <p class="absolute left-1/2 top-10 -translate-x-1/2 text-sm font-semibold text-red-600">
          {message}
        </p>
        <input
          type="submit"
          value="Done"
          class="mt-12 w-full max-w-xs cursor-pointer rounded-2xl bg-white/40 py-2 text-2xl font-semibold text-black shadow-lg md:mt-16 md:text-3xl"
        />
      </div>
    </form>
  </div>
</div>

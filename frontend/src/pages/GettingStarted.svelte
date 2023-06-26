<script lang="ts">
  import { useMutation } from '@sveltestack/svelte-query'
  import { z } from 'zod'
  import UserAPI from '@/api/user'
  import { navigate } from 'svelte-routing'
  import { fade } from 'svelte/transition'

  import { getEmojiURLBySlug, getInitEmoji, openEmojiSelector } from '@/utils/emojis'
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
      navigate('/home', { replace: true, state: {} })
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

<div class="w-full min-h-screen bg-[#C0C2CC] flex items-center justify-center">
  <div
    transition:fade
    class="bg-white/30 mx-3 w-full max-w-xl rounded-3xl px-4 md:px-6 py-6 md:py-10 flex shadow-lg flex-col gap-4 md:gap-6 items-start"
  >
    <div class="flex flex-row gap-2 items-center justify-center w-full">
      <img
        src={getEmojiURLBySlug(data.app_emoji)}
        alt="getting-started-emoji"
        class="w-8 md:w-10 h-full object-contain"
      />
      <h1 class="text-2xl md:text-3xl font-extrabold text-tgray-600">{data.app_name}</h1>
    </div>

    <div>
      <p class="font-bold text-xl md:text-2xl text-tgray-600">
        Welcome! <br />
        Personalize your KanBan.
      </p>
      <p class="text-base sm:text-lg md:text-xl font-normal text-tgray-600">Don't worry, you can change these in the settings later.</p>
    </div>

    <form
      class="w-full"
      on:submit={handleSubmit}
    >
      <div class="flex flex-row gap-3 items-center w-full">
        <div class="flex flex-col gap-0 relative">
          <p class="text-[10px] text-tgray-200">Emoji</p>
          <EmojiButton
            bind:emojiSlug={data.app_emoji}
            emojiKey="getting-started"
            extraStyles="w-12 h-12 md:w-14 md:h-14 bg-white/50 shadow-lg"
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
            bind:value={data.app_name}
            class="text-2xl md:text-3xl w-full h-12 md:h-14 px-2 bg-white/50 rounded-lg shadow-lg"
          />
        </div>
      </div>

      <div class="w-full relative mx-auto flex items-center justify-center">
        <p class="absolute top-10 left-1/2 -translate-x-1/2 text-red-600 font-semibold text-sm">
          {message}
        </p>
        <input
          type="submit"
          value="Done"
          class="w-full max-w-xs mt-12 md:mt-16 cursor-pointer text-2xl md:text-3xl shadow-lg text-black font-semibold bg-white/40 rounded-2xl py-2"
        />
      </div>
    </form>
  </div>
</div>

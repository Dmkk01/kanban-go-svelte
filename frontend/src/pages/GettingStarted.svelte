<script lang="ts">
  import { z } from 'zod'
  import UserAPI from '../api/user'
  import { navigate } from 'svelte-routing'

  const schema = z.object({
    app_name: z.string().min(3, { message: 'App name must be at least 3 characters long' }),
    app_emoji: z.string().min(1, { message: 'App emoji must be at least 1 character long' }),
  })

  type GetStartedType = z.infer<typeof schema>

  const data: GetStartedType = {
    app_name: 'BanBan',
    app_emoji: '💗',
  }

  let message = ''

  const handleEmojiInput = (e: Event) => {
    const value = (e.target as HTMLInputElement).value

    data.app_emoji = value
  }

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

    const response = await UserAPI.gettingStarted(data.app_name, data.app_emoji)
    if (response.ok) {
      navigate('/home', { replace: true, state: {} })
    } else {
        console.log(await response.json())
      message = 'Something went wrong'

      setTimeout(() => {
        message = ''
      }, 3000)
    }

  }
</script>

<div class="w-full min-h-screen bg-[#C0C2CC] flex items-center justify-center">
  <div class="bg-white/30 w-full max-w-xl rounded-3xl px-6 py-10 flex shadow-lg flex-col gap-6 items-start">
    <h1 class="text-3xl font-extrabold text-tgray-600 mx-auto">{data.app_emoji} {data.app_name}</h1>

    <div>
      <p class="font-bold text-2xl text-tgray-600">
        Welcome! <br />
        Personalize your KanBan.
      </p>
      <p class="font-xl font-normal text-tgray-600">Don't worry, you can change these in the settings later.</p>
    </div>

    <form
      class="w-full"
      on:submit={handleSubmit}
    >
      <div class="flex flex-row gap-3 items-center w-full">
        <div class="flex flex-col gap-0">
          <label
            for="#app_emoji"
            class="text-[10px] text-tgray-200"
          >
            Emoji
          </label>
          <input
            type="text"
            id="app_emoji"
            value={data.app_emoji}
            on:input={handleEmojiInput}
            class="text-4xl w-14 h-14 px-1 bg-white/50 rounded-lg shadow-lg"
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
            class="text-3xl w-full h-14 px-2 bg-white/50 rounded-lg shadow-lg"
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
          class="w-full max-w-xs mt-16 cursor-pointer text-3xl shadow-lg text-black font-semibold bg-white/40 rounded-2xl py-2"
        />
      </div>
    </form>
  </div>
</div>

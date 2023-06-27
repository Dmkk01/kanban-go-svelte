<script lang="ts">
  import { useMutation } from '@sveltestack/svelte-query'
  import { z } from 'zod'
  import UserAPI from '@/api/user'
  import SettingsInput from './SettingsInput.svelte'
  import SettingsSubmit from '../../common/SubmitButton.svelte'

  export let email: string
  export let username: string

  const schema = z.object({
    username: z
      .string()
      .min(3, {
        message: 'Username must be at least 3 characters long',
      })
      .max(20, {
        message: 'Username must be at most 20 characters long',
      }),
    email: z.string().email({
      message: 'Invalid email',
    }),
  })

  type Schema = z.infer<typeof schema>

  let form: Schema = {
    username: username,
    email: email,
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

  const mutationEmail = useMutation((email: string) => UserAPI.updateEmail(email), {
    onSuccess: () => {
      updateIsSaved()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const mutationUsername = useMutation((username: string) => UserAPI.updateUserName(username), {
    onSuccess: () => {
      updateIsSaved()
    },
    onError: (err) => {
      updateMessage(err as string)
    },
  })

  const onSubmit = async (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(form)

    if (!result.success) {
      updateMessage('Invalid form')
      return
    }

    const { username, email } = result.data

    if (username) {
      $mutationUsername.mutate(username)
    }

    if (email) {
      $mutationEmail.mutate(email)
    }
  }
</script>

<div class="flex w-full flex-col gap-2">
  <div class="flex flex-row items-center gap-2 text-black">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="currentColor"
      class="h-6 w-6"
    >
      <path
        fill-rule="evenodd"
        d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z"
        clip-rule="evenodd"
      />
    </svg>
    <h2 class="text-lg font-bold">Personal Information</h2>
  </div>
  <form
    on:submit={onSubmit}
    class="flex flex-col gap-6 pt-6 lg:gap-10 lg:pt-10"
  >
    <SettingsInput
      label="Name"
      placeholder="Name"
      bind:value={form.username}
      labelWidth="sm"
    />
    <SettingsInput
      label="Email"
      placeholder="Email"
      type="email"
      bind:value={form.email}
      labelWidth="sm"
    />
    <SettingsSubmit
      {message}
      {isSaved}
    />
  </form>
</div>

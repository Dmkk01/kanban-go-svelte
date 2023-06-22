<script lang="ts">
  import { useMutation } from '@sveltestack/svelte-query'
  import { z } from 'zod'
  import UserAPI from '../../../api/user'
  import SettingsInput from './SettingsInput.svelte'
  import SettingsSubmit from './SettingsSubmit.svelte'

  const schema = z.object({
    currentPassword: z.string().min(8, {
      message: 'Password must be at least 8 characters long',
    }),
    newPassword: z.string().min(8, {
      message: 'Password must be at least 8 characters long',
    }),
    repeatPassword: z.string().min(8, {
      message: 'Password must be at least 8 characters long',
    }),
  })

  type Schema = z.infer<typeof schema>

  let form: Schema = {
    currentPassword: '',
    newPassword: '',
    repeatPassword: '',
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

  const mutationPassword = useMutation(
    ({ oldPassword, newPassword }: { oldPassword: string; newPassword: string }) => UserAPI.updatePassword(oldPassword, newPassword),
    {
      onSuccess: () => {
        updateIsSaved()
      },
      onError: (err) => {
        updateMessage(err as string)
      },
    }
  )

  const onSubmit = (e: Event) => {
    e.preventDefault()

    const result = schema.safeParse(form)

    if (!result.success) {
      updateMessage('Invalid input')
      return
    }

    if (form.newPassword !== form.repeatPassword) {
      updateMessage('Passwords do not match')
      return
    }

    $mutationPassword.mutate({
      oldPassword: form.currentPassword,
      newPassword: form.newPassword,
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
        fill-rule="evenodd"
        d="M12 1.5a5.25 5.25 0 00-5.25 5.25v3a3 3 0 00-3 3v6.75a3 3 0 003 3h10.5a3 3 0 003-3v-6.75a3 3 0 00-3-3v-3c0-2.9-2.35-5.25-5.25-5.25zm3.75 8.25v-3a3.75 3.75 0 10-7.5 0v3h7.5z"
        clip-rule="evenodd"
      />
    </svg>
    <h2 class="font-bold text-lg">Security</h2>
  </div>
  <form
    on:submit={onSubmit}
    class="flex flex-col gap-10 pt-10"
  >
    <SettingsInput
      label="Current Password"
      placeholder="Current Password"
      bind:value={form.currentPassword}
      labelWidth="lg"
      type="password"
    />
    <SettingsInput
      label="New Password"
      placeholder="New Password"
      bind:value={form.newPassword}
      labelWidth="lg"
      type="password"
    />

    <SettingsInput
      label="Repeat Password"
      placeholder="Repeat Password"
      bind:value={form.repeatPassword}
      labelWidth="lg"
      type="password"
    />
    <SettingsSubmit
      text="Update Password"
      {message}
      {isSaved}
    />
  </form>
</div>

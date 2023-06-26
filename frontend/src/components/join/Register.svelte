<script lang="ts">
  import { useMutation } from '@sveltestack/svelte-query'
  import { navigate } from 'svelte-routing/src/history'
  import AuthAPI from '@/api/auth'
  import InputField from './common/InputField.svelte'
  import { z } from 'zod'

  const schema = z.object({
    email: z.string().email({ message: 'Invalid email' }),
    username: z.string().min(3, { message: 'Username must be at least 3 characters long' }),
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
    repeatPassword: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
  })

  type RegisterSchema = z.infer<typeof schema>

  const data: RegisterSchema = {
    email: '',
    username: '',
    password: '',
    repeatPassword: '',
  }

  let message: string = ''

  const registerMutation = useMutation((data: RegisterSchema) => AuthAPI.register(data.email, data.username, data.password), {
    onSuccess: async (data, request) => {
      const login = await AuthAPI.login(request.email, request.password)

      localStorage.setItem('token', login.token)
      navigate('/home', { replace: true, state: {} })
    },
    onError: (err) => {
      message = err as string

      setTimeout(() => {
        message = ''
      }, 3000)
    },
  })

  const submitData = async (e: Event) => {
    e.preventDefault()

    if (data.password !== data.repeatPassword) {
      message = 'Passwords do not match'

      setTimeout(() => {
        message = ''
      }, 3000)

      return
    }

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

    $registerMutation.mutate(data)
  }
</script>

<div class="flex flex-col gap-0 w-full max-w-md">
  <h2 class="text-3xl text-tgray-600 font-bold mb-6 sm:mb-10 md:mb-16">Welcome!</h2>

  <form
    on:submit={submitData}
    class="flex flex-col gap-5 md:gap-8 w-full"
  >
    <InputField
      bind:value={data.username}
      placeholder="Username"
      type="text"
      emoji="🙂"
      required
    />
    <InputField
      bind:value={data.email}
      placeholder="Email"
      type="email"
      emoji="📧"
      required
    />
    <InputField
      bind:value={data.password}
      placeholder="Password"
      type="password"
      emoji="🔒"
      required
    />
    <InputField
      bind:value={data.repeatPassword}
      placeholder="Repeat Password"
      type="password"
      emoji="🔏"
      required
    />

    <div class="w-full relative">
      <p class="absolute w-full top-5 left-1/2 -translate-x-1/2 text-red-600 font-semibold text-sm">
        {message}
      </p>
      <input
        type="submit"
        value="Register"
        class="w-full cursor-pointer my-6 md:my-10 text-xl sm:text-2xl md:text-3xl shadow-lg text-black font-semibold bg-white/40 rounded-md py-2"
      />
    </div>
  </form>
  <div class="flex flex-row gap-1 text-sm md:text-base mx-auto font-semibold">
    <p class="text-tgray-600">Already have an account?</p>
    <a
      href="/?type=login"
      class="text-rose font-bold"
    >
      Login
    </a>
  </div>
</div>

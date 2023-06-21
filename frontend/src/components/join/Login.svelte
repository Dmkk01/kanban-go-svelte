<script lang="ts">
  import AuthAPI from '../../api/auth'
  import InputField from './common/InputField.svelte'
  import { z } from 'zod'
  import { navigate } from "svelte-routing"

  const schema = z.object({
    email: z.string().email({ message: 'Invalid email' }),
    password: z.string().min(8, { message: 'Password must be at least 8 characters long' }),
  })

  type LoginSchema = z.infer<typeof schema>

  const data: LoginSchema = {
    email: '',
    password: '',
  }

  let message: string = ''

  const submitData = async (e: Event) => {
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

    const response = await AuthAPI.login(data.email, data.password)
    const json = await response.json()

    if (response.ok) {
      localStorage.setItem('token', json.token)
      navigate('/home', { replace: true, state: {} })
    } else {
      message = json.message

      setTimeout(() => {
        message = ''
      }, 3000)
    }
  }
</script>

<div class="flex flex-col gap-0 w-full max-w-md">
  <h2 class="text-3xl text-tgray-600 font-bold mb-16">Welcome back!</h2>

  <form
    on:submit={submitData}
    class="flex flex-col gap-8 w-full"
  >
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

    <div class="w-full relative">
      <p class="absolute w-full top-5 left-1/2 -translate-x-1/2 text-red-600 font-semibold text-sm">
        {message}
      </p>
      <input
        type="submit"
        value="Login"
        class="w-full my-10 cursor-pointer text-3xl shadow-lg text-black font-semibold bg-white/40 rounded-2xl py-2"
      />
    </div>
  </form>
  <div class="flex flex-row gap-1 text-base mx-auto font-semibold">
    <p class="text-tgray-600">Don't have an account yet?</p>
    <a
      href="/?type=register"
      class="text-rose font-bold"
    >
      Sign up
    </a>
  </div>
</div>

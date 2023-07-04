<script lang="ts">
  import Login from '@/components/join/Login.svelte'
  import Register from '@/components/join/Register.svelte'
  import { fade } from 'svelte/transition'
  const query = new URLSearchParams(window.location.search)

  const currentPage: 'login' | 'register' = query.get('type') === 'register' ? 'register' : 'login'
</script>

<div class="flex min-h-screen w-full items-center justify-center bg-[#C0C2CC]">
  <div
    transition:fade
    class="mx-4 flex w-full max-w-xl flex-col items-center gap-6 rounded-xl bg-white/30 px-4 py-6 shadow-lg md:gap-8 md:px-6 md:py-10"
  >
    <h1 class="text-xl font-extrabold text-tgray-600 sm:text-2xl md:text-3xl">💗 BanBan</h1>
    {#if currentPage === 'login'}
      <Login />
    {:else}
      <Register />
    {/if}
  </div>
</div>

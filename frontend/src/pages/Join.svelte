<script lang="ts">
  import Login from '@/components/join/Login.svelte'
  import Register from '@/components/join/Register.svelte'
  import { fade } from 'svelte/transition'
  const query = new URLSearchParams(window.location.search)

  const currentPage: 'login' | 'register' = query.get('type') === 'register' ? 'register' : 'login'
</script>

<div class="w-full min-h-screen bg-[#C0C2CC] flex items-center justify-center">
  <div
    transition:fade
    class="bg-white/30 mx-4 w-full max-w-xl rounded-xl px-4 md:px-6 py-6 md:py-10 flex shadow-lg flex-col gap-6 md:gap-8 items-center"
  >
    <h1 class="text-xl sm:text-2xl md:text-3xl font-extrabold text-tgray-600">💗 BanBan</h1>
    {#if currentPage === 'login'}
      <Login />
    {:else}
      <Register />
    {/if}
  </div>
</div>

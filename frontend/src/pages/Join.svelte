<script lang="ts">
  import Login from '../components/join/Login.svelte'
  import Register from '../components/join/Register.svelte'
  const query = new URLSearchParams(window.location.search)

  const currentPage: 'login' | 'register' = query.get('type') === 'register' ? 'register' : 'login'
</script>

<div class="w-full min-h-screen bg-[#C0C2CC] flex items-center justify-center">
  <div class="bg-white/30 w-full max-w-xl rounded-3xl px-6 py-10 flex shadow-lg flex-col gap-8 items-center">
    <h1 class="text-3xl font-extrabold text-tgray-600">💗 BanBan</h1>
    {#if currentPage === 'login'}
      <Login />
    {:else}
      <Register />
    {/if}
  </div>
</div>

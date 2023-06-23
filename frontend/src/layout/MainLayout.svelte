<script lang="ts">
  import BoardDrawer from '@/components/drawers/board/BoardDrawer.svelte'
  import Sidebar from '@/components/home/sidebar/Sidebar.svelte'
  import store from '@/store'
</script>

<div class="w-full h-screen flex flex-row bg-[#C0C2CC] relative">
  <Sidebar />

  <div class="px-4 py-2 h-screen w-full overflow-hidden">
    <slot />
  </div>
  {#if $store.boardDrawer.isOpen}
    <BoardDrawer />
  {/if}
</div>

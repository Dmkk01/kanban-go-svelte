<script lang="ts">
  import Sidebar from '../components/home/sidebar/Sidebar.svelte'
</script>

<div class="w-full h-screen flex flex-row bg-[#C0C2CC]">
  <Sidebar />

  <div class="px-4 py-2 min-h-screen w-full">
    <slot />
  </div>
</div>

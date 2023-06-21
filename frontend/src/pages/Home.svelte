<script lang="ts">
  import { onMount } from 'svelte'
  import UserAPI from '../api/user'
  import { navigate } from 'svelte-routing/src/history'
  import Sidebar from '../components/Home/Sidebar.svelte'
  import BoardsAPI from '../api/board'
  import { useQuery } from '@sveltestack/svelte-query'
</script>

<div class="w-full h-screen flex flex-row">
  <Sidebar />

  <div>Board</div>
</div>

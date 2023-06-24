<script lang="ts">
  import UserAPI from '@/api/user'
  import MainLayout from '@/layout/MainLayout.svelte'
  import { useQuery } from '@sveltestack/svelte-query'
  import { navigate } from 'svelte-routing/src/history'

  const settings = useQuery('settings', UserAPI.getUserSettings, {
    refetchOnMount: true,
    onSuccess: async (data) => {
      if (data.primary_board_id) {
        navigate(`/home/board/${data.primary_board_id}`, { replace: true, state: {} })
      }
    },
  })
</script>

<MainLayout>
  <div>
    {#if $settings.isLoading}
      <p>Loading...</p>
    {/if}
  </div>
</MainLayout>

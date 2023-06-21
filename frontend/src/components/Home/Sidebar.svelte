<script lang="ts">
  import { useQuery } from '@sveltestack/svelte-query'
  import UserAPI from '../../api/user'
  import { navigate } from 'svelte-routing/src/history'
  import BoardsAPI from '../../api/board'

  const settings = useQuery('settings', UserAPI.getUserSettings, {
    onSuccess: async (data) => {
      if (data.app_name === '') {
        navigate('/getting-started')
      }
    },
  })

  const boards = useQuery('boards', BoardsAPI.getBoards)
</script>

<div class="w-full max-w-xs border" />

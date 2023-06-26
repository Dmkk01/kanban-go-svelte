<script lang="ts">
  import { useQuery } from '@sveltestack/svelte-query'
  import SettingsCustomization from '@/components/home/settings/SettingsCustomization.svelte'
  import SettingsInfo from '@/components/home/settings/SettingsInfo.svelte'
  import SettingsSecurity from '@/components/home/settings/SettingsSecurity.svelte'
  import MainLayout from '@/layout/MainLayout.svelte'
  import UserAPI from '@/api/user'
  import BoardsAPI from '@/api/board'

  const userSettings = useQuery('user-settings', async () => await UserAPI.getUserSettings())
  const boards = useQuery('boards', async () => await BoardsAPI.getBoards())

  const user = useQuery('user', async () => await UserAPI.getUser())
</script>

<MainLayout>
  <div class="flex flex-col gap-4 px-4 py-4 h-screen overflow-y-auto">
    <div class="flex flex-row gap-2 items-center text-2xl font-bold">
      <p>⚙️ Settings</p>
    </div>
    {#if $user.data && $userSettings.data && $boards.data}
      <div class="grid grid-cols-2 2xl:grid-cols-3 gap-6">
        <div class="flex-1 w-full max-w-xl h-full 2xl:h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsInfo
            username={$user.data.username}
            email={$user.data.email}
          />
        </div>
        <div class="flex-1 w-full max-w-xl h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsSecurity />
        </div>
        <div class="flex-1 w-full max-w-xl h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsCustomization
            emoji={$userSettings.data.app_emoji}
            name={$userSettings.data.app_name}
            primary_board={$userSettings.data.primary_board_id}
            date_format={$userSettings.data.date_format}
            boards={$boards.data}
          />
        </div>
      </div>
    {/if}
  </div>
</MainLayout>

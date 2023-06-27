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
  <div class=" flex h-screen flex-col gap-4 overflow-y-auto px-4 py-4">
    <div class="flex flex-row items-center gap-2 text-2xl font-bold">
      <p>⚙️ Settings</p>
    </div>
    {#if $user.data && $userSettings.data && $boards.data}
      <div class="grid grid-cols-1 gap-6 pb-28 lg:grid-cols-2 2xl:grid-cols-3">
        <div class="h-full w-full max-w-xl flex-1 rounded-lg bg-white/20 px-4 py-6 shadow-lg 2xl:h-fit">
          <SettingsInfo
            username={$user.data.username}
            email={$user.data.email}
          />
        </div>
        <div class="h-fit w-full max-w-xl flex-1 rounded-lg bg-white/20 px-4 py-6 shadow-lg">
          <SettingsSecurity />
        </div>
        <div class="h-fit w-full max-w-xl flex-1 rounded-lg bg-white/20 px-4 py-6 shadow-lg">
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

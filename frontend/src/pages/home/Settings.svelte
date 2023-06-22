<script lang="ts">
  import { useQuery } from '@sveltestack/svelte-query'
  import SettingsCustomization from '../../components/home/settings/SettingsCustomization.svelte'
  import SettingsInfo from '../../components/home/settings/SettingsInfo.svelte'
  import SettingsSecurity from '../../components/home/settings/SettingsSecurity.svelte'
  import MainLayout from '../../layout/MainLayout.svelte'
  import UserAPI from '../../api/user'

  const userSettings = useQuery('user-settings', async () => await UserAPI.getUserSettings())

  const user = useQuery('user', () => UserAPI.getUser())
</script>

<MainLayout>
  <div class="flex flex-col gap-4 px-4 py-4">
    <div class="flex flex-row gap-2 items-center text-2xl font-bold">
      <p>⚙️</p>
      <p>Settings</p>
    </div>
    {#if !$user.isLoading}
      <div class="flex flex-row gap-6 flex-wrap justify-around">
        <div class="flex-1 max-w-xl h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsInfo
            username={$user.data.username}
            email={$user.data.email}
          />
        </div>
        <div class="flex-1 max-w-xl h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsSecurity />
        </div>
        <div class="flex-1 max-w-xl h-fit bg-white/20 shadow-lg rounded-lg px-4 py-6">
          <SettingsCustomization />
        </div>
      </div>
    {/if}
  </div>
</MainLayout>

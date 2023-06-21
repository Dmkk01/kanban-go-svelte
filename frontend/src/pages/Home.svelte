<script lang="ts">
  import { onMount } from "svelte"
  import UserAPI from "../api/user";
  import { navigate } from "svelte-routing/src/history"


    onMount(async () => {
        const response = await UserAPI.getUserSettings()
        const data = await response.json() as UserSettings

        if (data.app_name === "") {
            navigate('/getting-started')
        }
    })
</script>

<div class="text-6xl text-red-500">Home</div>

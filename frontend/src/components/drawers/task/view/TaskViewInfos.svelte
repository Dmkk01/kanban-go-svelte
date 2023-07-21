<script lang="ts">
  import UserAPI from '@/api/user'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import { useQuery } from '@sveltestack/svelte-query'

  export let created_at: string
  export let due_date: string
  export let time_needed: number

  const userSettings = useQuery('user-settings', async () => await UserAPI.getUserSettings())

  const formatDate = (dateString: string) => {
    if (!$userSettings.data) return ''
    const format = $userSettings.data.date_format
    const date = new Date(dateString)

    const day = date.getDate()
    const month = date.getMonth() + 1
    const year = date.getFullYear()

    switch (format) {
      case 'DD/MM/YYYY':
        return `${day}/${month}/${year}`
      case 'MM/DD/YYYY':
        return `${month}/${day}/${year}`
      case 'YYYY/MM/DD':
        return `${year}/${month}/${day}`
      default:
        return `${day}/${month}/${year}`
    }
  }
</script>

<div class="flex w-full flex-row justify-between">
  <div class="flex flex-col gap-1">
    <h3 class="text-base font-bold text-tgray-600">Created date</h3>
    <div class="ml-6 flex flex-row gap-3">
      <img
        src={getEmojiURLBySlug('apple-calendar')}
        alt="Created date"
        class="h-5 w-5"
      />
      <p class="text-base font-medium">
        {formatDate(created_at)}
      </p>
    </div>
  </div>

  {#if !due_date.includes('0001-01-01')}
    <div class="flex flex-col gap-1">
      <h3 class="text-base font-bold text-tgray-600">Due date</h3>
      <div class="ml-6 flex flex-row gap-3">
        <img
          src={getEmojiURLBySlug('apple-calendar')}
          alt="Due date"
          class="h-5 w-5"
        />
        <p class="text-base font-medium">
          {formatDate(due_date)}
        </p>
      </div>
    </div>
  {/if}

  {#if time_needed > 0}
    <div class="flex flex-col gap-1">
      <h3 class="text-base font-bold text-tgray-600">Approx. time</h3>
      <div class="ml-6 flex flex-row gap-3">
        <img
          src={getEmojiURLBySlug('apple-alarm-clock')}
          alt="Approx. time"
          class="h-5 w-5"
        />
        <p class="text-base font-medium">
          {time_needed} min
        </p>
      </div>
    </div>
  {/if}
</div>

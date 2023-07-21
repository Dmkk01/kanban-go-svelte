<script lang="ts">
  import store from '@/store'
  import { getEmojiURLBySlug } from '@/utils/emojis'
  import CloseButton from '@/components/common/CloseButton.svelte'

  export let data: { task: Task; column: ColumnBoard }
  export let closeDrawer: () => void
</script>

<div class="flex flex-row items-center justify-between gap-2">
  <div class="flex flex-row items-center gap-4">
    <h2 class="text-xl font-bold text-tgray-600">
      {data.task.title ?? ''}
    </h2>
    <div class="flex flex-row items-center gap-2 rounded-md border border-tgray-600 px-3 py-1">
      <img
        src={getEmojiURLBySlug(data.column.emoji)}
        alt={data.column.name}
        class="h-5 w-5"
      />
      <p class="text-base font-medium">
        {data.column.name}
      </p>
    </div>
  </div>
  <button
    type="button"
    class="text-tgray-600"
    on:click={() => {
      closeDrawer()
      $store.taskDrawer.ids.board = data.task.board_id ?? 0
      $store.taskDrawer.ids.column = data.task.column_id ?? 0
      $store.taskDrawer.ids.task = data.task.id ?? 0
      $store.taskDrawer.isOpen = true
    }}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2"
      stroke="currentColor"
      class="h-6 w-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10"
      />
    </svg>
  </button>
  <CloseButton onClickHandler={closeDrawer} />
</div>

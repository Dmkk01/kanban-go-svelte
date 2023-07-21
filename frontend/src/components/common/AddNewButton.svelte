<script lang="ts">
  export let onClickHandler: () => void
  export let text: string
</script>

<div>
  <button
    type="button"
    on:click={onClickHandler}
    class="flex w-fit flex-row items-center gap-2 rounded-md border border-tgray-200 px-2 py-1"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2.5"
      stroke="currentColor"
      class="h-5 w-5"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 4.5v15m7.5-7.5h-15"
      />
    </svg>
    <p class="text-sm font-medium">{text}</p>
  </button>
</div>

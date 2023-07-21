<script lang="ts">
  export let handleSubmit: (e: Event) => void
  export let tag: string
</script>

<div class="flex flex-col gap-1">
  <p class="text-sm font-medium text-tgray-600">Create a new tag</p>
  <form
    on:submit|preventDefault={handleSubmit}
    class="flex flex-row items-center gap-3"
  >
    <input
      type="text"
      class="text-md w-full rounded-md border border-tgray-200 px-2 py-1.5"
      bind:value={tag}
      placeholder="Tag name"
      required
    />
    <button
      type="submit"
      class="flex aspect-square h-7 w-7 items-center justify-center rounded-md border border-tgray-600"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2.5"
        stroke="currentColor"
        class="h-4 w-4"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M4.5 12.75l6 6 9-13.5"
        />
      </svg>
    </button>
  </form>
</div>

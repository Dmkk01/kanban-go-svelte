<script lang="ts">
  export let text: string = 'Save'
  export let message: string = ''
  export let isSaved: boolean = false
  export let extraStyles: string = ''
  export let textStyles: string = ''
</script>

<button
  type="submit"
  class={`relative mx-auto flex min-w-[15rem] max-w-xs items-center justify-center gap-2 rounded-lg border border-tgray-600 bg-white/40 px-6 py-2 ${extraStyles}`}
>
  {#if message}
    <p class="absolute -left-16 -right-16 -top-5 text-sm font-semibold text-red-500">
      {message}
    </p>
  {/if}
  {#if isSaved}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      class="h-6 w-6"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
      />
    </svg>
  {/if}
  <p class={`relative text-lg font-semibold ${textStyles}`}>
    {text}
  </p>
</button>

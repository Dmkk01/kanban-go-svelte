<script lang="ts">
  export let placeholder: string = 'placeholder'
  export let type: string = 'text'
  export let value: string = ''
  export let emoji: string = '🙂'
  export let required: boolean = false

  function handleInput(e: Event) {
    value = (e.target as HTMLInputElement).value
  }
</script>

<div class="flex flex-row bg-white/20 p-1.5 rounded-md gap-2 shadow-md items-center">
  <div class="p-0.5 h-8 md:h-10 w-10 md:w-12 flex items-center justify-center text-xl md:text-2xl bg-white/50 rounded-md">
    {emoji}
  </div>
  <input
    class="w-full bg-transparent text-base text-tgray-600 px-2 py-1 md:py-2 placeholder-tgray-200"
    {type}
    {required}
    {placeholder}
    on:input={handleInput}
  />
</div>

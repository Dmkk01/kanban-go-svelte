<script lang="ts">
  export let placeholder: string = 'placeholder'
  export let type: string = 'text'
  export let value: string = ''
  export let label: string = ''
  export let labelWidth: 'sm' | 'md' | 'lg' = 'md'

  function handleInput(e: Event) {
    value = (e.target as HTMLInputElement).value
  }
</script>

<div class="flex flex-row items-center gap-2 rounded-md p-1.5">
  <label
    for={label + 'input'}
    class={`text-tgray-400 px-2 py-2 text-sm font-semibold ${labelWidth === 'sm' ? 'w-28' : labelWidth === 'md' ? 'w-40' : 'w-52'}`}
  >
    {label}
  </label>
  <input
    id={label + 'input'}
    class="w-full rounded-lg bg-white/20 px-4 py-2 text-base text-tgray-600 placeholder-tgray-200 shadow-lg"
    {type}
    required
    {placeholder}
    on:input={handleInput}
    {value}
  />
</div>

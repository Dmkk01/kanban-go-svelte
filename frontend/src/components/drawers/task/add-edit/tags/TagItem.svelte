<script lang="ts">
  import { createEventDispatcher } from 'svelte'

  export let title: string = ''
  export let color: string = ''
  export let isNew: boolean = false
  export let isCurrent: boolean = false

  let isOptionsOpen = false
  let updateTitle = title

  const colorOptions = ['#A2E8DD', '#DFFFD6', '#ABE4FF', '#FFDAAB', '#FFDDFF', '#D4DCFF', '#E8D7FF', '#E8F4C8']

  const dispatch = createEventDispatcher<{
    'create-new': void
    'add-new': void
    'delete-current': void
    'change-color': string
    'full-delete': void
    'update-title': string
  }>()

  const handleClickText = () => {
    if (!isCurrent) {
      dispatch('add-new')
    }

    isOptionsOpen = false
  }

  const handleDelete = () => {
    dispatch('delete-current')
  }

  const handleFullDelete = () => {
    dispatch('full-delete')
  }

  const handleChangeColor = (colorOption: string) => {
    dispatch('change-color', colorOption)
  }

  $: {
    dispatch('update-title', updateTitle)
  }
</script>

<button
  type="button"
  class="px-2 py-0.5 w-fit relative rounded-md flex flex-row gap-0 items-center z-30"
  style="background-color: {color}"
>
  {#if !isNew && !isCurrent}
    <button
      type="button"
      class="-ml-2 cursor-help"
      on:click={() => {
        isOptionsOpen = !isOptionsOpen
      }}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="currentColor"
        class="w-4 h-4"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
        />
      </svg>
    </button>

    {#if isOptionsOpen}
      <div class="bg-white flex flex-col gap-2 absolute -left-40 w-36 rounded-md border border-tgray-600 shadow-lg">
        <input
          type="text"
          class="bg-transparent mx-2 mt-2 border border-tgray-600 rounded-md px-2 py-0.5 text-sm z-50"
          bind:value={updateTitle}
        />

        <button
          type="button"
          on:click={handleFullDelete}
          class="flex flex-row gap-2 items-center px-2"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-6 h-6"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
            />
          </svg>

          <p class="text-sm">Delete</p>
        </button>

        <hr class="border-tgray-600" />

        <div class="grid grid-cols-4 grid-rows-2 place-items-center px-2 pb-2 gap-2">
          {#each colorOptions as colorOption}
            <button
              type="button"
              class="w-5 h-5 flex items-center justify-center"
              style="background-color: {colorOption}"
              on:click={() => handleChangeColor(colorOption)}
            >
              {#if colorOption === color}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="2.5"
                  stroke="currentColor"
                  class="w-6 h-6 text-black"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M4.5 12.75l6 6 9-13.5"
                  />
                </svg>
              {/if}
            </button>
          {/each}
        </div>
      </div>
    {/if}
  {/if}

  <p
    on:click={handleClickText}
    on:keypress={handleClickText}
    class="text-sm font-semibold select-none"
  >
    {title}
  </p>

  {#if isCurrent}
    <button
      type="button"
      class="ml-1"
      on:click={handleDelete}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="currentColor"
        class="w-4 h-4"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M6 18L18 6M6 6l12 12"
        />
      </svg>
    </button>
  {/if}
</button>

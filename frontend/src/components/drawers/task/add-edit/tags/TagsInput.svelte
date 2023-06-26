<script lang="ts">
  import BoardsAPI from '@/api/board'
  import { useQuery } from '@sveltestack/svelte-query'
  import { fade } from 'svelte/transition'
  import TagItem from './TagItem.svelte'
  import { createEventDispatcher, onMount } from 'svelte'

  export let boardID: number = 0

  const dispatch = createEventDispatcher<{
    'new-added': TagType[]
    updated: TagType[]
    deleted: TagType[]
  }>()

  const tags = useQuery('board-tags', () => BoardsAPI.getTagsBoard(boardID), {
    refetchOnWindowFocus: false,
    onSuccess: (data) => {
      tagOptions = data.map((tag) => {
        return {
          id: tag.id,
          title: tag.title,
          color: tag.color,
        }
      })

      tagOptions = [
        {
          color: '#A2E8DD',
          id: 1,
          title: 'Tag 1',
          isNew: true,
        },
        {
          color: '#ABE4FF',
          id: 2,
          title: 'Tag 4',
          isNew: true,
        },
        {
          color: '#E8F4C8',
          id: 3,
          title: 'Tag Tag',
          isNew: true,
        },
      ]
    },
  })
  let currentTags: TagType[] = []
  let tagOptions: TagType[] = []

  const colorOptions = ['#A2E8DD', '#DFFFD6', '#ABE4FF', '#FFDAAB', '#FFDDFF', '#D4DCFF', '#E8D7FF', '#E8F4C8']
  const getDefaultTag = () => {
    return {
      title: '',
      color: colorOptions[Math.floor(Math.random() * colorOptions.length)],
      isNew: true,
      id: Math.floor(Math.random() * 1000),
    }
  }
  let newTag = getDefaultTag()
  let isFocus = true

  const createNew = () => {
    currentTags = [...currentTags, newTag]
    tagOptions = [...tagOptions, newTag]
    newTag = getDefaultTag()
  }

  onMount(() => {
    $tags.refetch()
  })
</script>

<div class="w-full relative">
  <div class="flex flex-row flex-wrap gap-4 w-full border border-tgray-600 rounded-md h-auto p-1">
    {#each currentTags as tag (`current-tag-${tag.id}`)}
      <TagItem
        title={tag.title}
        color={tag.color}
        isCurrent={true}
        on:delete-current={() => {
          currentTags = currentTags.filter((t) => t.id !== tag.id)
        }}
      />
    {/each}

    <input
      type="text"
      class="text-sm font-bold text-tgray-600 w-fit bg-transparent"
      bind:value={newTag.title}
      on:focus={() => (isFocus = true)}
      on:blur={() => (isFocus = false)}
    />
  </div>
  {#if isFocus}
    <div
      transition:fade
      class="bg-white/50 top-10 absolute border border-tgray-600 rounded-md backdrop-blur-lg p-3 flex flex-col w-full gap-2"
    >
      <p class="text-sm font-medium">Select an option or type to create one</p>
      {#each tagOptions as tag (`option-tag-${tag.id}`)}
        <TagItem
          bind:title={tag.title}
          color={tag.color}
          on:add-new={() => {
            if (!currentTags.find((t) => t.id === tag.id)) {
              currentTags = [...currentTags, tag]
            }
          }}
          on:change-color={(e) => {
            currentTags = currentTags.map((t) => {
              if (t.id === tag.id) {
                t.color = e.detail
              }
              return t
            })

            tagOptions = tagOptions.map((t) => {
              if (t.id === tag.id) {
                t.color = e.detail
              }
              return t
            })
          }}
          on:full-delete={() => {
            tagOptions = tagOptions.filter((t) => t.id !== tag.id)
            currentTags = currentTags.filter((t) => t.id !== tag.id)
          }}
          on:update-title={(e) => {
            tagOptions = tagOptions.map((t) => {
              if (t.id === tag.id) {
                t.title = e.detail
              }
              return t
            })

            currentTags = currentTags.map((t) => {
              if (t.id === tag.id) {
                t.title = e.detail
              }
              return t
            })
          }}
        />
      {/each}

      {#if newTag.title !== ''}
        <div class="flex flex-row gap-2 items-center">
          <p class="text-sm font-medium">Create</p>
          <TagItem
            title={newTag.title}
            color={newTag.color}
            isNew={true}
            on:create-new={createNew}
          />
        </div>
      {/if}
    </div>
  {/if}
</div>

<script lang="ts">
  import { getEmojiURLBySlug } from '@/utils/emojis'

  export let links: LinkTask[] = []
</script>

{#if links.length > 0}
  <div class="flex flex-col gap-1">
    <h3 class="text-base font-bold text-tgray-600">Links</h3>
    <div class="my-2 ml-6 flex flex-col gap-2">
      {#each links as link (link.id)}
        <div class="flex flex-row gap-4">
          <img
            src={getEmojiURLBySlug(link.emoji)}
            alt={link.url}
            class="h-5 w-5"
          />
          <a
            href={link.url}
            target="_blank"
            class="text-base font-medium underline"
          >
            {link.url}
          </a>
        </div>
      {/each}
    </div>
  </div>
{/if}

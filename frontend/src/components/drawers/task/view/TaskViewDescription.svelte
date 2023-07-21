<script lang="ts">
  export let description: string = ''
</script>

{#if description.length > 0}
  <div class="flex flex-col gap-1">
    <h3 class="text-base font-bold text-tgray-600">Description</h3>
    <p class="ml-6 text-base font-medium text-tgray-600">
      {description}
    </p>
  </div>
{/if}

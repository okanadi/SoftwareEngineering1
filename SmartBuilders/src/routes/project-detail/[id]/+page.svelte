<script lang="ts">
  import type { PageData } from './$types';
  import ProjectDetails from '$lib/components/ProjectDetails.svelte';

  // The `data` prop is passed from your +page.ts load function
  export let data: PageData;
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 flex justify-center">
  <div class="container mx-auto p-4 md:p-8">
    {#if data.project}
      <ProjectDetails project={data.project} />
    {:else if data.error}
      <div class="text-center p-8 bg-red-500/10 border border-red-500/30 rounded-2xl">
        <h1 class="text-2xl font-bold text-red-400">Error</h1>
        <p class="text-white/80 mt-2">{data.error}</p>
      </div>
    {:else}
      <p class="text-center">Loading project details...</p>
    {/if}
  </div>
</main>

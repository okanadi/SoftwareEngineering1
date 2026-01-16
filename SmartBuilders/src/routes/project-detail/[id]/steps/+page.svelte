<script lang="ts">
  import type { PageData } from './$types';

  export let data: PageData;

  function goBack() {
    history.back();
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32">
  <div class="max-w-6xl mx-auto">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-4xl font-bold text-white">Project Steps</h1>
      <button
        on:click={goBack}
        class="px-6 py-2 bg-white/20 hover:bg-white/30 text-white rounded-lg font-semibold transition"
      >
        ← Back
      </button>
    </div>

    {#if data.steps && data.steps.length > 0}
      <div class="space-y-4">
        {#each data.steps as step (step.id)}
          <div class="bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-6">
            <div class="flex justify-between items-start mb-4">
              <div>
                <h2 class="text-2xl font-bold text-white">{step.title}</h2>
                <p class="text-white/70 text-sm mt-1">{step.description}</p>
              </div>
              <span class="px-3 py-1 rounded-full text-sm font-semibold bg-blue-500/30 text-blue-300">
                {step.progress}
              </span>
            </div>

            <div class="grid grid-cols-2 gap-4 mt-4">
              <div>
                <p class="text-white/60 text-sm">Start Date</p>
                <p class="text-white font-semibold">
                  {step.start_date ? new Date(step.start_date).toLocaleDateString() : 'Not set'}
                </p>
              </div>
              <div>
                <p class="text-white/60 text-sm">End Date</p>
                <p class="text-white font-semibold">
                  {step.end_date ? new Date(step.end_date).toLocaleDateString() : 'Not set'}
                </p>
              </div>
            </div>

            <div class="mt-4 text-xs text-white/50">
              Created: {step.created_at ? new Date(step.created_at).toLocaleString() : 'Unknown'}
            </div>
          </div>
        {/each}
      </div>
    {:else if data.error}
      <div class="text-center p-8 bg-red-500/10 border border-red-500/30 rounded-2xl">
        <h2 class="text-2xl font-bold text-red-400">No Steps Found</h2>
        <p class="text-white/80 mt-2">{data.error}</p>
      </div>
    {:else}
      <div class="text-center p-8">
        <p class="text-white/80 text-lg">Loading project steps...</p>
      </div>
    {/if}
  </div>
</main>

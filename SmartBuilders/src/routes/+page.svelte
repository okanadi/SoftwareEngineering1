<script lang="ts">
  import Card from '$lib/components/Card.svelte';
  import { goto } from '$app/navigation';

  let projectId = '';
  let managerId = '';
  let lastname = '';
  let errorMessage = '';

  function handleManagerIdSubmit() {
    if (managerId.trim()) {
      errorMessage = '';
      goto(`/projects?managerId=${encodeURIComponent(managerId)}`);
    } else {
      errorMessage = 'Please enter a Manager ID.';
    }
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 lg:pt-40">
  <h1 class="text-5xl font-bold mb-4">Willkommen bei SmartBuilders!</h1>

  {#if errorMessage}
    <div class="max-w-7xl mx-auto mb-6 p-3 rounded-lg bg-red-500/20 text-red-300 text-sm text-center">
      {errorMessage}
    </div>
  {/if}

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto p-10">

  <Card
    title="Meine Projekte anschauen"
    inputPlaceholder="Manager ID eingeben"
    bind:inputValue={managerId}
    onArrowClick={handleManagerIdSubmit}
  />
  
  <div class="bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-6 flex flex-col justify-between">
    <h2 class="text-xl font-bold text-white mb-4">Gehe direkt zum Projekt!</h2>
    <div class="space-y-4 mb-4">
      <input
        type="text"
        placeholder="Projekt-ID"
        bind:value={projectId}
        class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 transition-colors"
      />
      <input
        type="text"
        placeholder="Nachname"
        bind:value={lastname}
        class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 transition-colors"
      />
    </div>
    <button
      on:click={() => {
        if (projectId.trim() && lastname.trim()) {
          errorMessage = '';
          goto(`/project-detail/${projectId}`);
        } else {
          errorMessage = 'Please enter both Project ID and Lastname.';
        }
      }}
      class="self-end p-3 bg-white/10 hover:bg-white/20 rounded-full text-white transition-colors"
      aria-label="Go to project"
    >
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
      </svg>
    </button>
  </div>

  <Card
    title="Erstelle ein neues Projekt!"
    onArrowClick={() => goto('/create-project')}
  />
</div>
</main>

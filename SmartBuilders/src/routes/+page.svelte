<script lang="ts">
  import Card from '$lib/components/Card.svelte';
  import { goto } from '$app/navigation';

  let projectId = '';
  let managerId = '';

  function handleManagerIdSubmit() {
    if (managerId.trim()) {
      goto(`/projects?managerId=${encodeURIComponent(managerId)}`);
    } else {
      alert('Please enter a Manager ID.');
    }
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 lg:pt-40">
  <h1 class="text-5xl font-bold mb-4">Willkommen bei SmartBuilders!</h1>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto p-10">

  <Card
    title="Meine Projekte anschauen"
    inputPlaceholder="Manager ID eingeben"
    bind:inputValue={managerId}
    onArrowClick={handleManagerIdSubmit}
  />
  
     <Card
       title="Gehe direkt zum Projekt!"
       inputPlaceholder="Projekt-ID"
      bind:inputValue={projectId}
      onArrowClick={() => {
        if (projectId) {
          goto(`/project-detail/${projectId}`);
        } else {
          alert('Please enter a Project ID.');
        }
      }}
     />

  <Card
    title="Erstelle ein neues Projekt!"
    onArrowClick={() => goto('/create-project')}
  />
</div>
</main>



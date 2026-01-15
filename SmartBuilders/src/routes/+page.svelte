<script lang="ts">
  import Card from '$lib/components/Card.svelte';
  import { goto } from '$app/navigation';

  let projectId = '';

  function handleRole(role: string) {
    if (role === 'kunde') {
      goto('/customer/projects');
    } else if (role === 'firma') {
      goto('/company/projects');
    } else if (role === 'mitarbeiter') {
      goto('/employee/projects');
     }
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 lg:pt-40">
  <h1 class="text-5xl font-bold mb-4">Willkommen bei SmartBuilders!</h1>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto p-10">

  <Card
    title="Organisiere jetzt deine Projekte!"
    buttons={[
      { label: 'Kunde', name: 'kunde' },
      { label: 'Firma', name: 'firma' },
      { label: 'Mitarbeiter', name: 'mitarbeiter' }
    ]}
    onButtonClick={handleRole}
          onArrowClick={() => {
        // The main action is handled by the buttons, but you could
        // define a default behavior for the arrow as well.
        // For example, go to a generic projects page if no role is selected.
        goto('/customer/projects');
      }}
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



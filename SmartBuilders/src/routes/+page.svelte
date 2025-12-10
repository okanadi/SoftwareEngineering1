<script lang="ts">
  import Card from '$lib/components/Card.svelte';

  function handleRole(role: String) {
    console.log("Chosen role:", role);
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
    onArrowClick={() => console.log("Weiter")}
  />

  <Card
    title="Gehe direkt zum Projekt!"
    inputPlaceholder="Projekt-ID"
    onArrowClick={() => console.log("Gehe zum Projekt")}
  />

  <Card
    title="Erstelle ein neues Projekt!"
    inputPlaceholder="Projekt-ID"
    onArrowClick={() => console.log("Neues Projekt")}
  />
</div>
</main>



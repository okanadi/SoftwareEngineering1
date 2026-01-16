<script lang="ts">
  import type { PageData } from './$types';
  import ProjectDetails from '$lib/components/ProjectDetails.svelte';
  import CreateProjectStepModal from '$lib/components/CreateProjectStepModal.svelte';
  import EditProjectModal from '$lib/components/EditProjectModal.svelte';
  import ProjectHistory from '$lib/components/ProjectHistory.svelte';
  import { goto } from '$app/navigation';

  // The `data` prop is passed from your +page.ts load function
  export let data: PageData;

  let showStepModal = false;
  let showEditModal = false;

  function openStepModal() {
    showStepModal = true;
  }

  function openEditModal() {
    showEditModal = true;
  }

  function goToSteps() {
    goto(`/project-detail/${data.project.id}/steps`);
  }

  function handleStepCreated(event: CustomEvent) {
    console.log('Project step created:', event.detail);
    showStepModal = false;
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 flex justify-center">
  <div class="container mx-auto p-4 md:p-8">
    {#if data.project}
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-white mb-6">Project Details</h1>
        <div class="flex gap-3 flex-wrap">
          <button
            on:click={openEditModal}
            class="px-6 py-2 bg-yellow-600 hover:bg-yellow-700 text-white rounded-lg font-semibold transition"
          >
            Edit
          </button>
          <button
            on:click={goToSteps}
            class="px-6 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg font-semibold transition"
          >
            View Steps
          </button>
          <button
            on:click={openStepModal}
            class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-semibold transition"
          >
            + Add Project Step
          </button>
        </div>
      </div>
      <ProjectDetails project={data.project} />
      <ProjectHistory projectId={data.project.id} />
      <EditProjectModal project={data.project} bind:isOpen={showEditModal} />
      <CreateProjectStepModal projectId={data.project.id} bind:isOpen={showStepModal} on:stepCreated={handleStepCreated} />
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

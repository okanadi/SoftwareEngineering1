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

  function handleDownload() {
    if (!data.project) return;
    const link = document.createElement('a');
    link.href = `http://13.49.46.226:8080/api/v1/projects/export/${data.project.id}`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }

  const timelineSteps = [
    { value: 'geplant', label: 'Planned' },
    { value: 'in_arbeit', label: 'in_arbeit' },
    { value: 'fertiggestellt', label: 'fertig' }
  ];

  $: currentStepIndex = data.project ? timelineSteps.findIndex(s => s.value === data.project.progress) : 0;
  $: progressPercentage = timelineSteps.length > 1 ? (Math.max(0, currentStepIndex) / (timelineSteps.length - 1)) * 100 : 0;
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
            on:click={handleDownload}
            class="px-6 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-lg font-semibold transition"
          >
            Download ZIP
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

      <!-- Example Timeline -->
      <div class="mb-8 bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-6">
        <h2 class="text-2xl font-bold text-white mb-8">Project Progress</h2>
        <div class="relative mx-4">
          <!-- Progress Line -->
          <div class="absolute top-1/2 left-0 w-full h-1 bg-white/10 -translate-y-1/2 rounded z-0"></div>
          <div 
            class="absolute top-1/2 left-0 h-1 bg-blue-500 -translate-y-1/2 rounded z-0 transition-all duration-500"
            style="width: {progressPercentage}%"
          ></div>

          <div class="relative flex justify-between w-full">
            {#each timelineSteps as step, index}
              <div class="flex flex-col items-center relative z-10">
                <div 
                  class="w-8 h-8 rounded-full flex items-center justify-center font-bold ring-4 ring-gray-800/50 transition-colors duration-300
                  {index < currentStepIndex ? 'bg-green-500 text-white' : 
                   index === currentStepIndex ? 'bg-blue-500 text-white' : 
                   'bg-gray-700 text-white/50'}"
                >
                  {#if index < currentStepIndex}
                    ✓
                  {:else}
                    {index + 1}
                  {/if}
                </div>
                <div class="absolute top-10 text-center w-32">
                  <p class="text-white font-semibold text-sm mt-1">{step.label}</p>
                  <p class="text-white/50 text-xs">
                    {index < currentStepIndex ? 'Completed' : 
                     index === currentStepIndex ? 'Current' : 'Pending'}
                  </p>
                </div>
              </div>
            {/each}
          </div>
        </div>
        <div class="h-16"></div>
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

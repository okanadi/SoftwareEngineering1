<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let projectId: string;
  export let isOpen = false;

  const dispatch = createEventDispatcher();

  let title = '';
  let description = '';
  let startDate = '';
  let endDate = '';
  let isLoading = false;
  let errorMessage = '';

  async function handleSubmit() {
    if (!title.trim()) {
      errorMessage = 'Title is required';
      return;
    }

    isLoading = true;
    errorMessage = '';

    try {
      const payload = {
        project_id: projectId,
        title,
        description,
        start_date: startDate ? new Date(startDate).toISOString() : '',
        end_date: endDate ? new Date(endDate).toISOString() : ''
      };

      const response = await fetch('http://13.49.46.226:8080/api/v1/project-steps/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        throw new Error(`Server responded with status ${response.status}`);
      }

      const result = await response.json();
      dispatch('stepCreated', result);
      closeModal();
    } catch (error: any) {
      errorMessage = error.message || 'Failed to create project step';
      console.error('Error creating project step:', error);
    } finally {
      isLoading = false;
    }
  }

  function closeModal() {
    isOpen = false;
    title = '';
    description = '';
    startDate = '';
    endDate = '';
    errorMessage = '';
  }
</script>

{#if isOpen}
  <div class="fixed inset-0 bg-black/60 flex items-center justify-center z-50">
    <div class="bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-8 max-w-md w-full mx-4">
      <h2 class="text-2xl font-bold text-white mb-6">Create New Project Step</h2>

      <form on:submit|preventDefault={handleSubmit} class="space-y-4">
        <!-- Title -->
        <div>
          <label for="title" class="block text-white/80 text-sm font-semibold mb-2">Title *</label>
          <input
            id="title"
            type="text"
            bind:value={title}
            placeholder="Enter step title"
            class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50"
            disabled={isLoading}
          />
        </div>

        <!-- Description -->
        <div>
          <label for="description" class="block text-white/80 text-sm font-semibold mb-2">Description</label>
          <textarea
            id="description"
            bind:value={description}
            placeholder="Enter step description"
            rows="3"
            class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50"
            disabled={isLoading}
          />
        </div>

        <!-- Start Date -->
        <div>
          <label for="startDate" class="block text-white/80 text-sm font-semibold mb-2">Start Date</label>
          <input
            id="startDate"
            type="date"
            bind:value={startDate}
            class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white focus:outline-none focus:border-white/50"
            disabled={isLoading}
          />
        </div>

        <!-- End Date -->
        <div>
          <label for="endDate" class="block text-white/80 text-sm font-semibold mb-2">End Date</label>
          <input
            id="endDate"
            type="date"
            bind:value={endDate}
            class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white focus:outline-none focus:border-white/50"
            disabled={isLoading}
          />
        </div>

        <!-- Error Message -->
        {#if errorMessage}
          <div class="bg-red-500/20 border border-red-500/50 text-red-300 p-3 rounded-lg text-sm">
            {errorMessage}
          </div>
        {/if}

        <!-- Buttons -->
        <div class="flex gap-3 pt-4">
          <button
            type="button"
            on:click={closeModal}
            disabled={isLoading}
            class="flex-1 px-4 py-2 bg-white/10 hover:bg-white/20 disabled:bg-white/5 text-white rounded-lg font-semibold transition"
          >
            Cancel
          </button>
          <button
            type="submit"
            disabled={isLoading}
            class="flex-1 px-4 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-600/50 text-white rounded-lg font-semibold transition"
          >
            {isLoading ? 'Creating...' : 'Create Step'}
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

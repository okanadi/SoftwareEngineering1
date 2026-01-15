<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';

  interface Project {
    id: string;
    manager_id: string;
    customer_lastname: string;
    address: string;
    description: string;
    start_date: string;
    end_date: string;
  }

  let projects: Project[] = [];
  let managerId = '';
  let isLoading = false;
  let errorMessage = '';
  let hasSearched = false;

  onMount(async () => {
    // Get manager ID from URL parameter
    const urlManagerId = $page.url.searchParams.get('managerId');
    if (urlManagerId) {
      managerId = urlManagerId;
      await fetchProjectsByManagerId(urlManagerId);
    }
  });

  async function fetchProjectsByManagerId(id: string) {
    if (!id.trim()) {
      errorMessage = 'Please enter a Manager ID';
      return;
    }

    isLoading = true;
    errorMessage = '';
    hasSearched = true;

    try {
      const response = await fetch(`http://13.49.46.226:8080/api/v1/projects/getByManagerID/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      });

      if (!response.ok) {
        throw new Error(`Server responded with status ${response.status}`);
      }

      const data = await response.json();
      projects = Array.isArray(data) ? data : [];
      
      if (projects.length === 0) {
        errorMessage = 'No projects found for this Manager ID';
      }
    } catch (error: any) {
      errorMessage = error.message || 'Failed to fetch projects';
      console.error('Error fetching projects:', error);
      projects = [];
    } finally {
      isLoading = false;
    }
  }

  function handleSearch() {
    fetchProjectsByManagerId(managerId);
  }

  function viewProject(projectId: string) {
    goto(`/projects/${projectId}`);
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32">
  <div class="max-w-6xl mx-auto">
    <h1 class="text-5xl font-bold mb-8 text-white">My Projects</h1>

    <!-- Search Bar -->
    <div class="mb-8 bg-white/10 p-6 rounded-lg">
      <label for="managerId" class="block text-white mb-2 font-semibold">Manager ID</label>
      <div class="flex gap-2">
        <input
          id="managerId"
          type="text"
          bind:value={managerId}
          placeholder="Enter your Manager ID (e.g., 474984e8-3f0f-4f21-9ef1-67f60633b5ec)"
          class="flex-1 px-4 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50"
          on:keydown={(e) => e.key === 'Enter' && handleSearch()}
        />
        <button
          on:click={handleSearch}
          disabled={isLoading}
          class="px-6 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-600/50 text-white rounded-lg font-semibold transition"
        >
          {isLoading ? 'Searching...' : 'Search'}
        </button>
      </div>
    </div>

    <!-- Error Message -->
    {#if errorMessage && hasSearched}
      <div class="mb-6 bg-red-500/20 border border-red-500/50 text-red-300 p-4 rounded-lg">
        {errorMessage}
      </div>
    {/if}

    <!-- Projects Grid -->
    {#if hasSearched && !isLoading && projects.length > 0}
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {#each projects as project (project.id)}
          <button
            on:click={() => viewProject(project.id)}
            class="text-left bg-white/10 border border-white/20 rounded-lg p-6 hover:bg-white/20 transition"
          >
            <h3 class="text-xl font-bold text-white mb-2">{project.customer_lastname}</h3>
            <p class="text-white/70 text-sm mb-3">
              <span class="font-semibold">Address:</span> {project.address}
            </p>
            <p class="text-white/70 text-sm mb-3">
              <span class="font-semibold">Description:</span> {project.description}
            </p>
            <p class="text-white/70 text-sm mb-3">
              <span class="font-semibold">Start:</span> {new Date(project.start_date).toLocaleDateString()}
            </p>
            <p class="text-white/70 text-sm mb-4">
              <span class="font-semibold">End:</span> {new Date(project.end_date).toLocaleDateString()}
            </p>
            <div class="text-xs text-white/50 truncate">
              ID: {project.id}
            </div>
          </button>
        {/each}
      </div>
    {/if}

    <!-- Empty State -->
    {#if hasSearched && !isLoading && projects.length === 0 && !errorMessage}
      <div class="text-center py-12">
        <p class="text-white/70 text-lg">No projects found</p>
      </div>
    {/if}

    <!-- Initial State -->
    {#if !hasSearched}
      <div class="text-center py-12">
        <p class="text-white/70 text-lg">Enter your Manager ID above to view your projects</p>
      </div>
    {/if}

    <!-- Loading State -->
    {#if isLoading}
      <div class="text-center py-12">
        <p class="text-white/70 text-lg">Loading projects...</p>
      </div>
    {/if}
  </div>
</main>

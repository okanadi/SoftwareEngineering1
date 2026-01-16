<script lang="ts">
  export let projectId: string;

  interface Photo {
    id: string;
    s3_key: string;
    file_type: string;
    url: string;
  }

  interface HistoryEntry {
    id: string;
    status: string;
    note: string;
    user_name: string;
    timestamp: string;
    photos?: Photo[];
  }

  interface ProjectStep {
    id: string;
    title: string;
    history: HistoryEntry[];
  }

  let allHistory: Array<HistoryEntry & { stepTitle: string }> = [];
  let isLoading = true;
  let errorMessage = '';

  async function loadHistory() {
    try {
      const response = await fetch(
        `http://13.49.46.226:8080/api/v1/project-steps/getAllByProjectID/${projectId}`
      );

      if (!response.ok) {
        throw new Error(`Failed to fetch steps: ${response.statusText}`);
      }

      const steps: ProjectStep[] = await response.json();
      
      // Flatten all history from all steps
      allHistory = [];
      steps.forEach(step => {
        if (step.history && Array.isArray(step.history)) {
          step.history.forEach(entry => {
            allHistory.push({
              ...entry,
              stepTitle: step.title
            });
          });
        }
      });

      // Sort by timestamp descending (newest first)
      allHistory.sort((a, b) => 
        new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime()
      );
    } catch (error: any) {
      errorMessage = error.message || 'Failed to load history';
      console.error('Error loading history:', error);
      allHistory = [];
    } finally {
      isLoading = false;
    }
  }

  // Load history when component mounts
  if (typeof window !== 'undefined') {
    loadHistory();
  }
</script>

<div class="mt-12">
  <h2 class="text-3xl font-bold text-white mb-6">Project History</h2>

  {#if isLoading}
    <div class="text-center py-8">
      <p class="text-white/70">Loading history...</p>
    </div>
  {:else if errorMessage}
    <div class="bg-red-500/20 border border-red-500/50 text-red-300 p-4 rounded-lg">
      {errorMessage}
    </div>
  {:else if allHistory.length === 0}
    <div class="text-center py-8 bg-white/5 border border-white/20 rounded-lg">
      <p class="text-white/70">No history entries yet</p>
    </div>
  {:else}
    <div class="space-y-4">
      {#each allHistory as entry (entry.id)}
        <div class="bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-6">
          <div class="flex justify-between items-start mb-3">
            <div>
              <span class="inline-block px-3 py-1 rounded-full text-sm font-semibold bg-blue-500/30 text-blue-300">
                {entry.status}
              </span>
              <p class="text-white/70 text-sm mt-2">Step: {entry.stepTitle}</p>
              <p class="text-white/70 text-sm">By: {entry.user_name}</p>
            </div>
            <p class="text-white/60 text-sm">
              {new Date(entry.timestamp).toLocaleString()}
            </p>
          </div>

          {#if entry.note}
            <p class="text-white mb-4">
              <span class="text-white/60">Note:</span> {entry.note}
            </p>
          {/if}

          {#if entry.photos && entry.photos.length > 0}
            <div class="mt-4">
              <p class="text-white/60 text-sm mb-2">Attachments ({entry.photos.length})</p>
              <div class="grid grid-cols-2 md:grid-cols-3 gap-2">
                {#each entry.photos as photo (photo.id)}
                  <a
                    href={photo.url}
                    target="_blank"
                    rel="noreferrer"
                    class="bg-white/5 border border-white/20 rounded p-2 text-center hover:bg-white/10 transition"
                  >
                    <p class="text-white/70 text-xs truncate">{photo.file_type}</p>
                    <p class="text-blue-300 text-xs hover:underline">View</p>
                  </a>
                {/each}
              </div>
            </div>
          {/if}
        </div>
      {/each}
    </div>
  {/if}
</div>

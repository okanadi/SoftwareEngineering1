<script lang="ts">
  import type { Project } from "$lib/types"; // Assuming you have a Project type defined, otherwise we can use a generic object.

  export let project: Project | Record<string, any>;

  // Helper function to format the key for display
  function formatLabel(key: string): string {
    return key
      .split('_')
      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  }

  // Helper function to format the value for display
  function formatValue(key: string, value: any): string {
    if (!value) return 'N/A';

    // Format dates into a more readable format
    if (key.includes('_date') || key.includes('_at')) {
      try {
        return new Intl.DateTimeFormat('en-US', {
          year: 'numeric',
          month: 'long',
          day: 'numeric',
          hour: '2-digit',
          minute: '2-digit'
        }).format(new Date(value));
      } catch (e) {
        return String(value); // Fallback for invalid dates
      }
    }
    
    // Replace underscores in progress status
    if (key === 'progress') {
        return String(value).replace('_', ' ');
    }

    return String(value);
  }
</script>

<div class="bg-white/5 backdrop-blur-md shadow-xl p-6 rounded-2xl border border-white/20">
  <h3 class="text-2xl font-bold mb-6 text-white">Project Information</h3>
  <div class="space-y-4">
    {#each Object.entries(project) as [key, value]}
      <div class="grid grid-cols-1 md:grid-cols-3 gap-2 border-b border-white/10 pb-3">
        <dt class="font-semibold text-white/70">{formatLabel(key)}</dt>
        <dd class="md:col-span-2 text-white break-words">{formatValue(key, value)}</dd>
      </div>
    {/each}
  </div>
</div>

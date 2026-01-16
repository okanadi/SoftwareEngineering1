<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let title: string;
  export let fields: Array<{
    name: string;
    label: string;
    type?: 'text' | 'email' | 'password' | 'date' | 'textarea';
    placeholder?: string;
    required?: boolean;
    disabled?: boolean;
  }> = [];
  export let values: Record<string, string> = {};
  export let isOpen = false;
  export let isLoading = false;
  export let submitButtonText = 'Submit';
  export let errorMessage = '';
  export let successMessage = '';
  export let cancelButtonText = 'Cancel';

  const dispatch = createEventDispatcher();

  function handleSubmit(e: Event) {
    e.preventDefault();
    dispatch('submit', values);
  }

  function handleCancel() {
    dispatch('cancel');
    errorMessage = '';
    successMessage = '';
  }
</script>

{#if isOpen}
  <div class="fixed inset-0 bg-black/60 flex items-center justify-center z-50 pt-20">
    <div class="bg-white/5 backdrop-blur-md shadow-xl border border-white/20 rounded-lg p-8 max-w-md w-full mx-4 max-h-[90vh] overflow-y-auto">
      <h2 class="text-2xl font-bold text-white mb-6">{title}</h2>

      <form on:submit={handleSubmit} class="space-y-4">
        <slot />
        {#each fields as field (field.name)}
          <div>
            <label for={field.name} class="block text-white/80 text-sm font-semibold mb-2">
              {field.label}
              {#if field.required}
                <span class="text-red-400">*</span>
              {/if}
            </label>
            {#if field.type === 'textarea'}
              <textarea
                id={field.name}
                bind:value={values[field.name]}
                placeholder={field.placeholder || ''}
                disabled={isLoading || field.disabled}
                class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 disabled:opacity-50"
              ></textarea>
            {:else}
              <input
                id={field.name}
                type={field.type || 'text'}
                bind:value={values[field.name]}
                placeholder={field.placeholder || ''}
                disabled={isLoading || field.disabled}
                class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 disabled:opacity-50"
              />
            {/if}
          </div>
        {/each}

        {#if errorMessage}
          <div class="p-3 rounded-lg bg-red-500/20 border border-red-500/50 text-red-200 text-sm">
            {errorMessage}
          </div>
        {/if}

        {#if successMessage}
          <div class="p-3 rounded-lg bg-green-500/20 border border-green-500/50 text-green-200 text-sm">
            {successMessage}
          </div>
        {/if}

        <div class="flex gap-3 pt-4">
          <button
            type="button"
            on:click={handleCancel}
            disabled={isLoading}
            class="flex-1 px-4 py-2 rounded-lg bg-white/10 border border-white/20 text-white font-medium hover:bg-white/20 disabled:opacity-50 transition-colors"
          >
            {cancelButtonText}
          </button>
          <button
            type="submit"
            disabled={isLoading}
            class="flex-1 px-4 py-2 rounded-lg bg-blue-600 text-white font-medium hover:bg-blue-700 disabled:opacity-50 transition-colors"
          >
            {#if isLoading}
              Loading...
            {:else}
              {submitButtonText}
            {/if}
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

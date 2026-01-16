<script lang="ts">
  import Modal from './Modal.svelte';
  import { createEventDispatcher } from 'svelte';

  interface ProjectStep {
    id: string;
    title: string;
    description: string;
    progress: string;
    start_date?: string;
    end_date?: string;
  }

  export let step: ProjectStep;
  export let isOpen = false;
  export let userId: string | undefined = undefined;

  const dispatch = createEventDispatcher();

  let isLoading = false;
  let errorMessage = '';
  let successMessage = '';
  
  let progress = '';
  let note = '';
  let files: FileList | undefined;

  $: if (step && isOpen) {
    progress = step.progress || '';
    note = '';
    files = undefined;
  }

  async function handleSubmit() {
    if (!progress.trim()) {
      errorMessage = 'Progress is required';
      return;
    }

    isLoading = true;
    errorMessage = '';
    successMessage = '';

    try {
      const formData = new FormData();
      formData.append('new_status', progress);

      // prefer explicit prop, otherwise fallback to localStorage
      const uid = userId || (typeof window !== 'undefined' ? localStorage.getItem('userId') || '' : '');
      if (uid) {
        formData.append('user_id', uid);
      }
      formData.append('note', note);

      // backend expects the file field name to be "photo"
      if (files && files.length > 0) {
        formData.append('photo', files[0]);
        formData.append('file_name', files[0].name);
        formData.append('file_content_type', files[0].type);
      }

      const response = await fetch(`http://13.49.46.226:8080/api/v1/project-steps/updateProgress/${step.id}`, {
        method: 'POST',
        body: formData
      });

      if (!response.ok) {
        // try to surface backend message
        const text = await response.text().catch(() => null);
        console.error('Update step failed', response.status, text);
        errorMessage = text || `Server responded with status ${response.status}`;
        return;
      }

      // success
      try {
        const body = await response.json().catch(() => null);
        console.debug('Update response body:', body);
      } catch (err) {
        // ignore
      }

      successMessage = 'Step progress updated successfully!';
      dispatch('stepUpdated', { id: step.id, progress });
      setTimeout(() => closeModal(), 800);
    } catch (error: any) {
      console.error('Error updating step:', error);
      errorMessage = error.message || 'Failed to update step progress';
    } finally {
      isLoading = false;
    }
  }

  function closeModal() {
    isOpen = false;
    errorMessage = '';
    successMessage = '';
  }

  function handleCancel() {
    closeModal();
  }
</script>

<Modal
  title="Edit Project Step"
  bind:isOpen
  bind:isLoading
  {errorMessage}
  {successMessage}
  submitButtonText="Update Step"
  on:submit={handleSubmit}
  on:cancel={handleCancel}
>
  <div>
    <label for="progress" class="block text-white/80 text-sm font-semibold mb-2">
      Progress <span class="text-red-400">*</span>
    </label>
    <input
      id="progress"
      type="text"
      bind:value={progress}
      placeholder="Enter progress percentage (e.g., 50%)"
      disabled={isLoading}
      class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 disabled:opacity-50"
    />
  </div>

  <div>
    <label for="note" class="block text-white/80 text-sm font-semibold mb-2">Note</label>
    <textarea
      id="note"
      bind:value={note}
      placeholder="Add a note about this update"
      disabled={isLoading}
      class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 disabled:opacity-50"
    ></textarea>
  </div>

  <div>
    <label for="file" class="block text-white/80 text-sm font-semibold mb-2">Attachment</label>
    <input
      id="file"
      type="file"
      bind:files
      disabled={isLoading}
      class="w-full px-3 py-2 rounded-lg bg-white/5 border border-white/20 text-white placeholder-white/50 focus:outline-none focus:border-white/50 disabled:opacity-50 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-500/30 file:text-blue-300 hover:file:bg-blue-500/40"
    />
  </div>
</Modal>

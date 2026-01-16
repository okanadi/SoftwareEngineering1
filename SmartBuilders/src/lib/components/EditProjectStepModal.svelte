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

  const dispatch = createEventDispatcher();

  let isLoading = false;
  let errorMessage = '';
  let successMessage = '';
  let values = {
    progress: ''
  };

  const fields: Array<{
    name: string;
    label: string;
    type?: 'text' | 'email' | 'password' | 'date' | 'textarea';
    placeholder?: string;
    required?: boolean;
    disabled?: boolean;
  }> = [
    {
      name: 'progress',
      label: 'Progress',
      type: 'text',
      placeholder: 'Enter progress percentage (e.g., 50%)',
      required: true
    }
  ];

  $: if (step && isOpen) {
    values = {
      progress: step.progress || ''
    };
  }

  async function handleSubmit(e: CustomEvent) {
    const formValues = e.detail;

    if (!formValues.progress.trim()) {
      errorMessage = 'Progress is required';
      return;
    }

    isLoading = true;
    errorMessage = '';
    successMessage = '';

    try {
      const formData = new FormData();
      formData.append('new_status', formValues.progress);
      formData.append('user_id', 'user-placeholder'); // TODO: Get from auth
      formData.append('note', '');

      const response = await fetch(`http://13.49.46.226:8080/api/v1/project-steps/updateProgress/${step.id}`, {
        method: 'POST',
        body: formData
      });

      if (!response.ok) {
        throw new Error(`Server responded with status ${response.status}`);
      }

      successMessage = 'Step progress updated successfully!';
      setTimeout(() => {
        closeModal();
      }, 1500);
      
      dispatch('stepUpdated', { id: step.id, progress: formValues.progress });
    } catch (error: any) {
      errorMessage = error.message || 'Failed to update step progress';
      console.error('Error updating step:', error);
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
  {fields}
  bind:values
  bind:isOpen
  bind:isLoading
  {errorMessage}
  {successMessage}
  submitButtonText="Update Progress"
  on:submit={handleSubmit}
  on:cancel={handleCancel}
/>

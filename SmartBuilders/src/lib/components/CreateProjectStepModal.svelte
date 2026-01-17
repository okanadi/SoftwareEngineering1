<script lang="ts">
  import Modal from './Modal.svelte';
  import { createEventDispatcher } from 'svelte';

  export let projectId: string;
  export let isOpen = false;

  const dispatch = createEventDispatcher();

  let isLoading = false;
  let errorMessage = '';
  let values = {
    title: '',
    description: '',
    startDate: '',
    endDate: ''
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
      name: 'title',
      label: 'Title',
      type: 'text',
      placeholder: 'Enter step title',
      required: true
    },
    {
      name: 'description',
      label: 'Description',
      type: 'textarea',
      placeholder: 'Enter step description'
    },
    {
      name: 'startDate',
      label: 'Start Date',
      type: 'date'
    },
    {
      name: 'endDate',
      label: 'End Date',
      type: 'date'
    }
  ];

  async function handleSubmit(e: CustomEvent) {
    const formValues = e.detail;

    if (!formValues.title.trim()) {
      errorMessage = 'Title is required';
      return;
    }

    isLoading = true;
    errorMessage = '';

    try {
      const payload = {
        project_id: projectId,
        title: formValues.title,
        description: formValues.description,
        start_date: formValues.startDate ? new Date(formValues.startDate).toISOString() : '',
        end_date: formValues.endDate ? new Date(formValues.endDate).toISOString() : ''
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
    values = {
      title: '',
      description: '',
      startDate: '',
      endDate: ''
    };
    errorMessage = '';
  }

  function handleCancel() {
    closeModal();
  }
</script>

<Modal
  title="Create New Project Step"
  {fields}
  bind:values
  bind:isOpen
  bind:isLoading
  bind:errorMessage
  submitButtonText="Create Step"
  on:submit={handleSubmit}
  on:cancel={handleCancel}
/>

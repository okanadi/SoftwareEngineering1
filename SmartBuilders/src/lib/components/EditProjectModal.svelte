<script lang="ts">
  import Modal from './Modal.svelte';
  import type { Project } from '$lib/types';
  import { createEventDispatcher } from 'svelte';

  export let project: Project;
  export let isOpen = false;

  const dispatch = createEventDispatcher();

  let isLoading = false;
  let errorMessage = '';
  let successMessage = '';
  let values = {
    customer_lastname: '',
    address: '',
    description: '',
    start_date: '',
    end_date: ''
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
      name: 'customer_lastname',
      label: 'Customer Lastname',
      type: 'text',
      placeholder: 'Enter customer lastname',
      required: true
    },
    {
      name: 'address',
      label: 'Address',
      type: 'text',
      placeholder: 'Enter address',
      required: true
    },
    {
      name: 'description',
      label: 'Description',
      type: 'textarea',
      placeholder: 'Enter project description'
    },
    {
      name: 'start_date',
      label: 'Start Date',
      type: 'date'
    },
    {
      name: 'end_date',
      label: 'End Date',
      type: 'date'
    }
  ];

  $: if (project && isOpen) {
    values = {
      customer_lastname: project.customer_lastname,
      address: project.address,
      description: project.description,
      start_date: project.start_date ? new Date(project.start_date).toISOString().split('T')[0] : '',
      end_date: project.end_date ? new Date(project.end_date).toISOString().split('T')[0] : ''
    };
  }

  async function handleSubmit(e: CustomEvent) {
    const formValues = e.detail;

    if (!formValues.customer_lastname.trim() || !formValues.address.trim()) {
      errorMessage = 'Customer lastname and address are required';
      return;
    }

    isLoading = true;
    errorMessage = '';
    successMessage = '';

    try {
      const payload = {
        customer_lastname: formValues.customer_lastname,
        address: formValues.address,
        description: formValues.description,
        start_date: formValues.start_date ? new Date(formValues.start_date).toISOString() : '',
        end_date: formValues.end_date ? new Date(formValues.end_date).toISOString() : ''
      };

      const response = await fetch(`http://13.49.46.226:8080/api/v1/projects/editProject/${project.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        throw new Error(`Server responded with status ${response.status}`);
      }

      successMessage = 'Project updated successfully!';
      setTimeout(() => {
        closeModal();
      }, 1500);
    } catch (error: any) {
      errorMessage = error.message || 'Failed to update project';
      console.error('Error updating project:', error);
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
  title="Edit Project"
  {fields}
  bind:values
  bind:isOpen
  bind:isLoading
  {errorMessage}
  {successMessage}
  submitButtonText="Update Project"
  on:submit={handleSubmit}
  on:cancel={handleCancel}
/>

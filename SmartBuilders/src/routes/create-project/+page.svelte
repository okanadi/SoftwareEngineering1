<script lang="ts">
  import DynamicCard from '$lib/components/DynamicCard.svelte';
  import { goto } from '$app/navigation';
  import type { Field } from '$lib/types';

  let projectData = {
    manager_id: '',
    customer_lastname: '',
    address: '',
    description: '',
    start_date: '',
    end_date: ''
  };

const fields: Field[] = [
  { name: 'manager_id', label: 'Manager ID', type: 'text' },
  { name: 'customer_lastname', label: 'Customer Lastname', type: 'text' },
  { name: 'address', label: 'Address', type: 'text' },
  { name: 'description', label: 'Description', type: 'textarea' },
  { name: 'start_date', label: 'Start Date', type: 'date' },
  { name: 'end_date', label: 'End Date', type: 'date' }
];

 const buttons = [{ label: 'create project', name: 'create-project' }];

  async function handleSubmit() {
    try {
      const response = await fetch('http://localhost:8080/api/v1/projects/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          ...projectData,
          start_date: new Date(projectData.start_date).toISOString(),
          end_date: new Date(projectData.end_date).toISOString()
        })
      });

      if (response.ok) {
        const result = await response.json();
        alert('Project created successfully!');
        goto(`/project-detail/${result.id}`);
      } else {
        const error = await response.text();
        alert(`Failed to create project: ${error}`);
      }
    } catch (error) {
      console.error('Error creating project:', error);
      alert('An error occurred while creating the project.');
    }
  }
</script>
<main class="container mx-auto p-4 pt-24 md:pt-32 flex justify-center">
  <div class="w-full max-w-2xl">
    <h1 class="text-5xl font-bold mb-8 text-center text-white">Erstelle ein neues Projekt</h1>
    <DynamicCard {fields} {buttons} onSubmit={handleSubmit}/>
  </div>
</main>
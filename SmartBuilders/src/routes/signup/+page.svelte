<!-- src/routes/signup/+page.svelte -->
<script lang="ts">
  import DynamicForm from '$lib/components/DynamicCard.svelte';
  import type { Field } from '$lib/types';

  const fields: Field[] = [
    { label: 'Anrede', name: 'title', type: 'select', options: ['Herr', 'Frau'], placeholder: 'Anrede wählen' },
    { label: 'Name', name: 'name', type: 'text' },
    { label: 'Rolle', name: 'role', type: 'select', options: ['Kunde', 'Firma', 'Mitarbeiter'], placeholder: 'Rolle wählen' },
    { label: 'Firma', name: 'company', type: 'text', conditional: values => values.role === 'Firma' },
    { label: 'Adresse', name: 'address', type: 'text' },
    { label: 'Email', name: 'email', type: 'email' },
    { label: 'Password', name: 'password', type: 'password' }
  ];

  function handleSignIn(values: Record<string, any>) {
    console.log('SignIn Werte:', values);
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 lg:pt-40 ">
  <h1 class="text-5xl font-bold mb-4">Sign up</h1>  
  <div class="flex justify-center">
    <DynamicForm {fields} buttons={[{ label: 'Sign In', name: 'signin' }]} onSubmit={handleSignIn} />
  </div>
  
</main>


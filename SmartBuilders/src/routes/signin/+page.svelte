<!-- src/routes/signin/+page.svelte -->
<script lang="ts">
  import DynamicForm from '$lib/components/DynamicCard.svelte';
  import type { Field } from '$lib/types';

  const fields: Field[] = [
    { label: 'Email', name: 'email', type: 'email', placeholder: 'Deine Email' },
    { label: 'Password', name: 'password', type: 'password', placeholder: 'Passwort' }
  ];

  function handleSignIn(values: Record<string, any>) {
    console.log('SignIn Werte:', values);
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 lg:pt-40 ">
  <h1 class="text-5xl font-bold mb-4">Sign In</h1>  
  <div class="flex justify-center">
    <DynamicForm {fields} buttons={[{ label: 'Sign In', name: 'signin' }]} onSubmit={handleSignIn} />
  </div>
  
</main>


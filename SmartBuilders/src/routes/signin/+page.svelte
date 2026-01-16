<script lang="ts">
  import DynamicCard from '$lib/components/DynamicCard.svelte';
  import type { Field } from '$lib/types';
  import { goto } from '$app/navigation';

  let errorMessage = '';
  let successMessage = '';
  let isLoading = false;

  const fields: Field[] = [
    {
      label: 'Email',
      name: 'email',
      type: 'email',
      placeholder: 'max.mustermann@mail.com'
    },
    {
      label: 'Passwort',
      name: 'password',
      type: 'password',
      placeholder: '••••••••'
    }
  ];

  const buttons = [{ label: 'Sign In', name: 'signin' }];

  async function handleSubmit(values: Record<string, any>) {
    isLoading = true;
    errorMessage = '';
    successMessage = '';

    const { email, password } = values;

    if (!email || !password) {
      errorMessage = 'Please fill out all fields.';
      isLoading = false;
      return;
    }

    try {
      const response = await fetch('http://13.49.46.226:8080/api/v1/users/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || `Server responded with status ${response.status}`);
      }
      
      const responseData = await response.json();
      
      successMessage = 'Signin successful! Redirecting...';
      console.log('User signed in:', responseData);

      setTimeout(() => {
        goto('/projects');
      }, 2000);

    } catch (error: any) {
      errorMessage = error.message || 'An unknown error occurred during signin.';
      console.error('Signin failed:', error);
    } finally {
      isLoading = false;
    }
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 flex justify-center">
  <div class="w-full max-w-2xl">
    <h1 class="text-5xl font-bold mb-8 text-center text-white">Sign In</h1>
    
    <DynamicCard {fields} {buttons} onSubmit={handleSubmit} />

    <!-- Messages -->
    {#if errorMessage}
      <div class="bg-red-500/20 text-red-300 p-3 rounded-lg mt-4 text-sm text-center">
        {errorMessage}
      </div>
    {/if}
    {#if successMessage}
      <div class="bg-green-500/20 text-green-300 p-3 rounded-lg mt-4 text-sm text-center">
        {successMessage}
      </div>
    {/if}
    {#if isLoading}
        <div class="text-white/80 p-3 rounded-lg mt-4 text-sm text-center">
            Signing in...
        </div>
    {/if}

     <div class="text-center mt-6">
      <a href="/signup" class="text-sm text-white/70 hover:text-white">Don't have an account? Sign up</a>
    </div>
  </div>
</main>

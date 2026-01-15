<script lang="ts">
  import DynamicCard from '$lib/components/DynamicCard.svelte';
  import type { Field } from '$lib/types';
  import { goto } from '$app/navigation';

  let errorMessage = '';
  let successMessage = '';
  let isLoading = false;

  const fields: Field[] = [
    {
      label: 'Anrede',
      name: 'salutation',
      type: 'select',
      options: ['Herr', 'Frau', 'Divers'],
      placeholder: 'Anrede auswählen'
    },
    {
      label: 'Vorname',
      name: 'firstName',
      type: 'text',
      placeholder: 'Max'
    },
    {
      label: 'Nachname',
      name: 'lastName',
      type: 'text',
      placeholder: 'Mustermann'
    },
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
    },
    {
      label: 'Passwort bestätigen',
      name: 'confirmPassword',
      type: 'password',
      placeholder: '••••••••'
    },
    {
      label: 'Rolle',
      name: 'role',
      type: 'select',
      options: ['admin', 'innendienst', 'handwerker'],
      placeholder: 'Rolle auswählen'
    }
  ];

  const buttons = [{ label: 'Sign Up', name: 'signup' }];

  async function handleSubmit(values: Record<string, any>) {
    isLoading = true;
    errorMessage = '';
    successMessage = '';

    const { firstName, lastName, email, password, confirmPassword, role } = values;

    if (!firstName || !lastName || !email || !password || !role) {
      errorMessage = 'Please fill out all fields.';
      isLoading = false;
      return;
    }

    if (password !== confirmPassword) {
      errorMessage = 'Passwords do not match.';
      isLoading = false;
      return;
    }

    const name = `${firstName} ${lastName}`;

    try {
      const response = await fetch('http://13.49.46.226:8080/api/v1/users/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, email, password, role })
      });
      
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(responseData.error || `Server responded with status ${response.status}`);
      }
      
      successMessage = 'Signup successful! Redirecting to signin...';
      console.log('User created:', responseData);

      setTimeout(() => {
        goto('/signin');
      }, 2000);

    } catch (error: any) {
      errorMessage = error.message || 'An unknown error occurred during signup.';
      console.error('Signup failed:', error);
    } finally {
      isLoading = false;
    }
  }
</script>

<main class="container mx-auto p-4 pt-24 md:pt-32 flex justify-center">
  <div class="w-full max-w-2xl">
    <h1 class="text-5xl font-bold mb-8 text-center text-white">Create Account</h1>
    
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
            Signing up...
        </div>
    {/if}

     <div class="text-center mt-6">
      <a href="/signin" class="text-sm text-white/70 hover:text-white">Already have an account? Sign in</a>
    </div>
  </div>
</main>

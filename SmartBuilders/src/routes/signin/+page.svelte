<!-- src/routes/signin/+page.svelte -->
<script>
  let email = '';
  let password = '';

  function handleLogin() {
    console.log('Logging in', { email, password });
    // API-Call hier
  }
</script>

<div class="max-w-md mx-auto p-6 bg-white rounded shadow">
  <h2 class="text-2xl font-bold mb-4">Sign In</h2>
  <form on:submit|preventDefault={handleLogin} class="space-y-4">
    <input type="email" bind:value={email} placeholder="Email" class="input" />
    <input type="password" bind:value={password} placeholder="Password" class="input" />
    <button type="submit" class="btn-primary w-full">Login</button>
  </form>
</div>

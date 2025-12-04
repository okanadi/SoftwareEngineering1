<script>
  import { onMount } from 'svelte';
  import logo from '$lib/assets/smartbuilders-logo.svg';
  let menuOpen = false;

  function toggleMenu() {
    menuOpen = !menuOpen;
  }

  function handleSignIn() {
    console.log('Sign In clicked');
  }

  function handleSignUp() {
    console.log('Sign Up clicked');
  }
</script>

<nav class="fixed top-0 left-0 w-full z-50 h-22 md:h-30 lg:h-38 flex items-center border-white/30 bg-white/5 backdrop-blur-md 
         shadow-xl">
  <div class="container mx-auto flex items-center justify-between p-4">
    <!-- Logo -->
    <div class="text-xl font-bold flex items-center">
      <img src={logo} alt="Logo" class="h-20 w-20 md:h-28 md:w-28 lg:h-32 lg:w-32 mr-20" />
    </div>

    <!-- Desktop Buttons -->
    <div class="hidden md:flex space-x-4">
      <button on:click={handleSignIn} class="bg-white/30 hover:bg-white/40 transition px-4 py-2 rounded">
        Sign In
      </button>
      <button on:click={handleSignUp} class="bg-white/10 hover:bg-white/20 transition px-4 py-2 rounded">
        Sign Up
      </button>
    </div>
  

    <!-- Hamburger Icon -->
    <div class="md:hidden">
      <button on:click={toggleMenu} class="focus:outline-none">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          {#if menuOpen}
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          {:else}
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          {/if}
        </svg>
      </button>
    </div>
  </div>

  <!-- Mobile Menu -->
  {#if menuOpen}
    <div class="md:hidden bg-white shadow-md">
      <button on:click={handleSignIn} class="block w-full text-left px-4 py-2 hover:bg-blue-100">Sign In</button>
      <button on:click={handleSignUp} class="block w-full text-left px-4 py-2 hover:bg-green-100">Sign Up</button>
    </div>
  {/if}
</nav>

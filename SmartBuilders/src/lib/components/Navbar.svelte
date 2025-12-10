<script lang="ts">
  import { onMount } from 'svelte';
  import logo from '$lib/assets/smartbuilders-logo.svg';
  import { goto } from '$app/navigation';
  
  let menuOpen = false;

  function toggleMenu() {
    menuOpen = !menuOpen;
  }

  function handleSignIn() {
    goto('/signin');
  }

  function handleSignUp() {
    goto('/signup');
  }
  function handleHome() {
    goto('/');
  }
</script>

<nav class="fixed top-0 left-0 w-full z-50 h-20 md:h-24 lg:h-28 flex items-center p-10 rounded-3xl border border-white/30 bg-white/5 backdrop-blur-md">
  <div class="container mx-auto flex items-center justify-between p-4">
    <!-- Logo -->
    <div class="text-xl font-bold flex items-center">
      <button on:click={handleHome} class="p-0 bg-transparent border-0">
        <img src={logo} alt="Logo" class="h-16 w-16 md:h-19 md:w-19 lg:h-20 lg:w-20 mr-10" />
      </button>
    </div>

    <!-- Desktop Buttons -->
    <div class="hidden md:flex space-x-4">
      <button on:click={handleSignIn} class="bg-gray-700 hover:bg-gray-500 transition px-4 py-1 rounded">
        Sign In
      </button>
      <button on:click={handleSignUp} class="bg-gray-900 hover:bg-gray-700 transition px-4 py-1 rounded">
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

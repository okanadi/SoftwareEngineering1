<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/components/Navbar.svelte';
	
	let { children } = $props();
</script>

<svelte:head>
</svelte:head>

<Navbar />

<main class="mt-16"> <!-- mt-16, damit Navbar nicht überlappt -->

</main>

{@render children()}

<style lang="postcss">
  @reference "tailwindcss";
  :global(html) {
	background-color: #212121;
    color: #f2f2f2;
  }
</style>

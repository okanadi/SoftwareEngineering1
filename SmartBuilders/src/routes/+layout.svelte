<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';
	
	let { children } = $props();
</script>

<svelte:head>
</svelte:head>

<div class="relative min-h-screen bg-cover bg-center bg-no-repeat"
     style="background-image: url('/assets/SmartBuildersBackgroundGradient.jpg');">

	<!-- overlay -->
	<div class="absolute inset-0 bg-black/70"></div>
	<!-- content -->
	<div class="relative z-10">
		<div class="relative z-20">
			<Navbar/>
		</div>
		
		<main class="relative z-10">
			{@render children()}
		</main>

		<div>
			<Footer/>
		</div>
		
	</div>
</div>

<style lang="postcss">
  @reference "tailwindcss";
  :global(html) {
	background-color: #212121;
    color: #fff;
  }
</style>

<script lang="ts">
  export let title = "";
  export let description = "";
  export let buttons: { label: string; name: string }[] = [];
  export let inputPlaceholder: string | null = null;
  export let showArrow: boolean = true;

  export let onButtonClick = (name: string) => {};
  export let onArrowClick = () => {};

  export let inputValue = "";
</script>

<div
  class="p-10 rounded-3xl border border-white/30 bg-white/5 backdrop-blur-md 
         shadow-xl h-full flex flex-col justify-between text-white"
>
  <div>
    <h2 class="text-4xl font-bold leading-tight mb-6">{title}</h2>

    {#if buttons.length > 0}
      <div class="flex gap-3 flex-wrap mb-8">
        {#each buttons as btn}
          <button
            class="px-5 py-2 rounded-lg bg-white/20 hover:bg-white/30 transition"
            on:click={() => onButtonClick(btn.name)}
          >
            {btn.label}
          </button>
        {/each}
      </div>
    {/if}

    {#if inputPlaceholder}
      <input
        bind:value={inputValue}
        placeholder={inputPlaceholder}
        class="w-full bg-white/10 border border-white/30 rounded-lg px-4 py-2 text-white
               focus:outline-none focus:ring-2 focus:ring-white/40"
      />
    {/if}
  </div>

  {#if showArrow}
    <button
      on:click={onArrowClick}
      class="mt-10 text-4xl self-start opacity-80 hover:opacity-100 transition"
    >
      →
    </button>
  {/if}
</div>

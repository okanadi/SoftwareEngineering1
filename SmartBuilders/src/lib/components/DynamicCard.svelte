<script lang="ts">
  import type { Field } from '$lib/types';

  export let title: string = ''; 
  export let fields: Field[] = [];
  export let buttons: { label: string; name: string }[] = [];
  export let submitLabel: string = 'Submit'; 
  export let onSubmit: (values: Record<string, any>) => void = () => {};
  export let data: Record<string, any> = {}; 

  let values: Record<string, any> = {};
</script>


<form
  class="p-8 flex flex-col gap-4 bg-white/5 rounded-3xl border border-white/30 backdrop-blur-md shadow-xl text-white"
  on:submit|preventDefault={() => onSubmit(values)}
>
  <div class="grid gap-4 md:grid-cols-2">
    {#each fields as field}
      {#if !field.conditional || field.conditional(values)}
        {#if field.type === 'select'}
          <label class="flex flex-col">
            {field.label}
            <select
              bind:value={values[field.name]}
              class="mt-1 bg-white/10 border border-white/30 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-2 focus:ring-white/40"
            >
              <option value="" disabled selected hidden>{field.placeholder ?? field.label}</option>
              {#each field.options as option}
                <option value={option}>{option}</option>
              {/each}
            </select>
          </label>
        {:else}
          <label class="flex flex-col">
            {field.label}
            <input
              type={field.type ?? 'text'}
              bind:value={values[field.name]}
              placeholder={field.placeholder ?? field.label}
              class="mt-1 bg-white/10 border border-white/30 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-2 focus:ring-white/40"
            />
          </label>
        {/if}
      {/if}
    {/each}

    {#if buttons.length > 0}
      <div class="flex gap-3 mt-4 flex-wrap">
        {#each buttons as btn}
          <button
            type="submit"
            class="px-5 py-2 rounded-lg bg-white/20 hover:bg-white/30 transition"
          >
            {btn.label}
          </button>
        {/each}
      </div>
    {/if}
  </div>
</form>

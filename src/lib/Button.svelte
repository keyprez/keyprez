<script lang="ts">
  import { Loader } from '$lib';

  export let text: string;
  export let type: string;
  export let loading = false;
  export let onClick;
  export let disabled = false;
  export let extraClassNames: string;
</script>

<button
  class="relative flex justify-center w-full px-20 py-6 bg-black {extraClassNames} {disabled
    ? 'opacity-50'
    : 'hover:bg-teal-900 dark:hover:bg-teal-900'} rounded-lg md:max-w-xs dark:bg-teal-800"
  {type}
  on:click={onClick}
  {disabled}
>
  {#if loading}
    <Loader />
  {:else}
    <span>{text}</span>
  {/if}
</button>

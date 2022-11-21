<script>
  import { Loader } from '$lib';

  export let text;
  export let type;
  export let loading = false;
  export let onClick;
  export let disabled = false;
</script>

<button
  class="relative flex justify-center w-full px-20 py-6 bg-black {disabled
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

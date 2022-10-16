<script>
  import { Loader } from '$lib';

  export let text;
  export let type;
  export let loading = false;
  export let onClick;
</script>

<button
  class="relative w-full md:max-w-xs flex justify-center bg-black rounded-lg py-6 px-20 hover:bg-teal-900"
  {type}
  on:click={onClick}
>
  {#if loading}
    <Loader />
  {:else}
    <span>{text}</span>
  {/if}
</button>

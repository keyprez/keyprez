<script>
  import { Loader } from '$lib';

  export let text;
  export let type;
  export let loading = false;
  export let onClick;
</script>

<button {type} on:click={onClick}>
  {#if loading}
    <Loader />
  {:else}
    <span>{text}</span>
  {/if}
</button>

<style>
  button {
    width: 100%;
    display: flex;
    justify-content: center;
  }

  @media (min-width: 600px) {
    button {
      max-width: 200px;
    }
  }
</style>

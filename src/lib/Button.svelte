<script>
  export let text;
  export let type;
  export let onClick;

  let loading = false;

  const resolver = async (promise, loading) => {
    loading(true);
    await promise;
    loading(false);
  };
</script>

{#if typeof onClick !== 'undefined'}
  <button
    type={type || 'button'}
    class={loading ? 'button-loading' : ''}
    on:click|preventDefault={() => resolver(onClick(), (status) => (loading = status))}
  >
    <span>{text}</span>
  </button>
{:else}
  <button type={type || 'button'} class={loading ? 'button-loading' : ''}>
    <span>{text}</span>
  </button>
{/if}

<style type="text/scss">
  @import 'src/variables';

  .button-loading span {
    opacity: 0;
    transition: all 0.2s;
    visibility: hidden;
  }

  .button-loading::after {
    animation: spinner 1s ease infinite;
    border: 0.3rem solid transparent;
    border-radius: 50%;
    border-top-color: $color-secondary-light;
    bottom: 0;
    content: '';
    height: 1rem;
    left: 0;
    margin: auto;
    position: absolute;
    right: 0;
    top: 0;
    width: 1rem;
  }

  @keyframes spinner {
    from {
      transform: rotate(0turn);
    }

    to {
      transform: rotate(1turn);
    }
  }
</style>

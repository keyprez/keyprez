<script>
  export let text;
  export let type;
  export let loading = false;
  export let onClick;
</script>

<button {type} class:loading on:click={onClick}>
  <span>{text}</span>
</button>

<style type="text/scss">
  @import 'src/variables';

  .loading span {
    opacity: 0;
    transition: all 0.2s;
    visibility: hidden;
  }

  .loading::after {
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

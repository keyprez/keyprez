<script context="module">
  export const prerender = true;
</script>

<svelte:head>
  <title>Corgi</title>
</svelte:head>

<div>
  <h1>CORGI</h1>
  <p>
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi eget bibendum elit. Nullam blandit gravida vehicula.
    Suspendisse pretium quam lectus, sed elementum dolor iaculis ac. Pellentesque vel tellus sed ante euismod maximus.
    Suspendisse pharetra ornare ex eget pellentesque. Suspendisse tincidunt augue urna, quis facilisis quam bibendum
    eget. Nulla at eros elementum, posuere mi a, suscipit dolor. Nullam lobortis vulputate dolor, quis laoreet neque
    finibus at. Maecenas aliquam scelerisque leo, et bibendum nulla consequat vitae. Suspendisse lacus sapien, auctor
    sed vestibulum at, maximus quis eros. Etiam rhoncus elementum tempor. Suspendisse porttitor nisl dolor, non euismod
    risus malesuada sit amet. Integer id tempus risus.
  </p>
</div>

<style type="text/scss">
  @import 'src/variables';

  h1 {
    color: $color-primary-dark;
  }

  p {
    color: $color-secondary-dark;
  }
</style>

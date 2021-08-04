<script context="module">
  import { fade } from 'svelte/transition';

  export const prerender = true;
</script>

<svelte:head>
  <title>Contact</title>
</svelte:head>

<div in:fade>
  <h1>Contact us</h1>
  <p>
    Curabitur quis facilisis sapien. Cras luctus elit in ante tincidunt aliquet. Praesent interdum euismod felis eget
    condimentum. Nam cursus pulvinar lacus at ultricies. Aliquam tempor consequat est, eu iaculis ipsum imperdiet non.
    Duis placerat aliquam justo, id sodales risus tempus eget. Quisque ornare, nunc a laoreet sagittis, diam sapien
    varius odio, nec posuere ligula neque eu nunc. Nulla hendrerit luctus turpis, sed egestas enim iaculis sit amet. Nam
    nec nunc quis nibh rutrum condimentum. Vestibulum risus risus, tempor in nisl sit amet, luctus gravida mauris. Donec
    vestibulum est quis blandit tincidunt. Curabitur lorem ex, porta id pretium eu, scelerisque sit amet lacus. Sed est
    ante, aliquet ut semper in, sagittis pretium leo. Nam gravida pharetra ex quis gravida. Suspendisse arcu lacus,
    sodales in est sit amet, faucibus euismod enim.
  </p>
  <p>
    Integer nec odio nec mi ornare auctor in ut libero. Pellentesque vel finibus sem. Donec vel lorem lorem. Fusce eu
    purus lobortis risus sagittis rutrum. Mauris gravida efficitur posuere. Praesent a libero ex. Nunc non sodales
    mauris. Ut commodo lectus sed est vestibulum, eu pulvinar mauris venenatis. Aenean tempor sapien blandit ex egestas
    rhoncus. Suspendisse sagittis dui nisl, non elementum neque eleifend in. Fusce cursus neque sed arcu fringilla
    pharetra. Fusce lacinia, augue id ultrices dictum, nibh nibh placerat mi, et tristique eros justo ut nisi. In tempus
    pretium hendrerit. Duis ac malesuada nulla. Nullam pellentesque purus id eros sagittis, egestas mollis sapien
    dignissim.
  </p>
  <p>
    Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl leo,
    sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac. Proin
    eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam finibus
    leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
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

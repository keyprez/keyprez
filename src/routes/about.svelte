<script context="module">
  import { fade } from 'svelte/transition';
  import { browser, dev } from '$app/env';

  // we don't need any JS on this page, though we'll load
  // it in dev so that we get hot module replacement...
  export const hydrate = dev;

  // ...but if the client-side router is already loaded
  // (i.e. we came here from elsewhere in the app), use it
  export const router = browser;

  // since there's no dynamic data here, we can prerender
  // it so that it gets served as a static asset in prod
  export const prerender = true;
</script>

<script>
  import SvelteSeo from 'svelte-seo';
</script>

<SvelteSeo title="Keyprez - About us" />

<div in:fade>
  <h1>About us</h1>
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

<style lang="scss">
  @import 'src/variables';

  h1 {
    color: $color-primary-dark;
  }

  p {
    color: $color-secondary-dark;
  }
</style>

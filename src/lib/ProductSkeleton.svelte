<script lang="ts">
  import { fade } from 'svelte/transition';
  import { Loader } from '$lib';
</script>

<div class="text-center w-full flex-[1_0_100%] md:flex-[1_0_40%] lg:flex-[1_0_20%]" transition:fade={{ duration: 10 }}>
  <div class="flex items-center justify-center">
    <img class="object-cover opacity-10 w-full rounded-lg " src={`corgi.jpg`} alt="Skeleton" />
    <span class="absolute opacity-100">
      <Loader />
    </span>
  </div>
  <p class="text-black dark:text-white opacity-10">⌨️</p>
  <h2 class="dark:text-teal-500 opacity-10">$$$</h2>
</div>

<script>
  export let product;
  $: ({ name, description, price, slug } = product);
</script>

<div class="text-center w-full flex-[1_0_100%] md:flex-[1_0_40%] lg:flex-[1_0_20%]">
  <a class="flex flex-col justify-center items-center group" data-sveltekit-prefetch href={`/product/${slug}`}>
    <img
      class="object-cover opacity-100 w-full rounded-lg ease-linear duration-300 group-hover:opacity-20"
      src={`${name.toLowerCase()}.jpg`}
      alt={name}
    />
    <span class="opacity-0 absolute ease-linear duration-300 group-hover:opacity-100"
      ><img class="max-w-[15rem] min-w-[13rem] max-h-[10rem]" src={`/${name.toLowerCase()}.svg`} alt={name} /></span
    >
  </a>
  <p class="text-black">{description.toUpperCase()}</p>
  <h2>&#36;<span>{price}</span></h2>
</div>

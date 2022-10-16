<script>
  import { fade } from 'svelte/transition';
  import SvelteSeo from 'svelte-seo';
  import { Button } from '$lib';

  const handleSubmit = (e) => {
    const form = new FormData(e.target);
    const formData = Object.fromEntries(form);

    // Send formData to keyprez API to handle email sendout
  };
</script>

<SvelteSeo title="Keyprez - Contact" />

<div class="flex flex-col md:flex-row gap-8" in:fade>
  <div class="flex-[40%]">
    <h1 class="text-center text-2xl mb-3 text-black">Contact us</h1>
    <form class="flex flex-col justify-center items-center gap-2" method="POST" on:submit|preventDefault={handleSubmit}>
      <textarea
        class="rounded-lg p-4 resize-none w-full"
        placeholder="Type your message..."
        name="message"
        rows="15"
        maxlength="1000"
        required
      />
      <div class="flex flex-col xl:flex-row gap-2 w-full">
        <input class="w-full rounded-lg py-6 px-4" placeholder="Name" type="text" name="name" />
        <input class="w-full rounded-lg py-6 px-4" placeholder="Email" type="email" name="email" required />
      </div>
      <div class="w-full md:w-auto">
        <Button type="submit" text="SUBMIT" />
      </div>
    </form>
  </div>
  <div class="flex-[60%]">
    <h1 class="text-center text-2xl mb-3">FAQ</h1>
    <ul>
      <li class="text-black text-2xl mb-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit.?</li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
      <li class="text-black text-2xl mb-4">
        Morbi eget bibendum elit. Nullam blandit gravida vehicula. Suspendisse pretium quam lectus, sed elementum dolor
        iaculis ac?
      </li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
      <li class="text-black text-2xl mb-4">
        Pellentesque vel tellus sed ante euismod maximus. Suspendisse pharetra ornare ex eget pellentesque?
      </li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
    </ul>
  </div>
</div>

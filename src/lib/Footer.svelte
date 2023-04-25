<script>
  import { onMount } from 'svelte';
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import { Button } from '$lib';

  import { products, productError } from '../store';
  import { fetchData, fetchProducts } from '../utils';

  let displayedEmail = '';
  let loading = false;
  let showError = false;
  let showSuccess = false;
  let hasSubscription = false;
  let inputRef;
  let focus = () => (!$form.email || $errors.email) && inputRef.focus();

  onMount(async () => {
    const { error, data } = await fetchProducts();
    if (error) $productError = error;
    if (data) $products = data;
  });

  const onSubmit = async ({ email }) => {
    loading = true;

    const res = await fetchData({ email }, '/newsletters');

    loading = false;

    if (res.status === 200 || res.status === 201) {
      showSuccess = true;
      displayedEmail = email;
      hasSubscription = res.status === 200;
      return setTimeout(() => (showSuccess = false), 5000);
    } else {
      showError = true;
      return setTimeout(() => (showError = false), 5000);
    }
  };

  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues: { email: '' },
    validationSchema: yup.object().shape({
      email: yup.string().email().required(),
    }),
    onSubmit,
  });
</script>

<footer class="flex flex-col gap-8 justify-center px-4 md:px-12 py-12 md:py-40">
  {#if showError}
    <h2 class="text-center text-2xl">There was an error ❌ Please try again later</h2>
  {:else if hasSubscription}
    <h2 class="text-center text-2xl ">
      <strong class="text-black dark:text-teal-800">{displayedEmail}</strong> has been already subscribed to our newsletter
      😊
    </h2>
  {:else if showSuccess}
    <h2 class="text-center text-2xl">Thank you for subscribing to our newsletter 🎉</h2>
  {:else}
    <h2 class="text-center text-2xl">
      <strong class="text-black dark:text-teal-800">SUBSCRIBE</strong> to latest news and updates
    </h2>
  {/if}
  <form class="flex flex-col md:flex-row gap-2 justify-center" on:submit={handleSubmit}>
    <div class="relative flex w-full max-w-full md:max-w-sm">
      <input
        bind:this={inputRef}
        class="py-6 px-4 text-black dark:text-white border-2 dark:bg-transparent rounded-lg shadow-xl w-full {$errors.email
          ? 'border-red-500'
          : 'border-transparent dark:border-teal-800'}"
        type="text"
        name="email"
        placeholder="Type your email"
        bind:value={$form.email}
        on:keyup={handleChange}
        on:blur={handleChange}
      />
      <small class="absolute right-2 bottom-1 text-red-500 {$errors.email ? 'block' : 'hidden'}">{$errors.email}</small>
    </div>
    <Button type="submit" text="SUBSCRIBE" {loading} onClick={focus} />
  </form>
  <div class="flex flex-row gap-4 justify-center">
    <a href="https://github.com/keyprez/keyprez" target="blank"
      ><img class="w-6 hover:scale-125" src="/github_logo.svg" alt="GitHub" /></a
    >
    <a href="https://github.com/keyprez/keyprez" target="blank"
      ><img class="w-6 hover:scale-125" src="/instagram_logo.svg" alt="Instagram" /></a
    >
    <a href="https://github.com/keyprez/keyprez" target="blank"
      ><img class="w-6 hover:scale-125" src="/reddit_logo.svg" alt="Reddit" /></a
    >
    <a href="https://github.com/keyprez/keyprez" target="blank"
      ><img class="w-6 hover:scale-125" src="/meetup_logo.svg" alt="Meetup" /></a
    >
  </div>
</footer>

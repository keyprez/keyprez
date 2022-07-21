<script>
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import SvelteSeo from 'svelte-seo';
  import { page } from '$app/stores';

  import { Button, Checkbox } from '$lib';
  import { capitalize, createSessionId, getCustomerStripeId, getProductBySlug, redirectToCheckout } from '/src/utils';

  const productSlug = $page.url.pathname.replace('/product/', '');

  let loading = false;
  let error = false;
  let inputRef;
  let focus = () => inputRef.focus();

  const onSubmit = async ({ email }) => {
    try {
      loading = true;
      const { PriceID } = await getProductBySlug(productSlug);
      if (!PriceID) throw new Error(`Could not get product by slug`);

      const customerStripeId = await getCustomerStripeId(email);
      if (!customerStripeId) throw new Error(`Could not get customerStripeId`);

      const sessionId = await createSessionId(PriceID, customerStripeId);
      if (!sessionId) throw new Error(`Could not get sessionId`);

      const success = await redirectToCheckout(sessionId);
      if (!success) throw new Error(`Could not redirect to stripe checkout`);
    } catch (err) {
      error = true;
      console.error(err);
    } finally {
      loading = false;
    }
  };

  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues: { email: '', terms: false, subscribe: false },
    validationSchema: yup.object().shape({
      email: yup.string().email().required(),
      terms: yup.boolean().oneOf([true], 'terms must be accepted'),
      subscribe: yup.boolean(),
    }),
    onSubmit,
  });
</script>

<div class="container">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then { Name, Description, Price }}
    <SvelteSeo title={`Keyprez - ${capitalize(Name)}`} {Description} />
    <img src={`/${Name.toLowerCase()}.jpg`} alt={Name} />
    <div class="text">
      <h1>{Description.toUpperCase()}</h1>
      <p>
        Curabitur quis facilisis sapien. Cras luctus elit in ante tincidunt aliquet. Praesent interdum euismod felis
        eget condimentum. Nam cursus pulvinar lacus at ultricies. Aliquam tempor consequat est, eu iaculis ipsum
        imperdiet non. Duis placerat aliquam justo, id sodales risus tempus eget. Quisque ornare, nunc a laoreet
        sagittis, diam sapien varius odio, nec posuere ligula neque eu nunc. Nulla hendrerit luctus turpis, sed egestas
        enim iaculis sit amet. Nam nec nunc quis nibh rutrum condimentum. Vestibulum risus risus, tempor in nisl sit
        amet, luctus gravida mauris. Donec vestibulum est quis blandit tincidunt. Curabitur lorem ex, porta id pretium
        eu, scelerisque sit amet lacus. Sed est ante, aliquet ut semper in, sagittis pretium leo. Nam gravida pharetra
        ex quis gravida. Suspendisse arcu lacus, sodales in est sit amet, faucibus euismod enim.
      </p>
      <p>
        Integer nec odio nec mi ornare auctor in ut libero. Pellentesque vel finibus sem. Donec vel lorem lorem. Fusce
        eu purus lobortis risus sagittis rutrum. Mauris gravida efficitur posuere. Praesent a libero ex. Nunc non
        sodales mauris. Ut commodo lectus sed est vestibulum, eu pulvinar mauris venenatis. Aenean tempor sapien blandit
        ex egestas rhoncus. Suspendisse sagittis dui nisl, non elementum neque eleifend in. Fusce cursus neque sed arcu
        fringilla pharetra. Fusce lacinia, augue id ultrices dictum, nibh nibh placerat mi, et tristique eros justo ut
        nisi. In tempus pretium hendrerit. Duis ac malesuada nulla. Nullam pellentesque purus id eros sagittis, egestas
        mollis sapien dignissim.
      </p>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
      <form class="form" on:submit={handleSubmit}>
        <div class="formInputWrapper">
          <input
            bind:this={inputRef}
            name="email"
            class="formInput {$errors.email ? 'errorInput' : ''}"
            type="text"
            placeholder="Type your email"
            bind:value={$form.email}
            on:keyup={handleChange}
            on:blur={handleChange}
          />
          <small class="error {$errors.email ? 'errorVisible' : ''}">{$errors.email}</small>
        </div>
        <div class="checkboxes">
          <div>
            <div class="checkboxWrapper">
              <Checkbox name="terms" label="I accept terms of agreement" bind:checked={$form.terms} />
              <small class="error errorCheckbox {!$form.terms ? 'errorVisible' : ''}">{$errors.terms}</small>
            </div>
            <Checkbox name="subscribe" label="Subscribe to newsletter" bind:checked={$form.subscribe} />
          </div>
        </div>
        <Button type="submit" text={error ? 'ERROR 😞' : 'BUY'} {loading} onClick={focus} />
      </form>
      <h2>&#36;<span>{Price}</span></h2>
    </div>
  {/await}
</div>

<style lang="scss">
  @import 'src/variables';

  .container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
  }

  img {
    flex: 60%;
    object-fit: cover;
    width: 100%;
    border-radius: $border-radius-large;
  }

  .text {
    align-items: center;
    display: flex;
    flex: 40%;
    flex-direction: column;
  }

  .form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;

    &InputWrapper {
      position: relative;
      display: flex;
      width: 100%;
    }
  }

  .checkboxes {
    margin: 0.5rem 0 1rem 0;
    display: flex;
    align-self: baseline;
  }

  .checkboxWrapper {
    position: relative;
  }

  .error {
    display: none;
    position: absolute;
    right: 0.3em;
    bottom: 0.3em;
    color: $color-warning;

    &Checkbox {
      bottom: 0;
      left: 105%;
      width: 100%;
    }

    &Visible {
      display: block;
    }

    &Input {
      border-color: $color-warning;
    }
  }

  @media (min-width: 1400px) {
    .container {
      flex-direction: row;
      gap: 2rem;
    }

    h1 {
      margin-top: 0;
    }
  }

  @media (min-width: 600px) {
    .form {
      width: 60%;
      max-width: 400px;
    }

    .checkboxes {
      align-self: center;
    }
  }
</style>

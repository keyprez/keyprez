<script context="module">
  import { fade } from 'svelte/transition';

  export const prerender = true;

  const handleSubmit = (e) => {
    const form = new FormData(e.target);
    const formData = Object.fromEntries(form);

    // Send formData to keyprez API to handle email sendout
  };
</script>

<script>
  import SvelteSeo from 'svelte-seo';
</script>

<SvelteSeo title="Keyprez - Contact" />

<div class="container" in:fade>
  <div class="form-container">
    <h1>Contact us</h1>
    <form class="form" method="POST" on:submit|preventDefault={handleSubmit}>
      <textarea placeholder="Type your message..." name="message" rows="15" maxlength="1000" required />
      <div class="inputs">
        <input placeholder="Name" type="text" name="name" />
        <input placeholder="Email" type="email" name="email" required />
      </div>
      <button type="submit">SUBMIT</button>
    </form>
  </div>
  <div class="faq">
    <h1>FAQ</h1>
    <ul>
      <li>Lorem ipsum dolor sit amet, consectetur adipiscing elit.?</li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
      <li>
        Morbi eget bibendum elit. Nullam blandit gravida vehicula. Suspendisse pretium quam lectus, sed elementum dolor
        iaculis ac?
      </li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
      <li>Pellentesque vel tellus sed ante euismod maximus. Suspendisse pharetra ornare ex eget pellentesque?</li>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>
    </ul>
  </div>
</div>

<style lang="scss">
  @import 'src/variables';

  .container {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .form-container {
    flex: 40%;
  }

  .form {
    align-items: center;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  input {
    box-sizing: border-box;
    width: 100%;
  }

  textarea {
    box-sizing: border-box;
  }

  .inputs {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    width: 100%;
  }

  button {
    margin: 1rem;
  }

  .faq {
    flex: 60%;
  }

  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
  }

  li {
    color: $color-tertiary;
    font-size: 150%;
  }

  p {
    margin-bottom: 2rem;
  }

  @media (min-width: 768px) {
    .inputs {
      flex-direction: row;
    }
  }

  @media (min-width: 1200px) {
    .container {
      flex-direction: row;
    }
  }
</style>

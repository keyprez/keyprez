<script>
  import '/src/app.css';
  import Header from '$lib/Header.svelte';
  import { onMount } from 'svelte';
  import github_logo from '/static/github_logo.svg';
  import instagram_logo from '/static/instagram_logo.svg';
  import reddit_logo from '/static/reddit_logo.svg';
  import meetup_logo from '/static/meetup_logo.svg';

</script>

<Header />

<main>
  <slot />
</main>

<footer>
  <a href="https://github.com/keyprez/keyprez" target="blank"><img src={github_logo} alt="GitHub" /></a>
  <a href="https://github.com/keyprez/keyprez" target="blank"><img src={instagram_logo} alt="Instagram" /></a>
  <a href="https://github.com/keyprez/keyprez" target="blank"><img src={reddit_logo} alt="Reddit" /></a>
  <a href="https://github.com/keyprez/keyprez" target="blank"><img src={meetup_logo} alt="Meetup" /></a>
</footer>

<style>
  main {
    flex: 1;
    display: flex;
    flex-direction: column;
    padding: 1rem;
    width: 100%;
    max-width: 1024px;
    margin: 0 auto;
    box-sizing: border-box;
  }

  footer {
    display: flex;
    gap: 1.5rem;
    justify-content: center;
    padding: 3rem;
    background-color: var(--primary-color);
  }

  img {
    width: 1.5rem;
  }

</style>

<script>
	import '../app.css';
	// import "../index.css"; // tailwind output file (important for hot reloading)
	import Header from '$lib/Header.svelte';
	import { theme } from './store';
	import { onMount } from 'svelte';

	onMount(() => {
		if ($theme === 'dark') {
			document.querySelector('html').classList.add('dark');
		}
	});

</script>

<Header />

<main>
	<slot />
</main>

<footer>
	<p>
		Check out the code on <a href="https://github.com/keyprez/keyprez" target="blank">GitHub</a>
	</p>
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
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 40px;
		background-color: var(--primary-color);
	}

	footer a {
		font-weight: bold;
	}

	@media (min-width: 480px) {
		footer {
			padding: 40px 0;
		}
	}

</style>

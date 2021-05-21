<script context="module">
	export const prerender = true;

</script>

<svelte:head>
	<title>Products</title>
</svelte:head>

<div class="content">
	<h1>Products</h1>
	<p>
		Phasellus magna velit, facilisis nec fringilla nec, molestie faucibus mi. Nulla bibendum neque
		nec felis interdum laoreet. Nullam eu urna leo. Duis suscipit a sem in sollicitudin. In hac
		habitasse platea dictumst. Suspendisse fermentum nibh ipsum, quis ultrices mi commodo congue.
		Suspendisse dapibus at massa ut placerat. Donec laoreet vulputate elit, eu suscipit ipsum congue
		semper. Nulla eu sapien lorem. Maecenas erat sem, pulvinar a fermentum in, aliquam non est. Nunc
		auctor libero diam, sodales blandit diam tempor consequat. Phasellus euismod massa sapien, ac
		elementum mauris finibus vitae. Curabitur consectetur sagittis porttitor.
	</p>
	<img src="/static/rhino.png" alt="Rhino" />
</div>

<style>
	.content {
		width: 100%;
		text-align: justify;
		max-width: var(--column-width);
		margin: var(--column-margin-top) auto 0 auto;
	}

	img {
		width: 100%;
		margin: 1em 0;
	}

</style>

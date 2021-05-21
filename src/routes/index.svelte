<script context="module">
	export const prerender = true;

</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<div class="content">
	<h1 class="text-red-500">Welcome</h1>
	<p>
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi eget bibendum elit. Nullam
		blandit gravida vehicula. Suspendisse pretium quam lectus, sed elementum dolor iaculis ac.
		Pellentesque vel tellus sed ante euismod maximus. Suspendisse pharetra ornare ex eget
		pellentesque. Suspendisse tincidunt augue urna, quis facilisis quam bibendum eget. Nulla at eros
		elementum, posuere mi a, suscipit dolor. Nullam lobortis vulputate dolor, quis laoreet neque
		finibus at. Maecenas aliquam scelerisque leo, et bibendum nulla consequat vitae. Suspendisse
		lacus sapien, auctor sed vestibulum at, maximus quis eros. Etiam rhoncus elementum tempor.
		Suspendisse porttitor nisl dolor, non euismod risus malesuada sit amet. Integer id tempus risus.
	</p>

	<div class="gallery">
		<div class="image-container">
			<a href="/products"
				><img src="/static/pcb_bg.jpg" alt="PCB" />
				<h2>PCB's</h2></a
			>
		</div>
		<div class="image-container">
			<a href="/products"
				><img src="/static/case_bg.png" alt="Case" />
				<div class="text"><h2>Cases</h2></div></a
			>
		</div>
		<div class="image-container">
			<a href="/products"
				><img src="/static/plate_bg.png" alt="Plate" />
				<h2>Plates</h2></a
			>
		</div>
		<div class="image-container">
			<a href="/products"
				><img src="/static/kit_bg.jpg" alt="DIY kits" />
				<h2>DIY kits</h2></a
			>
		</div>
		<div class="image-container">
			<a href="/products"
				><img src="/static/accessories_bg.png" alt="Accessories" />
				<h2>Accessories</h2></a
			>
		</div>
		<div class="image-container">
			<a href="/products"
				><img src="/static/swag_bg.jpg" alt="Swag" />
				<h2>Swag</h2></a
			>
		</div>
	</div>
</div>

<style style lang="postcss">
	.content {
		width: 100%;
		text-align: justify;
		max-width: var(--column-width);
		margin: var(--column-margin-top) auto 0 auto;
	}

	.gallery {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
	}

	.image-container {
		position: relative;
		display: inline-block;
	}

	.image-container h2 {
		position: absolute;
		@apply text-center;
		@apply mx-auto;
		left: 0;
		right: 0;
		top: 40%;
	}

	img {
		flex: 1 1 auto;
		margin: 1em;
		transition: 0.3s linear;
		opacity: 0.4;
	}

	img:hover {
		transform: scale(1.1);
		opacity: 1;
		color: var(--accent-color);
	}

</style>

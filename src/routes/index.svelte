<script context="module">
  import pcb_bg from '/static/pcb_bg.jpg';
  import case_bg from '/static/case_bg.png';
  import plate_bg from '/static/plate_bg.png';
  import kit_bg from '/static/kit_bg.jpg';
  import accessories_bg from '/static/accessories_bg.png';
  import swag_bg from '/static/swag_bg.jpg';

  export const prerender = true;

</script>

<svelte:head>
  <title>Home</title>
</svelte:head>

<div class="content">
  <div class="gallery">
    <div class="image-container">
      <a href="/products"
        ><img src={pcb_bg} alt="PCB" />
        <h2>PCB's</h2></a
      >
    </div>
    <div class="image-container">
      <a href="/products"
        ><img src={case_bg} alt="Case" />
        <div class="text"><h2>Cases</h2></div></a
      >
    </div>
    <div class="image-container">
      <a href="/products"
        ><img src={plate_bg} alt="Plate" />
        <h2>Plates</h2></a
      >
    </div>
    <div class="image-container">
      <a href="/products"
        ><img src={kit_bg} alt="DIY kits" />
        <h2>DIY kits</h2></a
      >
    </div>
    <div class="image-container">
      <a href="/products"
        ><img src={accessories_bg} alt="Accessories" />
        <h2>Accessories</h2></a
      >
    </div>
    <div class="image-container">
      <a href="/products"
        ><img src={swag_bg} alt="Swag" />
        <h2>Swag</h2></a
      >
    </div>
  </div>
</div>

<style>
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

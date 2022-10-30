<script>
  const toggle = () => {
    const list = window.document.documentElement.classList;
    list.toggle('dark');
    localStorage.setItem('theme', [...list].includes('dark') ? 'dark' : 'light');
  };
</script>

<button class="relative w-10 flex justify-center items-center" on:click={toggle}>
  <img class="absolute w-5" src="/keycap_simple.svg" alt="keycap" />
  <img class="absolute block dark:hidden" src="/sun.svg" alt="sun" />
</button>

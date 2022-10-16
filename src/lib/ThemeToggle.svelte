<script>
  const toggle = () => {
    const list = window.document.documentElement.classList;
    list.toggle('dark');
    localStorage.setItem('theme', [...list].includes('dark') ? 'dark' : 'light');
  };
</script>

<button on:click={toggle}>Toggle</button>

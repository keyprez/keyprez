<footer>
  <h2><strong class="subscribe">SUBSCRIBE</strong> to latest news and updates</h2>
  <form class="form">
    <input type="text" placeholder="Your e-mail" />
    <button>SUBSCRIBE</button>
  </form>
  <div class="links">
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="github_logo.svg" alt="GitHub" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="instagram_logo.svg" alt="Instagram" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="reddit_logo.svg" alt="Reddit" /></a>
    <a href="https://github.com/keyprez/keyprez" target="blank"><img src="meetup_logo.svg" alt="Meetup" /></a>
  </div>
</footer>

<style type="text/scss">
  @import 'src/variables';

  footer {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    justify-content: center;
    letter-spacing: 2px;
    padding: 10rem 1rem;
  }

  h2,
  strong {
    color: inherit;
    font-weight: 1000;
    margin: 0;
  }

  .subscribe {
    color: $color-tertiary;
  }

  .form {
    display: flex;
    flex-direction: row;
    gap: 0.5rem;
    justify-content: center;
  }

  input {
    width: 20rem;
  }

  .links {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    justify-content: center;
  }

  img {
    width: 1.5rem;
  }

  img:hover {
    transform: scale(1.3);
  }
</style>

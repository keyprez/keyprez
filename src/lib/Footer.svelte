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
    background: $color-secondary-dark;
    border: 0;
    color: $color-tertiary;
    font-size: inherit;
    height: 4rem;
    outline: none;
    padding: 0 0.5rem;
    width: 20rem;
  }

  button {
    background: $color-tertiary;
    border: 0;
    color: $color-secondary-light;
    font-size: 1rem;
    font-weight: 700;
    letter-spacing: inherit;
    transition: 0.1s linear;
    width: 10rem;
  }

  button:hover {
    background: transparent;
    border: 1px solid $color-secondary-light;
    color: $color-secondary-light;
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

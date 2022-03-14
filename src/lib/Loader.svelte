<img src="/keycap_simple.svg" alt="keycap" />

<style>
  img {
    transition: all 0.2s;
    height: 1.2rem;

    -webkit-animation: pulse 3s infinite ease-in-out;
    -o-animation: pulse 3s infinite ease-in-out;
    -ms-animation: pulse 3s infinite ease-in-out;
    -moz-animation: pulse 3s infinite ease-in-out;
    animation: pulse 3s infinite ease-in-out;
  }

  @-webkit-keyframes pulse {
    0% {
      -webkit-transform: scale(1.5);
      opacity: 0.3;
    }
    50% {
      -webkit-transform: scale(2);
      opacity: 1;
    }
    100% {
      -webkit-transform: scale(1.5);
      opacity: 0.3;
    }
  }

  @keyframes pulse {
    0% {
      transform: scale(1.5);
      opacity: 0.3;
    }
    50% {
      transform: scale(2);
      opacity: 1;
    }
    100% {
      transform: scale(1.5);
      opacity: 0.3;
    }
  }
</style>

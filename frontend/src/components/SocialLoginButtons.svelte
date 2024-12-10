<script lang="ts">

    type Provider = 'Google' | 'GitHub' | 'Apple';

    // Логика для обработки кликов на кнопки (если потребуется)
    const handleSocialLogin = (provider: Provider) => {
      console.log(`Logging in with ${provider}`);
    };
  </script>
  
  <div class="but-log">
    <button class="button-login" on:click={() => handleSocialLogin('Google')}>
      <img src="./images/Google Icon.svg" alt="Google" height="22" width="22" />
      Google
    </button>
    <button class="button-login" on:click={() => handleSocialLogin('GitHub')}>
      <img src="./images/Github Icon.svg" alt="GitHub" height="27" width="27" />
      GitHub
    </button>
    <button class="button-login" on:click={() => handleSocialLogin('Apple')}>
      <img src="./images/Apple Icon.svg" alt="Apple" height="25" width="25" />
      Apple
    </button>
  </div>
  
  <style>
    .but-log {
      display: flex;
      flex-direction: column;
      gap: 10px;
      padding: 1rem;
      text-align: center;
      align-items: center;
      width: 450px;
      margin-left: auto;
      margin-right: auto;
      padding-right: 0.2rem;
    }
  
    .button-login {
      color: black;
      background-color: #ffffff;
      width: 104%;
      height: 3.5rem;
      border-radius: 10px;
      cursor: pointer;
      font-size: 20px;
      border-style: none;
      margin: 0.1rem;
      display: flex;
      justify-content: center;
      align-items: center;
      opacity: 1;
    }
  
    .button-login:hover {
      box-shadow: 0px 0px 15px 1px #ff7b00;
    }
  
    .button-login img {
      align-items: center;
      margin: 3px;
      background-color: transparent;
      margin-right: 7px;
    }
  </style>
  
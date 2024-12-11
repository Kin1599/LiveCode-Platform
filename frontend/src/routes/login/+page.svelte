<script>
  import LoginForm from "../../components/LoginForm.svelte";
  import SocialLoginButtons from "../../components/SocialLoginButtons.svelte";
  import RegForm from "../../components/RegForm.svelte";

  let showLoginForm = true;

  function toggleLoginForm() {
    showLoginForm = !showLoginForm;
  }

</script>

<div>
  {#if showLoginForm}
    <LoginForm {toggleLoginForm}/>
  {:else}
    <RegForm {toggleLoginForm}/>
  {/if}
  <h2>Войти с помощью</h2>
  <SocialLoginButtons />
</div>

<style>
  h2 {
    font-size: 2rem;
    text-align: center;
    position: relative;
    color: #ffffff;
  }
</style>

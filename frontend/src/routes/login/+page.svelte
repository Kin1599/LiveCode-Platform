<script>
  import LoginForm from "../../components/LoginForm.svelte";
  import SocialLoginButtons from "../../components/SocialLoginButtons.svelte";

</script>

<div>
  <LoginForm />
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

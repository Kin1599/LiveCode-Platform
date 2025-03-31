<script lang="ts">
    import SendServer from "../api/api.js";  
    
    let email = "";
    let password = "";
    let message = "";

    export let toggleLoginForm = () => {};
  
    const handleLogin = async () => {
        message = "";
        try {
            const response = await SendServer.login(email, password);
            console.log(response);
            if (response.status === 200) {
                localStorage.setItem("token", response.data.access_token);
                localStorage.setItem("refreshToken", response.data.refresh_token);
                message = "Успешный вход";
                window.location.assign("/#");
            } else {
                message = "Ошибка при входе: " + response.statusText;
            }
        } catch (error) {
            console.error("Login error:", error);
            message = "Произошла ошибка при входе";
        }
    };
</script>
  
<form>
    <div class="form-widget">
        <h2 class="header" style="margin-bottom: 0.5rem;">Войдите в аккаунт</h2>
        <div class="form-group">
            <input
                required
                id="email"
                class="form-control"
                type="email"
                placeholder="Почта или имя пользователя"
                bind:value={email}
            />
        </div>
        <div class="form-group">
            <input
                required
                id="password"
                class="form-control"
                type="password"
                placeholder="Пароль"
                bind:value={password}
            />
        </div>
        <div>
            <button type="button" on:click={handleLogin}>Войти</button>
        </div>
      
        <div class="link">
            <a href="/forgot-pass">Забыли пароль?</a>
        </div>

        <div class="register">
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <div>Нет аккаунта? <span class="register-link" on:click={toggleLoginForm}>Зарегайтесь!</span></div>
        </div>
  
        {#if message}
            <p>{message}</p>
        {/if}
    </div>
</form>
  
<style>
    .form-widget {
        display: flex;
        flex-direction: column;
        gap: 10px;
        padding: 3rem;
        text-align: center;
        margin-top: 8rem;
        width: 450px;
        margin-left: auto;
        margin-right: auto;
    }
  
    input {
        width: 100%;
        height: 50px;
        background-color: #ffffff;
        border-radius: 10px;
        position: relative;
        padding-left: 1rem;
    }
  
    button {
        color: white;
        background-color: #ff7b00;
        border-style: solid;
        width: 104%;
        height: 3.5rem;
        border-radius: 10px;
        cursor: pointer;
        font-weight: 700;
        font-size: 1.2rem;
    }
  
    button:hover {
        background: transparent;
        border-width: 2px;
        border-color: #ff7b00;
        color: #ff7b00;
    }
  
    .link {
        margin-top: 0.5rem;
        align-self: flex-start;
    }
  
    .link a {
        color: #ffffff;
        text-decoration: none;
    }
  
    .link a:hover {
        color: #ff7b00;
    }

    .register{
        text-align: left;
    }

    .register-link{
        color: #ff7b00;
        cursor: pointer;
    }
</style>
  

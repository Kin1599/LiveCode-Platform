<script lang="ts">
    import SendServer from "../api/api.js";

    export let toggleLoginForm = () => {};
    
    let email = "";
    let password = "";
    let message = "";

    const handleRegister = async () => {
        message = "";
        try {
            const sessionData = await SendServer.register(email, password);
            localStorage.setItem("token", sessionData.token);

            const userInfo = await SendServer.getUserInfo(sessionData.token);
            localStorage.setItem("user", JSON.stringify(userInfo));

            message = `Добро пожаловать, ${userInfo.Nickname}!`;
            window.location.assign("/#");
        } catch (error) {
            console.error("Register error:", error);
            message = "Произошла ошибка при регистрации, попробуйте ещё раз";
        }
    };
</script>
  
<form class="register-form">
    <h2 class="header">Регистрация</h2>
    
    <div class="input-group">
        <input
            required
            id="email"
            class="input email"
            type="email"
            placeholder="Почта или имя пользователя"
            bind:value={email}
        />
    </div>
        
    <div class="input-group">
        <input
            required
            id="password"
            class="input password"
            type="password"
            placeholder="Пароль"
            bind:value={password}
        />
    </div>
    
    <div class="actions">
        <button type="button" class="btn primary" on:click={handleRegister}>Создать аккаунт</button>
        <button type="button" class="btn back" on:click={toggleLoginForm}>Назад</button>
    </div>
  
    {#if message}
        <p>{message}</p>
    {/if}
</form>
  
<style lang="scss">
    @use "../styles/colors.scss" as *;

    .register-form {
        display: flex;
        flex-direction: column;
        gap: 10px;
        padding: 3rem;
        text-align: center;
        margin: 8rem auto 0;
        width: 450px;
    }

    .header {
        font-size: 1.5rem;
        margin-bottom: 1.5rem;
    }

    .input-group {
        display: flex;
        flex-direction: column;
        gap: 5px;
    }

    .actions{
        display: flex;
        flex-direction: column;
        gap: 15px;
    }
  
    input {
        width: 100%;
        height: 50px;
        background-color: #ffffff;
        border-radius: 10px;
        padding-left: 1rem;
        position: relative;
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
        border: 2px solid #ff7b00;
        color: #ff7b00;
    }

    .back-btn{
        color: white;
        background-color: transparent;
        border: 2px solid #ff7b00;
        color: #ff7b00;
        width: 104%;
        height: 3.5rem;
        border-radius: 10px;
        cursor: pointer;
        font-weight: 700;
        font-size: 1.2rem;
    }

    .back-btn:hover{
        background-color: #ff7b00;
        color: #ffffff;
    }
</style>
  

<script>
    import SendServer from "../api/api.js";
  
    let email = "";
    let password = "";
    let message = "";

    export let toggleLoginForm = () => {};
  
    const handleRegister = async () => {
        message = "";
        try {
            const response = await SendServer.register(email, password);
            if (response.status === 200) {
                localStorage.setItem("token", response.data.token);
                message = "Аккаунт успешно создан!";
                window.location.assign("/#");
            } else {
                message = "Ошибка при регистрации: " + response.statusText;
            }
        } catch (error) {
            console.error("Register error:", error);
            message = "Произошла ошибка при регистрации, попробуйте ещё раз";
        }
    };
</script>
  
<form>
    <div class="form-widget">
        <h2 class="header">Регистрация</h2>
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
        <div class="actions">
            <button type="button" on:click={handleRegister}>Создать аккаунт</button>
            <button class="back-btn" type="button" on:click={toggleLoginForm}>Назад</button>
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
  

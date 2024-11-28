<script>
    let username = ''
    let password = ''
  
  </script>

<form>
    <div class="form-widget">
        <h2 class="header" style="margin-bottom: 0.5rem;">Войдите в аккаунт</h2>
        <div class="form-group">
            <input required id='email' class="form-control" type="email" placeholder="Почта или имя пользователя" bind:value={username}/>
        </div>
        <div class="form-group">
            <input required id='password' class="form-control" type="password" placeholder="Пароль" bind:value={password}/>
        </div>
        <div>
            <button type="submit">Войти</button>
        </div>

        <div class="link">
            <a href="#">Забыли пароль?</a>
        </div>
    </div>
</form>

<h2>Войти с помощью</h2>
<div class="but-log">
    <button class="button-login"><img src="./google-logo.jpg" alt="Google" height="22" width="22">Google</button>
    <button class="button-login"><img src="./github-logo1.png" alt="GitHub" height="30" width="50">GitHub</button>
    <button class="button-login"><img src="./Apple-logo1.png" alt="Apple" height="20" width="20">Apple</button>
</div>

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

    form input {
        outline: none;
        border-style: none;
    }

    form input:focus {
        outline: 0;
        border-style: 10px solid #FF7B00;
    }

    h2 {
        font-size: 2rem;
        text-align: center;
        position: relative;
        color:#ffffff; 
    }

    input {
        width: 100%;
        height: 50px;
        background-color: #ffffff;
        border-radius: 15px;
        position: relative;
        padding-left: 1rem;
    }

    form button {
        color: white;
        background-color: #FF7B00;
        border-style: solid;
        width: 104%;
        height: 3.5rem;
        border-radius: 15px;
        cursor: pointer;
    }

    form button:hover {
        background: transparent;
        border-width: 2px;
        border-color: #FF7B00;
        color: #FF7B00;
    }

    form a {
        color: #ffffff;
        text-decoration: none;
    }

    .link {
        margin-top: 0.5rem;
        align-self: flex-start;
    }

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
        border-radius: 15px;
        cursor: pointer;
        font-size: 20px;
        border-style: none;
        margin: 0.1rem;
    }

    .button-login:hover {
        box-shadow: 0px 0px 15px 1px #FF7B00;
    }

    .button-login img {
        align-items: center;
        margin: 3px;
    }
</style>

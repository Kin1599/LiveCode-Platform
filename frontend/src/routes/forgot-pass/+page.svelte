<script>
  let email = "";
  let step = "email"; // состояние формы
  let confirmationCode = Array(6).fill(""); // массив для хранения кода
  let password = ""; // первый пароль
  let confirmPassword = ""; // повторный пароль
  let errorMessage = ""; // сообщение об ошибке

  function nextStep() {
    if (step === "email") {
      step = "confirmation";
    } else if (step === "confirmation") {
      step = "passwordReset";
    }
  }

  function goBack() {
    if (step === "confirmation") {
      step = "email";
    } else if (step === "passwordReset") {
      step = "confirmation";
    }
  }

  function handlePasswordReset() {
    if (password === confirmPassword) {
      // Если пароли совпадают, переходим дальше
      errorMessage = "";
      window.location.href = "/login";
    } else {
      errorMessage = "Пароли не совпадают. Попробуйте снова.";
    }
  }

  /**
   * @param {{ key: string; }} event
   * @param {number} index
   */
  function handleArrowNavigation(event, index) {
    const inputs = document.querySelectorAll(".code-input");

    if (event.key === "ArrowRight") {
      if (index < inputs.length - 1) {
        // @ts-ignore
        inputs[index + 1].focus();
      }
    } else if (event.key === "ArrowLeft") {
      if (index > 0) {
        // @ts-ignore
        inputs[index - 1].focus();
      }
    }
  }
</script>

<form>
  <div class="form-widget">
    {#if step === "email"}
      <h2 class="header" style="margin-bottom: 0.5rem;">
        Восстановление пароля
      </h2>
      <div class="form-group">
        <input
          required
          id="email"
          class="form-control"
          type="email"
          placeholder="Введите почту или номер телефона"
          bind:value={email}
        />
      </div>
      <div>
        <button type="button" on:click={nextStep}
          >Получить код подтверждения</button
        >
      </div>
      <div>
        <a href="/login"
          ><button type="button" class="back-btn">Назад</button></a
        >
      </div>
    {/if}

    {#if step === "confirmation"}
      <h2 class="header" style="margin-bottom: 0.5rem;">
        Введите код подтверждения
      </h2>
      <div>
        {#each confirmationCode as _, index}
          <input
            id={`code-${index}`}
            class="code-input"
            type="text"
            maxlength="1"
            on:keydown={(e) => handleArrowNavigation(e, index)}
          />
        {/each}
      </div>
      <div>
        <button type="button" on:click={nextStep}>Далее</button>
      </div>
      <div>
        <button type="button" class="back-btn" on:click={() => (step = "email")}
          >Назад</button
        >
      </div>
    {/if}

    {#if step === "passwordReset"}
      <h2 class="header" style="margin-bottom: 0.5rem;">
        Введите новый пароль
      </h2>
      <div class="form-group">
        <input
          required
          class="form-control"
          type="password"
          placeholder="Новый пароль"
          bind:value={password}
        />
      </div>
      <div class="form-group">
        <input
          required
          class="form-control"
          type="password"
          placeholder="Ещё разок"
          bind:value={confirmPassword}
        />
      </div>
      {#if errorMessage}
        <p class="error-message">{errorMessage}</p>
      {/if}
      <div>
        <button type="button" on:click={handlePasswordReset}
          >Сменить пароль</button
        >
      </div>
      <div>
        <button type="button" class="back-btn" on:click={goBack}>Назад</button>
      </div>
    {/if}
  </div>
</form>

<style>
  form {
    margin-top: 15rem;
    margin-bottom: auto;
    justify-content: center;
  }

  .form-widget {
    display: flex;
    flex-direction: column;
    justify-content: center;
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
    border-style: 10px solid #ff7b00;
  }

  h2 {
    font-size: 2rem;
    text-align: center;
    position: relative;
    color: #ffffff;
  }

  input {
    width: 100%;
    height: 50px;
    background-color: #ffffff;
    border-radius: 10px;
    position: relative;
    padding-left: 1rem;
  }

  form button {
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

  form button:hover {
    background: transparent;
    border-width: 2px;
    border-color: #ff7b00;
    color: #ff7b00;
  }

  .back-btn {
    background: transparent;
    border-width: 2px;
    border-color: #ff7b00;
    color: #ff7b00;
  }

  .back-btn:hover {
    background: transparent;
    border-width: 2px;
    border-color: #7d3c00;
    color: #7d3c00;
  }

  .code-input {
    width: 2rem;
    height: 4rem;
    margin: 0.2rem;
    text-align: center;
    font-size: 1.2rem;
    background-color: transparent;
    border-radius: 10px;
    border: 1px solid #7e7e7e;
    color: #ffffff;
    margin-bottom: 1rem;
  }

  .code-input:focus {
    outline: none;
    border-color: #ff7b00;
  }

  .error-message {
    color: red;
    font-size: 0.9rem;
    margin-top: 0.5rem;
  }
</style>

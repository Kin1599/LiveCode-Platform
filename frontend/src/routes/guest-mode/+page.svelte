<script>
  let name = "";
  let surname = "";

  function handleSubmit() {
    console.log(`Name: ${name}, Surname: ${surname}`);
    window.location.href = "/#";
  }
</script>

<form>
  <div class="form-widget">
    <h2 class="header">Гостевой вход</h2>
    <div class="form-body">
      <div class="avatar">VN</div>
      <div class="inputs">
        <input
          required
          class="form-control"
          type="text"
          bind:value={name}
          placeholder="Your name"
        />
        <input
          required
          class="form-control"
          type="text"
          bind:value={surname}
          placeholder="Your surname"
        />
      </div>
    </div>
    <div style="margin-top: 1rem;">
      <button type="button" on:click={handleSubmit}>Войти</button>
    </div>
    <div>
      <a href="/login">
        <button type="button" class="back-btn">Назад</button>
      </a>
    </div>
  </div>
</form>

<style>
  form {
    margin-top: auto;
    margin-bottom: auto;
  }

  .form-widget {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 1rem;
    padding: 2rem;
    text-align: center;
    margin-top: 8rem;
    width: 40rem;
    margin-left: auto;
    margin-right: auto;
  }

  .form-body {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 1rem;
  }

  .avatar {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    background-color: #333;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 24px;
    color: #fff;
    font-weight: bold;
    margin-right: 2rem;
  }

  .inputs {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin: 3rem 0;
  }

  form input {
    outline: none;
    background-color: #162832;
    border-radius: 0.3125rem;
    border: 1px solid rgba(106, 106, 106, 0.4);
    font-size: 1.2rem;
    font-weight: 200;
    color: #fff;
    width: 100%;
    height: 50px;
    padding-left: 1rem;
  }

  form input:focus {
    outline: 0;
    border: 2px solid #ff7b00;
  }

  input::placeholder {
    color: #d4d4d4;
  }

  h2 {
    font-size: 2rem;
    color: #ffffff;
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
    border: 2px solid #ff7b00;
    color: #ff7b00;
  }

  .back-btn {
    background: transparent;
    border: 2px solid #ff7b00;
    color: #ff7b00;
  }

  .back-btn:hover {
    background: transparent;
    border-color: #7d3c00;
    color: #7d3c00;
  }
</style>

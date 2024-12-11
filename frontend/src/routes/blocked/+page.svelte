<script lang="ts">
   
</script>
  
<div class="blocked">
  <div class="blocked-content">
    <h1 class="blocked-title">Упс...</h1>
    <div class="blocked-description">
      <p class="description">Кажется, вас кикнули...(</p>
      <p class="description">Чтобы вернуться обратно, свяжитесь с владельцем сессии</p>
    </div>
    <button class="blocked-button" onclick={() => window.location.assign("/")}>Вернуться на главную</button>
  </div>
  <img class="blocked-image" src="./images/blocked.svg" alt="blocked">
</div>

<style>
  .blocked {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    height: 100vh;
    text-align: center;
    position: relative;
  }

  .blocked-image{
    position: absolute;
    left: 0;
    top: 50px;
    z-index: -1;
  }

  .blocked-content > *:not(:last-child){
    margin-bottom: 50px;
  }

  .blocked-title{
    font-size: 96px;
    font-weight: 500;
    line-height: 112.5px;
    color: #FF7B00;
  }

  .blocked-description{
    font-size: 36px;
    font-weight: 300;
    line-height: 42.19px;
    width: 635px;
  }

  .blocked-description > *:not(:last-child){
    margin-bottom: 25px;
  }

  .blocked-button{
    background-color: #162832;
    border-radius: 10px;
    font-size: 32px;
    line-height: 37.5px;
    font-weight: 500;
    color: #fff;
    border: none;
    width: 100%;
    height: 88px;
    cursor: pointer;
  }
</style>
  
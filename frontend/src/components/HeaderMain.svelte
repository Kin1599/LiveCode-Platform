<script>
  export let searchQuery = "";
  export let toggleVisibility;
</script>
  
<header class="header">
  <div class="left-section">
    <div class="icons">
      <button on:click={toggleVisibility}><img src="./images/icon-sidebar.svg" alt="" /></button>
    </div>
  </div>

  <div class="search-container">
    <img src="./images/icon-search.svg" alt="" />
    <input
      class="search-input"
      type="text"
      placeholder="Найти команду..."
      bind:value={searchQuery}
    />
  </div>

  <div class="right-section">
    <button><img src="./images/icon-plus.svg" alt="" /></button>
    <button><img src="./images/icon-noti.svg" alt="" /></button>
    <button><img src="./images/icon-more.svg" alt="" /></button>
    <div class="avatar"></div>
  </div>
</header>
  
<style>
    /* верхний блок */
  .header {
    grid-area: header;
    padding: 10px 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #444;
    max-height: 5.5rem;
    text-align: center;
  }

  .left-section {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .icons {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1.2rem;
  }

  .icons button {
    background-color: #0b1419;
    border-color: transparent;
    cursor: pointer;
  }

  .icons img {
    width: 1.5rem;
    height: 1.5rem;
  }

  .search-container {
    display: flex;
    justify-content: center;
    gap: 5px;
    padding: 10px;
    padding-left: 27px;
    flex-grow: 1;
    max-width: 900px;
    border: 1px solid #7e7e7e;
    border-radius: 10px;
    transition: border 0.2s;
  }

  .search-container:hover, .search-container:focus-within {
    border: 1px solid #ff7b00;
  }

  .search-input {
    flex-grow: 1;
    border: none;
    outline: none;
    color: #ffffff;
    font-size: 20px;
    background-color: #0b1419;
  }

  .search-input::placeholder {
    color: #7e7e7e;
  }

  .right-section {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .right-section button {
    background-color: #0b1419;
    border-color: transparent;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: #444;
    border-radius: 50%;
  }
</style>
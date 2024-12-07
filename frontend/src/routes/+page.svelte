<script>
  let folders = [
    { name: "Shared with me", type: "shared", files: [] },
    { name: "Unnamed (1)", type: "folder", files: [] },
  ];

  //для количества реп
  let repls = 1;

  // для выбранного пункта меню
  let selected = "Repls";
  // для изменения состояния выбранного пункта меню
  /**
   * @param {string} item
   */
  function selectItem(item) {
    selected = item;
  }

  let showBlocks = true;

  // для переключения состояния видимости бокового блока
  function toggleVisibility() {
    showBlocks = !showBlocks;
  }

  let searchQuery = "";

  let user = "username";
</script>

<div class="layout">
  <header class="header {showBlocks ? '' : 'expanded'}">
    <div class="left-section">
      <div class="icons">
        <button on:click={toggleVisibility}
          ><img src="./images/icon-sidebar.svg" alt="" /></button
        >
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

  <aside class="sidebar {showBlocks ? '' : 'hidden'}">
    <ul style="margin-top: 3rem;">
      <li
        class:selected={selected === "Repls"}
        on:click={() => selectItem("Repls")}
      >
        <img src="./images/folder-orange.svg" alt="" />Repls
      </li>
      <li
        class:selected={selected === "Настройки"}
        on:click={() => selectItem("Настройки")}
      >
        <img src="./images/icon-settings.svg" alt="" />Настройки
      </li>
    </ul>
  </aside>

  <main
    class="main"
    class:expanded={!showBlocks}
    style="margin-left: 4rem; margin-top: 2.5rem; margin-right: 4rem"
  >
    {#if selected === "Repls"}
      <div class="main-header">
        <div style="display: flex; align-items: center;">
          <div class="title">
            <img src="./images/icon-folder.svg" alt="" />
            <h2>Repls</h2>
          </div>
          <span class="info">({repls}/3 Repls)</span>
        </div>

        <div class="actions">
          <button class="import"
            ><img src="./images/icon-gitnub.svg" alt="" />Импортировать из
            GitHub</button
          >
          <button class="create">+ Создать</button>
        </div>
      </div>

      <button class="new-folder" style="margin-bottom: 20px;"
        ><img src="./images/icon-new-folder.svg" alt="" />Новая папка</button
      >

      <!-- Список папок -->
      <div class="folder-list">
        {#each folders as folder}
          <div class="folder">
            {folder.name}
          </div>
        {/each}
      </div>
    {:else if selected === "Настройки"}
      <h2>Настройки</h2>
      <p>Здесь будет контент для настроек...</p>
    {/if}
  </main>

  <div class="user-panel {showBlocks ? '' : 'hidden'}">
    <div class="user-info">
      <div class="avatar"></div>
      <div><a href="/login">{user}</a></div>
    </div>
  </div>
</div>

<style>
  .layout {
    display: grid;
    grid-template-areas:
      "user-panel header"
      "sidebar main";
    grid-template-columns: 17.5rem 1fr;
    grid-template-rows: 80px 1fr;
    height: 100vh;
  }

  button {
    cursor: pointer;
    border-radius: 10px;
    color: #fff;
  }

  .hidden {
    width: 0;
    opacity: 0;
    overflow: hidden;
  }

  .expanded {
    width: 100%;
    transform: translateX(-12rem);
  }

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
    background: transparent;
    padding: 0.3rem;
    padding-left: 1.3rem;
    flex-grow: 1;
    max-width: 900px;
    border: 1px solid #7e7e7e;
    margin: 0.7rem 1rem;
    border-radius: 10px;
  }

  .search-container:hover {
    border: 1px solid #ff7b00;
  }

  .search-input {
    flex-grow: 1;
    border: none;
    outline: none;
    padding: 0.5rem;
    color: #ffffff;
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
    border-color: transparent;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: #444;
    border-radius: 50%;
  }

  /* боковой блок */
  .sidebar {
    grid-area: sidebar;
    padding: 20px;
    border-right: 1px solid #444;
  }

  .sidebar ul {
    list-style-type: none;
    padding: 0;
  }

  .sidebar li {
    padding: 15px;
    font-size: 18px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .sidebar li img {
    background: transparent;
  }

  .sidebar li:hover {
    background-color: #444;
  }

  .sidebar .selected {
    background-color: #162832;
  }

  /* основной блок */
  .main {
    grid-area: main;
    padding: 20px;
  }

  .main-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  .title {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1rem;
  }

  .main-header h2 {
    font-size: 3rem;
    margin: 0;
  }

  .info {
    color: #fff600;
    padding: 0.5rem 0.6rem;
    border-radius: 4px;
    font-size: 1.25rem;
    margin-left: 2rem;
    background-color: #534700;
  }

  .actions {
    display: flex;
    gap: 0.6rem;
  }

  .import {
    background-color: #444;
    border: none;
    padding: 10px 15px;
    font-size: 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.6rem;
  }

  .import img {
    background: transparent;
  }

  .create {
    background-color: #ff7b00;
    border: none;
    padding: 10px 15px;
    font-size: 1.1rem;
  }

  .new-folder {
    margin-top: 2em;
    background: transparent;
    border-width: 2px;
    border-color: #ff7b00;
    color: #ff7b00;
    padding: 0.5rem 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.6rem;
    font-weight: 600;
  }

  .folder-list {
    display: flex;
    gap: 15px;
  }

  .folder {
    background-color: #333;
    padding: 15px;
    flex: 1;
    text-align: left;
    cursor: pointer;
    border: 1px solid transparent;
    transition: border 0.2s;
  }

  .folder:hover {
    border: 1px solid #ff7b00;
  }

  /* блок пользователя */
  .user-panel {
    grid-area: user-panel;
    padding: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-bottom: 1px solid #444;
    border-right: 1px solid #444;
  }

  .user-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .user-info .avatar {
    margin-right: 3rem;
  }

  a {
    text-decoration: none;
    color: #fff;
  }
</style>

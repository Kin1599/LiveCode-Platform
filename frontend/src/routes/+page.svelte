<script>
  // @ts-nocheck

  let templates = [
    { name: "Python", author: "misplit", language: "python" },
    { name: "Hello world!", author: "misplit", language: "python" },
    { name: "Fibonachi", author: "misplit", language: "python" },
  ];

  let folders = [
    { name: "Shared with me", type: "shared", files: [] },
    { name: "Unnamed (1)", type: "folder", files: [] },
  ];

  let repls = 1; //для количества реп
  let selected = "Repls"; //для выбранного пункта меню
  let showBlocks = true; //для показа/скрытия бокового блока
  let searchQuery = ""; //для поиска
  let username = "username"; //для ника

  let file = {
    name: "QuintessentialDarkvioletCertifications",
    date: "5 days ago",
    size: "203.57 MiB",
    visibility: "Public",
  };

  /**
   * @type {null}
   */
  let openedFolder = null;

  // Для изменения состояния выбранного пункта меню
  /**
   * @param {string} item
   */
  function selectItem(item) {
    selected = item;
    openedFolder = null; // Закрыть папку при переключении на другой пункт
  }

  // Для переключения состояния видимости бокового блока
  function toggleVisibility() {
    showBlocks = !showBlocks;
  }

  // Для открытия папки
  /**
   * @param {any} folder
   */
  function openFolder(folder) {
    openedFolder = folder;
  }

  function getLanguageIcon(language){
    switch(language) {
      case "python":
        return "./images/python-icon.svg"
      case "javascript":
        return "./images/javascript-icon.svg"
      case "golang":
        return "./images/golang-icon.svg"
    }
  }
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
          <button class="create" on:click={() => selectItem("create-repl")}>+ Создать</button>
        </div>
      </div>

      <div style="margin-bottom: 20px; font-size: 1.25rem;">
        All {#if openedFolder}
          / {openedFolder.name}{/if}
      </div>

      {#if !openedFolder}
        <button class="new-folder" style="margin-bottom: 20px;"
          ><img src="./images/icon-new-folder.svg" alt="" />Новая папка</button
        >
        <!-- Список папок -->
        <div class="folder-list">
          {#each folders as folder}
            <div class="folder" on:click={() => openFolder(folder)}>
              {folder.name}
            </div>
          {/each}
        </div>
      {/if}

      <!-- Содержимое открытой папки -->
      {#if openedFolder}
        <div class="folder-contents">
          <button style="border: none;" on:click={() => (openedFolder = null)}>
            <img src="./images/icon-back.svg" alt="" /></button
          >
          <!-- <ul>
          {#each openedFolder.files as file}
            <li>{file}</li>
          {/each}
        </ul> -->

          <div class="folder-item">
            <div class="folder-info">
              <div class="folder-header">
                <img
                  src="./images/python-icon.svg"
                  alt="icon"
                  class="folder-icon"
                />
                <div class="folder-details">
                  <a class="folder-name" href="/code-input">{file.name}</a>
                  <span class="folder-meta"> {file.date} • {file.size}</span>
                </div>
              </div>
              <div class="folder-visibility">
                🌐 {file.visibility}
              </div>
            </div>
            <div class="folder-actions">
              <button class="more-btn">•••</button>
            </div>
          </div>
        </div>
      {/if}
    {:else if selected === "Настройки"}
      <h2>Настройки</h2>
      <p>Здесь будет контент для настроек...</p>
    {:else if selected === "create-repl"}
    <h2 class="create-title">Create a new Repl</h2>
    <div class="create-container">
      <div>
        <h3 class="templates-title">Templates</h3>
        <div class="search-container search-templates">
          <img src="./images/icon-search.svg" alt="" />
          <input
            class="search-input"
            type="text"
            placeholder="Search"
          />
        </div>
        <ul class="template-list">
          {#each templates as template}
          <li class="template-item">
            <img class="template-icon" src={getLanguageIcon(template.language)} alt={template.language}/>
            <div class="template-info">
              <p class="info-name">{template.name}</p>
              <p class="info-author">{template.author}</p>
            </div>
          </li>
          {/each}
        </ul>
      </div>
      <div class="create-params">
        <div class="create-input-label">
          <p class="create-label">Title</p>
          <input class="create-input" type="text" placeholder="Name">
        </div>
        <div class="params-public"> 
          <p class="create-label">Public</p>
          <div class="public-info">
            <div class="public-info-icon">
              <img src="./images/global.svg" alt="global">
            </div>
            <div class="public-info-text">
              <p class="text-title">Private join link</p>
              <p class="text-description">Anyone with this link can edit files</p>
            </div>
          </div>
        </div>
        <button class="createBtn">+ Создать</button>
      </div>
    </div>
    {/if}
  </main>

  <div class="user-panel {showBlocks ? '' : 'hidden'}">
    <div class="user-info">
      <div class="avatar"></div>
      <div><a href="/login">{username}</a></div>
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

  .search-templates {
    margin: 0px;
    width: 100%;
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

  .search-templates > input{
    font-size: 20px;
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

  .folder-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #162832; /* Цвет фона */
    padding: 15px 20px;
    border-radius: 5px;
    margin: 10px;
  }

  .folder-info {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    background: transparent;
  }

  .folder-header {
    display: flex;
    align-items: center;
    background: transparent;
    margin-bottom: 1rem;
  }

  .folder-icon {
    width: 24px;
    height: 24px;
    margin-right: 10px;
    background: transparent;
  }

  .folder-details {
    display: flex;
    align-items: center;
    font-size: 1rem;
    background: transparent;
  }

  .folder-name {
    font-weight: bold;
    background: transparent;
  }

  .folder-meta {
    font-size: 0.875rem;
    opacity: 0.7;
    margin-left: 8px;
    background: transparent;
  }

  .folder-visibility {
    font-size: 0.875rem;
    margin-top: 5px;
    opacity: 0.8;
    background: transparent;
  }

  .folder-actions {
    display: flex;
    align-items: center;
    gap: 10px;
    background: transparent;
  }

  .more-btn {
    background: none;
    border: none;
    cursor: pointer;
    opacity: 0.8;
  }

  .more-btn:hover {
    opacity: 1;
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

  /* Create REPL */

  .create-title{
    font-size: 48px;
    font-weight: 600;
    line-height: 56.25px;
    margin-bottom: 30px;
  }

  .create-container{
    display: flex;
    justify-content: space-between;
  }

  .templates-title{
    font-size: 20px;
    font-weight: 500;
    line-height: 23.44px;
    margin-bottom: 22px;
  }

  .template-list{
    margin-top: 5px;
  }

  .template-item{
    background-color: #162832;
    height: 60px;
    padding: 5px 10px;
    display: flex;
    align-items: center;
    gap: 10px;
    width: 100%;
    cursor: pointer;
  }

  .template-item:first-child{
    border-radius: 5px 5px 0 0;
  }

  .template-item:last-child{
    border-radius: 0 0 5px 5px;
  }

  .template-icon{
    width: 40px;
    height: 40px;
    background-color: transparent;
  }

  .template-info{
    background-color: transparent;
  }

  .template-info > *:not(:last-child){
    margin-bottom: 2px;
  }

  .info-name{
    font-size: 20px;
    line-height: 23.44px;
    background-color: transparent;
  }

  .info-author{
    color: #7E7E7E;
    background-color: transparent;
  }

  .create-input-label > *:not(:last-child){
    margin-bottom: 22px;
  }

  .create-label{
    font-size: 20px;
    line-height: 23.44px;
  }

  .create-input{
    background-color: #162832;
    border-radius: 5px;
    width: 100%;
    padding: 10px;
    border: 1px solid #6A6A6A66;
    font-size: 20px;
  }

  .params-public{
    margin: 63px 0 130px 0;
  }

  .public-info{
    margin-top: 29px;
    display: flex;
    align-items: center;
    gap: 24px; 
  }

  .public-info-icon{
    border-radius: 50%;
    height: 66px;
    width: 66px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #162832;
  }

  .public-info-icon > img{
    background-color: transparent;
  }

  .public-info-text > *:not(:last-child){
    margin-bottom: 5px;
  }

  .public-info-text > .text-title{
    color: #FFF;
    font-size: 24px;
    font-weight: 500;
  }

  .public-info-text > .text-description{
    color: #7E7E7E;
  }

  .createBtn{
    width: 100%;
    background-color: #ff7b00;
    border: none;
    padding: 10px 15px;
    font-size: 1.1rem;
  }
</style>

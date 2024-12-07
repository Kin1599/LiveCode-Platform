<script lang="ts">
  import CodeMirror from "svelte-codemirror-editor";
  // import { javascript } from "@codemirror/lang-javascript";
  import { python } from "@codemirror/lang-python";

  let value: string = ""; //для codemirror

  type Folders = {
    [key: string]: string[]; // Ключ — строка, значение — массив строк (файлов)
  };

  let files: string[] = ["file.txt", "main.py"];
  let folders: Folders = {}; // Инициализация папок
  let selectedFile: string = "main.py";
  let showMenu: boolean = false;

  const selectFile = (file: any) => {
    selectedFile = file;
  };

  const createFile = () => {
    const fileName = prompt("Введите имя файла:");
    if (fileName) {
      files.push(fileName);
    }
  };

  const createFolder = () => {
    const folderName = prompt("Введите имя папки:");
    if (folderName) {
      folders[folderName] = []; // Создаем новую папку с пустым массивом файлов
    }
  };

  const toggleMenu = () => {
    showMenu = !showMenu;
  };

  const closeMenu = () => {
    showMenu = false;
  };

  const handleMenuOption = (option: string) => {
    alert(`Selected option: ${option}`);
    closeMenu();
  };

  // для верхнего меню
  let projectName = "the name of project";
  let lastModified = "5 days ago";
  let size = "203.57 MB";
  let tabs = ["main.py", "script.js", "script.js", "script.js"];
  let searchQuery: string = "";

  let isSidebarVisible: boolean = true;

  // для переключения состояния видимости бокового блока
  function toggleSidebar(): void {
    isSidebarVisible = !isSidebarVisible;
  }
</script>

<div class="container">
  <div class="header">
    <div class="left-section">
      <div class="icons">
        <button on:click={toggleSidebar}
          ><img src="./images/icon-sidebar.svg" alt="" /></button
        >
        <a href="/#"><img src="./images/icon-home.svg" alt="" /></a>
      </div>
      <div style="margin-left: 1.3rem;">{projectName}</div>
      <div class="prj-settings">{lastModified}</div>
      <div class="prj-settings">{size}</div>
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
      <button class="run-button">Run</button>
      <button class="invite-button"
        ><img src="./images/icon-invite.svg" alt="" />Пригласить</button
      >
      <button><img src="./images/icon-plus.svg" alt="" /></button>
      <button><img src="./images/icon-noti.svg" alt="" /></button>
      <button><img src="./images/icon-more.svg" alt="" /></button>
      <div class="avatar"></div>
    </div>
  </div>

  <div class="content">
    <div class="sidebar" class:hidden={isSidebarVisible ? "" : "hidden"}>
      <div class="sidebar-header">
        <h3>Files</h3>
        <div>
          <button on:click={createFile}
            ><img src="./images/icon-file-add.svg" alt="" /></button
          >
          <button on:click={createFolder}
            ><img src="./images/icon-folder-add.svg" alt="" /></button
          >
          <button class="menu-button" on:click={toggleMenu}
            ><img src="./images/3-poits.svg" alt="" /></button
          >
          {#if showMenu}
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <div class="menu" on:mouseleave={closeMenu}>
              <button on:click={() => handleMenuOption("Option 1")}
                >Option 1</button
              >
              <button on:click={() => handleMenuOption("Option 2")}
                >Option 2</button
              >
              <button on:click={() => handleMenuOption("Option 3")}
                >Option 3</button
              >
            </div>
          {/if}
        </div>
      </div>

      <ul class="file-list">
        {#each files as file}
          <li
            class:selected={selectedFile === file}
            on:click={() => selectFile(file)}
          >
            {file}
          </li>
        {/each}

        {#each Object.keys(folders) as folder}
          <div class="folder">
            {folder}
            <ul class="folder-files">
              {#each folders[folder] as file}
                <li
                  class:selected={selectedFile === file}
                  on:click={() => selectFile(file)}
                >
                  {file}
                </li>
              {/each}
            </ul>
          </div>
        {/each}
      </ul>
    </div>
    <div class="main">
      <div class="tabs">
        {#each tabs as tab}
          <div class="tab">{tab}</div>
        {/each}
      </div>
      <div class="code-input">
        <CodeMirror
          bind:value
          lang={python()}
          styles={{
            "&": {
              maxWidth: "100%",
              height: "49.5rem",
            },
          }}
        />
      </div>
    </div>
  </div>
</div>

<style>
  .container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }

  /* верхний блок */
  .header {
    width: 100%;
    max-height: 5.5rem;
    text-align: center;
    padding: 20px;
    box-sizing: border-box;
    border-bottom: 1px solid #444;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.5rem 1rem;
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

  .prj-settings {
    color: #7e7e7e;
  }

  .search-container {
    display: flex;
    justify-content: center;
    background: transparent;
    border-radius: 10px;
    padding: 0.3rem;
    padding-left: 1.3rem;
    flex-grow: 1;
    max-width: 400px;
    border: 1px solid #7e7e7e;
    margin: 0.7rem 1rem;
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
    border-radius: 10px;
    border-color: transparent;
    cursor: pointer;
    color: #ffffff;
  }

  .run-button {
    padding: 0.7rem 1rem;
    background-color: #218b01;
    border: none;
    cursor: pointer;
    width: 5.5rem;
  }

  .run-button:hover {
    background-color: #1a5d05;
  }

  .run-button:active {
    background-color: #3d3d3d;
  }

  .invite-button {
    background-color: #162832;
    padding: 0.5rem 0.75rem;
    width: 8rem;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.5rem;
  }

  .invite-button:hover {
    background-color: #0b1419;
    background: transparent;
    border-width: 2px;
    border-color: #6a6a6a66;
  }

  .invite-button img {
    background: transparent;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: #444;
    border-radius: 50%;
  }

  .content {
    display: flex;
    flex: 1;
  }

  /* боковой блок */
  .sidebar {
    width: 17.5rem;
    padding: 20px;
    box-sizing: border-box;
    border-right: 1px solid #444;
  }

  .sidebar.hidden {
    width: 0;
    opacity: 0;
    overflow: hidden;
  }

  .sidebar h3 {
    font-size: 1.25rem;
  }

  .sidebar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    margin-right: 1.5rem;
  }

  .sidebar-header button {
    border-color: transparent;
    cursor: pointer;
  }

  .menu-button img {
    width: 15px;
    height: 15px;
  }

  .file-list {
    list-style: none;
    padding-right: 2rem;
  }

  .file-list li {
    padding: 8px;
    border-radius: 10px;
    cursor: pointer;
    font-size: 1rem;
  }

  .file-list li:hover {
    background-color: #333;
  }

  .file-list li:active {
    background-color: #444;
    color: #ff7b00;
  }

  .folder {
    margin-top: 10px;
    color: #ff7b00;
  }

  .folder-files {
    margin-left: 10px;
  }

  /* основной блок */
  .main {
    flex: 1;
    padding: 20px;
    box-sizing: border-box;
    background-color: #162832;
  }

  .tabs {
    display: flex;
    gap: 0.5rem;
  }

  .tab {
    padding: 0.25rem 0.5rem;
    background-color: #333;
    border: 1px solid #444;
    border-radius: 3px;
  }
</style>

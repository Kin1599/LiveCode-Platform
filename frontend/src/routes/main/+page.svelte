// не доделано

<script>
  let folders = [
    { name: "Shared with me", type: "shared", files: [] },
    { name: "Unnamed (1)", type: "folder", files: [] }
  ];


  //переменная для количества папок
  let repls = 1

  // Состояние для выбранного пункта меню
  let selected = 'Repls';
  // Функции для изменения состояния выбранного пункта меню
  /**
   * @param {string} item
   */
  function selectItem(item) {
    selected = item;
  }
</script>
  
<div class="layout">
  <header class="header">
    Верхний блок
  </header>


  <aside class="sidebar">
    <ul style="margin-top: 3rem;">
      <li class:selected={selected === 'Repls'} on:click={() => selectItem('Repls')}>Repls</li>
      <li class:selected={selected === 'Настройки'} on:click={() => selectItem('Настройки')}>Настройки</li>
    </ul>
  </aside>


  <main class="main" style="margin-left: 4rem; margin-top: 2.5rem; margin-right: 4rem">
    
      {#if selected === 'Repls'}
      <div class="main-header">
        <div style="display: flex; align-items: center;">
          <h2>Repls</h2>
          <span class="info">({repls}/3 Repls)</span>
        </div>
  
        <div class="actions">
          <button class="secondary">Импортировать из GitHub</button>
          <button>+ Создать</button>
        </div>
      </div>

      <!-- Действия -->
      <button class="new-folder" style="margin-bottom: 20px;">Новая папка</button>

      <!-- Список папок -->
      <div class="folder-list">
        {#each folders as folder}
          <div class="folder">
            {folder.name}
          </div>
        {/each}
      </div>
      {:else if selected === 'Настройки'}
        <h2>Настройки</h2>
        <p>Здесь будет контент для настроек...</p>
      {/if}
    

  </main>

  

  <div class="user-panel">
    Панель пользователя
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


  /* верхний блок */
  .header {
    grid-area: header;
    padding: 10px 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #444;
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

  .main-header h2 {
    font-size: 3rem;
    margin: 0;
  }

  .info {
    color: #FFF600;
    padding: 0.5rem 0.6rem;
    border-radius: 4px;
    font-size: 1.25rem;
    margin-left: 2rem;
    background-color: #534700;
  }

  .actions {
    display: flex;
    gap: 10px;
  }

  .actions button {
    background-color: #FF7B00;
    color: #fff;
    border: none;
    padding: 10px 15px;
    font-size: 1.2rem;
    border-radius: 10px;
    cursor: pointer;
  }

  .actions button.secondary {
    background-color: #444;
  }

  .new-folder {
    margin-top: 2em;
    background: transparent;
    border-width: 2px;
    border-color: #FF7B00;
    color: #FF7B00;
    padding: 0.5rem 2rem;
    border-radius: 10px;
    cursor: pointer;
  }

  .folder-list {
    display: flex;
    gap: 15px;
  }

  .folder {
    background-color: #333;
    padding: 15px;
    flex: 1;
    text-align: center;
    border-radius: 8px;
    cursor: pointer;
    border: 1px solid transparent;
    transition: border 0.2s;
  }

  .folder:hover {
    border: 1px solid #FF7B00;
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
</style>

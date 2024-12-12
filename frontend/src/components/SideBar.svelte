<script lang="ts">
    type Folders = {
      [key: string]: string[];
    };
    
    export let files: string[] = [];
    export let folders: Folders = {};
    export let selectedFile: string = "";
    export let showMenu: boolean = false;
    export let selectedFolder: string | null = null;
  
    const selectFile = (file: string) => {
      selectedFile = file;
      selectedFolder = null;
    };
  
    const createFile = () => {
      const fileName = prompt("Введите имя файла:");
      if (fileName) {
        if (selectedFolder) {
          // Если выбрана папка, добавляем файл в неё
          if (!folders[selectedFolder].includes(fileName)) {
            folders[selectedFolder] = [...folders[selectedFolder], fileName];
            selectedFile = fileName; // Устанавливаем файл как выбранный
          } else {
            alert("Файл с таким именем уже существует в этой папке.");
          }
        } else {
          // Если папка не выбрана, добавляем файл в общий список
          if (!files.includes(fileName)) {
            files = [...files, fileName];
            selectedFile = fileName; // Устанавливаем файл как выбранный
          } else {
            alert("Файл с таким именем уже существует.");
          }
        }
      }
    };
  
    const createFolder = () => {
      const folderName = prompt("Введите имя папки:");
      if (folderName) {
        if (!folders[folderName]) {
          folders[folderName] = []; // Инициализируем пустую папку
          selectedFolder = folderName;
        } else {
          alert("Папка с таким именем уже существует.");
        }
      }
    };

    // Выбор папки
    const selectFolder = (folderName: string) => {
      selectedFolder = selectedFolder === folderName ? null : folderName; // Переключение
      selectedFile = ""; // Сброс выбора файла
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
</script>
  
<div class="sidebar">
    <div class="sidebar-header">
        <h3>Files</h3>
        <div class="menu-btns">
            <button class="active-button" on:click={createFile}
              ><img src="../images/icon-file-add.svg" alt="" /></button
            >
            <button class="active-button" on:click={createFolder}
              ><img src="../images/icon-folder-add.svg" alt="" /></button
            >
            <button class="menu-button" on:click={toggleMenu}
              ><img src="../images/3-poits.svg" alt="" /></button
            >
            {#if showMenu}
              <!-- svelte-ignore a11y_no_static_element_interactions -->
              <div class="menu" on:mouseleave={closeMenu}>
                <button on:click={() => handleMenuOption("Option 1")}>Option 1</button>
                <button on:click={() => handleMenuOption("Option 2")}>Option 2</button>
                <button on:click={() => handleMenuOption("Option 3")}>Option 3</button>
              </div>
            {/if}
        </div>
    </div>
  
    <ul class="file-list">
        {#each files as file}
            <li
                class:selected={selectedFile === file}
            > <button on:click={() => selectFile(file)}>
              {file}
            </button>
                
            </li>
        {/each}
  
        {#each Object.keys(folders) as folder}
          <li class="folder">
            <div class:selected={selectedFolder === folder}>
              <button on:click={() => selectFolder(folder)}><strong>{folder}</strong></button>
            </div>
            {#if selectedFolder === folder}
              <ul class="folder-files">
                {#each folders[folder] as file}
                  <li class:selected={selectedFile === file}>
                    <button on:click={() => selectFile(file)}>{file}</button>
                  </li>
                {/each}
              </ul>
            {/if}
          </li>
        {/each}
      </ul>
</div>
  
<style>
    .sidebar {
        width: 17.5rem;
        padding: 20px;
        box-sizing: border-box;
        border-right: 1px solid #444;
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

    .menu-btns{
      display: flex;
      align-items: center;
      gap: 5px;
    }

    .active-button{
      background-color: #0B1419;
    }

    .menu-button img {
        width: 15px;
        height: 15px;
    }

    .menu-button {
      background-color: #0B1419;
    }

    .file-list {
        list-style: none;
        padding-right: 2rem;
    }

    .file-list li {
        padding: 8px;
        border-radius: 10px;
        cursor: pointer;
    }

    .file-list li:hover {
        background-color: #333;
    }
    
    .file-list li:active {
        background-color: #444;
        color: #ff7b00;
    }

    li button {
      background: none;
      border: none;
      color: #fff;
      font-size: 1rem;
      cursor: pointer;
    }
  
    .folder {
        margin-top: 10px;
        color: #ff7b00;
    }
  
    .folder-files {
        margin-left: 10px;
    }
</style>
    

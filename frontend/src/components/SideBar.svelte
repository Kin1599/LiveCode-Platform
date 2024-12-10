<script lang="ts">
    type Folders = {
      [key: string]: string[];
    };
    
    export let files: string[] = [];
    export let folders: Folders = {};
    export let selectedFile: string = "";
    export let showMenu: boolean = false;
  
    const selectFile = (file: string) => {
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
        folders[folderName] = [];
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
</script>
  
<div class="sidebar">
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
</style>
    
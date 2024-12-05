<script lang="ts">
    import CodeMirror from "svelte-codemirror-editor";
    // import { javascript } from "@codemirror/lang-javascript";
    import { python } from "@codemirror/lang-python";
   
    let value = ""; //для codemirror

    type Folders = {
        [key: string]: string[]; // Ключ — строка, значение — массив строк (файлов)
    };

    let files: string[] = ["file.txt", "main.py"]; // Список файлов
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
</script>


<div class="container">
    <div class="header">
        Верхний блок (Header)
    </div>
    <div class="content">
        <div class="sidebar">
            <div class="sidebar">
                <div class="sidebar-header">
                  <h3>Files</h3>
                  <div >
                    <button on:click={createFile}><img src="./images/icon-file-add.svg" alt=""></button>
                    <button on:click={createFolder}><img src="./images/icon-folder-add.svg" alt=""></button>
                    <button class="menu-button" on:click={toggleMenu}><img src="./images/3-poits.svg" alt=""></button>
                    {#if showMenu}
                      <!-- svelte-ignore a11y_no_static_element_interactions -->
                      <div class="menu" on:mouseleave={closeMenu}>
                        <div on:click={() => handleMenuOption("Option 1")}>Option 1</div>
                        <div on:click={() => handleMenuOption("Option 2")}>Option 2</div>
                        <div on:click={() => handleMenuOption("Option 3")}>Option 3</div>
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
        </div>
        <div class="main">
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

    .header {
        width: 100%;
        text-align: center;
        padding: 20px;
        box-sizing: border-box;
        border-bottom: 1px solid #444;
    }

    .content {
        display: flex;
        flex: 1;
    }

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

    .menu-button img{
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
        color: #FF7B00;
    }

    .folder {
        margin-top: 10px;
        color: #FF7B00;
        font-weight: bold;
    }

    .folder-files {
        margin-left: 10px;
    }

    .main {
        flex: 1;
        padding: 20px;
        box-sizing: border-box;
    }
</style>

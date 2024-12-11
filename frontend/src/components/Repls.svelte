<script lang="ts">
  import FolderContent from "../components/FolderContent.svelte";

  type Folder = {
    name: string;
    type: string;
    files: any[];  // Modify this to a more specific type if needed
  };

  export let folders: Folder[];
  export let openedFolder: Folder | null = null;
  export let repls: number;
  export let openFolder: (folder: Folder) => void;
  export let selectItem: (item: string) => void;
</script>
    
<div class="main-header">
  <div style="display: flex; align-items: center;">
    <div class="title">
      <img src="./images/icon-folder.svg" alt="" />
      <h2>Repls</h2>
    </div>
    <span class="info">({repls}/3 Repls)</span>
  </div>

  <div class="actions">
    <button class="import">
      <img src="./images/icon-gitnub.svg" alt="" />Импортировать из GitHub
    </button>
    <button class="create" on:click={() => selectItem("create-repl")}>+ Создать</button>
  </div>
</div>
  
<div style="margin-bottom: 20px; font-size: 1.25rem;">
  All {#if openedFolder} / {openedFolder.name}{/if}
</div>

{#if !openedFolder}
  <button class="new-folder" style="margin-bottom: 20px;">
    <img src="./images/icon-new-folder.svg" alt="" />Новая папка
  </button>
  <div class="folder-list">
    {#each folders as folder}
      <button class="folder" on:click={() => openFolder(folder)}>
        {folder.name}
      </button>
    {/each}
  </div>
{/if}

{#if openedFolder}
  <div class="file-list">
      {#each openedFolder.files as file}
      <FolderContent {file} onBack={() => (openedFolder = null)} />
      {/each}
  </div>
{/if}
  
<style>
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
    cursor: default;
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
    cursor: default;
  }

  .actions {
    display: flex;
    gap: 20px;
  }

  .import {
    background-color: #444;
    color: #fff;
    font-size: 20px;
    border-radius: 10px;
    border: none;
    padding: 10px 15px;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.6rem;
    height: 55px;
    cursor: pointer;
  }

  .import:hover {
    background-color: #7E7E7E;
  }

  .create {
    background-color: #ff7b00;
    border-style: solid;
    border-radius: 10px;
    color: #fff;
    font-size: 20px;
    padding: 10px 15px;
    cursor: pointer;
  }

  .create:hover {
    background: transparent;
    border-width: 2px;
    border-color: #ff7b00;
    color: #ff7b00;
  }

  .new-folder {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 5px;
    margin-top: 30px;
    background: transparent;
    border: 2px solid #ff7b00;
    border-radius: 10px;
    color: #ff7b00;
    padding: 0 36px; 
    font-weight: 600;
    font-size: 20px;
    height: 43px;
    cursor: pointer;
  }


  .new-folder:hover {
    color: #7d3c00;
    border: 2px solid #7d3c00;
  }
  
  .folder-list {
    display: flex;
    gap: 15px;
  }

  .folder {
    border: none;
    color: #fff;
    font-size: 20px;
    background-color: #333;
    padding: 16px;
    flex: 1;
    text-align: left;
    cursor: pointer;
    border: 1px solid transparent;
    transition: border 0.2s;
  }

  .folder:hover {
    border: 1px solid #ff7b00;
  }
</style>

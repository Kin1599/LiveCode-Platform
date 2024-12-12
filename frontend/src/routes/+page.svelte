<script lang="ts">
  import SendServer from "../api/api.js";

  import CreateRepl from "../components/CreateRepl.svelte";
  import HeaderMain from "../components/HeaderMain.svelte";
  import Repls from "../components/Repls.svelte";
  import Settings from "../components/Settings.svelte";
  import SideBarMain from "../components/SideBarMain.svelte";
  import { onMount } from "svelte";
  
  interface Folder {
    name: string;
    type: string;
    files: Array<any>;
  }

  let templates = [
    { name: "Python", author: "misplit", language: "python" },
    { name: "Hello world!", author: "misplit", language: "python" },
    { name: "Fibonachi", author: "misplit", language: "python" },
  ];

  let file = {
    name: "QuintessentialDarkvioletCertifications",
    date: "5 days ago",
    size: "203.57 MiB",
    visibility: "Public",
  };

  let folders = [
    { name: "Shared with me", type: "shared", files: [file] },
    { name: "Unnamed (1)", type: "folder", files: [] },
  ];

  let repls = 1; //для количества реп
  let selected = "Repls"; //для выбранного пункта меню
  let showBlocks = true; //для показа/скрытия бокового блока
  let searchQuery = ""; //для поиска
  let username = "username"; //для ника
  let userID = "";

  let openedFolder: Folder | null = null;

  async function fetchUserInfo() {
    try {
      const token = localStorage.getItem("token");
      if (token){
        const response = await SendServer.getUserInfo(token);
        username = response.Nickname || "Guest";
        userID = response.ID;
      }
    } catch (error) {
      console.error("Error fetching user info:", error);
    }
  }

  onMount(() => {
    fetchUserInfo();
  });

  function selectItem(item: string) {
    selected = item;
    openedFolder = null;
  }

  // Для переключения состояния видимости бокового блока
  function toggleVisibility() {
    showBlocks = !showBlocks;
  }

  // Для открытия папки
  function openFolder(folder: Folder) {
    openedFolder = folder;
  }

  function getLanguageIcon(language: string){
    switch(language) {
      case "python":
        return "./images/python-icon.svg"
      case "javascript":
        return "./images/javascript-icon.svg"
      case "golang":
        return "./images/golang-icon.svg"
      default:
        return "./images/python-icon.svg"
    }
  }


  // для создания новой папки
  function createNewFolder(name: string) {
    let newFolderName = name || "Unnamed";
    let folderExists = folders.some(folder => folder.name === newFolderName);
    if (folderExists) {
      let counter = 1;
      while (folders.some(folder => folder.name === `${newFolderName} (${counter})`)) {
        counter++;
      }
      newFolderName = `${newFolderName} (${counter})`;
    }
    folders.push({ name: newFolderName, type: "folder", files: [] });
    openFolder({ name: newFolderName, type: "folder", files: [] });
  }

  async function createNewSession(sessionData: {
    owner_id: string;
    editable: boolean;
    title: string;
    language: string;
    max_users: number;
  }) {
    try {
      const response = await SendServer.createSession(sessionData);
      if (response) {
        repls += 1;
        console.log("Repl created successfully");

        let unnamedFolder = folders.find(folder => folder.name === "Unnamed (1)");
        if (!unnamedFolder) {
          unnamedFolder = { name: "Unnamed (1)", type: "folder", files: [] };
          folders.push(unnamedFolder);
        }

        unnamedFolder.files.push({
            name: sessionData.title,
            date: new Date().toLocaleDateString(),
            size: "0 B",
            visibility: "Public",
          });

        window.location.assign(response.url);
      } else {
        console.error("Error creating repl:", response.message);
      }
    } catch (error) {
      console.error("Error creating session:", error);
    }
  }
</script>

<div class="layout">
  <HeaderMain {searchQuery} {toggleVisibility}/>

  <SideBarMain {selected} {selectItem} {showBlocks}/>

  <main
    class="main"
    class:expanded={!showBlocks}
    style="margin-left: 4rem; margin-top: 2.5rem; margin-right: 4rem"
  >
    {#if selected === "Repls"}
      <Repls {folders} {openedFolder} {openFolder} {repls} {selectItem} {createNewFolder}/>
    {:else if selected === "Настройки"}
      <Settings />
    {:else if selected === "create-repl"}
      <CreateRepl {templates} {getLanguageIcon} {createNewSession} owner_id={userID}/>
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

  .hidden {
    width: 0;
    opacity: 0;
    overflow: hidden;
  }

  .expanded {
    width: 100%;
    transform: translateX(-12rem);
  }  


  /* основной блок */
  .main {
    grid-area: main;
    padding: 20px;
  }

  
  /* блок пользователя */
  .user-panel {
    grid-area: user-panel;
    padding: 20px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #444;
    border-right: 1px solid #444;
  }

  .user-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 10px;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: #444;
    border-radius: 50%;
  }

  a {
    font-size: 20px;
    font-weight: 300px;
    text-decoration: none;
    color: #fff;
  }

</style>

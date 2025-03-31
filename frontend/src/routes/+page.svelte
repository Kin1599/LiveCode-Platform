<script lang="ts">
  import { onMount } from "svelte";
  import SendServer from "../api/api.js";

  import CreateRepl from "../components/CreateRepl.svelte";
  import HeaderMain from "../components/HeaderMain.svelte";
  import Repls from "../components/Repls.svelte";
  import Settings from "../components/Settings.svelte";
  import SideBarMain from "../components/SideBarMain.svelte";
  
  interface FileItem {
    name: string;
    date: string;
    size: string;
    visibility: string;
  }

  interface Folder {
    name: string;
    type: string;
    files: FileItem[];
  }

  let templates = [
    { name: "Python", author: "misplit", language: "python" },
    { name: "Hello world!", author: "misplit", language: "python" },
    { name: "Fibonachi", author: "misplit", language: "python" },
  ];

  let folders: Folder[] = [
    { name: "Shared with me", type: "shared", files: [] },
    { name: "Unnamed (1)", type: "folder", files: [] },
  ];

  let repls = 1; 
  let selected = "Repls"; 
  let showBlocks = true; 
  let searchQuery = ""; 
  let username = "Guest"; 
  let userID = "";

  let openedFolder: Folder | null = null;

  async function fetchUserInfo() {
    try {
      const token = localStorage.getItem("token");
      if (token){
        const response = await SendServer.getUserInfo(token);
        username = response.Nickname || "Guest";
        userID = response?.ID || "";
      }
    } catch (error) {
      console.error("Error fetching user info:", error);
    }
  }

  onMount(fetchUserInfo);

  function selectItem(item: typeof selected) {
    selected = item;
    openedFolder = null;
  }

  function toggleVisibility() {
    showBlocks = !showBlocks;
  }

  function openFolder(folder: Folder) {
    openedFolder = folder;
  }

  function getLanguageIcon(language: string){
    const icons: Record<string, string> = {
      python: "./images/python-icon.svg",
      javascript: "./images/javascript-icon.svg",
      golang: "./images/golang-icon.svg",
    };
    return icons[language] || icons.python;
  }

  function createNewFolder(name: string) {
    let newFolderName = name || "Unnamed";
    let counter = 1;

    while (folders.some(folder => folder.name === newFolderName)) {
      newFolderName = `${name} (${counter++})`;
    }

    const newFolder = { name: newFolderName, type: "folder", files: [] };
    folders = [...folders, newFolder];
    openFolder(newFolder);
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
        repls++;
        console.log("Repl created successfully");

        const unnamedFolder = folders.find(
          folder => folder.name === "Unnamed (1)") || 
          { name: "Unnamed (1)", type: "folder", files: [] };
        
        unnamedFolder.files.push({
          name: sessionData.title,
          date: new Date().toLocaleDateString(),
          size: "0 B",
          visibility: "Public",
        });

        folders = [...folders];
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
  >
    {#if selected === "Repls"}
      <Repls {folders} {openedFolder} {openFolder} {repls} {selectItem} {createNewFolder}/>
    {:else if selected === "Настройки"}
      <Settings />
    {:else if selected === "create-repl"}
        {#if userID}
          <CreateRepl {templates} {getLanguageIcon} {createNewSession} owner_id={userID}/>
        {:else}
          <p>Loading...</p>
        {/if}
    {/if}
  </main>

  <div class="user-panel" class:hidden={!showBlocks}>
    <div class="user-info">
      <div class="avatar"></div>
      <a href="/login">{username}</a>
    </div>
  </div>
</div>

<style lang="scss">
  @use "../styles/colors.scss" as *;
  @use "../styles/fonts.scss" as *;

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

  .main {
    grid-area: main;
    padding: 20px;
  }
  
  .user-panel {
    grid-area: user-panel;
    padding: 20px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid $border-color;
    border-right: 1px solid $border-color;
  }

  .user-info {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: $avatar-bg;
    border-radius: 50%;
  }

  a {
    font-size: 20px;
    font-weight: 300px;
    text-decoration: none;
    color: $text-color;
  }

</style>

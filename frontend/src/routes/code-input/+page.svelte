<script lang="ts">
  import HeaderCode from "../../components/HeaderCode.svelte";
  import TextEditor from "../../components/CodeWindow.svelte";
  import SideBar from "../../components/SideBar.svelte";
  import Tabs from "../../components/Tabs.svelte";

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

  async function getNickname(): Promise<string> {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve("Стандартное Имя");
      }, 1000);
    });
  }

  let ws: WebSocket; // WebSocket-соединение
  let userId: string = generateUserId();
  let userNickname: string = "";
  const userColor = generateColor();

  function generateUserId(): string {
    return Math.random().toString(36).substring(2) + Date.now().toString(36);
  }

  function generateColor(): string {
    return `#${Math.floor(Math.random() * 16777215).toString(16)}`;
  }

  function connect() {
    ws = new WebSocket("ws://217.114.2.64/ws");

    ws.onopen = () => {
      console.log("WebSocket соединение установлено");
      ws.send(
        JSON.stringify({
          type: "init",
          userId,
          color: userColor,
          nickname: userNickname,
        })
      );
    };

    ws.onmessage = (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      if (data.userId !== userId && data.type === "update") {
        value = data.text; // обновление значения при получении данных
      }
    };

    ws.onclose = () => {
      console.log("Соединение закрыто, переподключение...");
      setTimeout(connect, 1000);
    };

    ws.onerror = (error) => {
      console.error("Ошибка WebSocket:", error);
    };
  }

  // Отправка изменений на сервер при изменении текста
  $: {
    if (ws?.readyState === WebSocket.OPEN) {
      const message = {
        type: "update",
        text: value,
        userId,
        color: userColor,
      };
      ws.send(JSON.stringify(message));
    }
  }

  getNickname().then((nickname) => {
    userNickname = nickname;
    connect(); // Подключение к WebSocket только после получения ника
  });
</script>

<div class="container">
  <HeaderCode
    {projectName}
    {lastModified}
    {size}
    {searchQuery}
    {toggleSidebar}
  />

  <div class="content">
    <div class="sidebar {isSidebarVisible ? '' : 'hidden'}">
      <SideBar
        {files}
        {folders}
        bind:selectedFile
        bind:showMenu
      />
    </div>
    
    <div class="main">
      <Tabs {tabs} />
      <div class="code-input">
        <TextEditor bind:value />
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
    transition: width 0.3s, opacity 0.3s;
  }

  .sidebar.hidden {
    width: 0;
    opacity: 0;
    overflow: hidden;
  }


  /* основной блок */
  .main {
    flex: 1;
    padding: 20px;
    background-color: #162832;
  }
</style>

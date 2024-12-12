<script lang="ts">
  import HeaderCode from "../../../components/HeaderCode.svelte";
  import TextEditor from "../../../components/CodeWindow.svelte";
  import SideBar from "../../../components/SideBar.svelte";
  import Tabs from "../../../components/Tabs.svelte";
  import ChatWindow from "../../../components/ChatWindow.svelte";
  import { page } from "$app/stores";
  import { onMount } from "svelte";
    import SendServer from "../../../api/api";

  $: sessionId = $page.params.sessionId;

  let value: string = ""; //для codemirror
  let messages: { user: string, text: string }[] = []; 
  let currentMessage: string = ""; // для чата

  type Folders = {
    [key: string]: string[];
  };

  let files: string[] = ["main.py", "script.js", "index.html"];
  let folders: Folders = {};
  let selectedFile: string = "main.py";
  let showMenu: boolean = false;

  const closeMenu = () => {
    showMenu = false;
  };

  const handleMenuOption = (option: string) => {
    alert(`Selected option: ${option}`);
    closeMenu();
  };

  // для верхнего меню
  let projectName = "the name of project";
  let language = "python";
  let lastModified = "5 days ago";
  let size = "203.57 MB";
  let tabs = [
    {
      section: "1",
      tabs: [
        { id: 1, name: "main.py" },
        { id: 2, name: "script.js" },
        { id: 3, name: "index.html" },
      ],
      activeTab: 1
    },
    {
      section: "2",
      tabs: [
        { id: 1, name: "Chat" },
      ],
      activeTab: 1
    },
  ];
  let searchQuery: string = "";

  let isSidebarVisible: boolean = true;

  async function getSessionInfo(sessionId: string) {
    try {
      const response = await SendServer.getSessionInfo(sessionId);
      console.log(response);
      if (response) {
        projectName = response.Title;
        language = response.Language;
        lastModified = response.CreatedAt;
      } 
    } catch (error) {
      console.error("Error fetching session info:", error);
    }
  }

  onMount(() => {
    getSessionInfo(sessionId);
  });

  // для переключения состояния видимости бокового блока
  function toggleSidebar(): void {
    isSidebarVisible = !isSidebarVisible;
  }

  const handleTabClick = (section: string, event: CustomEvent<{ tabId: number }>) => {
    const tabSection = tabs.find(tab => tab.section === section);
    if (tabSection){
      tabSection.activeTab = event.detail.tabId; 
    }    
  };

  async function getNickname(): Promise<string> {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve("Стандартное Имя");
      }, 1000);
    });
  }

  let ws: WebSocket; // WebSocket-соединение
  let chatWs: WebSocket;
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
    ws = new WebSocket(`ws://217.114.2.64/ws?session_id=${sessionId}`);

    ws.onopen = () => {
      console.log("Connected to WebSocket server with session_id:", sessionId);
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
    connectChat();
  });

  function connectChat(){
    chatWs = new WebSocket(`ws://217.114.2.64/chat?userId=${userId}`);

    chatWs.onopen = () => {
      console.log("Websocket соединение для чата установлено");
    };

    chatWs.onmessage = (event: MessageEvent) => {
      const data = JSON.parse(event.data);
      messages = [...messages, { user: data.user, text: data.text }];
    };

    chatWs.onclose = () => {
      console.log("Соединение для чата закрыто, переподключение...");
      setTimeout(connectChat, 1000);
    };

    chatWs.onerror = (error) => {
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
      }; 
      ws.send(JSON.stringify(message)); 
    } 
  }

  function sendMessage(message: string) { 
    if (chatWs?.readyState === WebSocket.OPEN) { 
      const chatMessage = { 
        user: userNickname, 
        text: message 
      }; 
    chatWs.send(JSON.stringify(chatMessage)); 
    messages = [...messages, chatMessage]; 
    currentMessage = ""; 
    } 
  }
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
      <div class="code-section">
        <div class="tabs">
          <Tabs tabs={tabs[0].tabs} activeTab={tabs[0].activeTab} onTabClick={(event) => handleTabClick("1", event)} /> 
        </div>
        <div class="code-window">
          <TextEditor bind:value />
        </div>
      </div>  
      <div class="chat-section">
        <div class="tabs">
          <Tabs tabs={tabs[1].tabs} activeTab={tabs[1].activeTab} onTabClick={(event) => handleTabClick("2", event)} /> 
        </div>
        <div class="chat-window">
          <ChatWindow {messages} {currentMessage} {sendMessage} />
        </div>
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
    display: flex;
    background-color: #6A6A6A66;
  }

  .code-section{
    display: flex;
    flex-direction: column;
    border-radius: 0px 10px 0px 0px;
    flex: 1;
  }

  .code-window {
    flex: 1;
    box-sizing: border-box;
  }

  .chat-section{
    display: flex;
    flex-direction: column;
    border-radius: 10px 0px 0px 0px;
    flex: 1;
    overflow: hidden;
  }

  .chat-window{
    flex: 1;
    box-sizing: border-box;
    overflow-y: auto;
  }
</style>

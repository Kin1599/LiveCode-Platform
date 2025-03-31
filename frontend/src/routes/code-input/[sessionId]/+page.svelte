<script lang="ts">
  import { onMount } from "svelte";
  import { page } from "$app/stores";
  import HeaderCode from "../../../components/HeaderCode.svelte";
  import TextEditor from "../../../components/CodeWindow.svelte";
  import SideBar from "../../../components/SideBar.svelte";
  import Tabs from "../../../components/Tabs.svelte";
  import ChatWindow from "../../../components/ChatWindow.svelte";
  import SendServer from "../../../api/api.js";

  interface Message {
    user: string;
    text: string;
  }

  interface Tab {
    id: number;
    name: string;
  }

  interface TabSection {
    section: string;
    tabs: Tab[];
    activeTab: number;
  }

  interface Folders {
    [key: string]: string[];
  }

  $: sessionId = $page.params.sessionId;

  let value: string = ""; //для codemirror
  let messages: Message[] = []; 
  let currentMessage: string = ""; // для чата
  let files: string[] = ["main.py", "script.js", "index.html"];
  let folders: Folders = {};
  let selectedFile: string = "main.py";
  let showMenu: boolean = false;

  // для верхнего меню
  let projectName: string = "the name of project";
  let language: string = "python";
  let lastModified: string = "5 days ago";
  let size: string = "203.57 MB";
  let tabs: TabSection[] = [
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

  // WebSocket
  let ws: WebSocket | null = null; // WebSocket-соединение
  let chatWs: WebSocket | null = null;
  let userNickname: string = "";

  const generateUserId = (): string => {
    return Math.random().toString(36).substring(2) + Date.now().toString(36);
  }

  const generateColor = (): string => {
    return `#${Math.floor(Math.random() * 16777215).toString(16)}`;
  }

  const userId: string = generateUserId();
  const userColor: string = generateColor();

  const closeMenu = (): void => {
    showMenu = false;
  };

  const handleMenuOption = (option: string): void => {
    alert(`Selected option: ${option}`);
    closeMenu();
  };

  // для переключения состояния видимости бокового блока
  const toggleSidebar = (): void => {
    isSidebarVisible = !isSidebarVisible;
  }

  const handleTabClick = (section: string, event: CustomEvent<{ tabId: number }>): void => {
    const tabSection = tabs.find(tab => tab.section === section);
    if (tabSection){
      tabSection.activeTab = event.detail.tabId; 
    }    
  };

  const getNickname = (): Promise<string> =>
    new Promise((resolve) => setTimeout(() => resolve("Стандартное Имя"), 1000));

  const getSessionInfo = async (sessionId: string): Promise<void> => {
    try {
      const response = await SendServer.getSessionInfo(sessionId);
      if (response) {
        projectName = response.Title;
        language = response.Language;
        lastModified = response.CreatedAt;
      } 
    } catch (error) {
      console.error("Error fetching session info:", error);
    }
  }

  const connect = (): void => {
    ws = new WebSocket(`ws://localhost:8080/ws?session_id=${sessionId}`);

    ws.onopen = () => {
      console.log("Connected to WebSocket server with session_id:", sessionId);
      ws?.send(
        JSON.stringify({
          type: "init",
          userId: userId,
          nickname: userNickname,
          sessionId: sessionId,
        })
      );
    };

    ws.onmessage = (event: MessageEvent) => {
      try {
        const data = JSON.parse(event.data);
        console.log("Raw message on /ws:", event.data);
        console.log("Parsed message on /ws:", data);
        if (data.type === "history") {
          console.log("History received on /ws:", data.history || "No history field");
          // Можно обработать историю, если нужно
        } else if (data.type === "update" && data.userId !== userId) {
          value = data.text;
        }
      } catch (error) {
        console.error("Error parsing message on /ws:", error);
        ws?.close(1000, "Error parsing message");
      }
    };

    ws.onclose = (event: CloseEvent) => {
      console.log("Соединение /ws закрыто:", event.code, event.reason, "wasClean:", event.wasClean);
      setTimeout(connect, 1000);
    };

    ws.onerror = (error) => {
      console.error("Ошибка WebSocket /ws:", error);
    };
  }

  const connectChat = (): void => {
    chatWs = new WebSocket(`ws://localhost:8080/chat?session_id=${sessionId}`);

    chatWs.onopen = () => {
      console.log("Websocket соединение для чата установлено");
      chatWs?.send(
        JSON.stringify({
          type: "init",
          userId: userId,
          nickname: userNickname,
          sessionId: sessionId,
        })
      );
    };

    chatWs.onmessage = (event: MessageEvent) => {
      try {
        const data = JSON.parse(event.data);
        if (data.type === "history") {
          console.log("History received on /chat:", data.history || "No history field");
        } else if (data.type === "chat") {
          if (data.userId !== userId) {
            const userName = data.nickname || data.userId || "Unknown";
            messages = [...messages, { user: userName, text: data.text }];
          }
        }
      } catch (error) {
        console.error("Error parsing message on /chat:", error);
        chatWs?.close(1000, "Error parsing message"); 
      }
    };

    chatWs.onclose = (event: CloseEvent) => {
      console.log("Соединение /chat закрыто:", event.code, event.reason, "wasClean:", event.wasClean);
      setTimeout(connectChat, 1000);
    };

    chatWs.onerror = (error) => {
      console.error("Ошибка WebSocket /chat:", error);
    };
  }

  const sendMessage = (message: string): void => {
    if (chatWs?.readyState === WebSocket.OPEN) {
      const chatMessage = { 
        type: "chat",
        userId: userId,
        user: userNickname, 
        text: message 
      };
      chatWs.send(JSON.stringify(chatMessage));
      messages = [...messages, { user: userNickname, text: message }];
      currentMessage = "";
    }
  };

  // Отправка изменений на сервер при изменении текста 
  $: { 
    if (ws?.readyState === WebSocket.OPEN) { 
      const message = { 
        type: "update",
        text: value, 
        userId: userId,
        color: userColor, 
      }; 
      ws.send(JSON.stringify(message)); 
    } 
  }

  onMount(async () => {
    await getSessionInfo(sessionId);
    userNickname = await getNickname();
    connect();
    connectChat(); 
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
    <div class="sidebar" class:hidden={!isSidebarVisible}>
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
          <Tabs 
            tabs={tabs[0].tabs} 
            activeTab={tabs[0].activeTab} 
            onTabClick={(event) => handleTabClick("1", event)} 
          /> 
        </div>
        <div class="code-window">
          <TextEditor bind:value />
        </div>
      </div>  
      <div class="chat-section">
        <div class="tabs">
          <Tabs 
            tabs={tabs[1].tabs} 
            activeTab={tabs[1].activeTab} 
            onTabClick={(event) => handleTabClick("2", event)} 
          /> 
        </div>
        <div class="chat-window">
          <ChatWindow {messages} {currentMessage} {sendMessage} />
        </div>
      </div>          
    </div>
  </div>
</div>

<style lang="scss">
  @use "../../../styles/colors.scss" as *; 

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
    border-right: 1px solid $border-color;
    transition: width 0.3s, opacity 0.3s;
  
    &.hidden {
      width: 0;
      opacity: 0;
      overflow: hidden;
    }
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
    border-radius: 0 10px 0 0;
    flex: 1;
  }

  .code-window {
    flex: 1;
    box-sizing: border-box;
  }

  .chat-section{
    display: flex;
    flex-direction: column;
    border-radius: 10px 0 0 0;
    flex: 1;
    overflow: hidden;
  }

  .chat-window{
    flex: 1;
    box-sizing: border-box;
    overflow-y: auto;
  }
</style>

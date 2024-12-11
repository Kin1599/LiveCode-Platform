<!-- Sidebar.svelte -->
<script>
    let { 
        selected, 
        selectItem,
        showBlocks
    } = $props();
    
</script>
  
<aside class="sidebar {showBlocks ? '' : 'hidden'}">
    <div class="sidebar-list">
        <button
            class:selected={selected === "Repls"}
            onclick={() => selectItem("Repls")}
        >
            <img src="./images/folder-orange.svg" alt="" />Repls
        </button>
        <button
            class:selected={selected === "Настройки"}
            onclick={() => selectItem("Настройки")}
        >
            <img src="./images/icon-settings.svg" alt="" />Настройки
        </button>
    </div>
</aside>

<style>
    .sidebar {
        grid-area: sidebar;
        padding-top: 60px;
        border-right: 1px solid #444;
    }

    .hidden {
        width: 0;
        opacity: 0;
        overflow: hidden;
    }

    .sidebar-list{
        display: flex;
        flex-direction: column;
        padding: 0;
    }

    .sidebar button {
        background-color: transparent;
        border: none;
        color: #fff;
        padding: 15px;
        font-size: 20px;
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .sidebar button:hover {
        background-color: #444;
    }

    .sidebar .selected {
        background-color: #162832;
    }
</style>
  
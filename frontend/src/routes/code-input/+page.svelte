<script lang="ts">
    import CodeMirror from "svelte-codemirror-editor";
    import { javascript } from "@codemirror/lang-javascript";
    import { python } from "@codemirror/lang-python";
   

    let value = "";
</script>


<div class="container">
    <div class="header">
        Верхний блок (Header)
    </div>
    <div class="content">
        <div class="sidebar">
            Боковой блок (Sidebar)
        </div>
        <div class="main">
            Основной блок (Main Content)
            <CodeMirror bind:value lang={python()} />

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
        flex: 1; /* Контент заполняет оставшееся пространство */
    }

    .sidebar {
        width: 17.5rem;
        padding: 20px;
        box-sizing: border-box;
        border-right: 1px solid #444;
    }

    .main {
        flex: 1; /* Основной блок растягивается */
        padding: 20px;
        box-sizing: border-box;
    }
</style>

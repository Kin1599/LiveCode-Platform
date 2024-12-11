<script>
    export let file;
    export let onBack;
</script>
  
<div class="folder-contents">
    <button class="back-btn" on:click={onBack}>
        <img src="./images/icon-back.svg" alt="" />
    </button>
    <a class="folder-item" href="/code-input">
        <div class="folder-info">
            <div class="folder-header">
                <img src="./images/python-icon.svg" alt="icon" class="folder-icon" />
                <div class="folder-details">
                    <span class="folder-name">{file.name}</span>
                    <span class="folder-meta"> {file.date} • {file.size}</span>
                </div>
            </div>
            <div class="folder-visibility">
              🌐 {file.visibility}
            </div>
        </div>
        <div class="folder-actions">
            <button class="more-btn">•••</button>
        </div>
    </a>
</div>
  
<style>
    .back-btn{
        background-color: transparent;
        border: none;
        cursor: pointer;
        margin-bottom: 40px;
    }

    .folder-contents {
        padding: 15px;
    }

    .folder-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        background-color: #162832;
        padding: 15px 20px;
        border-radius: 5px;
        margin: 10px;
        text-decoration: none;
        color: #fff;
    }

    .folder-info {
        display: flex;
        flex-direction: column;
        flex-grow: 1;
        font-size: 20px;
    }

    .folder-info > *:not(:last-child) {
        margin-bottom: 20px;
    }

    .folder-header {
        display: flex;
        align-items: center;
        background: transparent;
        margin-bottom: 1rem;
    }

    .folder-icon {
        width: 24px;
        height: 24px;
        margin-right: 10px;
        background: transparent;
    }

    .folder-details {
        display: flex;
        align-items: center;
        font-size: 1rem;
        background: transparent;
    }

    .folder-name {
        font-weight: bold;
        background: transparent;
    }

    .folder-meta {
        opacity: 0.7;
        margin-left: 8px;
        background: transparent;
    }

    .folder-visibility {
        opacity: 0.8;
        background: transparent;
    }

    .folder-actions {
        display: flex;
        align-items: center;
        gap: 10px;
        background: transparent;
    }

    .more-btn {
        background: none;
        border: none;
        cursor: pointer;
        opacity: 0.8;
    }

    .more-btn:hover {
        opacity: 1;
    }
</style>

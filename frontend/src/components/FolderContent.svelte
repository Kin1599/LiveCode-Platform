<script>
    export let file;
    export let onBack;

    let isPopupVisible = false;

    function togglePopup() {
    isPopupVisible = !isPopupVisible;
    }

    /**
   * @param {string} action
   */
    function handleAction(action) {
    console.log(`Action selected: ${action}`);
    isPopupVisible = false;
    }
</script>
  
<div class="folder-contents">
    <button class="back-btn" on:click={onBack}>
        <img src="./images/icon-back.svg" alt="" />
    </button>
    <div class="folder-item">
        <div class="folder-info">
            <div class="folder-header">
                <img src="./images/python-icon.svg" alt="icon" class="folder-icon" />
                <div class="folder-details">
                    <a href="/code-input"><span class="folder-name">{file.name}</span></a>
                    <span class="folder-meta"> {file.date} • {file.size}</span>
                </div>
            </div>
            <div class="folder-visibility">
              🌐 {file.visibility}
            </div>
        </div>
        <div class="folder-actions">
            <button class="more-btn" on:click={togglePopup}><img src="./images/icon-points.svg" alt=""></button>
        </div>
    </div>
</div>

<!-- Попап меню -->
{#if isPopupVisible}
  <div class="popup-menu">
    <div class="popup-list">
      <button class="popup-item" on:click={() => handleAction('upload-file')}>Edit</button>
      <button class="popup-item" on:click={() => handleAction('upload-folder')}>Cover page</button>
      <button class="popup-item" on:click={() => handleAction('open-pane')}>Fork</button>
      <button class="popup-item" on:click={() => handleAction('download-zip')}>Move</button>
      <button class="popup-item" on:click={() => handleAction('show-hidden')}>Select</button>
      <button class="popup-item" on:click={() => handleAction('show-hidden')}>Delete</button>
    </div>
  </div>
{/if}
  
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

    .folder-details a {
        text-decoration: none;
        color: #fff;
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

  .popup-menu {
    position: absolute;
    top: 100%;
    right: 0;
    background-color: #162832;
    border: 1px solid #2f3e47;
    border-radius: 5px;
    padding: 10px;
    width: 10rem;
    z-index: 10;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }

  .popup-list {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .popup-item {
    background: none;
    border: none;
    padding: 8px 10px;
    cursor: pointer;
    border-radius: 3px;
    text-align: left;
    color: #fff;
    transition: background 0.2s;
  }

  .popup-item:hover {
    background-color: #2f3e47;
  }
</style>

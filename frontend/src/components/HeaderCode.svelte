<script lang="ts">
    import UserInvitations from "./UserInvitations.svelte";

    let showInvitation = $state(false);

    let { 
        projectName = "Project Name",
        lastModified = "Last Modified",
        size = "Size",
        searchQuery = "",
        toggleSidebar = () => {}
    } = $props();

    function toggleInvitation() {
      showInvitation = !showInvitation;
    }

</script>
  
<div class="header">
    <div class="left-section">
        <div class="icons">
            <button onclick={() => toggleSidebar()} >
                <img src="../images/icon-sidebar.svg" alt="Toggle Sidebar" />
            </button>
            <a href="/#"><img src="../images/icon-home.svg" alt="Home" /></a>
        </div>
        <div style="margin-left: 1.3rem;">{projectName}</div>
        <div class="prj-settings">{lastModified}</div>
        <div class="prj-settings">{size}</div>
    </div>

    <div class="search-container">
        <img src="../images/icon-search.svg" alt="Search Icon" />
        <input
            class="search-input"
            type="text"
            placeholder="Найти команду..."
            bind:value={searchQuery}
        />
    </div>

    <div class="right-section">
        <button class="run-button">Run</button>
        <button class="invite-button" onclick={toggleInvitation}>
            <img src="../images/icon-invite.svg" alt="Invite Icon" />Пригласить
        </button>
        <button class="active-button">
            <img src="../images/icon-plus.svg" alt="Add Icon" />
        </button>
        <button class="active-button">
            <img src="../images/icon-noti.svg" alt="Notifications Icon" />
        </button>
        <button class="active-button">
            <img src="../images/icon-more.svg" alt="More Options Icon" />
        </button>
        <div class="avatar"></div>
    </div>
    {#if showInvitation}
      <div class="modal">
        <div class="modal-content">
          <UserInvitations />
        </div>
      </div>
    {/if}
</div>

<style>
  .header {
    width: 100%;
    max-height: 5.5rem;
    text-align: center;
    padding: 20px;
    box-sizing: border-box;
    border-bottom: 1px solid #444;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0.5rem 1rem;
  }

  .left-section {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .icons {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 1.2rem;
  }

  .icons button {
    background-color: #0b1419;
    border-color: transparent;
    cursor: pointer;
  }

  .icons img {
    width: 1.5rem;
    height: 1.5rem;
  }

  .prj-settings {
    color: #7e7e7e;
  }

  .search-container {
    display: flex;
    justify-content: center;
    background: transparent;
    border-radius: 10px;
    padding: 0.3rem;
    padding-left: 1.3rem;
    flex-grow: 1;
    max-width: 400px;
    border: 1px solid #7e7e7e;
    margin: 0.7rem 1rem;
  }

  .search-container:hover {
    border: 1px solid #ff7b00;
  }

  .search-input {
    flex-grow: 1;
    border: none;
    outline: none;
    padding: 0.5rem;
    color: #ffffff;
    background-color: #0b1419;
  }

  .search-input::placeholder {
    color: #7e7e7e;
  }

  .right-section {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .right-section button {
    border-radius: 10px;
    border-color: transparent;
    cursor: pointer;
    color: #ffffff;
  }

  .active-button{
    background-color: #0b1419;
  }

  .run-button {
    padding: 0.7rem 1rem;
    background-color: #218b01;
    border: none;
    cursor: pointer;
    width: 5.5rem;
  }

  .run-button:hover {
    background-color: #1a5d05;
  }

  .run-button:active {
    background-color: #3d3d3d;
  }

  .invite-button {
    background-color: #162832;
    padding: 0.5rem 0.75rem;
    width: 8rem;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.5rem;
  }

  .invite-button:hover {
    background-color: #0b1419;
    background: transparent;
    border-width: 2px;
    border-color: #6a6a6a66;
  }

  .invite-button img {
    background: transparent;
  }

  .avatar {
    width: 2.5rem;
    height: 2.5rem;
    background-color: #444;
    border-radius: 50%;
  }

  .modal{
    position: absolute;
    background-color: #0b1419;
    top: 85px;
    right: 0;
    z-index: 1;
  }

  .modal-content{
    padding: 30px 25px;
  }
</style>

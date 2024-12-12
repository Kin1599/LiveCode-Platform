
<script>
    import UserItem from "./UserItem.svelte";

    let emailOrUsername = "";
    let users = [
        { name: 'User1', role: 'Editor' },
        { name: 'User2', role: 'Admin' },
        { name: 'User3', role: 'Viewer' },
    ];

    let inviteLink = window.location.href;

    function invite() {
        console.log(`Inviting ${emailOrUsername}`);
    }

    function copyToClipboard() {
        navigator.clipboard.writeText(inviteLink)
            .then(() => {
                alert('Link copied to clipboard!');
                console.log('Link copied to clipboard!');
            })
            .catch(err => {
                console.error('Failed to copy: ', err);
            });
    }
</script>

<div class="user-invitation">
    <div class="invite-header">
        <h2 class="invite-title">Multiplayer</h2>
    </div>
    <div class="invite-section">
        <input class="invite-input" type="text" bind:value={emailOrUsername} placeholder="Enter username or email" />
        <button class="invite-button" on:click={invite}>Пригласить</button>
    </div>
    <div class="user-list">
        {#each users as user}
            <UserItem user={user} />
        {/each}
    </div>
    <div class="invite-link">
        <div class="invite-link-header">
            <div class="invite-global">
                <div class="invite-global-icon">
                    <img src="../images/global.svg" alt="global">
                </div>
                <div class="invite-global-text">
                    <p class="text-title">Private join link</p>
                    <p class="text-description">Anyone with this link can edit files</p>
                </div>
            </div>
            <div class="toggle-switch">
                <input type="checkbox" id="toggle-edit-permission" />
                <label for="toggle-edit-permission" class="toggle-label"></label>
            </div>
        </div>
        <div class="invite-link-body">
            <input class="invite-input" type="text" bind:value={inviteLink} placeholder="https://misplitблабла"/>
            <button class="copy-button" on:click={copyToClipboard}>
                <img src="../images/link.svg" alt="copy">
                <p>Copy join link</p>
            </button>
        </div>
    </div>
</div>

<style>
    .user-invitation {
        background-color: #0B1419;
    }

    .user-invitation > *:not(:last-child){
        margin-bottom: 30px;
    }

    .invite-header{
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .invite-title{
        font-size: 36px;
        font-weight: 500;
        color: #FFF;
    }

    .invite-section{
        display: flex;
        gap: 15px;
        align-items: center;
    }

    .invite-input{
        background-color: #162832;
        border: 1px solid #6A6A6A66;
        border-radius: 5px;
        padding: 15px 10px;
        color: #7E7E7E;
        font-size: 20px;
        outline: none;
    }

    .invite-button{
        background-color: #162832;
        border-radius: 5px;
        color: #FFFFFF;
        font-size: 20px;
        padding: 15px 46px;
        border: none;
        cursor: pointer;
    }

    .user-list{
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .invite-link{
        border-top: 1px solid #6A6A6A66;
        padding: 15px 23px;
    }

    .invite-link-header{
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    
    .invite-global{
        display: flex;
        align-items: center;
        gap: 24px;
    }

    .invite-global-icon{
        height: 66px;
        width: 66px;
        background-color: #162832;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .invite-global-text{
        text-align: left;
    }

    .text-title{
        font-size: 24px;
        font-weight: 500;
        color: #fff;
    }

    .text-description{
        font-size: 14px;
        color: #7E7E7E;
        font-weight: 400;
    }

    .toggle-switch{
        position: relative;
    }

    .toggle-switch input{
        display: none;
    }

    .toggle-label{
        display: block;
        width: 40px;
        height: 20px;
        background-color: #ccc;
        border-radius: 20px;
        cursor: pointer;
        position: relative;
        transition: background-color 0.2s;
    }

    .toggle-label::after{
        content: "";
        display: block;
        width: 18px;
        height: 18px;
        background-color: #fff;
        border-radius: 50%;
        position: absolute;
        top: 1px;
        left: 1px;
        transition: transform 0.2s;
    }

    .toggle-switch input:checked + .toggle-label{
        background-color: #007bff;
    }

    .toggle-switch input:checked + .toggle-label::after{
        transform: translateX(20px);
    }

    .invite-link-body{
        margin-top: 15px;
        display: flex;
        gap: 15px;
    }

    .copy-button{
        background-color: #162832;
        border-radius: 5px;
        display: flex;
        align-items: center;
        gap: 5px;
        font-size: 20px;
        border: none;
        color: #fff;
        padding: 15px;
        cursor: pointer;
    }
</style>
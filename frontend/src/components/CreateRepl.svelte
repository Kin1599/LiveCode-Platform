<script lang="ts">

    interface SessionData{
        owner_id: string;
        editable: boolean;
        title: string;
        language: string;
        max_users: number;
    }

    interface Template {
        name: string;
        author: string;
        language: string;
    }

    type GetLanguageIcon = (language: string) => string;
    type CreateNewSession = (sessionData: SessionData) => void;

    export let templates: Template[] = [];
    export let getLanguageIcon: GetLanguageIcon;
    export let owner_id: string = "";
    export let createNewSession: CreateNewSession = (sessionData) => {};

    let replsName = '';
    let editable = true;
    let language = '';
    let maxUsers = 1;

    function handleCreateRepl(){
        if (replsName) {
            const sessionData = {
                owner_id: owner_id,
                editable: editable,
                title: replsName,
                language: 'python',
                max_users: maxUsers,
            }
            createNewSession(sessionData);
        }
    }
</script>
  
<div class="create-container">
    <div>
        <h2 class="create-title">Create a new Repl</h2>
        <h3 class="templates-title">Templates</h3>
        <div class="search-container search-templates">
            <img src="./images/icon-search.svg" alt="" />
            <input class="search-input" type="text" placeholder="Search" />
        </div>
        <ul class="template-list">
            {#each templates as template}
                <li class="template-item">
                    <img class="template-icon" src={getLanguageIcon(template.language)} alt={template.language} />
                    <div class="template-info">
                        <p class="info-name">{template.name}</p>
                        <p class="info-author">{template.author}</p>
                    </div>
                </li>
            {/each}
        </ul>
    </div>
    <div class="create-params">
        <div class="create-input-label">
            <p class="create-label">Title</p>
            <input class="create-input" type="text" placeholder="Name" bind:value={replsName}/>
        </div>
        <div class="create-input-label" style="margin-top: 3rem;">
            <p class="create-label">Max Users</p>
            <input class="create-input" type="text" placeholder="Max users" bind:value={maxUsers} />
        </div>
        <!-- <div class="params-public"> 
            <p class="create-label">Public</p>
            <div class="public-info">
                <div class="public-info-icon">
                    <img src="./images/global.svg" alt="global" />
                </div>
                <div class="public-info-text">
                    <p class="text-title">Private join link</p>
                    <p class="text-description">Anyone with this link can edit files</p>
                </div>
            </div>
        </div> -->
        <div class="params-editable">
            <p class="create-label">Editable</p>
            <label class="toggle-switch">
                <input class="toggle-checkbox" type="checkbox" bind:checked={editable}>
                <span class="slider"></span>
            </label>
        </div>
        <button class="createBtn" on:click={handleCreateRepl}>+ Создать</button>
    </div>
</div>
  
<style>
    .create-title{
        font-size: 48px;
        font-weight: 600;
        line-height: 56.25px;
        margin-bottom: 30px;
    }

    .create-container{
        display: grid;
        justify-content: center;
        grid-template-columns: repeat(2, minmax(0, 505px));
        gap: 124px;
    }

    .templates-title{
        font-size: 20px;
        font-weight: 500;
        line-height: 23.44px;
        margin-bottom: 22px;
    }

    .search-container {
        display: flex;
        align-items: center;
        padding: 5px 10px;
        max-width: 900px;
        border: 1px solid #7e7e7e;
        border-radius: 10px;
        height: 44px;
    }

    .search-templates {
        margin: 0px;
        width: 100%;
    }

    .search-container:hover {
        border: 1px solid #ff7b00;
    }

    .search-input {
        border: none;
        outline: none;
        padding: 10px;
        font-size: 20px;
        color: #ffffff;
        background-color: #0b1419;
    }

    .template-list{
        margin-top: 5px;
    }

    .template-item{
        background-color: #162832;
        height: 60px;
        padding: 5px 10px;
        display: flex;
        align-items: center;
        gap: 10px;
        width: 100%;
        cursor: pointer;
    }

    .template-item:first-child{
        border-radius: 5px 5px 0 0;
    }

    .template-item:last-child{
        border-radius: 0 0 5px 5px;
    }

    .template-icon{
        width: 40px;
        height: 40px;
        background-color: transparent;
    }

    .template-info{
        background-color: transparent;
    }

    .template-info > *:not(:last-child){
        margin-bottom: 2px;
    }

    .info-name{
        font-size: 20px;
        line-height: 23.44px;
        background-color: transparent;
    }

    .info-author{
        color: #7E7E7E;
        background-color: transparent;
    }

    .create-input-label > *:not(:last-child){
        margin-bottom: 22px;
    }

    .create-label{
        font-size: 20px;
        line-height: 23.44px;
    }

    .create-input{
        background-color: #162832;
        outline: none;
        border-radius: 5px;
        width: 100%;
        padding-left: 10px;
        border: 1px solid #6A6A6A66;
        font-size: 20px;
        color: #EBEBEB;
        height: 44px;
    }

    .params-public{
        margin: 63px 0 130px 0;
    }

    .public-info{
        margin-top: 29px;
        display: flex;
        align-items: center;
        gap: 24px; 
    }

    .public-info-icon{
        border-radius: 50%;
        height: 66px;
        width: 66px;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: #162832;
    }

    .public-info-icon > img{
        background-color: transparent;
    }
  
    .public-info-text > *:not(:last-child){
        margin-bottom: 5px;
    }
  
    .public-info-text > .text-title{
        color: #FFF;
        font-size: 24px;
        font-weight: 500;
    }
  
    .public-info-text > .text-description{
        color: #7E7E7E;
    }

    .toggle-switch{
        position: relative;
        display: inline-block;
        width: 60px;
        height: 34px;
    }

    .toggle-switch input {
        opacity: 0;
        width: 0;
        height: 0;
    }

    .slider {
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: #ccc;
        transition: 0.4s;
        border-radius: 34px;
    }

    .slider:before {
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: 0.4s;
        border-radius: 50%;
    }

    .toggle-checkbox:checked + .slider {
        background-color: #ff7b00;
    }

    .toggle-checkbox:checked + .slider:before {
        transform: translateX(26px);
    }

    .params-editable {
        display: flex;
        align-items: center;
        gap: 10px;
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .create-label {
        font-size: 20px;
        line-height: 23.44px;
    }
  
    .createBtn{
        width: 100%;
        background-color: #ff7b00;
        border-style: solid;
        border-radius: 10px;
        color: #fff;
        font-size: 20px;
        padding: 10px 15px;
        height: 55px;
    }

    .createBtn:hover {
        background: transparent;
        border-width: 2px;
        border-color: #ff7b00;
        color: #ff7b00;
        cursor: pointer;
    }
</style>

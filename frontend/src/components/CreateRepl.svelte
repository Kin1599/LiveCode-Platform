<script>
    export let templates;
    export let getLanguageIcon;
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
            <input class="create-input" type="text" placeholder="Name" />
        </div>
        <div class="params-public"> 
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
        </div>
        <button class="createBtn">+ Создать</button>
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
        color: #7E7E7E;
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
  
    .createBtn{
        width: 100%;
        background-color: #ff7b00;
        border: none;
        border-radius: 10px;
        color: #fff;
        font-size: 20px;
        padding: 10px 15px;
        height: 55px;
    }
</style>
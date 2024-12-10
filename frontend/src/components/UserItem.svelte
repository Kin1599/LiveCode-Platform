<script>
    let { user } = $props();
</script>

<div class="user-item">
    <div class="user-info">
        <img src={user.avatar ? user.avatar : "./images/avatar.png"} alt="avatar">
        <p>{user.name}</p>
    </div>
    <div class="user-actions">
        <select class="action-edit"> 
            <option value="can edit">can edit</option> 
            <option value="admin">admin</option> 
            <option value="viewer">viewer</option>
        </select>
        <button class="action-delete">
            <img src="./images/delete.svg" alt="delete">
        </button>
    </div>
</div>

<style>
    .user-item{
        display: flex;
        justify-content: space-between;
        align-items: center;
        background-color: #162832;
        border-radius: 5px;
        padding: 15px;
    }

    .user-info{
        display: flex;
        align-items: center;
        gap: 10px;
    }

    .user-actions{
        display: flex;
        align-items: center;
        gap: 20px;
    }

    .action-delete{
        background-color: transparent;
        border: none;
    }
</style>
<script lang="ts">
    export let messages: { user: string, text: string }[] = [];
    export let currentMessage: string = "";
    export let sendMessage: (message: string) => void;
  
    function handleSendMessage() {
        if (currentMessage.trim() !== "") {
            sendMessage(currentMessage);
            currentMessage = "";
        }
    }

    function handleKeyDown(event: KeyboardEvent) {
        // Проверяем, что нажата клавиша Enter (код 13)
        if (event.key === "Enter" && currentMessage.trim() !== "") {
            sendMessage(currentMessage);
            currentMessage = "";
        }
    }
</script>
  
<div class="chat-container">
    <div class="messages">
        {#each messages as { user, text }}
            <div class="message">
                <strong>{user}</strong>: {text}
            </div>
        {/each}
    </div>
    <div class="input-container">
        <input 
            class="message-input" 
            type="text" 
            bind:value={currentMessage} 
            placeholder="Sent message" 
            on:keydown={handleKeyDown}
        />
        <button class="send-button" on:click={handleSendMessage}>
            <img src="../images/send.svg" alt="send">
        </button>
    </div>
</div>
  
<style>
    .chat-container {
        display: flex;
        flex-direction: column;
        height: 100%;
        padding: 10px;
        box-sizing: border-box;
        background-color: #0B1419;
        color: #fff;
    }

    .messages {
        flex: 1;
        overflow-y: auto;
        margin-bottom: 10px;
    }

    .message {
        margin: 5px 0;
    }

    .input-container {
        display: flex;
        max-height: 66px;
        border: 1px solid #6A6A6A66;
        padding: 21px 10px;
        border-radius: 5px;
    }

    .message-input {
        flex: 1;
        border: none;
        outline: none;
        background-color: transparent;
        color: #7e7e7e;
        font-size: 20px;
        font-weight: 300px;
    }

    .send-button {
        background-color: transparent;
        border: none;
        cursor: pointer;
    }
</style>
  
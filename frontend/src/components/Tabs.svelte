<script lang="ts">
  export let tabs: { id: number, name: string }[] = [];
  export let activeTab: number = 0;
  export let onTabClick: (event: CustomEvent<{ tabId: number }>) => void;
</script>

<div class="tabs">
  {#each tabs as { id, name }}
    <button 
      class="tab {id === activeTab ? 'active' : ''}"
      on:click={() => {
        const event = new CustomEvent('tabClick', {
          detail: { tabId: id },  
        });
        onTabClick(event);  
      }}
    >
      {name}
    </button>
  {/each}
</div>
  
<style>
  .tabs {
    display: flex;
    background-color: #0B1419;
  }

  .tab {
    display: flex;
    align-items: center;
    width: 121px;
    height: 33px;
    padding: 0.5px 10px;
    background-color: #0B1419;
    border: none;
    border-right: 1px solid #6A6A6A66;
    color: #fff;
  }

  .tab:hover {
    background-color: #555;
    cursor: pointer;
  }

  .tab.active{
    background-color: #162832;
  }
</style>
  
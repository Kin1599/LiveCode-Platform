<script lang="ts">
    export let tabs: string[] = [];
  </script>
  
  <div class="tabs">
    {#each tabs as tab}
      <div class="tab">{tab}</div>
    {/each}
  </div>
  
  <style>
    .tabs {
      display: flex;
      gap: 0.5rem;
    }
  
    .tab {
      padding: 0.25rem 0.5rem;
      background-color: #333;
      border: 1px solid #444;
      border-radius: 3px;
    }
  </style>
  
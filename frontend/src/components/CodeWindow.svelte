<script lang="ts">
    import CodeMirror from "svelte-codemirror-editor";
    import { javascript } from "@codemirror/lang-javascript";
    export let value: string = "";
    export let lang = javascript(); 
</script>
  
<CodeMirror
  class="codemirror"
  bind:value
  lang={lang}
  styles={{
    "&": {
      maxWidth: "100%",
      height: "49.5rem",
    },
  }}
/>
  
<style>

</style>
  
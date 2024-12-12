<script lang="ts">
  import CodeMirror from "svelte-codemirror-editor";
  import { javascript } from "@codemirror/lang-javascript";
  import { autocompletion } from "@codemirror/autocomplete";
  import { highlightActiveLine, highlightActiveLineGutter } from "@codemirror/view"; 
  import { EditorView } from "@codemirror/view";
  import { lineNumbers } from "@codemirror/gutter"

  export let value: string = "";
  export let lang = javascript();

  let editorConfig = {
      tabSize: 4,
      extensions: [
          autocompletion(),
          highlightActiveLine(),
          highlightActiveLineGutter(),
          EditorView.theme({
              "&": {
                backgroundColor: "rgba(22, 40, 50, 1)",
                color: "#fff",
                fontSize: "20px" 
              }, 
              ".cm-content": {
                paddingTop: "0px",
              },
              ".cm-gutters": {
                backgroundColor: "rgba(22, 40, 50, 1)",
                borderColor: "rgba(22, 40, 50, 1)"
              },
              ".cm-gutterElement": { 
                color: "#7E7E7E",
              },
              ".cm-activeLine": { 
                backgroundColor: "rgb(38, 70, 88)" 
              },
              ".cm-activeLineGutter": {
                backgroundColor: "rgb(38, 70, 88)" 
              }
          })
      ],
  };
</script>

<CodeMirror
  class="codemirror"
  bind:value
  lang={lang}
  tabSize={editorConfig.tabSize} 
  extensions={editorConfig.extensions}
  styles={{
      "&": {
        maxWidth: "100%",
        height: "52.5rem",
      },
  }}
/>
  
<style>

</style>
<template>
  <div id="div1" style="width: 100%;">
    <p>请输入内容</p>
  </div>
</template>

<script>
import wangEditor from "wangeditor";
export default {
  name: "noteEditor",
  data() {
    return {
        editor:null
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.editor = new wangEditor("#div1");
      this.editor.config.height=480
      this.editor.create();
    },
  },
};
</script>
<style>

</style>
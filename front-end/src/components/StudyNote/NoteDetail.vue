<template>
  <div>
    <h1>{{ detail.title }}</h1>
  </div>
</template>
<script>
import bus from "../../assets/eventBus";
export default {
  data() {
    return {
      detail: {
        title: "暂无数据",
        content: "",
        text: "暂无数据",
        tags: "",
        create_time: "",
        id: 0,
        update_time: "",
      },
    };
  },
  mounted() {
    bus.$on("detail", function (detail) {
     
      this.detail = detail;
    });
  },
};
</script>
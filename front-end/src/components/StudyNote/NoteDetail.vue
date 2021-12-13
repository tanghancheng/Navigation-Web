<template>
  <div
    class="content"
    style="
      width: 90%;
      height: 680px;
      overflow-y: scroll;
      background-color: #fff;
      padding: 0px 25px;
    "
  >
    <h3 align="center">{{ detail.title }}</h3>
    <el-divider></el-divider>
    <div v-html="detail.content">{{ detail.content }}</div>
    <el-divider></el-divider>
    <div>
      <span> 
        <el-tag v-for="(tag,index) in tags" :key="index" type="warning">{{ tag }}
         </el-tag>
      </span> 
      <span style="float:right">{{ detail.create_time }}</span>
    </div>
    </br>
  </div>
</template>
<script>
export default {
  data() {
    return {
      detail: {
        title: "",
        content: "",
        text: "",
        tags: "",
        create_time: "",
        id: 0,
        update_time: "",
      },
      tags:[]
    };
  },
  created() {
    console.log(this.$route.params);
    this.$axios.get("/api/note/getOne/" + this.$route.params.id).then((res) => {
      console.log(res.data);
      this.detail = res.data;
      this.tags=this.detail.tags.split(",")
    });
  },
};
</script>
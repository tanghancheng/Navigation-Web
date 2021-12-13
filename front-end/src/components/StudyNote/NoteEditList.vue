<template>
  <div>
    <div class="content" style="width: 90%; height: 680px; overflow-y:scroll; background-color: #fff">
      <ul>
          <li  style="list-style: none" >
          <el-row>
              <el-col :span="3" :push="21"> 
                  <el-button  style="margin-bottom: 20px;" @click="toEdit(0)"
                  type="success"> 新增笔记
                   </el-button>
              </el-col>
           </el-row>
        </li>
        <li style="list-style: none" v-for="item in tableData">
          <div style="height: 90px; width: 100%">
            <el-row>
              <el-col :span="20">
                <h4>{{ item.title }}</h4>
                <span>{{ item.text|subString}}</span>
                </el-col>
              <el-col :span="2"></br><el-tag type="success">{{item.tags}}</el-tag></el-col>
                <el-col :span="2"> 
                    <el-button style="margin-top: 20px;margin-left: 20px;" 
                type="primary" icon="el-icon-edit" @click="toEdit(item.id)" circle>
                </el-button>
                </el-col>
            </el-row>
          </div>
        </li>
        <div style="text-align:center;margin:auto;">
          <el-button type="text" @click="load">加载更多</el-button>
        </div>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "noteEditList",
  data() {
    return {
      tableData: [
        {
          title: "暂无数据",
          content: "",
          text: "暂无数据",
          tags: "",
          create_time: "",
          id: 0,
          update_time: "",
        },
      ],
    };
  },
  methods: {
    load() {
      // this.tableData = this.tableData.concat(this.tableData);
      this.$axios.get("/api/note/getListByQuery").then((res) => {
        this.tableData = res.data;
      });
      console.log(this.tableData);
    },
    toEdit(id) {
      console.log(id);
      this.$router.push("/noteEditor/" + id);
    },
  },
  filters: {
    subString(content) {
      if (content != "" && content.length > 24) {
        return content.substr(0, 50) + "…";
      } else {
        return content;
      }
    },
  },
  created() {
    this.load();
  },
};
</script>
<style>
.content::-webkit-scrollbar {
  display: none;
}
</style>
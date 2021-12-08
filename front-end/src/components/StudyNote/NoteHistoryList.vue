<template>
  <div>
    <div class="content" style="width: 90%; height: 680px; overflow-y:scroll; background-color: #fff">
      <ul>
        <li style="list-style: none" v-for="item in tableData">
          <div style="height: 90px; width: 100%">
            <el-row>
              <el-col :span="22">
                <h4>{{ item.name }}</h4>
                <span>{{ item.address }}</span>
                </el-col>
              <el-col :span="2"></br><el-tag type="success">{{item.tag}}</el-tag></el-col>
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
  name: "noteHistoryList",
  data() {
    return {
      tableData: [
        {
          date: "2016-05-02",
          name: "王小虎",
          address: "上海市普陀区金沙江路 1518 弄",
          tag: "家",
        },
      ],
    };
  },
  methods: {
    load() {
      this.tableData = this.tableData.concat(this.tableData);
      console.log(this.tableData);
    },
  },
};
</script>
<style>
.content::-webkit-scrollbar {
  display: none;
}
</style>
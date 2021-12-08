<template>
  <div>
    <el-container style="height: 775px; border: 0px solid #eee">
      <el-header style="margin-bottom: 20px; margin: auto">
        <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          style="width: 900px"
        >
          <el-row :gutter="250">
            <el-col :span="8">
              <div>
                <el-menu-item index="1" @click="checkToNoteHistoryList">
                  以往的笔记
                </el-menu-item>
              </div>
            </el-col>
            <el-col :span="8">
              <div>
                <el-menu-item index="2" @click="checkToNoteEditor">
                  编写我的笔记
                </el-menu-item>
              </div>
            </el-col>
            <el-col :span="8">
              <div>
                <el-menu-item index="3">消息中心</el-menu-item>
              </div>
            </el-col>
          </el-row>
        </el-menu>
        <div class="line"></div>
      </el-header>
      <el-container>
        <Aside />
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Aside from "../../components/Aside.vue";

export default {
  components: {
    Aside,
  },
  data() {
    return {
      activeIndex: "1",
    };
  },

  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    checkToNoteEditor() {
      this.$router.push("/noteEditor");
    },
    checkToNoteHistoryList() {
      this.$router.push("/note");
    },
  },
};
</script>
<style>
.el-aside {
  color: #333;
}
</style>
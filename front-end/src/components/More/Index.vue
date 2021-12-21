<template>
  <div>
      <el-container style="height: 775px; border: 0px solid #eee">
        <el-header>
          <el-row :gutter="20">
            <el-col :span="23">
              <div class="grid-content bg-purple"></div>
            </el-col>
            <el-col :span="1">
              <div class="grid-content bg-purple">
              </div>
            </el-col>
          </el-row>
        </el-header>
        <el-container>
          <Aside />
          <el-main>
              <div style="background-color: aliceblue; width: 88%;
               height: 95%;text-align: center;">
               <span > 暂无更多  </span>
              </div>
          </el-main>
        </el-container>
      </el-container>
  </div>
</template>

<script>
import Aside from "../Aside.vue";

export default {
  components: {
    Aside,
  },
  data() {
    return {
      isCollapse: true,
    };
  },
  methods: {
  },
};
</script>

<style scoped>
.el-header {
  color: #333;
  line-height: 60px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}


</style>

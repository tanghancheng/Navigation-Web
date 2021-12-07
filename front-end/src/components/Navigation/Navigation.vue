<template>
  <div>
    <transition name="el-zoom-in-top">
      <el-container style="height: 775px; border: 0px solid #eee">
        <el-header>
          <el-row :gutter="20">
            <el-col :span="23"
              ><div class="grid-content bg-purple"></div
            ></el-col>
            <el-col :span="1"
              ><div class="grid-content bg-purple">
                <i class="el-icon-share"></i></div
            ></el-col>
          </el-row>
        </el-header>
        <el-container>
          <Aside />
            <router-view></router-view>
        </el-container>
      </el-container>
    </transition>
  </div>
</template>

<script>
import WebItem from "../WebItem.vue";
import Aside from "../Aside.vue";

export default {
  components: {
    WebItem,
    Aside,
  },
  data() {
    return {
      isCollapse: true,
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    transName(webItem) {
      return this.lang.key === "en" ? webItem.en_name : webItem.name;
    },
    moveSlow() {},
    checkToNavigtionList() {
      this.$router.push("/navigationList");
    },
  },
  created(){
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header {
  color: #333;
  line-height: 60px;
}
.el-aside {
  color: #333;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.my-border {
  box-shadow: 0 2px 12px 0 rgb(23 20 20);
  margin-right: 100px;
  min-height: 100px;
  border-radius: 4px;
}
</style>

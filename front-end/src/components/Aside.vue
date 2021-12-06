<template>
  <el-aside width="210px" >
    <el-radio-group
      v-model="isCollapse"
      style="margin-bottom: 10px; background-color: #fff0"
    >
      <el-radio-button :label="false"
        ><li class="el-icon-plus"></li
      ></el-radio-button>
      <el-radio-button :label="true" 
        ><li class="el-icon-minus"></li
      ></el-radio-button>
    </el-radio-group>
    <el-menu
      default-active="1-4-1"
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      :collapse="isCollapse"
      style="background-color: #fff0;border: 0px"
    >
      <el-menu-item index="1" @click="checkToSearchNavigtion">
        <i class="el-icon-location"></i>
        <span slot="title">导航页</span>
      </el-menu-item>
      <el-menu-item index="2">
        <i class="el-icon-menu"></i>
        <span slot="title">导航二</span>
      </el-menu-item>
      <el-menu-item index="3">
        <i class="el-icon-document"></i>
        <span slot="title">导航三</span>
      </el-menu-item>
      <el-menu-item index="4" @click="checkToNavigtionList">
        <i class="el-icon-setting"></i>
        <span slot="title">导航列表</span>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>

<script>
export default {
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
    checkToNavigtionList() {
      this.$router.push("/navigationList");
    },
    checkToSearchNavigtion() {
      this.$router.push("/navigation");
    },
  },
};
</script>
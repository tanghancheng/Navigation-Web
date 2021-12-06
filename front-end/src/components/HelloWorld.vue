<template>
  <div>
    <transition name="el-zoom-in-top">
      <el-container style="height: 775px; border: 1px solid #eee">
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
          <el-aside width="210px">
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
              style="background-color: #fff0"
            >
              <el-menu-item index="1">
                <i class="el-icon-location"></i>
                <span slot="title">导航二</span>
              </el-menu-item>
              <el-menu-item index="2">
                <i class="el-icon-menu"></i>
                <span slot="title">导航二</span>
              </el-menu-item>
              <el-menu-item index="3">
                <i class="el-icon-document"></i>
                <span slot="title">导航三</span>
              </el-menu-item>
              <el-menu-item index="4">
                <i class="el-icon-setting"></i>
                <span slot="title">导航四</span>
              </el-menu-item>
            </el-menu>
          </el-aside>
          <el-main>
            <Search />
            <el-row :gutter="20" style="margin-top: 100px">
              <el-col :span="8" v-for="(item, idx) in items" :key="idx"
                ><el-link :underline="false" :href="item.url" target="_blank"
                        >
                <div class="grid-content bg-purple my-border">
                  <el-row :gutter="20">
                    <el-col :span="2">
                      <img
                        :src="item.logo"
                        alt="加载"
                        width="60px"
                        style="border-radius: 50%"
                    /></el-col>
                    <el-col :span="22"
                      ><h3 style="text-align: center">
                        {{ item.title }}
                      </h3></el-col
                    >
                  </el-row>
                  <el-row :gutter="20">
                    <el-col :span="24"
                      ><span>{{ item.desc }}</span>
                        </br></el-col
                    > 
                  </el-row>
                </div></el-link
                      ></el-col
              >
            </el-row>
          </el-main>
        </el-container>
      </el-container>
    </transition>
  </div>
</template>

<script>
import Search from "../components/Search.vue";
import WebItem from "../components/WebItem.vue";
import itemsData from "../assets/data.json";
export default {
  components: {
    Search,
    WebItem,
  },
  data() {
    return {
      items: itemsData,
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
  },
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

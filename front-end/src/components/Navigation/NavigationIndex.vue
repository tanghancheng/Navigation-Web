<template>
      <el-main >
        <Search/>
            <el-row  :gutter="20" class="content-row"  v-loading="loading" style="margin-top: 80px;">
              <el-col :span="8"   v-for="(item, idx) in items" :key="idx"
                ><el-link :underline="false" :href="item.url" target="_blank" >
                <div class="grid-content  div-border">
                  <div>
                      <img
                        :src="item.logo"
                        alt="加载"
                        width="50px"
                        height="40px"
                        style="border-radius: 40%"
                    />
                     <span style="font-size:16px">{{ item.title }}</span>
                  </div></br>
                  <span>{{ item.desc }}</span>
                </div></el-link></el-col>
            </el-row>
    </el-main>
</template>
<script>
import Search from "../Search.vue";
import axios from "axios";
export default {
  components: {
    Search,
  },
  data() {
    return {
      items: [],
      loading: true
    };
  },
  methods: {
    initNavigaionData() {
      axios.get("/api/navigation/navigation").then((res) => {
        this.items = res.data;
        this.items = this.items.slice(0, 6);
        console.log(this.items);
        this.loading =false
      });
    },
  },
  created() {
    this.initNavigaionData();
  },
};
</script>
<style scoped >
.div-border {
  box-shadow: 0 2px 12px 0 rgb(23 20 20);
  margin-right: 50px;
  border-radius: 20px;
  margin-bottom: 30px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
  padding: 20px;
}

.content-row {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  align-content: middle;
}
</style>
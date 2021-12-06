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
          <el-main>
            <el-table :data="tableData" border style="width: 100%">
              <el-table-column fixed type="index" label="id" width="120">
              </el-table-column>
              <el-table-column prop="created_time" label="日期" width="250">
              </el-table-column>
              <el-table-column prop="url" label="url" width="120">
              </el-table-column>
              <el-table-column prop="logo" label="logo" width="300">
              </el-table-column>
              <el-table-column prop="title" label="标题" width="120">
              </el-table-column>
              <el-table-column fixed="right" label="操作" width="200">
                <template slot-scope="scope">
                  <el-popconfirm
                    confirm-button-text="好的"
                    cancel-button-text="不用了"
                    icon="el-icon-info"
                    icon-color="red"
                    @confirm="deleteClick(scope.row)"
                    title="确定要删除吗"
                  >
                    <el-button
                      slot="reference"
                      type="danger"
                      icon="el-icon-delete"
                    ></el-button>
                  </el-popconfirm>
                  <el-button
                    type="primary"
                    icon="el-icon-edit"
                    circle
                    @click="getForm(scope.row)"
                  ></el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-main>
        </el-container>
      </el-container>
    </transition>
    <el-dialog title="收货地址" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="url地址" :label-width="formLabelWidth">
          <el-input v-model="form.url" autocomplete="off"></el-input>
        </el-form-item>
         <el-form-item label="logo" :label-width="formLabelWidth">
          <el-input v-model="form.logo" autocomplete="off"></el-input>
        </el-form-item>
         <el-form-item label="标题" :label-width="formLabelWidth">
          <el-input v-model="form.title" autocomplete="off"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="updateNavigation()"
          >确 定</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Aside from "../Aside.vue";

var dataList = [];
import axios from "axios";
export default {
  components: { Aside },
  data() {
    return {
      isCollapse: true,
      tableData: dataList,
      dialogFormVisible: false,
      form: {
        url: "",
        id: "",
        title: "",
        logo: "",
      },
      formLabelWidth: "120px",
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
    deleteClick(row) {
      axios
        .delete("/api/navigation/navigation", {
          data: {
            id: row.id,
          },
        })
        .then((res) => {
          console.log(res.data);
          this.$router.go(0);
        });
    },
    initData() {
      axios.get("/api/navigation/navigation").then((res) => {
        this.tableData = res.data;
        console.log(this.tableData);
      });
    },
    getForm(row){
      this.dialogFormVisible=true
      this.form=row
    },
    updateNavigation(){
      console.log(this.form);
      console.log("异步前");
       axios.put("/api/navigation/navigation/"+this.form.id,{
           url:this.form.url,
           title:this.form.title,
           logo:this.form.logo,

       }).then((res) => {
         this.dialogFormVisible=false
         console.log(this.form);
        console.log(res.data);
        console.log(this.tableData);
      });
    }

  },
  created() {
    this.initData();
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
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>

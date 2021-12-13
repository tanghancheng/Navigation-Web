<template>
  <div>
    <el-main>
      <el-table :data="tableData" st v-loading="loading" border 
      style="width: 100%;overflow-y:scroll;height: 680px;" class="my-border content">
        <el-table-column fixed type="index" label="序号" width="120">
        </el-table-column>
           <el-table-column prop="logo" label="LOGO" width="100">
          <template slot-scope="scope">
            <img :src="scope.row.logo" alt="加载失败" width="60px">
            </img>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题" width="150">
        </el-table-column>
        <el-table-column prop="url" label="访问网址" width="220"> 
        </el-table-column>
        <el-table-column prop="desc" label="简介" width="200">
        </el-table-column>
        <el-table-column prop="weight" label="权重" width="100">
        </el-table-column>
        <el-table-column fixed="right" width="200">
          <template slot="header" slot-scope="scope">
            <el-button type="success" size="medium" @click="getForm(1)">新增
              </el-button>
          </template>
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
                icon="el-icon-delete">
              </el-button>
            </el-popconfirm>
            <el-button
              type="primary"
              icon="el-icon-edit"
              circle
              @click="getForm(scope.row)">
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
    <el-dialog title="修改导航栏" :visible.sync="dialogFormVisible">
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
        <el-form-item label="备注" :label-width="formLabelWidth">
          <el-input v-model="form.desc" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="权重" :label-width="formLabelWidth">
          <el-input type="number" v-model.number="form.weight" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="updateNavigation()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
var dataList = [];
import axios from "axios";
export default {
  components: {},
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
        desc: "",
        weight: 0,
      },
      loading: true,
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
          if (res.data) {
            this.$notify({
              title: "成功",
              message: "这是一条成功的提示消息",
              type: "success",
            });
          }
          this.$router.go(0);
        });
    },
    initData() {
      axios.get("/api/navigation/navigation").then((res) => {
        this.tableData = res.data.data;
        this.loading = false;
      });
    },
    getForm(row) {
      this.dialogFormVisible = true;
      if (row == 1) {
        this.form = {};
      } else {
        this.form = row;
      }
    },
    updateNavigation() {
      var data = {
        url: this.form.url,
        title: this.form.title,
        logo: this.form.logo,
        desc: this.form.desc,
        weight: this.form.weight,
      };
      if (!this.form.id) {
        axios.post("/api/navigation/navigation", data).then((res) => {
          this.dialogFormVisible = false;
          if (res.data) {
            this.$notify({
              title: "成功",
              message: "这是一条成功的提示消息",
              type: "success",
            });
          }
        });
      } else {
        axios
          .put("/api/navigation/navigation/" + this.form.id, data)
          .then((res) => {
            this.dialogFormVisible = false;
            if (res.data) {
              this.$notify({
                title: "成功",
                message: "这是一条成功的提示消息",
                type: "success",
              });
            }
          });
      }
      this.$router.go(0);
    },
  },
  created() {
    this.initData();
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
.content::-webkit-scrollbar {
  display: none;
}
</style>

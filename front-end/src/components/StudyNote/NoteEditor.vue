<template>
  <div style="width: 88%; height: 680px; background-color: #fff">
    <el-form ref="form" :model="form" label-width="30px" style="width: 97%">
      <el-form-item>
        <el-input
          v-model="form.title"
          placeholder="请输入标题"
          style="margin-top: 15px"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-tag
          :key="tag"
          v-for="tag in tags"
          closable
          :disable-transitions="false"
          @close="handleClose(tag)"
        >
          {{ tag }}
        </el-tag>
        <el-input
          class="input-new-tag"
          v-if="inputVisible"
          v-model="inputValue"
          ref="saveTagInput"
          size="small"
          @keyup.enter.native="handleInputConfirm"
          @blur="handleInputConfirm"
        >
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput"
          >+ New Tag</el-button
        >
      </el-form-item>
      <el-form-item>
        <div id="div1">
          <p>请输入内容</p>
        </div>
      </el-form-item>
      <el-form-item>
        <el-row>
          <el-col :span="4">
            <el-button type="success" @click="onSubmit(1)"
              >先保存一下</el-button
            >
          </el-col>
          <el-col :span="4">
            <el-button  @click="closeEditor">取消编辑</el-button>
          </el-col>
          <el-col :span="4" :push="13">
            <el-button type="primary" @click="onSubmit(0)"
              >保存并退出</el-button
            >
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import wangEditor from "wangeditor";

import axios from "axios";
export default {
  name: "noteEditor",
  components: {},
  data() {
    return {
      form: {
        title: "",
      },
      tags: ["标签一"],
      inputVisible: false,
      inputValue: "",
      detail: {
        id: 0,
      },
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    onSubmit(editStaus) {
      var data = {
        id: this.detail.id,
        title: this.form.title,
        tags: this.tags.toString(),
        content: this.editor.txt.html(),
        text: this.editor.txt.text(),
        edit_status: editStaus,
      };
      if (this.detail.id) {
        data.edit_status = editStaus;
        axios.put("/api/note/update/" + this.detail.id, data).then((res) => {
          if (res.status == 200) {
            this.$notify({
              title: "保存成功",
              message: "保存成功",
              type: "success",
            });
            this.$router.push("/noteEditList");
          }
        });
      } else {
        axios.post("/api/note/create", data).then((res) => {
          if (res.status == 200) {
            this.$notify({
              title: "保存成功",
              message: "保存成功",
              type: "success",
            });
            this.$router.push("/noteEditList");
          }
        });
      }
    },
    init() {
      this.editor = new wangEditor("#div1");
      this.editor.config.height = 400;
      this.editor.create();
    },
    handleClose(tag) {
      this.tags.splice(this.tags.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.tags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    closeEditor() {
      this.$router.push("/noteEditList");
    },
  },
  created() {
    console.log(this.$route.params);
    if (this.$route.params) {
      this.$axios
        .get("/api/note/getOne/" + this.$route.params.id)
        .then((res) => {
          console.log(res.data);
          this.detail = res.data;
          this.form.title = this.detail.title;
          this.tags = this.detail.tags.split(",");
          this.editor.txt.html(this.detail.content);
        });
    }
  },
};
</script>
<style>
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
</style>
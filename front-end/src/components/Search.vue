<template>
  <div class="home">
    <span class="time">{{currentTime}}</span>
    <div class="search_box">
      <div class="search_x">
        <input
          type="text"
          v-model="value"
          class="search_ipt"
          @keydown.enter="openlinks(value)"
          @blur="narrow"
          :style="{ width: iptwidth + 'px' }"
        />
        <div class="iconbut"><span class="el-icon-search " @mousemove="x" @click="openlinks(value)"></span></div>
      </div>
    </div>
    
  </div>
</template>

<script>
import "../assets/css/app.css";
// @ is an alias to /src
export default {
  data() {
    return {
      value: "",
      iptwidth: "0",
      listwidth: "0",
      data: {},
      timer: "",//定义一个定时器的变量
      currentTime:""
    };
  },
  methods: {
    x() {
      this.iptwidth = "450";
    },
    openlinks(val) {
      if(val){
        window.open(`https://www.baidu.com/s?wd=${val}`);
      }
    },
    narrow(){
      if(!this.value){
        this.iptwidth="0"
      }
    },
    check(i){
       const num = (i<10)?("0"+i) : i;
       return num;
    }
  },
  created() {
    var _this = this; //声明一个变量指向Vue实例this，保证作用域一致
    this.timer = setInterval(function() {
      // _this.currentTime = //修改数据date
      //   new Date().getHours() +
      //   ":" +
      //   new Date().getMinutes() 
      let hours=new Date().getHours();
      let minutes=new Date().getMinutes();
      if(hours<10){
        hours="0"+hours
      }
      if(minutes<10){
        minutes="0"+minutes
      }
      _this.currentTime=hours+":"+minutes
    }, 1000);
  },
beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer); // 在Vue实例销毁前，清除我们的定时器
    }
  }
};
</script>

<style>

</style>
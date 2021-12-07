<template>
<div>
 <el-container class="content" style="height: 775px ;border: 0px" >
     <!-- <router-link to="/navigtion"> -->
   <div class="center" @click="goToNavigtion" >
      <p>当很累很累的时候，你应该告诉自己你能坚持得住，不要轻易的否定自己。</p>
       <p>一时的输赢并不代表什么，重要的是把后面的路走好。</p>
       <p>当你强大，才会遇到比你更强大的；当你变好，你才配得起更好的。</p>
         <p>所以，不要轻易放弃一段坚持，即使没人为你鼓掌，也要优雅的谢幕。</p>
         <p>人活着不是靠泪水博得同情，而是靠汗水赢得掌声！</p>
     </div> 
    <!-- </router-link> -->
  </el-container>
  </div>
</template>

<script>

export default {
  name: 'index',
  data() {
    return {
      show:true
    }
  },
  methods:{
    goToNavigtion(){
      this.$router.push("/navigation")
    }
  },
  watch:{
      $route(){}
  }
}
</script>

<style>
.center{
  width: 50%;
  margin: 240px 0 0 100px ;
  font-size: 18px;
  
}
.content{
    width: 100%;
    height: auto;
    
}
h1{
  text-align: center;
}

</style>

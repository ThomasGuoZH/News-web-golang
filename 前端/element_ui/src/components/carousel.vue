<template>
    <el-carousel indicator-position="outside" height="360px">
    <el-carousel-item v-for="(image,index) in imageInfo" :key="index" >
      <img :src="image.src"  :alt="image.alt">
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
  props:{
    imageInfo:{
      type:Array,
      required:true
    }
  }
}
</script>

<style>
  .el-carousel {
    width:80%;
    display: block;
    left:10%;
  }
     .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0 ;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>
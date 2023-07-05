<template>
  <transition>
    <div class="back-to-top" @click="toTop">
      <el-button class="button" @click="toTop">
        <img src="../../assets/icons/arrow-up.svg" alt="回到顶端"> 
      </el-button>
    </div>
  </transition>
</template>
    
<script>


export default {
  name: "BackToTop",
  methods: {
    toTop() {
      window.scrollTo({ top: 0, behavior: 'smooth' })
    }
  }
}
</script>
    
<style scoped>
.back-to-top {
  cursor: pointer;
  position: fixed;
  right: 25px;
  bottom: 50px;
  z-index: 999;
}

.button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
.back-to-top img {
  height: 25px;
  width: 25px;
}

.animate__animated {
  animation-duration: 1.5s !important;
}
</style>

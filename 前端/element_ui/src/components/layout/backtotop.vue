<template>
     <transition >
     <div class="back-to-top" @click="toTop">
    </div>
     </transition>
    
    </template>
    
    <script>
    
    
    export default {
    name: "BackToTop",
     methods: {
        toTop() {
          window.scrollTo({top: 0, behavior: 'smooth'})
        }
      }
    }
    </script>
    
    <style scoped>
    
    .back-to-top {
      cursor: pointer;
      position: fixed;
      right: 25px;
      top: 650px;
      z-index: 999;
      width: 60px;
      height: 60px;
      background-color: #6410d2;  
      /* background-image: url('../assets/images/to-top04.svg');

        /* height: 100px;
        width: 100px;
        background-color: #89b4c2;
        box-shadow: 0 0 6px rgba(0,0,0, .12);
        text-align: center;
        line-height: 40px;
        color: #1989fa; */
    }
    
    .animate__animated{
      animation-duration: 1.5s !important;
    }
    
    </style>
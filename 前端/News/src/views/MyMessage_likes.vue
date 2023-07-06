<template>
    <div class="MyMessage_likes">
        <div class="routeback">
        <el-button class= 'back' @click="routeBack" >
            <img src="../assets/icons/back.svg" alt="返回">
        </el-button>
        </div>
        <div class="set">
            <el-row>
                <el-col :span="6">
                    <personal_nav />
                </el-col>
                <el-col :span="18">
                    <myRelikes />
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
import personal_nav from '../components/user/personal_nav.vue';
import myRelikes from '../components/user/MyRe_like.vue';
export default {
    components: {
        personal_nav,
        myRelikes
    },
    methods: {
    routeBack() {
      this.$router.push({ path: '/' })
      window.close();
    }
  },
}
</script>


<style scoped>
.MyMessage_likes {
  font-family: 'sihan-regular', sans-serif;
}
.set {
    display: block;
    width: 1200px;
    background-color: aliceblue;
    margin: 0 auto;
}
.routeback {
  margin-top: 15px;
  margin-left: 15px;
}

.back {
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
.back img {
  height: 25px;
  width: 25px;
}
</style>
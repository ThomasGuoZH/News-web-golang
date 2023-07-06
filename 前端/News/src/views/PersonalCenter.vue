<template>
    <div class="personalcenter">
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
                    <info />
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
import personal_nav from '../components/user/personal_nav.vue';
import info from '../components/user/MyInformation.vue';
export default {
    data() {
    return {
      channel: '',
      pathFromUrl: '',
    }
  },
    components: {
        personal_nav,
        info
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
  background-color: #F9F7F7;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
.back img {
  height: 25px;
  width: 25px;
}
.personalcenter {
  font-family: 'sihan-regular', sans-serif;
}
</style>
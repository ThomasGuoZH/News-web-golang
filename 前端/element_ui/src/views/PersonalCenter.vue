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
                    <myMessage />
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
import personal_nav from '../components/user/personal_nav.vue';
import myMessage from '../components/user/myMessage.vue';
export default {
    data() {
    return {
      items: [
        { name: '首页', url: '/home' },
        { name: '时政', url: '/politics' },
        { name: '科技', url: '/science' },
        { name: '娱乐', url: '/entertainment' },
        { name: '体育', url: '/sports' },
        { name: '军事', url: '/military' },
        { name: '教育', url: '/education' },
        { name: '个人中心',url: '/'}],
      channel: '',
      pathFromUrl: '',
    }
  },
    components: {
        personal_nav,
        myMessage
    },
    computed: {
    getchannel() {
      if (this.channel === '' || this.channel === 'home') {
        return '首页';
      } else if (this.channel === 'education') {
        return '教育';
      } else if (this.channel === 'science') {
        return '科技';
      } else if (this.channel === 'politics') {
        return '时政';
      } else if (this.channel === 'sports') {
        return '体育';
      } else if (this.channel === 'military') {
        return '军事';
      } else if (this.channel === 'entertainment') {
        return '娱乐';
      }
    },
    channelname() {
      return this.getchannel;
    }
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.pathFromUrl = from.fullPath;
      vm.channel = to.params.channel;
    })
  },
    methods: {
    routeBack() {
      this.$router.push({ path: this.pathFromUrl })
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
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
.back img {
  height: 25px;
  width: 25px;
}
</style>
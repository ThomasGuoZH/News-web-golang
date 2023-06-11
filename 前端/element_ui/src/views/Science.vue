<template>
    <div class="science">
        <navigation :items="items">
            <template v-slot:科技>
                <b>{{ items[2].name }}</b>
            </template>
        </navigation>
        <sidebox />
        <backtotop />
        <div class="main">
            <el-row gutter="20">
                <el-col span="16">
                    <div class="news-box">
                        <el-col span="24">
                            <div class="m1">
                                0
                            </div>
                        </el-col>
                        <el-col span="12">
                            <div class="m1-1">
                                1
                            </div>
                        </el-col>
                        <el-col span="12">
                            <div class="m1-1">
                                2
                            </div>
                        </el-col>
                    </div>
                </el-col>
                <el-col span="8">
                    <div class="side">
                        <el-col span="24">
                            <div class="m2">
                                3
                            </div>
                        </el-col>
                        <el-col span="24">
                            <div class="m2-1">
                                4
                            </div>
                        </el-col>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script>
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
export default {
    name: 'science',
    components: {
        navigation,
        sidebox,
        backtotop
    },
    data() {
        return {
            items: [
                { name: '首页', url: '/home' },
                { name: '时政', url: '/politics' },
                { name: '科技', url: '/science' },
                { name: '娱乐', url: '/entertainment' },
                { name: '体育', url: '/sports' },
                { name: '军事', url: '/military' },
                { name: '教育', url: '/education' }],
        }
    }
}
</script>

<style>
body {
    background-color: rgb(229, 242, 245);
}

* {
    margin: 0;
    padding: 0;
}
.el-col {
    margin-bottom: 20px;
    border-radius: 4px;
  } 
.el-col:last-child {
  margin-bottom: 0;
}
.main{
    width:1200px;
    margin:100px auto 0;
    background-color: rgb(247, 239, 239);
}
.m1{
    height: 600px;
    background-color:blueviolet;
}
.m1-1{
    height:200px;
    background-color:aquamarine;
}
.m2{
    height: 300px;
    background-color:bisque;
}
.m2-1{
    height: 500px;
    background-color: aqua;
}
</style>
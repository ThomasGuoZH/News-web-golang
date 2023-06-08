<template>
    <el-carousel :interval="4000" type="card" height="300px">
        <el-carousel-item v-for="(item, index) in picUrlList" :key="index">
            <img class="medium" :src="item.pic" router-link to="#">
            <a href="#" class="头条字体" target="_blank">{{ item.title }}</a>
        </el-carousel-item>
    </el-carousel>
</template>

<script>
import { getNewsList } from '../../api/getNewsList.js'
export default {
    data() {
        return {
            picUrlList: {}
        };
    },
    created: async function () {
        this.picUrlList = await getNewsList(6, '头条');
        console.log(this.picUrlList);
    }
}

</script>
<style>
.el-carousel {
    display: block;
    width: 800px;
    margin: auto;
}

.el-carousel-item img {
    display: block;
    width: 100%;
    height: 100%;
}

.medium {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.头条字体 {
    display: block;
    text-decoration: none;
    text-align: center;
    width: 100%;
    height: 20px;
    color: black;
}

.头条字体 a:hover {
    color: blue;
}
</style>
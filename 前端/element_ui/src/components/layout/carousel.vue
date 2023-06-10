<template>
    <el-carousel :interval="4000" type="card" height="300px">
        <el-carousel-item v-for="(item, index) in picUrlList" :key="index">
            <router-link :to="'/home/newspage/' + item.title" class="route">
                <img class="medium" :src="item.pic">
            </router-link>
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

.medium {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.route h5 {
    display: block;
    text-decoration: none;
    text-align: center;
    width: 100%;
    height: 20px;
    color: black;
}

.route h5 :hover {
    color: blue;
}
</style>
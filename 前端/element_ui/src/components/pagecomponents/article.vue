<template>
    <div id="articleContent">

        <p v-for="item in articlelist" v-html="item.content"></p>

    </div>
</template>
<script>
import { getNewsList } from '@/api/getNewsList.js';
export default {
    name: "articleContent",
    data() {

        return {
            articlelist: {},
        };
    },
    created: async function () {
        this.articlelist = await getNewsList(10, '体育')
        console.log(this.articlelist)
    }
};

</script>
<style scoped>
div {
    margin-top: 100px;
    width: 1000px;
    height: 300px;
}

p {
    font-size: 20px;
}
</style>
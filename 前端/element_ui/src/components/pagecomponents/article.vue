<template>
    <div id="articleContent">

        <p v-for="item in articlelist" v-html="item.content"></p>

    </div>
</template>
<script>
import { getArticleList } from '@/api/article';
export default {
    name: "articleContent",
    data() {

        return {
            articlelist: {},
        };
    },
    created: async function () {
        this.articlelist = await getArticleList(10, '体育')
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
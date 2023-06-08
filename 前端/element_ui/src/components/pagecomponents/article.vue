<template>
    <div id="articleContent">

        <p v-for="item in articlelist">{{ item.weburl }}</p>

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
    mounted: function () {
        this.articlelist = getArticleList(10, '体育')
    }
};

</script>
<style scoped>
div {
    margin-top: 100px;
    width: 400px;
    height: 300px;
}

p {
    font-size: 20px;
}
</style>
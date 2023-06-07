<template>
    <div id="article">
        <button @click="getarticle"></button>
        <div>{{ article }}</div>
    </div>
</template>
<script>
import { getarticleContent } from '@/api/getarticle';
export default {
    name: '#article',
    data() {
        return {
            article: [],
            channel: '体育'
        };
    },
    method: {
        getarticle(channel) {
            getarticleContent(channel).then(
                res => { this.article = res; }
            )


        }
    }
}

</script>
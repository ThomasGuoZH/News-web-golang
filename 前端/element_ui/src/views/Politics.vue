<template>
    <div class="politics">
        <navigation :items="items">
            <template v-slot:item-时政>
                <b>{{ items[1].name }}</b>
            </template>
        </navigation>
        <sidebox />
        <backtotop />
    </div>
</template>
<script>
import sidebox from '../components/layout/sidebox.vue'
import backtotop from '../components/layout/backtotop.vue'
import navigation from '../components/layout/nav.vue'
export default {
    name: 'politics',
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
</style>
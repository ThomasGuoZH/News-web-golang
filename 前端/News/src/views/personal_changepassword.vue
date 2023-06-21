<template>
    <div class="set">
        <el-row>
            <el-col :span="6">
                <personal_nav />
            </el-col>
            <el-col :span="18">
                <Changepassword />
            </el-col>
        </el-row>
    </div>
</template>

<script>
import personal_nav from '../components/user/personal_nav.vue';
import Changepassword from '../components/user/changepassword.vue';
export default {
    components: {
        personal_nav,
        Changepassword
    }
}
</script>


<style>
.set {
    display: block;
    width: 1200px;
    background-color: aliceblue;
    margin: 0 auto;
}</style>
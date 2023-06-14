<template>
  <router-view />
</template>

<script>

import router from './router';
import axios from 'axios'
export default {
  name: 'app',
  components: {
    router
  }
}
</script>



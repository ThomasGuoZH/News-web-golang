<template>
    <div class="news">
      <h1>{{ article.title }}</h1>
      <p>{{ article.date }}</p>
      <p>{{ article.src }}</p>
      <div class="picture">
        <img :src="article.pic"  alt="image">
      </div>
      <div class="content">{{ article.content }}</div>
    </div>
  </template>
  
  <script>
  // import { getNewsList } from '@/api/getNewsList.js';
  export default {
    data() {
      return {
        article: {
          title: '官方：贝林厄姆加盟皇马 转会费总价1.3亿欧元',
          date: '2023-06-08',
          src: '新浪体育讯',
          pic: 'https://n.sinaimg.cn/sports/transform/361/w649h512/20230608/22b2-18b5cc1aeb65892bc3319b622b0e42f7.png',
          content: '<p class=\"art_p\">6月8日,多特蒙德官方宣布,球队中场贝林厄姆正式转会皇马,转会费总价为1.3亿,分为1亿固定转会费,以及3000万的浮动条款。',
        },
      };
    },
    // created: async function () {
    //     this.articlelist = await getNewsList(10, '体育')
    //     console.log(this.articlelist)
    // }
  };
  </script>
  
  <style scoped>
  .news {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
  }
  
  h1 {
    font-size: 24px;
    margin-bottom: 10px;
  }
  
  p {
    font-size: 14px;
    margin-bottom: 20px;
    color: #888;
  }
  
  .content {
    font-size: 16px;
    line-height: 1.6;
  }
  </style>
  
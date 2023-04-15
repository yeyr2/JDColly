<template>
  <div v-loading="loading">
    <div class="nav">
      <div class="content">
        <a href="/home">首页</a> > <a href="#">{{ arr.key }}</a> >
        {{ arr.shopProduct_id }}
      </div>
    </div>
    <div class="main">
      <div class="turn">
        <div class="echart">
          <button @click="goto">点此查看该商品评分</button>
        </div>
        <div class="jd">
          <button @click="open">点此跳转至京东商城</button>
        </div>
      </div>
      <div class="name">
        <h3>{{ arr.shopName }}</h3>
      </div>
      <div class="info">
        <div class="price">￥ {{ arr.shopPrice }}</div>
        <div class="title">{{ arr.shopTitle }}</div>
      </div>
      <div class="imag">
        <div class="img1"><img :src="arr.shopImgSrc" alt="商品图片" /></div>
        <div class="img2">
          <img :src="shopComments.analyze_word" alt="词云图片" />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "CommodityDetails",
  data() {
    return {
      arr: null,
      shopurl: "",
      shopProduct_id: "",
      analyze_word: "",
      shopComments: null,
      // loading: true,
    };
  },
  async created() {
    const arr = JSON.parse(window.localStorage.getItem("shopInfo"));
    this.arr = arr;
    this.arr = JSON.parse(window.localStorage.getItem("shopInfo"));
    const { data: res } = await this.$http.get("/new/comment", {
      params: {
        id: localStorage.getItem("id"),
        productId: this.arr.shopProduct_id,
        isColly: "0",
        startTime: "1639315501",
        lastTime: parseInt(new Date().getTime() / 1000) + "",
        token: localStorage.getItem("token"),
      },
    });
    this.shopComments = res.value;
    console.log(res.value);
    window.localStorage.setItem("shopComments", JSON.stringify(res.value));
  },
  methods: {
    open() {
      this.$confirm("前往京东商城前建议查看左侧的商品评分可视化图标", "提示", {
        confirmButtonText: "我已查看，确认前往",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$message({
            type: "success",
            message: "已前往!",
          });
          window.open(this.arr.shopURL, "_self");
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "暂不前往",
          });
        });
    },
    goto() {
      this.$router.push("/Echarts");
    },
  },
};
</script>
<style scoped>
.nav {
  width: 100%;
  background-color: #f2f2f2;
}
.nav .content {
  width: 70%;
  margin: 0 auto;
  font: 16px/200% tahoma, arial, Microsoft YaHei, Hiragino Sans GB,
    "\u5b8b\u4f53", sans-serif;
  color: #666;
}
.main {
  width: 60%;
  margin: 10px auto;
}
.main .turn {
  overflow: hidden;
  width: 100%;
  height: 70px;
}
.main .turn .echart {
  float: left;
}
.main .turn .echart button {
  display: block;
  width: 200px;
  height: 50px;
  color: black;
  font: 16px/200% tahoma, arial, Microsoft YaHei, Hiragino Sans GB,
    "\u5b8b\u4f53", sans-serif;
  background-color: aquamarine;
  /* background-color: red; */
}
.main .turn .echart button:hover {
  /* color: #ffffff; */
  font-size: 20px;
}
.main .turn .jd {
  float: right;
}
.main .turn .jd button {
  display: block;
  width: 200px;
  height: 50px;
  font: 16px/200% tahoma, arial, Microsoft YaHei, Hiragino Sans GB,
    "\u5b8b\u4f53", sans-serif;
}
.main .turn .jd button a {
  font: 16px/200% tahoma, arial, Microsoft YaHei, Hiragino Sans GB,
    "\u5b8b\u4f53", sans-serif;
}
.main .turn .jd button:hover {
  color: red;
}
.main .name {
  width: 100%;
  text-align: center;
}
.main .info {
  margin-top: 20px;
  height: 50px;
  line-height: 50px;
  width: 100%;
  overflow: hidden;
  padding: 0 100px;
  /* background-color: aqua; */
}
.main .info .price {
  float: left;
  color: red;
}
.main .info .title {
  float: right;
  color: red;
}
.main .imag {
  width: 100%;
  margin-top: 30px;
}
.main .imag .img1 {
  float: left;
  width: 400px;
  height: 400px;
}
.main .imag .img1 img {
  width: 400px;
  height: 400px;
  /* background-color: aqua; */
}
.main .imag .img2 {
  float: right;
  width: 450px;
  height: 450px;
}
</style>
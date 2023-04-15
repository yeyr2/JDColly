<template>
  <div>
    <div class="main">
      <ul>
        <li v-for="(item, index) in shopsarr" :key="index">
          <div class="box" @click="goto(index)">
            <div class="image">
              <img :src="item.shopImgSrc" alt="商品图片" />
            </div>
            <div class="name">{{ item.shopName }}</div>
            <div class="price">￥{{ item.shopPrice }}</div>
            <!-- <div class="isSort" v-if="isRank">{{ rank }}</div> -->
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>
<script>
export default {
  name: "commoditycenter",
  data() {
    return {
      shopsarr: [],
      isRank: true,
      rank: "",
      shopProduct_id: "",
    };
  },
  created() {
    this.shopsarr = JSON.parse(window.localStorage.getItem("shopsData"));
    console.log(this.shopsarr);
  },
  mounted() {
    let shopsData = JSON.parse(window.localStorage.getItem("shopsData"));
    this.shopProduct_Id = shopsData.shopProduct_id;
  },
  methods: {
    goto(index) {
      window.localStorage.setItem(
        "shopInfo",
        JSON.stringify(this.shopsarr[index])
      );
      this.$router.push(`/CommodityDetails`);
    },
  },
};
</script>
<style scoped>
.title {
  margin: 10px auto;
  width: 77vw;
  text-align: center;
}
.main {
  width: 79vw;
  margin: 10px auto;
  overflow: hidden;
}
.main ul li {
  position: relative;
  float: left;
  width: 15vw;
  /* height: 38vh; */
  height: 300px;
  background-color: aliceblue;
  margin: 10px 10px 20px 0;
  box-sizing: border-box;
  background-color: #fff;
  box-shadow: 5px 5px 10px #888888;
  transition: all 0.5s;
}
.main ul li:hover {
  top: -10px;
  opacity: 0.4;
}
.main ul li img {
  position: relative;
  width: 60%;
  height: 170px;
  top: 20px;
  left: 50%;
  margin-left: -30%;
  overflow: hidden;
  opacity: 1;
}
/* .main ul li img:hover {
  opacity: 0.5;
} */
.main ul li .name {
  width: 80%;
  margin: 40px auto 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
}
.main ul li .price {
  width: 40%;
  text-align: center;
  color: red;
}
.main ul li .isSort {
  position: absolute;
  width: 45px;
  height: 45px;
  background-color: rgb(0, 213, 255);
  top: 0;
  right: 0;
  text-align: center;
}
</style>
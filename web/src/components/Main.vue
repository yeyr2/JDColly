<template>
  <div>
    <el-container class="allBox">
      <el-header>
        <div class="navbar">
          <div class="nav_left"></div>
          <div class="nav_title">购物助手————商品评论评分</div>
          <div class="nav_right">
            <el-button
              @click="logout"
              icon="el-icon-switch-button"
              class="outBtn"
              type="danger"
              plain
              round
              >退出</el-button
            >
          </div>
        </div>
      </el-header>
      <el-main>
        <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b"
          @select="handleSelect"
        >
          <el-menu-item index="0" @click="toHome()">首页</el-menu-item>
          <el-menu-item index="1" @click="search(0)">iPhone</el-menu-item>
          <el-menu-item index="2" @click="search(1)">iPad</el-menu-item>
          <el-menu-item index="3" @click="search(2)">Watch</el-menu-item>
          <el-menu-item index="4" @click="search(3)">AirPods</el-menu-item>

          <el-menu-item index="5" style="margin-right: 100px">
            <i class="el-icon-search" />
            <el-input
              v-model="searchContent"
              size="mini"
              placeholder="搜索"
              @keyup.native.enter="searchColly()"
            />
          </el-menu-item>
          <el-menu-item
            index="6"
            :key="searchContent"
            @click="torecords()"
            style="float: right"
          >
            <el-tooltip content="搜索记录" effect="dark" placement="bottom">
              <i class="el-icon-caret-bottom">搜索记录</i>
            </el-tooltip>
          </el-menu-item>
          <el-menu-item index="7" @click="toperson()" style="float: right">
            <el-tooltip content="个人中心" effect="dark" placement="bottom">
              <i class="el-icon-user-solid">个人中心</i>
            </el-tooltip>
          </el-menu-item>
        </el-menu>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "mainHome",
  data() {
    return {
      shopsarr: {},
      searchContent: "",
      activeIndex: "0",
      keyWord: "",
      keyArr: ["iPhone", "iPad", "Watch", "AirPods"],
      iscolly: ["0", "1"],
    };
  },
  created() {
    this.$router.push(`/home`).catch(() => {});
  },
  methods: {
    // 首页
    toHome() {
      this.$router.push(`/home`);
    },
    // 非爬取搜索
    async search(index) {
      const { data: data } = await this.$http.get(`/new/colly`, {
        withCredentials: true,
        params: {
          id: localStorage.getItem("id"),
          key: this.keyArr[index],
          isColly: "1",
          token: localStorage.getItem("token"),
        },
      });
      console.log(data);
      window.localStorage.setItem("shopsData", JSON.stringify(data.value));
      this.$router.push(`CommodityCenter`);
    },
    // 爬取搜索
    async searchColly() {
      const { data: data } = await this.$http.get("/new/colly", {
        params: {
          id: localStorage.getItem("id"),
          key: this.searchContent,
          isColly: "0",
          token: localStorage.getItem("token"),
        },
      });
      console.log(data.value);
      this.shopsarr = data.value;
      window.localStorage.setItem("shopsData", JSON.stringify(data.value));
      this.$router.push(`CommodityCenter`);
    },
    // 搜索记录
    torecords() {
      this.$router.push(`/SearchRecords`);
    },
    // 个人中心
    async toperson() {
      const res = await this.$http.get("/new/userinfo", {
        params: {
          id: localStorage.getItem("id"),
          token: localStorage.getItem("token"),
        },
      });
      window.localStorage.setItem("userinfo", JSON.stringify(res.data.value));
      this.$router.push(`/PersonalCenter`);
    },
    // 退出
    logout() {
      this.$confirm("是否确认退出?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$message({
            type: "success",
            message: "退出成功!",
          });
          sessionStorage.removeItem("activePath");
          this.$router.push("/login");
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消退出",
          });
        });
    },
    handleSelect() {},
  },
};
</script>

<style scoped>
.el-header {
  padding: 0;
  margin: 0;
  transition: all 0.5s;
}
.navbar {
  width: 100%;
  height: 100%;
  background-color: #6777ef;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  color: #fff;
  color: black;
  cursor: pointer;
  padding: 0px 15px;
  box-sizing: border-box;
}

.nav_left {
  display: flex;
}
.outBtn {
  height: 40px;
}
.nav_title {
  font-size: 25px;
  letter-spacing: 3px;
  font-weight: 600;
}
.nav_right {
  display: flex;
  align-items: center;
}
.el-main {
  width: 100%;
  padding: 0 0;
  /* background-color: blue; */
  /* background-color: #f2f3f7; */
}
.el-menu {
  width: 80%;
  padding-left: 10%;
  box-sizing: border-box;
}
.el-menu-demo {
  width: 100%;
}
.el-menu .el-menu-item {
  margin-right: 4%;
}
</style>
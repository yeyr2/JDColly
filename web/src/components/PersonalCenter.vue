<template>
  <div>
    <div class="main">
      <div class="top">
        <!-- 头像 -->
        <div class="image">
          <img src="../assets/img/headSculpture.jpg" alt="" />
        </div>
        <!-- 昵称 -->
        <!-- <div class="nickname">哆啦A梦</div> -->
        <div class="nickname">{{ this.info.nickname }}</div>
      </div>
      <div class="info">
        <el-descriptions class="margin-top" :column="1" :size="size" border>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-place"></i>
              昵称
            </template>
            {{ this.info.nickname }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-user"></i>
              账号
            </template>
            {{ this.info.username }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-lock"></i>
              密码
            </template>
            <!-- {{ this.info }} -->
            ***
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-tickets"></i>
              性别
            </template>
            {{ this.info.sex }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-phone"></i>
              手机号
            </template>
            {{ this.info.phoneNumber }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-message"></i>
              邮箱
            </template>
            {{ this.info.email }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-location"></i>
              家庭住址
            </template>
            {{ this.info.address }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-s-check"></i>
              紧急联系人
            </template>
            {{ this.info.emergencyContact }}
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "PersonalCenter",
  data() {
    return {
      size: "",
    };
  },
  created() {
    this.info = JSON.parse(window.localStorage.getItem("userinfo"));
    console.log(this.info);
  },
  method: {},
};
</script>
<style>
.main {
  width: 60%;
  margin: 0 auto 0;
}
/* 头像部分 */
.main .top {
  width: 100%;
  height: 200px;
}
.main .top .image {
  width: 150px;
  height: 150px;
  margin: 0 auto;
  border-radius: 50% 50%;
}
.main .top .image img {
  width: 100%;
  height: 100%;
  border-radius: 50% 50%;
}
.main .top .nickname {
  width: 100px;
  height: 30px;
  text-align: center;
  line-height: 30px;
  margin: 0 auto;
}
/* 信息部分 */
.main .info {
  width: 800px;
  height: 400px;
  margin: 0 auto;
  /* background-color: aquamarine; */
}
</style>
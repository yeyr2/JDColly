<template>
  <div class="container">
    <div class="login">
      <h2>商品评分系统</h2>
      <el-form
        class="login_form"
        :model="loginForm"
        :rules="loginFormRules"
        ref="loginFormRef"
      >
        <el-form-item prop="username">
          <span>账号：</span>
          <input v-model="loginForm.username" placeholder="123" />
        </el-form-item>
        <el-form-item prop="password">
          <span>密码：</span>
          <input
            v-model="loginForm.password"
            type="password"
            placeholder="123"
          />
        </el-form-item>
        <div class="btn">
          <el-button type="primary" @click="login" class="btn_login"
            >登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;录</el-button
          >
          <el-button type="primary" class="btn_register"
            >还没有账号，前往注册</el-button
          >
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  name: "loginIn",
  data() {
    return {
      loginForm: {
        username: "",
        password: "",
      },
      loginFormRules: {
        username: [
          {
            required: true,
            message: "请输入账号(账号为手机号)",
            trigger: "blur",
          },
          {
            min: 3,
            max: 20,
            message: "用户名长度在3 到 20 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 3,
            max: 15,
            message: "密码长度在 3 到 15 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          `/new/login`,
          {
            username: this.loginForm.username,
            password: this.loginForm.password,
          },
          {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }
        );
        console.log(res);
        if (res.status_code !== 0) {
          return this.$message({
            type: "error",
            message: "登录失败！",
            duration: 3000,
          });
        }
        window.localStorage.setItem("token", res.token);
        window.localStorage.setItem("id", res.id);
        this.$message.success("登录成功！");
        this.$router.push("/main");
      });
    },
  },
};
</script>
<style scoped>
.container {
  padding: 0;
  margin: 0;
  height: 100vh;
  display: flex;
  justify-content: center;
  background-image: linear-gradient(#a18cd1 0%, #fbc2eb 100%);
  background-size: cover;
  flex: 1;
  align-items: center;
}

.login {
  text-align: center;
  margin: 0 auto;
  width: 600px;
  height: 520px;
  background-color: rgba(87, 86, 86, 0.2);
  border-radius: 25px;
  box-shadow: 5px 2px 35px -7px #ff9a9e;
}

.login h2 {
  margin-top: 40px;
  color: aliceblue;
  font-weight: 100;
}

.login_form {
  padding: 20px;
}

.login_form .el-form-item span {
  color: rgb(131, 220, 255);
  font-size: 18px;
  font-weight: 100;
}

.login_form input {
  background-color: transparent;
  width: 320px;
  padding: 2px;
  text-indent: 2px;
  color: white;
  font-size: 20px;
  height: 45px;
  margin: 30px 30px 30px 5px;
  padding-left: 20px;
  outline: none;
  border: 0;
  border-bottom: 1px solid rgb(131, 220, 255);
}

input::placeholder {
  color: #fbc2eb;
  font-weight: 100;
  font-size: 18px;
  font-style: italic;
}

.btn_login {
  background-color: rgba(255, 255, 255, 0.582);
  border: 1px solid rgb(190, 225, 255);
  padding: 10px;
  width: 240px;
  height: 60px;
  border-radius: 30px;
  font-size: 30px;
  color: rgb(100, 183, 255);
  font-weight: 100;
  margin-top: 15px;
  box-sizing: border-box;
  vertical-align: middle;
}
.btn_register {
  background-color: rgba(255, 255, 255, 0.582);
  border: 1px solid rgb(190, 225, 255);
  padding: 10px;
  width: 300px;
  height: 60px;
  border-radius: 30px;
  font-size: 20px;
  /* color: rgb(100, 183, 255); */
  color: #150f5a;
  font-weight: 100;
  margin-top: 15px;
  box-sizing: border-box;
  vertical-align: middle;
}
.login_btn:hover {
  box-shadow: 2px 2px 15px 2px rgb(190, 225, 255);
  background-color: transparent;
  color: white;
  /* 选择动画 */
  animation: login_mation 0.5s;
}

/* 定义动画 */
@keyframes login_mation {
  from {
    background-color: rgba(255, 255, 255, 0.582);
    box-shadow: 0px 0px 15px 2px rgb(190, 225, 255);
  }

  to {
    background-color: transparent;
    color: white;
    box-shadow: 2px 2px 15px 2px rgb(190, 225, 255);
  }
}
</style>
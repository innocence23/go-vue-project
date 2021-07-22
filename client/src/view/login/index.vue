<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
          <img class="login_panle_form_title_logo" :src="$GIN_VUE_ADMIN.appLogo" alt=""><p class="login_panle_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form
          ref="loginForm"
          :model="loginForm"
          :rules="rules"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入用户名">
              <i slot="suffix" class="el-input__icon el-icon-user" />
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <i
                slot="suffix"
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
              />
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                width="100%"
                height="100%"
                alt="请输入验证码"
                @click="loginVerify()"
              >
            </div>
          </el-form-item>
          <div />
          <el-form-item>
            <el-button
              type="primary"
              style="width: 100%;"
              @click="submitForm"
            >登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panle_right" />
      <div class="login_panle_foot">
        <div class="copyright">Copyright &copy; {{ curYear }} </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }
    return {
      curYear: 0,
      lock: 'lock',
      loginForm: {
        username: 'admin',
        password: '123456',
        captcha: '',
        captchaId: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
      },
      logVerify: '',
      picPath: ''
    }
  },
  created() {
    this.loginVerify()
    this.curYear = new Date().getFullYear()
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async login() {
      return await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async(v) => {
        if (v) {
          const flag = await this.login()
          if (!flag) {
            this.loginVerify()
          }
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          this.loginVerify()
          return false
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath
        this.loginForm.captchaId = ele.data.captchaId
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>

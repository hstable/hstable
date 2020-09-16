<template>
  <div class="login-wrapper">
    <div style="margin: 5vh 0; padding: 0 16px">
      <h1>校园用户登录</h1>
      <p class="vice-info">
        我们不会存储您的密码，登录时将使用校内接口验证密码并自动同步课表
      </p>
    </div>
    <van-form @submit="onSubmit">
      <field
        id="username"
        v-model="username"
        name="username"
        label="学号"
        placeholder="学号"
        :rules="[{ required: true, message: '请填写学号' }]"
      />
      <field
        id="password"
        v-model="password"
        type="password"
        name="password"
        label="密码"
        placeholder="密码"
        :rules="[{ required: true, message: '请填写密码' }]"
      />
      <div style="margin: 5vh 16px 0">
        <p class="vice-info" style="margin: 20px 0">
          浏览器开启隐私模式可能会导致登录状态的遗失
        </p>
        <van-button block hairline plain type="info" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>
    <p style="text-align: center; margin: 5vh 0; opacity: 0.65">
      <a href="https://github.com/hstable/hstable">HSTable 开源于 GitHub</a>
    </p>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data: () => ({
    username: '',
    password: '',
    Username: null,
    Password: null,
    alignInterval: null,
  }),
  mounted() {
    this.Username = document.querySelector('#username')
    this.Password = document.querySelector('#password')
    // ugly but no choice
    this.alignInterval = setInterval(this.checkAlign, 100)
  },
  beforeDestroy() {
    clearInterval(this.alignInterval)
  },
  methods: {
    checkAlign() {
      this.$nextTick(() => {
        if (this.Username.value !== this.username) {
          this.username = this.Username.value
        }
        if (this.Password.value !== this.password) {
          this.password = this.Password.value
        }
      })
    },
    onSubmit() {
      this.$axios({
        url: '/api/login',
        method: 'post',
        data: { account: this.username, password: this.password },
      })
        .then((res) => {
          const { data } = res
          if (data.code === 200) {
            localStorage.token = 'Bearer ' + data.token
            this.$cookies.set('token', 'Bearer ' + data.token, { expires: 30 })
            // console.log(document.cookie)
            this.$nextTick(() => {
              // console.log(2, document.cookie)
              this.$router.push('/')
            })
          }
        })
        .catch((err) => {
          if (err.response) {
            if (err.response.status === 401) {
              this.$Notify({
                type: 'warning',
                message: '密码错误，请检查后重试',
              })
              return
            }
          }
          this.$Notify({
            type: 'warning',
            message: err.message,
          })
        })
    },
  },
}
</script>

<style lang="scss" scoped>
.login-wrapper {
  padding: 5vh 5vw 0;
}

.vice-info {
  font-size: 1.2em;
  margin: 10px 0;
  color: rgba(0, 0, 0, 0.5);
}
</style>

<template>
  <div class="login-wrapper">
    <div style="margin: 5vh 0; padding: 0 16px">
      <h1>校园用户登录</h1>
      <p class="vice-info">我们不存储您的密码，首次登录将自动同步课表</p>
    </div>
    <van-form @submit="onSubmit">
      <field
        v-model="username"
        name="学号"
        label="学号"
        placeholder="学号"
        :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <field
        v-model="password"
        type="password"
        name="密码"
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
import { Form as VanForm, Field, Button as VanButton } from 'vant'
export default {
  name: 'Login',
  components: { VanForm, Field, VanButton },
  data: () => ({
    username: '',
    password: '',
  }),
  methods: {
    onSubmit() {},
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

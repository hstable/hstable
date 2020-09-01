<template>
  <div>
    <tabbar
      v-model="active"
      style="z-index: 10"
      active-color="#07c160"
      inactive-color="#000"
    >
      <tabbar-item icon="home-o">课表</tabbar-item>
      <tabbar-item icon="setting-o">我的</tabbar-item>
    </tabbar>
    <div
      style="
        padding-bottom: 100px;
        background-color: #f5f5f5;
        min-height: 100vh;
      "
    >
      <course-calender v-show="active === 0"></course-calender>
      <div v-show="active === 1" style="padding: 50px 0 100px">
        <div style="display: flex; justify-content: center">
          <van-image
            round
            width="150px"
            height="150px"
            :src="require('~/assets/avatar.jpg')"
          />
        </div>
        <cell-group title="当前">
          <cell title="学号" value="20S22222" />
          <cell title="姓名" value="王五" />
          <cell title="课表同步时间" value="2020年8月31日" />
        </cell-group>
        <cell-group title="便民">
          <cell title="校园地图" is-link />
        </cell-group>
        <cell-group title="操作">
          <cell title="同步课表" is-link />
        </cell-group>
        <cell-group title="状态">
          <cell title="登出" is-link @click="handleLogout" />
        </cell-group>
      </div>
    </div>
  </div>
</template>

<script>
import {
  Tabbar,
  TabbarItem,
  Cell,
  CellGroup,
  Image as VanImage,
  Grid,
  GridItem,
} from 'vant'
import CourseCalender from '~/components/courseCalendar'

export default {
  components: {
    CourseCalender,
    Tabbar,
    TabbarItem,
    Cell,
    CellGroup,
    VanImage,
    // eslint-disable-next-line vue/no-unused-components
    Grid,
    // eslint-disable-next-line vue/no-unused-components
    GridItem,
  },
  data: () => ({
    active: 0,
  }),
  methods: {
    handleLogout() {
      localStorage.removeItem('token')
      this.$router.push('/login')
    },
  },
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>

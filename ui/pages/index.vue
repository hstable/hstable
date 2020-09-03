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
      style="padding-bottom: 75px; background-color: #f5f5f5; min-height: 100vh"
    >
      <course-calender
        v-show="active === 0"
        :data="courseCalendarData"
      ></course-calender>
      <div v-show="active === 1" style="padding: 25px 0 0">
        <div style="display: flex; justify-content: center">
          <van-image
            round
            width="100px"
            height="100px"
            :src="require('~/assets/avatar.jpg')"
          />
        </div>
        <cell-group title="当前">
          <cell title="学号" :value="account" />
          <cell title="姓名" :value="name" />
          <cell title="课表同步时间" :value="latestUpdate" />
        </cell-group>
        <cell-group title="便民">
          <cell title="研究生教务平台" is-link url="http://jw.hitsz.edu.cn" />
          <cell title="信息门户" is-link url="http://portal.hitsz.edu.cn/" />
          <cell title="图书馆" is-link url="https://lib.utsz.edu.cn/" />
          <cell
            title="校园地图"
            is-link
            :url="require('~/assets/img/map.png')"
          />
        </cell-group>
        <cell-group title="操作">
          <cell title="同步课表" is-link @click="showSyncPopup = true" />
        </cell-group>
        <cell-group title="状态">
          <cell title="登出" is-link @click="handleLogout" />
        </cell-group>
        <p style="text-align: center; margin: 5vh 0; opacity: 0.65">
          <a href="https://github.com/hstable/hstable">HSTable 开源于 GitHub</a>
        </p>
      </div>
    </div>
    <van-popup
      v-model="showSyncPopup"
      position="top"
      :style="{ height: '175px' }"
    >
      <van-form style="padding: 5vw" @submit="handleSubmitSync">
        <field :value="this.account" name="学号" label="学号" disabled />
        <field
          v-model="password"
          type="password"
          name="密码"
          label="密码"
          placeholder="密码"
        />
        <van-button
          style="margin-top: 12px"
          block
          hairline
          type="primary"
          native-type="submit"
        >
          同步课表
        </van-button>
      </van-form>
    </van-popup>
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
  Popup as VanPopup,
  Form as VanForm,
  Field,
  Button as VanButton,
  Notify,
} from 'vant'
import dayjs from 'dayjs'
import CourseCalender from '~/components/courseCalendar'

export default {
  middleware: 'auth',
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
    VanPopup,
    VanForm,
    Field,
    VanButton,
  },
  data: () => ({
    active: 0,
    courseCalendarData: [],
    showSyncPopup: false,
    account: '',
    password: '',
    latestUpdate: '',
  }),
  computed: {
    name() {
      if (this.courseCalendarData.length) {
        return this.courseCalendarData[0].xm
      } else {
        return ''
      }
    },
  },
  created() {
    this.getCourses()
  },
  methods: {
    handleLogout() {
      localStorage.removeItem('token')
      this.$router.push('/login')
    },
    getCourses() {
      this.$axios({
        url: '/api/course',
      }).then((res) => {
        const { data } = res
        this.courseCalendarData = data.course.Course.yxkcList
        this.account = data.course.Student_number
        this.latestUpdate = dayjs(data.course.Last_sync).format(
          'YYYY-MM-DD HH:mm:ss'
        )
      })
    },
    handleSubmitSync() {
      this.$axios({
        url: '/api/course',
        method: 'put',
        data: { password: this.password },
      })
        .then((res) => {
          location.reload()
        })
        .catch(() => {
          Notify({ type: 'warning', message: '密码错误，请检查后重试' })
        })
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

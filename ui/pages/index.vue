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
      <client-only>
        <course-calender
          v-if="courseCalendarData && courseCalendarData.length"
          v-show="active === 0"
          :data="courseCalendarData"
        ></course-calender>
      </client-only>
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
        <field :value="account" name="学号" label="学号" disabled />
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
import dayjs from 'dayjs'
import { genAuth } from 'assets/js/cookieTool'
import CourseCalender from '~/components/courseCalendar'

// import { getFromCookie } from '~/assets/js/cookieTool'

function getCourses(axios, Authorization) {
  const headers = {}
  if (Authorization) {
    headers.Authorization = Authorization
  }
  return axios({
    url: '/api/course',
    headers,
  }).then((res) => {
    const { data } = res
    // console.log('length', data.course.Course.yxkcList.length)
    return {
      courseCalendarData: data.course.Course.yxkcList.map((x) => {
        x.sksj = ''
        return x
      }),
      account: data.course.Student_number,
      latestUpdate: dayjs(data.course.Last_sync).format('YYYY-MM-DD HH:mm:ss'),
    }
  })
}

export default {
  components: {
    CourseCalender,
  },
  async asyncData({ $axios, req }) {
    if (process.server && req) {
      // console.log(req.headers)
      return await getCourses($axios, genAuth(req))
    } else if (process.browser) {
      // console.log(document.cookie)
      return await getCourses($axios)
    }
    // 这里千万不能返回空{}，不然会再次进行更新
    // return {}
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
      if (this.courseCalendarData && this.courseCalendarData.length) {
        return this.courseCalendarData[0].xm
      } else {
        return ''
      }
    },
  },
  methods: {
    handleLogout() {
      this.$cookies.remove('token')
      localStorage.removeItem('token')
      this.$router.push('/login')
    },
    handleSubmitSync() {
      this.$Spin.show()
      this.$axios({
        url: '/api/course',
        method: 'put',
        data: { password: this.password },
      })
        .then((res) => {
          location.reload()
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
        .finally(() => {
          this.$Spin.hide()
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

<template>
  <div class="wrapper">
    <Row v-if="screenWidth > 1024" class="operation">
      <Col span="12">
        <Table
          ref="calendar"
          class="calendar"
          border
          disabled-hover
          :columns="calendarColumns"
          :data="calendar"
          :max-height="calendarHeight"
        ></Table>
      </Col>
      <Col span="12">
        <Table
          ref="calendarRaw"
          class="calendarRaw"
          stripe
          border
          :columns="columns"
          :data="data"
          size="default"
        ></Table>
      </Col>
    </Row>
    <div v-else>
      <Table
        ref="calendar"
        style="margin-bottom: 1.5rem"
        class="calendar"
        border
        disabled-hover
        :columns="calendarColumns"
        :data="calendar"
        :max-height="calendarHeight"
      ></Table>
      <Table
        ref="calendarRaw"
        class="calendarRaw"
        stripe
        border
        :columns="columns"
        :data="data"
        size="default"
      ></Table>
    </div>
    <div class="float-week" @click="showPicker = true">
      <span>第</span>{{ week }}<span>周</span>
    </div>
    <VanPopup v-model="showPicker" round position="bottom">
      <Picker
        show-toolbar
        :title="term"
        :columns="weeks"
        :default-index="week - 1"
        @cancel="showPicker = false"
        @confirm="onPickerConfirm"
      />
    </VanPopup>
  </div>
</template>

<script>
// import calendarMock from 'assets/js/calendarMock'
import dayjs from 'dayjs'

const tiptopMap = {
  一: 1,
  二: 2,
  三: 3,
  四: 4,
  五: 5,
  六: 6,
  七: 7,
  天: 7,
  日: 7,
}
const colorArray = [
  '#d50000',
  '#E91E63',
  '#9C27B0',
  '#673AB7',
  '#3F51B5',
  '#03A9F4',
  '#00BCD4',
  '#009688',
  '#4CAF50',
  '#FF5722',
  '#795548',
  '#607D8B',
  '#ffab00',
  '#37474f',
  '#00bfa5',
  '#c51162',
]

const MaxWeeks = 25
const MaxDurs = 14

function newMatrix() {
  const matrix = new Array(MaxWeeks)
  for (let week = 1; week < MaxWeeks; week++) {
    matrix[week] = new Array(8)
    for (let whatDay = 1; whatDay < 8; whatDay++) {
      matrix[week][whatDay] = new Array(MaxDurs)
    }
  }
  return matrix
}

function arrToInt(arr) {
  for (let i = 0; i < arr.length; i++) {
    arr[i] = parseInt(arr[i])
  }
  return arr
}

function generateFromRow(sksj) {
  const timeRoom = sksj.split(' ')
  const weeks = timeRoom[0].split('周')[0].split(',')
  const indexWhatDay = timeRoom[0].indexOf('星期')
  const whatDay = tiptopMap[timeRoom[0][indexWhatDay + '星期'.length]]
  const dur = timeRoom[0].split('节')[0].split('第')[1]
  const durFromTo = arrToInt(dur.split('-'))
  const matrix = newMatrix()
  for (const week of weeks) {
    const weekFromTo = arrToInt(week.split('-'))
    // console.log(weekFromTo, whatDay, durFromTo)
    if (weekFromTo.length === 1) {
      for (let i = durFromTo[0]; i <= durFromTo[1]; i++) {
        matrix[weekFromTo[0]][whatDay][i] = timeRoom[1]
      }
    } else {
      for (let i = weekFromTo[0]; i <= weekFromTo[1]; i++) {
        for (let j = durFromTo[0]; j <= durFromTo[1]; j++) {
          matrix[i][whatDay][j] = timeRoom[1]
          // console.log(i, whatDay, j, timeRoom[1])
        }
      }
    }
  }
  // console.log(matrix)
  return matrix
}

function mergedMatrix(a, b) {
  const c = newMatrix()
  for (let i = 1; i < MaxWeeks; i++) {
    for (let j = 1; j < 8; j++) {
      for (let k = 1; k < MaxDurs; k++) {
        c[i][j][k] = a[i][j][k] || b[i][j][k]
      }
    }
  }
  return c
}

export default {
  name: 'CourseCalender',
  // mixins: [calendarMock],
  props: {
    data: {
      type: Array,
      required: true,
      default: () => [],
    },
  },
  data() {
    return {
      calendarColumns: [
        {
          title: ' ',
          key: 'time',
          width: 50,
          align: 'center',
        },
        {
          title: '周一',
          key: 1,
          align: 'center',
        },
        {
          title: '周二',
          key: 2,
          align: 'center',
        },
        {
          title: '周三',
          key: 3,
          align: 'center',
        },
        {
          title: '周四',
          key: 4,
          align: 'center',
        },
        {
          title: '周五',
          key: 5,
          align: 'center',
        },
        {
          title: '周六',
          key: 6,
          align: 'center',
        },
        {
          title: '周日',
          key: 7,
          align: 'center',
        },
      ],
      weeks: [
        1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10,
        11,
        12,
        13,
        14,
        15,
        16,
        17,
        18,
        19,
      ],
      columns: [
        {
          title: '课程号',
          key: 'kcdm',
          align: 'center',
        },
        {
          title: '课程名',
          key: 'kcmc',
          align: 'center',
        },
        {
          title: '教师名',
          key: 'dgjsmc',
          align: 'center',
        },
        {
          title: '学分',
          key: 'xf',
          align: 'center',
        },
        {
          title: '上课时间',
          key: 'sksj',
          width: 150,
          align: 'center',
        },
      ],
      selected: '',
      calendar: [
        { time: '1' },
        { time: '2' },
        { time: '3' },
        { time: '4' },
        { time: '5' },
        { time: '6' },
        { time: '7' },
        { time: '8' },
        { time: '9' },
        { time: '10' },
        { time: '11' },
        { time: '12' },
      ],
      nodeMatrix: [],
      isStudent: false,
      calendarHeight: null,
      week: 1, // begin from 1
      matrix: newMatrix(),
      showPicker: false,
      termBeginTime: null,
    }
  },
  computed: {
    screenWidth() {
      return document.body.clientWidth
    },
    term() {
      if (this.data.length) {
        return this.data[0].xnxqmc
      } else {
        return ''
      }
    },
  },
  watch: {
    data(val) {
      this.updateDataView(val)
    },
  },
  mounted() {
    this.updateDataView()
  },
  methods: {
    updateDataView(val) {
      if (!val) {
        val = this.data
      }
      val.forEach((x) => {
        const result = x.kcxx.match(/<p>([^<]*?)<\/p>/g)
        if (result) {
          // result: ["<p>1-4,7-9周,星期四第5-6节 T6205</p>", "<p>6周,星期六第5-6节,T6205	1-5,7-9周,星期二第5-6节 T6205</p>"]
          result.forEach((r) => {
            r = r.substr('<p>'.length, r.length - '<p></p>'.length)
            const mat = generateFromRow(r)
            this.matrix = mergedMatrix(this.matrix, mat)
          })
          x.sksj = result
            .map((x) => x.substr('<p>'.length, x.length - '<p></p>'.length))
            .join('\t')
        } else {
          x.sksj = ''
        }
      })
      this.calendarHeight = document.body.clientHeight - 66
      val.forEach((x) => {
        if (x.ktxkjssj) {
          if (!this.termBeginTime || this.termBeginTime < dayjs(x.ktxkjssj)) {
            this.termBeginTime = dayjs(x.ktxkjssj)
          }
        }
      })
      this.termBeginTime = this.termBeginTime.subtract(7, 'day')
      this.week = Math.trunc(
        dayjs(Date.now() - this.termBeginTime) / 1000 / 3600 / 24 / 7 + 1
      )
      this.renderCalendar()
    },
    onPickerConfirm(val) {
      this.week = parseInt(val)
      this.showPicker = false
      this.renderCalendar()
    },
    renderCalendar({ raw = this.data, hover: { kcdm, sksj } = {} } = {}) {
      // raw是课程数组
      // clear
      // matrix下标从0开始
      const matrix = new Array(12)
      if (this.nodeMatrix.length) {
        for (let i = 0; i < this.nodeMatrix.length; i++) {
          for (let j = 0; j < this.nodeMatrix[j].length; j++) {
            this.nodeMatrix[i][j].firstChild.innerHTML = ''
            this.nodeMatrix[i][j].style = ''
          }
        }
      }
      for (let i = 0; i < 12; i++) {
        matrix[i] = new Array(7)
      }
      // 计算，并赋予颜色
      const colorArrayBak = Object.assign([], colorArray)

      function swap(arr, i, j) {
        ;[arr[i], arr[j]] = [arr[j], arr[i]]
      }

      let cnt = 0
      raw.forEach((obj, index) => {
        const x = Object.assign({}, obj)
        let ct = x.sksj.split('\t')
        // 每个不同的时间段
        if (!ct) {
          ct = [x.sksj]
        }
        ct = ct.filter((x) => !!x)
        for (let j = 0; j < ct.length; j++) {
          const indexWhatDay = ct[j].indexOf('星期')
          const whatDay = tiptopMap[ct[j][indexWhatDay + '星期'.length]]
          const timeArr = ct[j]
            .substr(indexWhatDay)
            .match(/\d+-\d+/g)[0]
            .split('-')
          const begin = parseInt(timeArr[0])
          const end = parseInt(timeArr[1])
          // 给课程初始化一个颜色
          let color
          if (
            kcdm === undefined ||
            sksj === undefined ||
            (obj.kcdm === kcdm && obj.sksj === sksj)
          ) {
            // 没有hover的元素，或hover的就是这个课程
            if (!raw[index]._color) {
              swap(
                colorArrayBak,
                cnt,
                Math.floor(Math.random() * (colorArrayBak.length - cnt) + cnt)
              )
              raw[index]._color = colorArrayBak[cnt++]
            }
            color = raw[index]._color
          } else {
            color = 'rgba(0,0,0,.2)' // 灰色
          }
          matrix[begin - 1][whatDay - 1] = {
            ...obj,
            color,
            _length: end - begin + 1,
          }
        }
      })
      // return
      // 开始渲染
      // eslint-disable-next-line no-unreachable
      let tr = document.querySelectorAll('.calendar .ivu-table-tbody tr')
      for (let i = 0; i < tr.length; i++) {
        const tds = tr[i].querySelectorAll('td:not(:first-child)')
        this.nodeMatrix.push(tds)
      }
      const rawTbody = document.querySelector('.calendarRaw .ivu-table-tbody')
      // 在.calendarRaw上挂特效函数
      rawTbody.onmouseleave = this.mouseLeaveTbody
      tr = rawTbody.querySelectorAll('tr')
      for (let i = 0; i < tr.length; i++) {
        tr[i].onmouseenter = () => this.mouseEnterTr(i)
      }

      for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[i].length; j++) {
          if (!matrix[i][j]) {
            continue
          }
          const room = this.matrix[this.week][j + 1][i + 1]
          if (!room) {
            continue
          }
          this.nodeMatrix[i][j].style.position = 'relative'
          this.nodeMatrix[i][j].firstChild.innerHTML = `
              <div class="node" style="
                height: ${matrix[i][j]._length}00%;
                background-color: ${matrix[i][j].color};
              ">
                <div>
                  <p style="font-weight:bold">${matrix[i][j].kcmc}</p>
                  <p style="font-weight:bold">[${matrix[i][j].dgjsmc}]</p>
                  <p style="font-weight:bold">@${room}</p>
                </div>
              </div>
              `
          this.nodeMatrix[i][j].style.color = 'white'
        }
      }
    },
    mouseEnterTr(line) {
      // console.log(line)
      this.renderCalendar({
        hover: {
          kcdm: this.$refs.calendarRaw.data[line].kcdm,
          sksj: this.$refs.calendarRaw.data[line].sksj,
        },
      })
    },
    mouseLeaveTbody() {
      // console.log('out')
      this.renderCalendar()
    },
  },
}
</script>

<style lang="scss">
.calendar {
  /*margin: 25px 32.5px;*/
  //.ivu-table-cell {
  //  height: 100%;
  //  padding: 0;
  //}
  td {
    text-align: center;
    height: 60px;
  }

  .node {
    position: absolute;
    top: 0;
    left: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    font-size: 0.8em;
    border-radius: 5px;
  }

  $border-color: #e8eaeca0;

  .ivu-table-border td,
  .ivu-table-border th {
    border-right: 1px solid $border-color;
  }

  .ivu-table td,
  .ivu-table th {
    border-bottom: 1px solid $border-color;
  }
}

.calendarRaw {
  .ivu-table-cell {
    padding: 0 5px !important;
  }
}

.float-week {
  height: 50px;
  width: 50px;
  position: fixed;
  right: 15px;
  bottom: 60px;
  background-color: #ff9462;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  font-size: 14px;
  border-radius: 50%;
  box-shadow: 0 0 5px 1px #ffb292;
  z-index: 10;
}
</style>

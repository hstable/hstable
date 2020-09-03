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
  Row,
  Col,
  Picker,
  Notify,
} from 'vant'
import 'vant/lib/index.css'
import Vue from 'vue'

// eslint-disable-next-line no-empty-pattern
export default ({}, inject) => {
  Vue.component('Tabbar', Tabbar)
  Vue.component('TabbarItem', TabbarItem)
  Vue.component('Cell', Cell)
  Vue.component('CellGroup', CellGroup)
  Vue.component('VanImage', VanImage)
  Vue.component('Grid', Grid)
  Vue.component('GridItem', GridItem)
  Vue.component('VanPopup', VanPopup)
  Vue.component('VanForm', VanForm)
  Vue.component('Field', Field)
  Vue.component('VanButton', VanButton)
  Vue.component('Row', Row)
  Vue.component('Col', Col)
  Vue.component('Picker', Picker)
  Vue.prototype.$Notify = Notify
  // Vue.use(Vant)
  inject('Notify', Notify)
}

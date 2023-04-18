import { createApp } from 'vue'
// import './style.css'

// 全局引入elementui-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'



// createApp(App).mount('#app')
const app = createApp(App)

app.use(ElementPlus)
app.mount('#app')

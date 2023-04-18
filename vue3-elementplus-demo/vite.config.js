import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(
      {
        // 开启 $ref
        refTransform: true,  
        reactivityTransform: true,
      }
    )
  ],
  server: {
    host: 'localhost', //本机ip
    port: 8080,  // 设置启动端口
    open: true,
    proxy: {
      '/user': {  //代理别名
        target: 'http://localhost:8090',   //代理目标值
      }
    },
  },
  base: "./"  // 打包相对路径
})

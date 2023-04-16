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
    port: 8080,  // 设置启动端口
  },
  base: "./"  // 打包相对路径
})

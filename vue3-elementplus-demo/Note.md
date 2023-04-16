
- 安装yarn
npm install -g yarn 
- 设置源
yarn config set registry https://registry.npm.taobao.org -g 
yarn config set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass -g
- yarn常用命令
```bash
安装yarn 
npm install -g yarn
安装成功后，查看版本号： 
yarn --version
创建文件夹 yarn 
md yarn
进入yarn文件夹 
cd yarn
初始化项目 
yarn init // 同npm init，执行输入信息后，会生成package.json文件
yarn的配置项： 
yarn config list // 显示所有配置项
yarn config get <key> //显示某配置项
yarn config delete <key> //删除某配置项
yarn config set <key> <value> [-g|--global] //设置配置项
安装包： 
yarn install //安装package.json里所有包，并将包及它的所有依赖项保存进yarn.lock
yarn install --flat //安装一个包的单一版本
yarn install --force //强制重新下载所有包
yarn install --production //只安装dependencies里的包
yarn install --no-lockfile //不读取或生成yarn.lock
yarn install --pure-lockfile //不生成yarn.lock
添加包（会更新package.json和yarn.lock）：

yarn add [package] // 在当前的项目中添加一个依赖包，会自动更新到package.json和yarn.lock文件中
yarn add [package]@[version] // 安装指定版本，这里指的是主要版本，如果需要精确到小版本，使用-E参数
yarn add [package]@[tag] // 安装某个tag（比如beta,next或者latest）
//不指定依赖类型默认安装到dependencies里，你也可以指定依赖类型：

yarn add --dev/-D // 加到 devDependencies
yarn add --peer/-P // 加到 peerDependencies
yarn add --optional/-O // 加到 optionalDependencies
//默认安装包的主要版本里的最新版本，下面两个命令可以指定版本：

yarn add --exact/-E // 安装包的精确版本。例如yarn add foo@1.2.3会接受1.9.1版，但是yarn add foo@1.2.3 --exact只会接受1.2.3版
yarn add --tilde/-T // 安装包的次要版本里的最新版。例如yarn add foo@1.2.3 --tilde会接受1.2.9，但不接受1.3.0
发布包

yarn publish
移除一个包 
yarn remove <packageName>：移除一个包，会自动更新package.json和yarn.lock
更新一个依赖 
yarn upgrade 用于更新包到基于规范范围的最新版本
运行脚本 
yarn run 用来执行在 package.json 中 scripts 属性下定义的脚本
显示某个包的信息 
yarn info <packageName> 可以用来查看某个模块的最新版本信息
缓存 
yarn cache 
yarn cache list # 列出已缓存的每个包 
yarn cache dir # 返回 全局缓存位置 
yarn cache clean # 清除缓存



```

- 创建vite项目
$ yarn create vite
yarn create v1.22.19
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...

success Installed "create-vite@4.2.0" with binaries:
      - create-vite
      - cva
√ Project name: ... vue3-elementui-demo
√ Select a framework: » Vue
√ Select a variant: » JavaScript

Scaffolding project in D:\study\project\golang\web\5-vue3\6-Elementui\vue3-elementui-demo...

Done. Now run:

  cd vue3-elementui-demo
  yarn
  yarn dev

Done in 31.36s.

- 安装依赖
$ cd vue3-elementui-demo/
$ yarn

- 启动服务
$ yarn dev
  VITE v4.2.1  ready in 1240 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h to show help

// 浏览器访问: http://localhost:5173/


- 上传github
```bash
echo"# vue3-demo1" >> README.md
git init   // 初始化本地仓库
git add README.md  // 添加文件到暂存区
git commit -m "first commit"  // 将暂存区的文件提交到本地仓库
git branch -M main // 选择分支
git remote add origin git@github.com:it0114/vue3-demol.git // 添加远程仓库

git push -u origin main  // 将本地仓库main分支文件提交到远程仓库origin分支


git remote add origin git@github.com;it0114/vue3-demol.git 
git branch -M main
git push -u origin main


```


- 安装element-plus
npm install element-plus --save
yarn add element-plus

- 使用
// 完全使用
// main.ts
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'

const app = createApp(App)

app.use(ElementPlus)
app.mount('#app')

- 测试
```bash
# main.js
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




# App.vue
<script setup>

</script>

<template>

  <!-- 测试 -->
  <el-row class="mb-4">
    <el-button>Default</el-button>
    <el-button type="primary">Primary</el-button>
    <el-button type="success">Success</el-button>
    <el-button type="info">Info</el-button>
    <el-button type="warning">Warning</el-button>
    <el-button type="danger">Danger</el-button>
  </el-row>
</template>

<style scoped>

</style>

```


- 打包
yarn build
vscode 安装插件: Live Server
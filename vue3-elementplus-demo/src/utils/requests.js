/**
 * axios 二次封装
 * 1> 全局配置
 * 2> 响应拦截
 * 3> 请求方法
 */

// 引入文件
import axios from 'axios'
import { ElMessage } from 'element-plus'

// 全局配置
const service = axios.create({
    // baseURL: "http://localhost:8090/", // 后端服务
    baseURL: "/user", // 解决跨域,需要设置代理
    timeout: 5000,  // 请求超时时间
})

// request 拦截器
// 可以自请求发送前对请求做一些处理
// 比如统一加token，对请求参数统一加密
service.interceptors.request.use(config => {
    config.headers['Content-Type'] = 'application/json;charset=utf-8';
// let user = localStorage.getItem("user")?JSON.parse(localStorage.getItem("user")):{};
// if(user){
//     config.headers['token'] = user.token;  // 设置请求头
// }

    return config
}, error => {
    return Promise.reject(error)
});


// 响应拦截
service.interceptors.response.use(res => {
    const {code,data,message} = res.data

    if (code === 200) {
        // 请求成功
        ElMessage.success(message)
        return data;
    } else if(code === 404) {
        // 请求失败
        ElMessage.error(message)
    }
})


// request 方法
function request(options) {

    // 请求方法判断
    options.method = options.method || "get"  
    // 方法名大写转小写
    if (options.method.toLowerCase() === "get") options.params = options.data
    return service(options)
}
["get","post","put","delete"].forEach(iten =>{
    request[iten] = (url, data) =>{
        return request({
            url, 
            data,
            method: iten, 
        })
    }
})

export default request
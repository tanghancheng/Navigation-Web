import axios from 'axios'
export function request(config, success, failure) {
    //创建axios实例
    const instance = axios.creat({
        baseURL: 'http://192.168.14.146:8088',
        timeout: 50000
    })
    //发送真正的网络请求
    instance(config.baseconfig)
        .then(res => {
            // console.log(res) //要回调出去
            success(res) //回调
        })
        .catch(err => {
            // console.log(err) //要回调出去
            failure(err) //回调
        })
}
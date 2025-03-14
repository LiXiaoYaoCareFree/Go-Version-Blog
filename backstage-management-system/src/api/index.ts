import axios from "axios";
import {Message} from "@arco-design/web-vue";
import {userStorei} from "@/stores/user_store";

export interface baseResponse<T> {
    code: number;
    msg: string;
    data: T
}





export const useAxios = axios.create({
    timeout: 6000,
    baseURL: "", // 在使用前端代理的情况下，这里必须留空，不然会跨域
})



useAxios.interceptors.request.use((config) => {
    const userStore = userStorei()
    config.headers.set("token", userStore.userInfo.token)
    return config
})



useAxios.interceptors.response.use((res) => {
    if (res.status === 200) {
        return res.data
    }
    return res
}, (res) =>{
    console.log(2, res)
    Message.error(res.message)
})
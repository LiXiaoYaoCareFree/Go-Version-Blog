import {type baseResponse, useAxios} from "@/api/index";

export interface emailLoginRequest {
    user_name: string
    password: string
}


export function emailLoginApi(data: emailLoginRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/email_login", data)
}

export interface userInfoType {
    "id": number
    "created_at": string
    "nick_name": string
    "user_name": string
    "avatar": string
    "email": string
    "tel": string
    "addr": string
    "token": string
    "ip": string
    "role": string
    "sign_status": string
    "integral": number
    "sign": string
    "link": string
}
export function userInfoApi():Promise<baseResponse<userInfoType>>{
    return  useAxios.get("/api/user_info")
}
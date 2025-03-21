package redis_site

import "Blog-Server/global"

const key = "blogx_site_flow"

func SetFlow() {
	v, _ := global.Redis.Get(key).Int()
	global.Redis.Set(key, v+1, 0)
}

func GetFlow() int {
	v, _ := global.Redis.Get(key).Int()
	return v
}

func ClearFlow() {
	global.Redis.Del(key)
}


<script setup lang="ts">

import {reactive, ref} from "vue";
import {emailLoginApi, type emailLoginRequest} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";
import {userStorei} from "@/stores/user_store";
import router from "@/router";


const userStore = userStorei()

const form = reactive<emailLoginRequest>({
  user_name: "",
  password: ""
})

const formRef = ref()
async function emailLogin() {
  const val = await formRef.value.validate()
  if (val)return
  const res = await emailLoginApi(form)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success("登陆成功")
  // 如何获取用户信息 1. 直接解析token   2. 调用户信息接口

  userStore.saveUserInfo(res.data)
  router.push({
    name: "web"
  })

}



</script>

<template>
  <div class="login_view">
    <div class="login_mask">
      <a-form ref="formRef" :model="form" :label-col-props="{span: 0}" :wrapper-col-props="{span:24}">
        <div class="title">用户登录</div>
        <a-form-item label="用户名" field="user_name" :rules="[{required:true, message:'请输入用户名'}]">
          <a-input placeholder="用户名" v-model="form.user_name">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item label="密码" field="password" :rules="[{required:true, message:'请输入密码'}]">
          <a-input placeholder="请输入密码" type="password" v-model="form.password">
            <template #prefix>
              <icon-lock />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" @click="emailLogin" long>登陆</a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<style lang="less">
.login_view {
  background: url("https://raw.githubusercontent.com/LiXiaoYaoCareFree/Go-Version-Blog/refs/heads/dev-BMS/backstage-management-system/screen_photo/1.jpg") 50% / cover no-repeat;
  position: relative;
  height: 100vh;
  overflow: hidden;

  .login_mask {
    width: 400px;
    height: 100vh;
    //background-color: rgba(white, 0.8);
    background: transparent;
    position: absolute;
    right: 700px;
    top: 110px;
    display: flex;
    justify-content: center;
    align-items: center;

    .arco-form {
      padding: 40px;

      .arco-btn-primary {
        background-color: rgb(var(--primary-4));
        border: 1px solid transparent;
      }

      .title {
        font-size: 26px;
        font-weight: 600;
        text-align: center;
        color: @primary-2;
        margin-bottom: 10px;
      }
    }
  }
}
</style>
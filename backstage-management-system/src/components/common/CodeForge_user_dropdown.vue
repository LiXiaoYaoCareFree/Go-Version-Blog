<script setup lang="ts">

import router from "@/router";

function handleSelect(val: string) {
  if (val === "logout") {
    return
  }
  router.push({name: val})
}

interface OptionType {
  name: string
  title: string
}


const options: OptionType[] = [
  {title: "个人信息", name: "userInfo"},
  {title: "用户列表", name: "userList"},
  {title: "系统信息", name: "settings"},
  {title: "注销退出", name: "logout"},
]




</script>

<template>
  <a-dropdown @select="handleSelect" :popup-max-height="false">
    <div class="CodeForge_user_dropdown_com">
      <a-avatar
          image-url="https://avatars.githubusercontent.com/u/188313914?v=4"></a-avatar>
      <span class="user_name">LiXiaoYaoCareFree</span>
      <icon-down/>
    </div>

    <template #content>
      <a-doption v-for="item in options" :value="item.name">{{ item.title}}</a-doption>
    </template>
  </a-dropdown>
</template>

<style lang="less">
.CodeForge_user_dropdown_com {
  cursor: pointer;
  .user_name {
    margin: 0 5px;
  }

  svg {
    margin-right: 0!important;
  }
}

</style>
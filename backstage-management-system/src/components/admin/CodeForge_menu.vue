<script setup lang="ts">
import type {Component} from "vue";
import {IconHome, IconUser, IconSettings} from "@arco-design/web-vue/es/icon";
import CodeForge_component from "@/components/common/CodeForge_component.vue";


interface MenuType {
  title: string
  name: string
  icon?: string|Component
  children?: MenuType[]
}

const menuList : MenuType[] = [
  {title: "首页", name: "admin", icon: IconHome},
  {title: "个人中心", name: "userCenter", icon: IconUser, children: [
      {title: "用户信息", name: "userInfo"},
    ]},
  {title: "用户管理", name: "userManage", icon: IconUser, children: [
      {title: "用户列表", name: "userList"},
    ]},
  {title: "组件管理", name: "componentManage", icon: IconUser, children: [
      {title: "组件列表", name: "componentList"},
    ]},
  {title: "系统设置", name: "settingsManage", icon: IconSettings, children: [
      {title: "系统信息", name: "settingsInfo"},
    ]},
]



</script>

<template>
<div class="CodeForge_menu">
  <a-menu
      show-collapse-button>
    <template v-for="menu in menuList">
      <a-menu-item :key="menu.name" v-if="!menu.children">
        <template #icon>
          <component :is="menu.icon"></component>
        </template>
        {{ menu.title }}
      </a-menu-item>
      <a-sub-menu :key="'sub-menu-' + menu.name" v-else :title="menu.title">
        <template #icon>
        <CodeForge_component :is="menu.icon"></CodeForge_component>
        </template>
        <a-menu-item :key="sub.name" v-for="sub in menu.children">
          {{ sub.title }}
          <template #icon>
            <component :is="sub.icon"></component>
          </template>
        </a-menu-item>
      </a-sub-menu>
    </template>
  </a-menu>
</div>
</template>

<style scoped>

</style>
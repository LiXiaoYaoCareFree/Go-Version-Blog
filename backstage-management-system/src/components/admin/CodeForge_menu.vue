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
  {title: "系统设置", name: "settingsManage", icon: IconSettings, children: [
      {title: "系统信息", name: "settingsInfo"},
    ]},
]



</script>

<template>
  <div class="CodeForge_menu ">
    <div class="CodeForge_menu_inner scrollbar">
      <a-menu
          show-collapse-button>
        <template v-for="menu in menuList">
          <a-menu-item :key="menu.name" v-if="!menu.children">
            <template #icon>
              <CodeForge_component :is="menu.icon"></CodeForge_component>
            </template>
            {{ menu.title }}
          </a-menu-item>
          <a-sub-menu :key="menu.name" v-else :title="menu.title">
            <template #icon>
              <CodeForge_component :is="menu.icon"></CodeForge_component>
            </template>
            <a-menu-item :key="sub.name" v-for="sub in menu.children">
              {{ sub.title }}
              <template #icon>
                <CodeForge_component :is="sub.icon"></CodeForge_component>
              </template>
            </a-menu-item>
          </a-sub-menu>
        </template>
      </a-menu>
    </div>
  </div>
</template>

<style lang="less">
.CodeForge_menu {
  height: calc(100vh - 90px);

  &:hover {
    .arco-menu-collapse-button{
      opacity: 1 !important;
    }

  }


  .CodeForge_menu_inner.scrollbar {
    height: 100%;
    overflow-y: auto;
    overflow-x: hidden;

    .arco-menu{
      position: static;
      height: 100%;
      .arco-menu-collapse-button{
        top: 50%;
        transform: translate(-50%, -50%);
        left: 240px;
        transition: all .3s;
        opacity: 0;
      }
    }
  }
}
</style>
<script setup lang="ts">
import type {Component} from "vue";
import {IconHome, IconUser, IconSettings} from "@arco-design/web-vue/es/icon";
import CodeForge_component from "@/components/common/CodeForge_component.vue";
import {ref} from "vue";
import {collapsed} from "@/components/admin/CodeForge_menu";
import router from "@/router";
import {useRoute} from "vue-router"

interface MenuType {
  title: string
  name: string
  icon?: string|Component
  children?: MenuType[]
}

const route = useRoute()

const menuList: MenuType[] = [
  {title: "首页", name: "home", icon: IconHome},
  {
    title: "个人中心", name: "userCenter", icon: IconUser, children: [
      {title: "个人信息", name: "userInfo", icon: "iconfont icon-yonghuxinxi-"}
    ]
  },
  {
    title: "用户管理", name: "userManage", icon: "iconfont icon-yonghuguanli", children: [
      {title: "用户列表", name: "userManage", icon: "iconfont icon-yonghuguanli_huaban"}
    ]
  },
  {
    title: "组件管理", name: "componentManage", icon: "iconfont icon-xitongpeizhi", children: [
      {title: "组件列表", name: "componentList", icon: IconSettings}
    ]
  },
  {
    title: "系统设置", name: "settingsManage", icon: "iconfont icon-xitongpeizhi", children: [
      {title: "系统信息", name: "settingsManage", icon: IconSettings}
    ]
  },
]

function menuItemClick(key: string) {
  router.push({
    name: key
  })
}

</script>

<template>
  <div class="CodeForge_menu" :class="{collapsed: collapsed}">
    <div class="CodeForge_menu_inner scrollbar">
      <a-menu
          @menu-item-click="menuItemClick"
          v-model:collapsed="collapsed"
          :default-selected="[route.name]"
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


  &.collapsed {
    .arco-menu-collapse-button{
      left: 48px !important;
    }
  }
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
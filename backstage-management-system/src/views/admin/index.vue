<script setup lang="ts">
import CodeForge_theme from "@/components/common/CodeForge_theme.vue";
import CodeForge_screen from "@/components/common/CodeForge_screen.vue";
import CodeForge_menu from "@/components/admin/CodeForge_menu.vue";
import {collapsed} from "@/components/admin/CodeForge_menu";
import CodeForge_breadcrumb from "@/components/admin/CodeForge_breadcrumb.vue";
import CodeForge_user_dropdown from "@/components/common/CodeForge_user_dropdown.vue";
import router from "@/router"
import CodeForge_tabs from "@/components/admin/CodeForge_tabs.vue";
import CodeForge_logo from "@/components/admin/CodeForge_logo.vue";


function goHome() {
  router.push({name: "web"})
}


</script>

<template>
<div class="CodeForge_admin">
  <div class="CodeForge_aside" :class="{collapsed: collapsed}">
    <CodeForge_logo></CodeForge_logo>
    <CodeForge_menu></CodeForge_menu>
  </div>
  <div class="CodeForge_main">
    <div class="CodeForge_head">

      <CodeForge_breadcrumb></CodeForge_breadcrumb>
      <div class="CodeForge_actions">
          <span title="去首页" @click="goHome"><icon-home/></span>
          <CodeForge_theme></CodeForge_theme>
          <CodeForge_screen></CodeForge_screen>
        <CodeForge_user_dropdown>

        </CodeForge_user_dropdown>
      </div>
    </div>
    <CodeForge_tabs></CodeForge_tabs>
    <div class="CodeForge_container scrollbar">
      <router-view class="CodeForge_base_view" v-slot="{Component}">
        <transition name="fade" mode="out-in">
          <component :is="Component"></component>
        </transition>
      </router-view>
    </div>
  </div>
</div>
</template>

<style lang="less">


.CodeForge_admin {
  display: flex;
  background-color: var(--color-bg-1);
  color: @color-text-1;


  .CodeForge_aside {
    width: 240px;
    height: 100vh;
    border-right: 1px solid @color-border-1;
    background-color:   @color-fill-1;
    transition: width 0.3s;


    .CodeForge_logo {
      width: 100%;
      height: 90px;
      border-bottom: 1px solid var(--color-neutral-2);
    }

    &.collapsed {
      width: 48px;

      &~.CodeForge_main{
        width:calc(100% - 48px);
      }

    }
  }

  .CodeForge_main {
    width: calc(100% - 240px);
    transition: width .3s;

    .CodeForge_head {
      width: 100%;
      height: 60px;
      border-bottom: 1px solid var(--color-neutral-2);
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 20px;


      .CodeForge_actions{
        display: flex;
        align-items: center;
        svg{
          font-size: 18px;
          cursor: pointer;
          margin-right: 10px;
        }
        .CodeForge_user_dropdown_com {
          padding-left: 10px;
        }
      }
    }

    .CodeForge_tabs {
      width: 100%;
      height: 30px;
      border-bottom: 1px solid var(--color-neutral-2);
    }

    .CodeForge_container {
      width: 100%;
      height: calc(100vh - 90px);
      overflow-y: auto;
      overflow-x: hidden;
      background-color: @color-fill-2;
      padding: 20px;


    }
    .CodeForge_base_view {
      background-color: var(--color-bg-1);
      border-radius: 5px;
      height: 1000px;

    }
  }
}

// 组件刚开始离开
.fade-leave-active{

}
// 组件离开结束
.fade-leave-to {
  transform: translateX(30px);
  opacity: 0;
}

// 组件刚开始进入
.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

// 组件进入完成
.fade-enter-to {
  transform: translateX(0);
  opacity: 1;
}

// 正在进入和离开
.fade-leave-active, .fade-enter-active {
  transition: all .3s ease-out;
}

</style>
<script setup lang="ts">
import CodeForge_theme from "@/components/common/CodeForge_theme.vue";
import CodeForge_screen from "@/components/common/CodeForge_screen.vue";
import CodeForge_menu from "@/components/admin/CodeForge_menu.vue";
import {collapsed} from "@/components/admin/CodeForge_menu";
import CodeForge_breadcrumb from "@/components/admin/CodeForge_breadcrumb.vue";
import CodeForge_user_dropdown from "@/components/common/CodeForge_user_dropdown.vue";
import router from "@/router"

function goHome() {
  router.push({name: "web"})
}


</script>

<template>
<div class="CodeForge_admin">
  <div class="CodeForge_aside" :class="{collapsed: collapsed}">
    <div class="CodeForge_logo">

    </div>
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
    <div class="CodeForge_tabs">


    </div>
    <div class="CodeForge_container">
      <router-view></router-view>
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
    transition: width 0.3s;

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
      height: calc(100% - 90px);
      overflow-y: auto;
      overflow-x: hidden;
      background-color: @color-fill-2;
    }
  }
}



</style>
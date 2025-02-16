<script setup lang="ts">


import CodeForge_theme from "@/components/common/CodeForge_theme.vue";
import CodeForge_screen from "@/components/common/CodeForge_screen.vue";
</script>

<template>
<div class="CodeForge_admin">
  <div class="CodeForge_aside">
    <div class="CodeForge_logo">

    </div>
    <div class="CodeForge_menu">

    </div>
  </div>
  <div class="CodeForge_main">
    <div class="CodeForge_head">
      <div class="CodeForge_breadcrumbs"></div>
      <div class="CodeForge_actions">
          <icon-home />
          <CodeForge_theme></CodeForge_theme>
          <CodeForge_screen></CodeForge_screen>
          <div class="CodeForge_user_info_action">

          </div>
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
  background-color: @color-fill-1;
  color: @color-text-1;

  .CodeForge_aside {
    width: 240px;
    height: 100vh;
    overflow-y: auto;
    overflow-x: hidden;
    border-right: 1px solid @color-border-1;
    background-color:   @color-fill-1;


    .CodeForge_logo {
      width: 100%;
      height: 90px;
      border-bottom: 1px solid var(--color-neutral-2);
    }
  }

  .CodeForge_main {
    width: calc(100% - 240px);


    .CodeForge_head {
      width: 100%;
      height: 60px;
      border-bottom: 1px solid var(--color-neutral-2);
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 20px;
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
    }
  }
}



</style>
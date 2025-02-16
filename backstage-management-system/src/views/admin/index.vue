<script setup lang="ts">

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
          <icon-sun-fill />
          <icon-moon-fill />
          <icon-fullscreen-exit />
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


  .CodeForge_aside {
    width: 240px;
    height: 100vh;
    overflow-y: auto;
    overflow-x: hidden;
    border-right: 1px solid var(--color-neutral-2);


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
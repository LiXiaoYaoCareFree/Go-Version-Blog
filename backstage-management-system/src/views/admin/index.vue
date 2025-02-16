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

</style>
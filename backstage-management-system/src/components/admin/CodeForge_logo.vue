<script setup lang="ts">
import {collapsed} from "@/components/admin/CodeForge_menu";

</script>

<template>
  <div class="CodeForge_logo" :class="{collapsed: collapsed}">
    <img src="https://raw.githubusercontent.com/LiXiaoYaoCareFree/Go-Version-Blog/refs/heads/dev-BMS/CodeForge.jpg" class="logo" alt="logo">
    <div class="slogan">
      <div>CodeForge</div>
      <div>代码乐坊</div>
    </div>
  </div>
</template>

<style lang="less">
.CodeForge_logo{
  display: flex;
  justify-content: center;
  align-items: center;

  &.collapsed{
    .slogan{
      transform: scale(0);
      transform-origin: left;
      opacity: 0;
    }

    .logo {
      transform: translateX(65px);

    }
  }

  .logo{
    width: 50px;
    height: 50px;
    border-radius: 50%;
    object-fit: cover;
    cursor: pointer;
    transition: all .3s;

  }
  .slogan{
    margin-left: 10px;
    transition: all .3s;
    div:nth-child(1){
      font-size: 22px;
      font-weight: 600;
    }
    div:nth-child(2){
      font-size: 12px;
      margin-top: 1px;
      font-weight: 550;
    }
  }
}

</style>
<script setup lang="ts">
import {IconClose} from "@arco-design/web-vue/es/icon";
import {useRoute} from "vue-router";
import router from "@/router";
import {ref, watch} from "vue";
import {Swiper, SwiperSlide} from 'swiper/vue';
import {onMounted} from "vue";

const route = useRoute();


interface TabType {
  name: string
  title: string
}


function check(item: TabType){
  router.push({
    name: item.name,
  })
  saveTabs()
}

function saveTabs() {
  localStorage.setItem("CodeForge_tabs", JSON.stringify(tabs.value))
}

const tabs = ref<TabType[]>([
  {title: "首页", name: "home"},
  {title: "个人信息", name: "userInfo"},
  {title: "用户列表", name: "userList"},
  {title: "系统信息", name: "settings"},
])

function removeItem(item: TabType){
  if (item.name === "home") {
    return
  }
  const index = tabs.value.findIndex((value)=>item.name === value.name)
  if (index !== -1) {
    // 判断删除的元素是否是当前界面
    if (item.name === route.name) {
      router.push({
        name: tabs.value[index - 1].name
      })
    }
    tabs.value.splice(index, 1);
  }
  saveTabs()
}


function removeAllItem(item: TabType){
  tabs.value = [
    {title: "首页", name: "home"},
  ]
  router.push({name: "home"})
  saveTabs()
}

function loadTabs() {
  const CodeForge_tabs =  localStorage.getItem("CodeForge_tabs")
  if (CodeForge_tabs){
    try{
      tabs.value = JSON.parse(CodeForge_tabs)
    } catch(e) {
      console.log(e)
    }
  }
}

loadTabs()

watch(()=>route.name,()=>{
  // 判断当前路由的名称，如果不在tabs里面就加进去
  const index = tabs.value.findIndex((value) => route.name === value.name)
  if (index === -1) {
    tabs.value.push({
      name: route.name as string,
      title: route.meta.title,
    })
  }
}, {immediate: true})

const slidesCount = ref(100)
onMounted(()=>{
  //可显示的总宽度
  const swiperDom = document.querySelector(".CodeForge_tabs_swiper") as HTMLDivElement
  const swiperWidth = swiperDom.clientWidth
  //实际的总宽度
  const wrapperDom = document.querySelector(".CodeForge_tabs_swiper .swiper-wrapper") as HTMLDivElement
  const wrapperWidth = wrapperDom.scrollWidth

  if (swiperWidth > wrapperWidth) {
    return
  }
  //如果实际的总宽度大于了显示的总宽度
  const slideList = document.querySelectorAll(".CodeForge_tabs_swiper .swiper-slide")
  let allWidth = 0
  let index = 1

  for (const slideListElement of slideList) {
    allWidth += (slideListElement.clientWidth + 20)
    index++
    if (allWidth >= swiperWidth) {
      break
    }
  }
  slidesCount.value = index

  const activeSlide = document.querySelector(".CodeForge_tabs_swiper .swiper_slide.active") as HTMLDivElement
  if (activeSlide.offsetLeft > swiperWidth) {
    const offsetLeft = swiperWidth - activeSlide.offsetLeft

    setTimeout(()=>{
      wrapperDom.style.transform = `translate3d(${offsetLeft}px, 0px, 0px)`
    }, 1000)
  }

})


</script>

<template>
  <div class="CodeForge_tabs">
    <swiper class="CodeForge_tabs_swiper" :slide-per-view="slidesCount">
      <swiper-slide v-for="item in tabs" :class="{active: route.name === item.name}">
        <div class="item" @click="check(item)" :class="{active: route.name === item.name}">
          {{ item.title }}
          <span class="close" @click.stop="removeItem(item)" title="删除" v-if="item.name != 'home'">
        <IconClose></IconClose>
      </span>
        </div>
      </swiper-slide>
    </swiper>

    <div class="item" @click="removeAllItem(item)">
      删除全部
    </div>
  </div>
</template>

<style lang="less">
.CodeForge_tabs {
  display: flex;
  align-items: center;
  padding: 0 10px;
  justify-content: space-between;

  .swiper {
    width: calc(100% - 100px);
    display: flex;
    overflow-y: hidden;
    overflow-x: hidden;


    .swiper-wrapper {
      display: flex;
      align-items: center;

      .swiper-slide {
        width: fit-content !important;
        flex-shrink: 0;
      }
    }
  }


  .item {
    padding: 3px 7px;
    background-color: var(--color-bg-1);
    border: @CodeForge_border;
    margin-right: 10px;
    cursor: pointer;
    border-radius: 5px;
    flex-shrink: 0;
    &:hover {
      background-color: var(--color-fill-1);
    }

    &.active {
      background-color: @primary-6;
      color: white;
    }
  }
}

</style>
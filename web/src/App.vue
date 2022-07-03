<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import { defineComponent } from 'vue'
import PomoItem from '@/components/Pomo.vue'
import TodoList from '@/components/Todo.vue'
import PanelGroup from '@/components/PanelGroup.vue'
import StatChart from '@/components/dataStatistics/StatChart.vue'
import PomoHistory from '@/components/dataStatistics/PomoHistory.vue'
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
const store = useMainStore()

let { analyzeComponent } = storeToRefs(store)
</script>



<template>
  <section class="container">
    <!-- 左容器 -->
    <section class="itemLeft">
      <PomoItem />
    </section>

    <section class="itemRight">
      <TodoList />
    </section>
    <!-- 右容器 -->

  </section>

  <section class="summary">
    <PanelGroup />

  </section>

  <section class="analyze">
    <div v-if="analyzeComponent === 'Stat'">
      <StatChart />
    </div>
    <div v-else-if="analyzeComponent === 'B'">
    </div>
    <div v-else-if="analyzeComponent === 'PomoHistory'">
      <PomoHistory />
    </div>
    <div v-else>
    </div>
  </section>

</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  margin-top: 60px;
  margin-left: 10%;
  margin-right: 10%;
}
</style>

<style lang="less" >
// 大容器的样式
.container {
  // 最大最小的宽度
  // min-width: 1200px;
  // max-width: 2048px;
  // padding: 0.125rem 0.125rem 0;
  // background-color: gray;
  display: flex;
  // 设置左右在页面的份数
  .itemLeft,
  .itemRight {
    flex: 1;
  }
}
</style>
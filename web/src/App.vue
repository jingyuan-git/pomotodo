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
    <section class="itemLeft">
      <PomoItem />
    </section>

    <section class="itemRight">
      <TodoList />
    </section>
  </section>

  <section class="panel">
    <PanelGroup />

  </section>

  <section class="analyze">
    <div v-if="analyzeComponent === 'Stat'">
      <StatChart />
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
.container {
  display: flex;
  .itemLeft,
  .itemRight {
    flex: 1;
  }
}
</style>
<template>
  <div class="demo-date-picker">
    <div class="block">
      <!-- <p>Component value: {{ value }}</p> -->
      <el-date-picker v-model="value"
                      type="daterange"
                      start-placeholder="Start date"
                      end-placeholder="End date"
                      :default-time="defaultTime" />
    </div>
  </div>

  <div>
    <slot></slot>
  </div>

</template>


<script lang="ts">
import * as echarts from 'echarts'
import { inject, onMounted, reactive, ref, watch, defineComponent } from 'vue'
import axios from 'axios'
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
import { isNil } from 'lodash'

export default defineComponent({
  setup() {
    const value = ref('')
    const defaultTime = ref([
      new Date(2000, 1, 1, 0, 0, 0),
      new Date(2000, 2, 1, 23, 59, 59),
    ])
    const store = useMainStore()

    let { pomoTotal } = storeToRefs(store)
    let chart: echarts.ECharts | null = null
    watch(
      value,
      (newValue: any, oldValue: any) => {
        console.log('value变化了', newValue, oldValue)
        store.counts(
          newValue[0].valueOf().toString(),
          newValue[1].valueOf().toString()
        )
      },
      { deep: true }
    )

    // onMounted(() => {
    //   chart = echarts.init(document.getElementById('main') as HTMLDivElement)
    //   // 调用请求
    //   chart.setOption({
    //     // title: {
    //     //   text: 'ECharts 入门示例',
    //     // },
    //     tooltip: {},
    //     type: 'category',
    //     xAxis: {
    //       data: pomoTotal.value.map((v) => v.date),
    //     },
    //     yAxis: {},
    //     series: [
    //       {
    //         name: 'number',
    //         type: 'line',
    //         data: pomoTotal.value.map((v) => v.number),
    //       },
    //     ],
    //   })
    // })

    return {
      value,
      defaultTime,
    }
  },
})
</script>
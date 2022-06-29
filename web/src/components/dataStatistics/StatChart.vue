<template>
  <Slot>
    <el-row>
      <el-col :span="24">
        <div class="grid-content ep-bg-purple-dark"
             id="main"> </div>

      </el-col>
    </el-row>
    <!-- TODO： Best workday Top tags  Best worktime  -->
    <el-row>
      <el-col :span="8">
        <div class="grid-content ep-bg-purple" />
      </el-col>
      <el-col :span="8">
        <div class="grid-content ep-bg-purple-light" />
      </el-col>
      <el-col :span="8">
        <div class="grid-content ep-bg-purple" />
      </el-col>
    </el-row>
  </Slot>

</template>

<script lang="ts">
import * as echarts from 'echarts'
import { inject, onMounted, reactive, ref, watch, defineComponent } from 'vue'
import axios from 'axios'
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
import { isNil } from 'lodash'
import Slot from '@/components/dataStatistics/dataStatisticSlot.vue'

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

    watch(
      pomoTotal,
      (newValue: any, oldValue: any) => {
        chart?.setOption({
          xAxis: {
            data: pomoTotal.value.map((v) => v.date),
          },
          series: [
            {
              name: 'number',
              type: 'line',
              data: pomoTotal.value.map((v) => {
                if (isNil(v.number)) {
                  return 0
                }
                return v.number
              }),
            },
          ],
        })
        console.log('chart set option')
      },
      { deep: true }
    )

    onMounted(() => {
      chart = echarts.init(document.getElementById('main') as HTMLDivElement)
      // 调用请求
      chart.setOption({
        // title: {
        //   text: 'ECharts 入门示例',
        // },
        tooltip: {},
        type: 'category',
        xAxis: {
          data: pomoTotal.value.map((v) => v.date),
        },
        yAxis: {},
        series: [
          {
            name: 'number',
            type: 'line',
            data: pomoTotal.value.map((v) => v.number),
          },
        ],
      })
    })

    return {
      value,
      defaultTime,
    }
  },
  components: {
    Slot,
  },
})
</script>

<style>
.demo-date-picker {
  display: flex;
  width: 100%;
  padding: 0;
  flex-wrap: wrap;
}
.demo-date-picker .block {
  padding: 30px 0;
  /* text-align: center; */
  border-right: solid 1px var(--el-border-color);
  flex: 1;
  /* float: right; */
  text-align: right;
}
.demo-date-picker .block:last-child {
  border-right: none;
}

.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 400px;
}
</style>
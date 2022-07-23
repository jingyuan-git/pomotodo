<template>
  <Slot>
    <el-row>
      <el-col :span="24">
        <div class="grid-content ep-bg-purple-dark"
             id="main"> </div>

      </el-col>
    </el-row>
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
import { onMounted, ref, watch, defineComponent } from 'vue'
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
import { isNil } from 'lodash'
import Slot from '@/components/dataStatistics/PomoStatistics.vue'

export default defineComponent({
  setup() {
    const store = useMainStore()

    let { pomoTotal } = storeToRefs(store)
    let chart: echarts.ECharts | null = null

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
      chart.setOption({
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
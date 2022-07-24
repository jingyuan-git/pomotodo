<template>
  <div class="demo-date-picker">
    <div class="block">
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
        console.log('value has changed', newValue, oldValue)
        store.counts(
          newValue[0].valueOf().toString(),
          newValue[1].valueOf().toString()
        )
      },
      { deep: true }
    )

    return {
      value,
      defaultTime,
    }
  },
})
</script>
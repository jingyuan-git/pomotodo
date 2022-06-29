<template>
  <Slot>
    <el-card class="card"
             shadow="never">
      <el-timeline>
        <el-timeline-item v-for="item in pomoOnePage"
                          :timestamp="item.date"
                          placement="top">
          <div v-if="item.number > 0">
            <el-card>
              <div v-for="(pomo, index) in item.Pomodoros">
                <p> {{time(pomo.startTime)}} - {{time(pomo.endTime)}} &nbsp; {{ pomo.title}} </p>
              </div>
            </el-card>
          </div>
        </el-timeline-item>
      </el-timeline>

    </el-card>

    <div class="demo-pagination-block">
      <el-pagination v-model:currentPage="currentPage"
                     v-model:page-size="pageSize"
                     layout="prev, pager, next, jumper"
                     :total="total"
                     :background="true"
                     @size-change="handleSizeChange"
                     @current-change="handleCurrentChange" />
    </div>
  </Slot>

</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  onMounted,
  reactive,
  ref,
  toRefs,
  watch,
} from 'vue'
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
import { PomoTotal } from '@/types/pomo'
import dayjs from 'dayjs'
import Slot from '@/components/dataStatistics/dataStatisticSlot.vue'
import { isNil } from 'lodash'

export default defineComponent({
  setup() {
    let data = reactive<{ pomos: PomoTotal[] }>({
      pomos: [],
    })
    const store = useMainStore()
    console.log('TTTTTTT', store.pomoTotal)
    let { pomoTotal } = storeToRefs(store)
    let { pomoOnePage } = storeToRefs(store)

    console.log('data', data)

    const time = computed(() => (date_example: string) => {
      return dayjs(date_example).format('HH:mm')
    })

    watch(
      pomoTotal,
      (newValue: any, oldValue: any) => {
        const a = JSON.parse(JSON.stringify(newValue))
        store.updatePomo(a)
        store.currentPomo(1, pageSize.value)
      },
      { deep: true }
    )

    onMounted(() => {
      // console.log(store.pomoTotal)
    })

    let currentPage = ref(1)
    let pageSize = ref(3)
    const handleSizeChange = (val: number) => {
      console.log(`${val} items per page`)
    }
    const handleCurrentChange = (val: number) => {
      console.log(`current page: ${val}`)
      store.currentPomo(val, pageSize.value)
    }

    return {
      ...toRefs(data),
      time,
      total: store.pomoFilter.length,
      listQuery: {
        page: 1,
        limit: 20,
      },
      currentPage,
      pageSize,
      handleSizeChange,
      handleCurrentChange,
      pomoOnePage,
    }
  },
  components: {
    Slot,
  },
  methods: {
    getList() {
      // Fetch data
    },
  },
})
</script>

<style>
</style>
<template>
  <ul class="infinite-list"
      style="overflow: auto">

    <li v-for="(item,index) in pomos"
        :key="index"
        class="infinite-list-item">

      {{ formatDate(item.startTime) }} - {{ formatTime(item.endTime) }} &nbsp;&nbsp;&nbsp; {{ item.title }} <br />
    </li>
  </ul>
</template>

<script  lang="ts">
import { defineComponent, ref, computed, watch } from 'vue'
import dayjs from 'dayjs'
import { Pomo } from '@/types/pomo'
import { formatDate, formatDay, formatTime } from '@/utils/format'

export default defineComponent({
  name: 'PomoList',
  props: ['pomos'],
  setup(props) {
    watch(
      props.pomos,
      (newValue: Pomo[], oldValue: Pomo[]) => {
        console.log('pomos is change', newValue, oldValue)
      },
      { deep: true }
    )

    return {
      dayjs,
      formatDate,
      formatDay,
      formatTime,
    }
  },
})
</script>

<style>
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}

.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  /* justify-content: center; */
  /* height: 50px; */
  /* background: var(--el-color-primary-light-9); */
  margin: 10px;
  /* color: var(--el-color-primary); */
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>
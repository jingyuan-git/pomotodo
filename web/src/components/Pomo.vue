<template>
  <section class="pomo-container">
    <Header :addToPomo='addToPomo'></Header>
    <el-divider />
    <List :pomos="pomos" />
  </section>

</template>

<script lang="ts">
import Header from '@/components/pomos/PomoHeader.vue'
import List from '@/components/pomos/PomoList.vue'
import { ref, reactive, toRefs, watch } from 'vue'
import { watchEffect, onMounted, defineComponent } from 'vue'
import { useTimer } from 'vue-timer-hook'
import { Pomo } from '@/types/pomo'
import { useMainStore } from '@/stores/index'
import axios from 'axios'

export default defineComponent({
  name: 'Pomo',

  setup() {
    let isShow = ref(true)

    const time = new Date()
    const timer = useTimer(time.setSeconds(time.getSeconds() + 5), false)
    // const timer = useTimer(time.setMinutes(time.getMinutes() + 25), false)
    const startPomo = () => {
      isShow.value = !isShow.value
      timer.start()
      console.log('开始')
    }

    const state = reactive<{ pomos: Pomo[] }>({
      pomos: [],
    })

    onMounted(() => {
      const sendPostRequest = async () => {
        try {
          const resp = await axios.post(
           '/api/v1/pomos/list',
            {},
            {
              headers: {
                'Content-Type': 'text/plain',
              },
              responseType: 'text',
            }
          )
          console.log('resp.data.data.lists', resp.data.data.lists)
          state.pomos = resp.data.data.lists
        } catch (err) {
          console.error(err)
        }
      }

      sendPostRequest()

      watchEffect(async () => {})
    })

    const store = useMainStore()
    const addToPomo = (pomo: Pomo) => {
      state.pomos.unshift(pomo)
      store.counts('', '')
    }
    return {
      ...toRefs(state),
      isShow,
      timer,
      startPomo,
      addToPomo,
    }
  },
  components: {
    Header,
    List,
  },
})
</script>

<style scoped>
.pomo-container {
  /* margin-left: 10%; */
  margin-right: 5%;
}
</style>
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
import { savePomos, readPomos } from '@/utils/localStorageUtils'
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
    watch(() => state.pomos, savePomos, { deep: true })

    watch(state, (newValue, oldValue) => {
      console.log('state的值变化了', newValue, oldValue)
    })

    onMounted(() => {
      const sendPostRequest = async () => {
        try {
          const resp = await axios.post(
            'http://localhost:8000/api/v1/pomos/list',
            {},
            {
              headers: {
                'Content-Type': 'text/plain',
              },
              responseType: 'text',
            }
          )
          console.log(resp.data.lists)
          console.log(resp.data.data)
          console.log('resp.data.data.lists', resp.data.data.lists)
          state.pomos = resp.data.data.lists
          console.log('resp.data', resp.data, state.pomos)
          // console.log(resp.data)
        } catch (err) {
          // Handle Error Here
          console.error(err)
        }
      }

      sendPostRequest()
      //TODO:
      // const resp = await axios
      //   .post('http://localhost:8000/api/v1/pomos/list', {}, config)
      //   .then((response: any) => {
      //     console.log(response.data)
      //   })
      watchEffect(async () => {})
    })

    const addToPomo = (pomo: Pomo) => {
      state.pomos.unshift(pomo)

      console.log('pomos', state.pomos)
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
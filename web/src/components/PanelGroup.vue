<template>
  <el-row :gutter="40"
          class="panel-group">
    <el-col :xs="12"
            :sm="12"
            :lg="12"
            class="card-panel-col">
      <div class="card-panel"
           @click="handleSetLineChartData('Stat')">
        <div class="card-panel-icon-wrapper icon-people">
          <img :src="user">
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            Pomodoro statistics
          </div>
        </div>
      </div>
    </el-col>
    <el-col :xs="12"
            :sm="12"
            :lg="12"
            class="card-panel-col">
      <div class="card-panel"
           @click="handleSetLineChartData('PomoHistory')">
        <div class="card-panel-icon-wrapper icon-money">
          <img :src="tag">
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">
            Pomodoro history
          </div>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { useMainStore } from '@/stores/index'
import { storeToRefs } from 'pinia'
import { defineComponent, onMounted, ref } from 'vue'
import { CountTo } from 'vue3-count-to'
import tag from '@/assets/images/home/tag.png'
import user from '@/assets/images/home/user.png'
import axios from 'axios'
export default defineComponent({
  components: {
    CountTo,
  },
  emits: ['handle-set-line-chart-data'],
  setup(_, { emit }) {
    const store = useMainStore()
    store.counts('', '')
    console.log(store.pomoTotal)

    const handleSetLineChartData = (type: string) => {
      // emit('handle-set-line-chart-data', type)
      store.$patch({ analyzeComponent: type })
      console.log(store.analyzeComponent)

      const sendPostRequest = async () => {
        try {
          const resp = await axios.post(
            '/api/v1/pomos/count',
            JSON.stringify({
              // startTime: pomoStartTime.value.toString(),
              // endTime: Date.now().toString(),
            }),
            {
              headers: {
                'Content-Type': 'text/plain',
              },
              responseType: 'text',
            }
          )
          console.log(
            'resp.data.data.lists',
            resp.data.data.lists,
            store.pomoTotal.map((v) => v.date)
          )

          store.pomoTotal = resp.data.data.lists
          store.$patch({ pomoTotal: resp.data.data.lists })
          console.log('currentValue patch', store.pomoTotal)
        } catch (err) {
          // Handle Error Here
          console.error(err)
        }
      }

      sendPostRequest()
    }

    // store.increment()

    const bgColor = ref('white')
    const myColor = ref('black')

    const mouseHandler = (flag: boolean) => {
      if (flag) {
        bgColor.value = 'pink'
        myColor.value = 'green'
      } else {
        bgColor.value = 'white'
        myColor.value = 'black'
      }
    }

    return {
      handleSetLineChartData,
      tag,
      user,
      mouseHandler,
      bgColor,
      myColor,
    }
  },
})
</script>

<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);
    height: 124px;
    background: #ffffff;
    border-radius: 10px;
    img {
      width: 60px;
      height: 60px;
      display: inline-block;
    }
    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3;
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }

    &:hover {
      color: lightcoral;
      background-color: rgba(137, 190, 78, 0.3);
    }
  }
}

@media (max-width: 550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    svg {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>

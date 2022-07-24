<template>
  <el-row>
    <el-button @click="startPomo()">
      <div v-if="isShow">Start Pomo</div>
      <div v-else>
        <span>{{ timer.hours }}</span>:<span>{{ timer.minutes }}</span>:<span>{{ timer.seconds }}</span>
      </div>
    </el-button>
  </el-row>

  <Dialog v-model="dialogVisible">
    <el-input v-model="pomoTitle" placeholder="pomodoro name" />
    <div id="wrap">
      <el-form-item id="wrap">
        <el-button id="button1" type="primary" @click="onSubmit">Create</el-button>
        <el-button id="button2" @click="dialogVisible = false, isShow = true">Cancel</el-button>
      </el-form-item>
    </div>
  </Dialog>
</template>

<script lang="ts">
import { ref } from 'vue'
import { watchEffect, onMounted, defineComponent } from 'vue'
import { useTimer } from 'vue-timer-hook'
import { Pomo } from '@/types/pomo'
import { nanoid } from 'nanoid'
import Dialog from '@/components/pomos/PomoDialog.vue'
import dayjs from 'dayjs'
import axios from 'axios'

export default defineComponent({
  name: 'Pomo',
  props: {
    addToPomo: {
      type: Function,
      required: true,
    },
  },
  setup(props) {
    let isShow = ref(true)
    let dialogVisible = ref(false)
    let pomoStartTime = ref(0)
    let pomoTitle = ref('')

    const time = new Date()
    const timer = useTimer(time.setMinutes(time.getMinutes() + 25), false)
    const startPomo = () => {
      // 
      if (isShow.value === false) {
        if (timer.isRunning) {
          dialogVisible.value = true
        }
        timer.pause()
      } else {
        isShow.value = !isShow.value
        reset()
        pomoStartTime.value = Date.now()
      }
    }

    const reset = () => {
      // Restarts to 25 minutes timer
      const time = new Date()
      time.setSeconds(time.getSeconds() + 1500)
      timer.restart(time.getTime())
      // timer.pause()
    }

    onMounted(() => {
      watchEffect(async () => {
        if (timer.isExpired.value) {
          dialogVisible.value = true
          timer.pause()
        }
      })
    })

    const onSubmit = () => {
      const pomo: Pomo = {
        id: nanoid(),
        title: pomoTitle.value,
        startTime: dayjs(pomoStartTime.value).format('YYYY-MM-DD HH:mm:ss'),
        endTime: dayjs(Date.now()).format('YYYY-MM-DD HH:mm:ss'),
      }

      axios
        .post(
          '/api/v1/pomos/create',
          JSON.stringify({
            id: pomo.id,
            title: pomo.title,
            startTime: pomoStartTime.value.toString(),
            endTime: Date.now().toString(),
          }),
          {
            headers: {
              'Content-Type': 'text/plain',
            },
            responseType: 'text',
          }
        )
        .then((response: any) => {
          console.log(response)
          props.addToPomo(pomo)
          console.log('submit!', props)
        })
        .catch(function (error) {
          console.log(error);
        })
      isShow.value = true
      dialogVisible.value = false
    }

    return {
      isShow,
      timer,
      startPomo,
      onSubmit,
      dialogVisible,
      pomoTitle,
    }
  },
  components: {
    Dialog,
  },
})
</script>

<style lang="less" scoped>
.el-button {
  width: 100%;
}

.app-container {
  display: grid;
  place-items: center;
}

.dialog-body {
  img {
    display: block;
    width: 200px;
    height: 200px;
    margin: 0 auto;
  }

  img+p {
    margin: 30px 0;
    color: #6266f5;
  }

  .text-button {
    margin-top: 15px;
  }
}

el-button {
  width: auto;
}

.wrap {
  display: flex;
}

#button1,
#button2 {
  display: inline-block;
  margin-top: 15px;
  width: 100px;
}

#div1,
#div2,
#div3 {
  display: inline-block;
  width: 50px;
}
</style>
<template>
  <div class="dialog"
       :class="{ 'is-show': visible }"
       :style="display ? 'display: none' : ''">

    <div class="dialog-modal"
         @click.self="closeDialog"></div>

    <div class="dialog-main">

      <div class="dialog-head">
        <button class="button icon-botton"
                @click="closeDialog">x</button>
      </div>

      <!-- content -->
      <div class="dialog-body">
        <slot></slot>
      </div>

    </div>

  </div>
</template>

<script>
import { defineComponent, onBeforeUnmount, ref, watch } from 'vue'
export default defineComponent({
  props: {
    modelValue: Boolean
  },
  emits: ['update:modelValue'],
  setup (props, ctx) {
    const visible = ref(props.modelValue)
    const display = ref(true)
    let timer = 0

    // close dialog
    const closeDialog = () => {
      clearTimeout(timer)
      visible.value = false
      timer = setTimeout(() => {
        display.value = false
        clearTimeout(timer)
      }, 300)
    }

    watch(() => props.modelValue, (currentValue) => {
      if (currentValue) {
        clearTimeout(timer)
        display.value = false
        timer = setTimeout(() => {
          visible.value = currentValue
          clearTimeout(timer)
        }, 20)
      } else {
        closeDialog()
      }
    })

    watch(() => visible.value, (currentValue) => {
      if (!currentValue) {
        ctx.emit('update:modelValue', currentValue)
      }
    })

    onBeforeUnmount(() => {
      clearTimeout(timer)
    })
    return {
      visible,
      display,
      closeDialog
    }
  }
})
</script>

<style lang="less" scoped>
@import '@/style/modal.less';
</style>

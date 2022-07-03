<template>
  <div class="header-wrap">
    <input type="text"
           v-model="todoItem"
           placeholder="Enter the tod and press enter to add"
           @keyup.enter="add">
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import { Todo } from '@/types/todo'
import { nanoid } from 'nanoid'
export default defineComponent({
  name: 'Head',
  props: {
    addToPlan: {
      type: Function,
      required: true,
    },
  },
  setup(props, context) {
    const todoItem = ref('')
    const add = () => {
      const text = todoItem.value
      if (!text.trim()) return
      const todo: Todo = {
        id: nanoid(),
        title: text.trim(),
        createdAt: Date.now().toString(),
        isCompleted: false,
      }
      // setup中不能直接调用props中的内容
      // 因此此处用到setup的props参数
      props.addToPlan(todo)
      todoItem.value = ''
    }
    return {
      todoItem,
      add,
    }
  },
})
</script>

<style scoped>
.header-wrap input {
  width: 100%;
  padding: 5px;
  border-radius: 5px;
  outline: none;
  border: 1px solid lightblue;
  box-shadow: 1px 1px 1px rgba(0, 0, 0, 0.2);
}
</style>
<template>
  <div>
    <el-input type="text"
              v-model="todoItem"
              placeholder="Input a new todo and press enter to add"
              @keyup.enter="add" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
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
</style>
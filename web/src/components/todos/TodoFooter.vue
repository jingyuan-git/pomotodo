<template>
  <div class="todo-footer">
    <div>
      <label><input type="checkbox"
               :isCheckAll="isCheckAll"></label>
      <span class="todo-tag"><span>已完成 {{count}}</span> /全部 {{todos.length}}</span>
    </div>
    <button class="btn btn-danger"
            @click="isClearCompleted">清除已完成任务</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { Todo } from '../../types/todo'
export default defineComponent({
  name: 'Footer',
  props: {
    todos: {
      type: Array as () => Todo[],
      required: true,
    },
    checkAll: {
      type: Function,
      required: true,
    },
    clearCompleted: {
      type: Function,
      required: true,
    },
  },
  setup(props, context) {
    const count = computed(() => {
      return props.todos.reduce(
        (pre: number, todo: Todo) => pre + (todo.isCompleted ? 1 : 0),
        0
      )
    })
    // 全选是否为true
    const isCheckAll = computed({
      get() {
        return count.value > 0 && count.value === props.todos.length
      },
      set(val) {
        //调用父组件的方法去修改全选状态
        props.checkAll(val)
      },
    })

    const isClearCompleted = () => props.clearCompleted

    return {
      count,
      isCheckAll,
      isClearCompleted,
    }
  },
})
</script>

<style scoped>
.todo-footer {
  display: flex;
  justify-content: space-between;
}
.btn-danger {
  color: #fff;
  background-color: #e95b47;
  border: none;
  border-radius: 5px;
  padding: 5px;
  cursor: pointer;
}
.todo-tag {
  margin-left: 15px;
}
</style>
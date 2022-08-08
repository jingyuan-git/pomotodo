<template>
  <li class="todo-item">
    <div>
      <input id="todo"
             type="checkbox"
             v-model="isCompleted">
      <span>{{todoitem.title}}</span>
    </div>
    <button class="delBtn"
            @click="del(todoitem.id)"> delete </button>
  </li>
</template>

<script lang="ts">
import { defineComponent, reactive, inject, computed } from 'vue'
import { Todo } from '@/types/todo'
export default defineComponent({
  name: 'ListItem',
  props: {
    todoitem: {
      type: Object as () => Todo,
      required: true,
    },
    index: {
      type: Number,
    },
  },
  setup(props, context) {
    let selectedItemList = reactive([])
    if (window.sessionStorage.getItem('selectedItemList')) {
      selectedItemList = JSON.parse(
        window.sessionStorage.getItem('selectedItemList') + ''
      )
      selectedItemList.push()
    }
    const delTodo: Function | undefined = inject('delTodo')
    const updateState: Function | undefined = inject('updateState')

    const del = (id: String | undefined) => {
      if (window.confirm('Confirm delete?')) {
        if (typeof delTodo === 'function') delTodo(id, props.index)
      }
    }
    const isCompleted = computed({
      get() {
        return props.todoitem.isCompleted
      },
      set(val) {
        if (typeof updateState === 'function') updateState(props.todoitem, val)
      },
    })

    return {
      isCompleted,
      del,
    }
  },
})
</script>

<style scoped>
.todo-item {
  border: 1px solid lightblue;
  list-style: none;
  width: 100%;
  padding: 5px;
  margin-left: -40px;
  border-radius: 5px;
  display: flex;
  justify-content: space-between;
}
.delBtn {
  background-color: #e95b47;
  color: #fff;
  border-radius: 5px;
  border: none;
  display: none;
}
.todo-item:hover {
  color: lightcoral;
  background-color: rgba(137, 190, 78, 0.3);
}
.todo-item:hover .delBtn {
  display: block;
}
</style>
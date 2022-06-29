<template>
  <!--
        此处样式我使用css hover来做
        视频中老师用到的两个方法 @mouseenter @mouseleave分别为鼠标进入和离开的回调
    -->
  <li class="todo-item">
    <div>
      <input type="checkbox"
             :isCompleted="isCompleted">
      <span>{{todoitem.title}}</span>
    </div>
    <button class="delBtn"
            @click="del(todoitem.id)">删除</button>
  </li>
</template>

<script lang="ts">
import { defineComponent, reactive, inject, computed } from 'vue'
// 引入接口
import { Todo } from '@/types/todo'
export default defineComponent({
  name: 'ListItem',
  props: {
    todoitem: {
      //函数返回的是Todo类型
      type: Object as () => Todo,
      required: true,
    },
    index: {
      type: Number,
    },
  },
  setup(props, context) {
    // 当复选框选中时，所选事件下标索引将添加到selectedItemList中
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
      if (window.confirm('确认删除吗？')) {
        if (typeof delTodo === 'function') delTodo(id, props.index)
      }
    }
    const isCompleted = computed({
      get() {
        console.log('is completed get', props.todoitem.isCompleted)
        return props.todoitem.isCompleted
      },
      set(val) {
        console.log('is completed set', props.todoitem.isCompleted)
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
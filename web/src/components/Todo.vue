<template>
  <div class="todo-wrap">
    <Header :addToPlan='addToPlan'></Header>
    <el-divider />
    <List :todos='todos'></List>
    <Footer :todos='todos'
            :clearCompleted='clearCompleted'
            :checkAll='checkAll' />
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  provide,
  reactive,
  toRefs,
  watch,
  onMounted,
} from 'vue'
import Header from '@/components/todos/TodoHeader.vue'
import List from '@/components/todos/TodoList.vue'
import Footer from '@/components/todos/TodoFooter.vue'
import { Todo } from '@/types/todo'
import axios from 'axios'

export default defineComponent({
  name: 'App',
  setup() {
    const state = reactive<{ todos: Todo[] }>({
      todos: [],
    })
    onMounted(() => {
      const sendPostRequest = async () => {
        try {
          const resp = await axios.post(
            '/api/v1/todos/list',
            {},
            {
              headers: {
                'Content-Type': 'text/plain',
              },
              responseType: 'text',
            }
          )
          console.log('todo resp.data.data.lists', resp.data.data.lists)
          state.todos = resp.data.data.lists
          console.log('resp.data', resp.data, state.todos)
        } catch (err) {
          console.error(err)
        }
      }

      sendPostRequest()
    })

    // used to add data to the head component
    const addToPlan = (todo: Todo) => {
      axios
        .post(
          '/api/v1/todos/create',
          JSON.stringify({
            id: todo.id,
            title: todo.title,
            createdAt: todo.createdAt,
            isCompleted: todo.isCompleted,
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
          state.todos.unshift(todo)
          console.log('addTodo!', todo)
        })
        .catch(function (error) {
          console.log(error)
        })
    }

    const delTodo: Function = (id: string, index: number) => {
      // todo:
      axios
        .post(
          '/api/v1/todos/delete',
          JSON.stringify({
            id: id,
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
          state.todos.splice(index, 1)
        })
        .catch(function (error) {
          console.log(error)
        })

      console.log('delTodo!', state, id)
    }

    const updateState = (todo: Todo, val: boolean) => {
      axios
        .post(
          '/api/v1/todos/update',
          JSON.stringify({
            id: todo.id,
            title: todo.title,
            createdAt: todo.createdAt,
            isCompleted: val,
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
          todo.isCompleted = val
          console.log('update Todo!', todo)
        })
        .catch(function (error) {
          console.log(error)
        })

      console.log(todo)
    }
    const checkAll = (val: boolean) => {
      state.todos.forEach((item) => {
        axios
          .post(
            '/api/v1/todos/update',
            JSON.stringify({
              id: item.id,
              title: item.title,
              createdAt: item.createdAt,
              isCompleted: val,
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
            item.isCompleted = val
          })
          .catch(function (error) {
            console.log(error)
          })
      })
    }
    const clearCompleted = () => {
      console.log('clearCompleted', clearCompleted)
      state.todos = state.todos.filter((item, index) => {
        if (item.isCompleted) {
          return false
        }
        return true
      })
    }
    provide('delTodo', delTodo)
    provide('updateState', updateState)
    return {
      ...toRefs(state),
      // add method to header component
      addToPlan,
      // add checkAll method for footer component
      checkAll,
      clearCompleted,
    }
  },
  components: {
    Header,
    List,
    Footer,
  },
})
</script>

<style>
.todo-wrap {
  /* margin-right: 5%; */
  /* width: 500px; */
  /* margin: 0 auto; */
  /* border: 1px solid lightblue; */
  /* border-radius: 5px; */
  /* padding: 20px; */
}

.todo-wrap h2 {
  text-align: center;
}
</style>

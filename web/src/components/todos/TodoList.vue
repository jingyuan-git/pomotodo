<template>
  <div class="todo-main">
    <ul class="todo-list">
      <ListItem v-for="(item,index) in todos"
                :todoitem='item'
                :index='index'
                :key="item.id" />
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import ListItem from './TodoItem.vue'
export default defineComponent({
  name: 'List',
  props: ['todos'],
  setup() {
    return {}
  },
  components: {
    ListItem,
  },
})
</script>

<style scoped>
.todo-list {
  width: 100%;
}
</style>
import { Todo } from '../types/todo'
import { Pomo } from '../types/pomo'
// 保存数据到浏览器缓存中
export function saveTodos(todos: Todo[]) {
    localStorage.setItem('todos_key', JSON.stringify(todos))
}
// 从浏览器中读取数据
export function readTodos(): Todo[] {
    return JSON.parse(localStorage.getItem('todos_key') || '[]')
}

export function savePomos(pomos: Pomo[]) {
    localStorage.setItem('pomos_key', JSON.stringify(pomos))
}

export function readPomos(): Pomo[] {
    return JSON.parse(localStorage.getItem('pomos_key') || '[]')
}
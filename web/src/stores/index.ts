import { defineStore } from 'pinia'
import type { PomoTotal } from '@/types/pomo'
import axios from 'axios'
import { isNil } from 'lodash'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore({
	id: "main",
	// a function that returns a fresh state
	state: () => ({
		analyzeComponent: '',
		pomoTotal: <PomoTotal[]>[],
		pomoFilter: <PomoTotal[]>[],
		pomoOnePage: <PomoTotal[]>[],
	}),
	// optional getters
	getters: {
	},
	// optional actions
	actions: {
		async counts(filterBegin: String, filterEnd: String) {
			const resp = await axios.post(
				'/api/v1/pomos/count',
				JSON.stringify({
					filterBegin: filterBegin,
					filterEnd: filterEnd,
				}),
				{
					headers: {
						'Content-Type': 'text/plain',
					},
					responseType: 'text',
				}
			)
			this.pomoTotal = resp.data.data.lists
			console.log("pomoTotal", this.pomoTotal)
		},
		updatePomo(pomoT: PomoTotal[]) {
			this.pomoFilter = pomoT.reverse().filter((v) => {
				if (isNil(v.number)) {
					return false
				} else {
					return true
				}
			})
		},
		currentPomo(currentPage: number, pageSize: number) {
			this.pomoOnePage = this.pomoFilter.slice((currentPage - 1) * pageSize, (currentPage) * pageSize)
			console.log("current pageSize", currentPage * pageSize, (currentPage + 1) * pageSize)
			console.log("currentPomo", this.pomoOnePage)
		}
	},
})
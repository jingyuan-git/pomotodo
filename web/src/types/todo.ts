//定义一个接口，约束state的数据类型
export interface Todo {
	id: string,
	title: string,
	createdAt: string,
	isCompleted: boolean
}
//定义一个接口，约束state的数据类型
export interface Pomo {
	id: string,
	title: string,
	startTime: string,
	endTime: string,
}

export interface PomoTotal {
	date: string;
	number: number;
	Pomodoros: Array<Pomo>;
}
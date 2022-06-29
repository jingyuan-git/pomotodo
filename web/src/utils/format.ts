import dayjs from "dayjs"

export const formatDate = (time: string) => {
	return dayjs(time).format('MM-DD HH:mm')
}

export const formatDay = (time: string) => {
	return dayjs(time).format('YY-MM-DD')
}

export const formatTime = (time: string) => {
	return dayjs(time).format('HH:mm')
}
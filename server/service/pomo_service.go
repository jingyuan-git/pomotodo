package service

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"server/models"
)

type Pomo struct {
	ID        string `json:"id,omitempty" gorm:"primarykey"`
	Title     string `json:"title,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`

	FilterBegin string `json:"filterBegin,omitempty"`
	FilterEnd   string `json:"filterEnd,omitempty"`
}

type PomoDto struct {
	ID        string `json:"id" gorm:"primarykey"`
	Title     string `json:"title,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

func (p *Pomo) CreatePomo() error {
	startTime, _ := formatTimeStamp(p.StartTime)
	endTime, _ := formatTimeStamp(p.EndTime)

	pomo := models.Pomodoro{
		ID:        p.ID,
		Title:     p.Title,
		StartTime: startTime,
		EndTime:   endTime,
		Date:      endTime.Format("2006.01.02"),
	}
	if err := models.CreatePomo(pomo); err != nil {
		return err
	}

	return nil
}

func (a *Pomo) GetAll() ([]*models.Pomodoro, error) {
	aa := make(map[string]interface{})
	pomos, err := models.GetPomos(aa)
	if err != nil {
		log.Default().Printf("fail to list all pomos, error: %+v \n", err)
		return nil, err
	}

	return pomos, nil
}

func (a *Pomo) CountPomos() ([]*models.CountPomosByDate, error) {
	pomos, err := models.CountPomos(a.getMaps())
	if err != nil {
		log.Default().Printf("fail to list all pomos, error: %+v \n", err)
		return nil, err
	}

	return pomos, nil
}

func (a *Pomo) CountAndTotalAmount() (int, int, error) {

	return 0, 0, nil
}

func (a *Pomo) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	filterBegin, _ := formatTimeStamp(a.FilterBegin)
	filterEnd, _ := formatTimeStamp(a.FilterEnd)

	if a.FilterBegin != "" {
		maps["filterBegin"] = filterBegin
	}
	if a.FilterEnd != "" {
		maps["filterEnd"] = filterEnd
	}
	fmt.Println("maps  ", maps, a.StartTime, a.EndTime)
	return maps
}

func formatTimeStamp(t string) (time.Time, error) {
	formatTime, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return time.UnixMilli(0), err
	}
	return time.UnixMilli(formatTime), nil
}

package models

import (
	// "fmt"
	"time"

	"gorm.io/gorm"
)

type Pomodoro struct {
	ID        string    `json:"id,omitempty" gorm:"primarykey"`
	Title     string    `json:"title,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Date      string    `json:"date,omitempty"`
}

type CountPomosByDate struct {
	Date      string     `json:"date" gorm:"primarykey"`
	Number    int        `json:"number,omitempty"`
	Pomodoros []Pomodoro `gorm:"foreignKey:Date;references:Date"`
}

// Create insert the value into database
func CreatePomo(p Pomodoro) error {
	if err := db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

// GetPomos gets a list of pomos based on paging constraints
func GetPomos(maps map[string]interface{}) ([]*Pomodoro, error) {
	results := []*Pomodoro{}
	db.Model([]*Pomodoro{}).Order("start_time desc").Find(&results)
	return results, nil
}

// GetPomos gets a list of pomos based on paging constraints
func CountPomos(maps map[string]interface{}) ([]*CountPomosByDate, error) {
	results := []*CountPomosByDate{}
	t1 := &gorm.DB{}
	if maps["filterBegin"] != nil && maps["filterEnd"] != nil {
		t1 = db.Table("pomodoros").Select("to_char((generate_series(min(start_time)::date,max(start_time)::date, '1 days')),'yyyy.MM.dd') as date").Where("start_time > ? and end_time < ? ", maps["filterBegin"], maps["filterEnd"]).Order("date")
	} else {
		t1 = db.Table("pomodoros").Select("to_char((generate_series(min(start_time)::date,max(start_time)::date, '1 days')),'yyyy.MM.dd') as date").Order("date")
	}

	db.Migrator().DropTable([]CountPomosByDate{})
	db.Table("count_pomos_by_dates").AutoMigrate(&CountPomosByDate{})

	t2 := db.Table("pomodoros").Select("date, count(*) as number").Group("date")
	db.Select("t1.date as date, COALESCE(t2.number,?) as number", 0).Table("(?) as t1 left join (?) as t2 on (t1.date = t2.date)", t1, t2).Find(&results)
	db.Table("count_pomos_by_dates").Create(&results).Preload("Pomodoros").Find(&results)
	db.Migrator().DropTable("count_pomos_by_dates")

	return results, nil
}

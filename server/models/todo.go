package models

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          string    `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt   time.Time `csv:"created_at" json:"createdAt,omitempty"`
	Title       string    `json:"title,omitempty"`
	IsCompleted bool      `json:"isCompleted,omitempty"`
}

type CountTodosByDate struct {
	Date   string `json:"date" gorm:"primarykey"`
	Number int    `json:"number,omitempty"`
	Todos  []Todo `gorm:"foreignKey:Date;references:Date"`
}

// Create insert the value into database
func CreateTodo(t Todo) error {
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

// GetOrders gets a list of orders based on paging constraints
func GetTodos(maps map[string]interface{}) ([]*Todo, error) {
	results := []*Todo{}
	db.Debug().Model([]*Todo{}).Order("created_at desc").Find(&results)
	return results, nil
}

// GetOrders gets a list of orders based on paging constraints
func CountTodos(maps map[string]interface{}) ([]*CountTodosByDate, error) {
	results := []*CountTodosByDate{}
	// aa := []*CountTodosByDate{}
	// db.Debug().Table("pomodoros").Select("date, count(*) as number").Group("date").Order("date").Scan(&results)
	if maps["filterBegin"] != nil && maps["filterEnd"] != nil {
		db.Debug().Table("pomodoros").Select("date, count(*) as number").Where("created_at > ? and created_at < ? ", maps["filterBegin"], maps["filterEnd"]).Group("date").Order("date").Scan(&results)
	} else {
		db.Debug().Table("pomodoros").Select("date, count(*) as number").Group("date").Order("date").Scan(&results)
	}

	db.Debug().Migrator().DropTable([]CountTodosByDate{})
	db.Debug().Table("count_pomos_by_dates").AutoMigrate(&CountTodosByDate{})
	db.Debug().Table("count_pomos_by_dates").Create(&results)

	db.Debug().Preload("Todos").Find(&results)

	db.Migrator().DropTable("count_pomos_by_dates")

	return results, nil
}

// Delete insert the value into database
func DeleteTodo(t Todo) error {
	fmt.Println(t)
	if err := db.Delete(t).Error; err != nil {
		return err
	}
	return nil
}

// Create insert the value into database
func UpdateTodo(t Todo) error {
	if err := db.Debug().Model(t).Updates(
		map[string]interface{}{
			"is_completed": t.IsCompleted,
			"title":        t.Title,
		}).Error; err != nil {
		return err
	}
	return nil
}

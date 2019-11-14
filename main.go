package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// abnormal type and reason
type AbnormalType struct {
	TypeId   uint   `gorm:"type:int(10) unsigned auto_increment;primary_key"`
	TypeName string `gorm:"type:varchar(32);not null;index:type_name"`
	Reason   string `gorm:"type:varchar(255)"`
}

// abnormal record
type AbnormalRecord struct {
	Id             uint      `gorm:"type:int(10) unsigned auto_increment;primary_key"`
	CreatedAt      time.Time `gorm:"default:current_timestamp;index:created_at"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp on update current_timestamp"`
	UserId         uint      `gorm:"type:int(10) unsigned;not null"`
	AbnormalTypeId uint      `gorm:"type:int(10) unsigned;not null;index:abnormal_type_id"`
	OperationId    uint      `gorm:"index:operation_id"`
	OperationType  string    `gorm:"type:varchar(32);index:operation_type"`
	RuleType       string    `gorm:"type:varchar(64)"`
	Desc           string    `gorm:"type:varchar(255)"`
	Extra          string    `gorm:"type:varchar(255)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:Jason_199213@/risk_control?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connect to mysql error")
		return
	}
	defer db.Close()

	db.AutoMigrate(&AbnormalType{})
	db.AutoMigrate(&AbnormalRecord{})
}

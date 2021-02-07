package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"story/models"
	"story/pkg/config"
)

// DB is a gorm db object
var DB *gorm.DB

// Setup connect to postgresql db
func Setup() {
	conf := config.GetConfig()
	db, err := gorm.Open("postgres",
		"host="+conf.Database.Host+
			" port="+conf.Database.Port+
			" user="+conf.Database.Username+
			" dbname="+conf.Database.Dbname+
			"  sslmode=disable password="+conf.Database.Password)
	if err != nil {
		log.Fatalf("An error occured when connect to postgresql."+
			"Connection credentials : "+
			"host:%s, "+
			"port:%s, "+
			"user:%s, "+
			"dbname:%s, "+
			"password:%s. "+
			"Original error:%v",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Username,
			conf.Database.Dbname,
			conf.Database.Password,
			err)
	}

	db.LogMode(false)
	db.AutoMigrate(&models.Story{},
		models.Token{},
		models.Event{})
	DB = db
}

// GetDB returns db object
func GetDB() *gorm.DB {
	return DB
}

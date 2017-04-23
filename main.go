package main

import (
	_ "go-shiba/routers"
	"github.com/astaxie/beego"
	"go-shiba/middlewares/mongodb"
	"go-shiba/conf"
	"github.com/goinggo/tracelog"
	"os"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	
	// Init mongo
	tracelog.Started("main", "Initializing Mongo")
	
	// init beego config file
	init_config()
	
	err := mongodb.Startup(conf.MainGoRoutine)
	if err != nil {
		tracelog.CompletedError(err, conf.MainGoRoutine, "initApp")
		os.Exit(1)
	}
	
	beego.Run()
	
	tracelog.Completed(conf.MainGoRoutine, "Website Shutdown")
	tracelog.Stop()
	
}

func init_config() {
	env := beego.BConfig.RunMode
	
	beego.Info("env is :", env)
	
	if env == "test" {
		beego.Info("env is test mode.")
		beego.LoadAppConfig("ini", "./conf/app-test.conf")
	} else if env == "prod" {
		beego.Info("env is prod mode.")
		beego.LoadAppConfig("ini", "./conf/app-prod.conf")
	} else {
		beego.Info("env is dev mode.")
		beego.LoadAppConfig("ini", "./conf/app.conf")
	}
	
	beego.Info("host: ", beego.AppConfig.String("mongo.hosts"))
}
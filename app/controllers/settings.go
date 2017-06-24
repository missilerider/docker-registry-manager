package controllers

import (
	"github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	funcs "github.com/snagles/docker-registry-manager/app/templatefuncs"
)

type SettingsController struct {
	beego.Controller
}

func (c *SettingsController) Get() {
	c.Data["activeLevel"] = logrus.GetLevel()
	c.Data["allLevels"] = logrus.AllLevels
	c.Data["logs"] = parseLogs()
	c.TplName = "settings.tpl"
}

// GetLiveStatistics returns stats on request information tracked by beego
func (c *SettingsController) GetLiveStatistics() {

	r := toolbox.StatisticsMap
	rs := r.GetMapData()

	// Convert beego times to seconds for sorted
	for _, req := range rs {
		var err error
		req["min_s"], err = funcs.StatToSeconds(req["min_time"].(string))
		if err != nil {
			logrus.Error(err)
		}
		req["max_s"], err = funcs.StatToSeconds(req["max_time"].(string))
		if err != nil {
			logrus.Error(err)
		}
		req["avg_s"], err = funcs.StatToSeconds(req["avg_time"].(string))
		if err != nil {
			logrus.Error(err)
		}
		req["total_s"], err = funcs.StatToSeconds(req["total_time"].(string))
		if err != nil {
			logrus.Error(err)
		}
	}

	c.Data["json"] = &rs
	c.ServeJSON()
}

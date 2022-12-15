package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"web_complier/core"
)

type runTpl struct {
	Image   string `json:"image"`
	Dir     string `json:"dir"`
	Ext     string `json:"ext"`
	Cmd     string `json:"cmd"`
	Timeout int64  `json:"timeout"`
	Memory  int64  `json:"memory"`
	CpuSet  string `json:"cpuset"`
}

func Run(lang string) runTpl {
	var tpl runTpl
	currentDir, _ := os.Getwd()
	path := fmt.Sprintf("%s/configs/lang/%s.json", currentDir, lang)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		core.ZLogger.Sugar().Error("Some error occured while reading file. Error: %s", err)
	}
	err = json.Unmarshal(file, &tpl)
	if err != nil {
		core.ZLogger.Sugar().Error("Error occured during unmarshaling. Error: %s", err.Error())
	}
	//fmt.Println(tpl.Image)
	//fmt.Printf("tpl Struct: %#v\n", tpl)
	return tpl
}

func LangExists(lang string) (bool, error) {
	currentDir, _ := os.Getwd()
	path := fmt.Sprintf("%s/configs/lang/%s.json", currentDir, lang)
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

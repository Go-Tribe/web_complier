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
	File    string `json:"file"`
	Cmd     string `json:"cmd"`
	Timeout int64  `json:"timeout"`
	Memory  int64  `json:"memory"`
	CpuSet  string `json:"cpuset"`
}

func Run(lang string) runTpl {
	var tpl runTpl
	lang = fmt.Sprintf("/Users/develop/Project/app/go/src/web_complier/configs/lang/%s.json", lang)
	core.ZLogger.Error("lang:" + lang)
	file, err := ioutil.ReadFile(lang)
	if err != nil {
		core.ZLogger.Sugar().Error("Some error occured while reading file. Error: %s", err)
	}
	err = json.Unmarshal(file, &tpl)
	if err != nil {
		core.ZLogger.Sugar().Error("Error occured during unmarshaling. Error: %s", err.Error())
	}
	fmt.Println(tpl.Image)
	fmt.Printf("tpl Struct: %#v\n", tpl)
	return tpl
}

func LangExists(lang string) (bool, error) {
	path := fmt.Sprintf("/Users/develop/Project/app/go/src/web_complier/configs/lang/%s.json", lang)
	core.ZLogger.Error("path:" + path)
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

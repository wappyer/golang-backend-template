package gen

import (
	"bytes"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/utils"
	"html/template"
	"log"
	"os"
	"testing"
)

// 生成DDD文件
func Test_gen_curl(t *testing.T) {
	modelName := "event_queue"
	modelDesc := "事件任务队列"

	// 模版数据
	modelNameUpper := utils.UnderscoreToUpperCamelCase(modelName)
	modelNameLower := utils.UnderscoreToLowercaseCamelCase(modelName)
	data := map[string]string{
		"ModelName":      modelName,
		"ModelNameUpper": modelNameUpper,
		"ModelDesc":      modelDesc,
	}
	//_createFileByTemplate(data, "./model.template", "../internal/domain/model/"+modelNameLower+".go") // model文件

	_createFileByTemplate(data, "./repository.template", "../internal/domain/repository/"+modelNameLower+".go") // repository文件

	//_createFileByTemplate(data, "./application.template", "../internal/application/"+modelNameLower+".go") // application文件
	//_createFileByTemplate(data, "./contractCmd.template", "../internal/application/contract/cmd/"+modelNameLower+".go")     // contractCmd文件
	//_createFileByTemplate(data, "./contractQuery.template", "../internal/application/contract/query/"+modelNameLower+".go") // contractQuery文件

	//_createFileByTemplate(data, "./controller.template", "../internal/api/controller/"+modelNameLower+".go")      // controller文件
	//_createFileByTemplate(data, "./controllerVo.template", "../internal/api/controller/vo/"+modelNameLower+".go") // controllerVo文件
}

func _createFileByTemplate(data map[string]string, templateFilePath, newFilePath string) {
	// template.ParseFiles实现模板文件的创建
	t1, err := template.ParseFiles(templateFilePath)
	if err != nil {
		panic(err)
	}

	//模板数据进行相关的数据绑定（即将m中数据输入到字节流中）
	var b1 bytes.Buffer
	err = t1.Execute(&b1, data)
	if err != nil {
		panic(err)
	}

	// 创建文件并且将模板数据写入到新文件中
	file, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("错误异常：", err)
		log.Println(err)
	}
	_, err = file.WriteString(b1.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

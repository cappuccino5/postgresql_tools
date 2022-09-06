package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"postgresql_tools/internal/models"
	"strings"
	"text/template"
)

func PararWithIdSeq(filepath string, outputFile string) error {
	var idSeq models.ResultByIdSeq
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &idSeq)
	if err != nil {
		return err
	}
	fmt.Println(idSeq)
	//按模版生成文件
	var fileContent []string
	for _, v := range idSeq.Record {

		out := TemplateText(models.IdSeqSetValSqlTpl, map[string]interface{}{
			"sequenceName": v.SequenceName,
			"tableName":    v.TableName,
		})
		fmt.Println(out)
		fileContent = append(fileContent, out)
	}
	OutputFileByTemplateResult(outputFile, fileContent)
	return nil
}

func TemplateText(tpl string, tplArgv map[string]interface{}) string {
	tpl1, err := template.New("x").Parse(tpl)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	_ = tpl1.Execute(&b, tplArgv)
	return b.String()
}

func OutputFileByTemplateResult(targetFile string, fileContent []string) {
	_, err := os.Stat(targetFile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Errorf("%v", err)
			return
		}
	} else {
		fmt.Println("reset file:", targetFile)
		os.RemoveAll(targetFile)
	}
	f, err := os.Create(targetFile)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	f.WriteString(strings.Join(fileContent, "\n"))
	f.Close()
	return
}

// deprecated
func GenerateFilesByTemplates(projectRootDirectory string, templates map[string]string, tmplArgs map[string]interface{}) {
	for filePath, fileTmp := range templates {
		filePath = filepath.Join(projectRootDirectory, filePath)
		dir, _ := filepath.Split(filePath)
		targetDir := dir
		if !strings.HasPrefix(dir, "/") {
			targetDir = filepath.Join(projectRootDirectory, dir)
		}
		err := os.MkdirAll(targetDir, 0755)
		if err != nil {
			fmt.Errorf("%v", err)
			continue
		}

		GenerateFileByTemplate(filePath, fileTmp, tmplArgs)
	}
}

func GenerateFileByTemplate(filePath string, fileTmp string, tmplArgs map[string]interface{}) {
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Errorf("%v", err)
			return
		}
	} else {
		fmt.Println("skip file already exist:", filePath)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	fileContent := TemplateText(fileTmp, tmplArgs)
	f.WriteString(fileContent)
	f.Close()

}

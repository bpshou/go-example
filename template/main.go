package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

func main() {
	TestCobraGenEngine()
	// TestGenEngine()
}

func TestCobraGenEngine() {
	c := &cobra.Command{}

	var basePath string
	var serviceName string

	apiCmd := &cobra.Command{
		Use:   "Api",
		Short: "Api 生成",
		Long:  `Api 生成数据`,
		Run: func(cmd *cobra.Command, args []string) {
			basePathFlag, _ := cmd.Flags().GetString("base_path")
			serviceNameFlag, _ := cmd.Flags().GetString("name")
			fmt.Println("basePath", basePath)
			fmt.Println("serviceName", serviceName)
			fmt.Println("basePathFlag", basePathFlag)
			fmt.Println("serviceNameFlag", serviceNameFlag)
			// TestGenEngine()
		},
	}
	apiCmd.PersistentFlags().StringVarP(&basePath, "base_path", "b", "", "主要的生成路径，handler和logic的上游路径")
	apiCmd.PersistentFlags().StringVarP(&serviceName, "name", "n", "", "服务名称必填，用于生成handler和logic的文件名")
	c.AddCommand(apiCmd)
	// 执行
	c.Execute()
}

func TestGenEngine() {
	codeGen := NewGenEngine().
		AddTemplate("handler", "template/handler.tpl").
		AddTemplate("logic", "template/logic.tpl")

	moduleName, err := codeGen.GetModuleName()
	if err != nil {
		log.Fatal(err)
	}

	serviceName := "jwt"

	handlerPath := moduleName + "/handler"
	logicPath := moduleName + "/logic"

	handlerPackage := "api"
	logicPackage := "api"

	handlerWriter, err := CreateFile("handler/" + handlerPackage + "/" + serviceName + "_handler.go")
	if err != nil {
		log.Fatal(err)
	}
	logicWriter, err := CreateFile("logic/" + logicPackage + "/" + serviceName + "_logic.go")
	if err != nil {
		log.Fatal(err)
	}

	defer handlerWriter.Close()
	defer logicWriter.Close()

	codeGen.GenCode(handlerWriter, "handler", map[string]string{
		"serviceName":        serviceName,
		"handlerPackage":     handlerPackage,
		"logicPackage":       logicPackage,
		"logicImportPackage": logicPath + "/" + logicPackage,
	})
	codeGen.GenCode(logicWriter, "logic", map[string]string{
		"serviceName":  serviceName,
		"logicPackage": logicPackage,
	})

	fmt.Println(handlerPath)
}

func CreateFile(path string) (*os.File, error) {
	// 获取文件所在的目录
	fileDir := filepath.Dir(path)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

type GenEngine struct {
	Tmpl *template.Template
}

func NewGenEngine() *GenEngine {
	return &GenEngine{
		Tmpl: template.New("name").Funcs(sprig.FuncMap()),
	}
}

// 添加模板
func (c *GenEngine) AddTemplate(name string, filename string) *GenEngine {
	// 读取文件流
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 解析模板
	c.Tmpl, err = c.Tmpl.New(name).Parse(string(fileBytes))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return c
}

// 生成代码
func (c *GenEngine) GenCode(writer io.Writer, name string, data any) error {
	if writer == nil {
		writer = os.Stdout
	}

	if err := c.Tmpl.ExecuteTemplate(writer, name, data); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// 获取项目模块名称
func (c *GenEngine) GetModuleName() (string, error) {
	filename := "go.mod"
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	mf, err := modfile.Parse(filename, data, nil)
	if err != nil {
		return "", err
	}
	return mf.Module.Mod.Path, nil
}

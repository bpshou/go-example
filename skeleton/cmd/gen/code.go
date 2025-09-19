package gen

import (
	"fmt"
	"gin_app/app/core/tmpl"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/mod/modfile"
)

type GenCode struct {
	GenEngine *tmpl.GenEngine
}

func NewGenCode() *GenCode {
	return &GenCode{
		GenEngine: tmpl.NewGenEngine().
			AddTemplate("handler", "cmd/template/handler.tpl").
			AddTemplate("logic", "cmd/template/logic.tpl").
			AddTemplate("types", "cmd/template/types.tpl"),
	}
}

type HandlerValue struct {
	ServiceName        string
	ServicePackage     string
	LogicImportPackage string
}

func (g *GenCode) GenHandler(writer io.Writer, value HandlerValue) error {
	return g.GenEngine.GenCode(writer, "handler", value)
}

type LogicValue struct {
	ServiceName    string
	ServicePackage string
}

func (g *GenCode) GenLogic(writer io.Writer, value LogicValue) error {
	return g.GenEngine.GenCode(writer, "logic", value)
}

type TypesValue struct {
	ServiceName string
}

func (g *GenCode) GenTypes(writer io.Writer, value TypesValue) error {
	return g.GenEngine.GenCode(writer, "types", value)
}

type ConfigValue struct {
	ServicePackage     string
	ServiceName        string
	HandlerPath        string
	LogicPath          string
	TypesPath          string
	LogicImportPackage string
}

// 生成项目所用配置内容
func (g *GenCode) GenConfig(name string) ConfigValue {
	nameSlice := strings.Split(name, ".")
	if len(nameSlice) < 2 {
		log.Fatal("服务名称格式错误，请使用包名.服务名格式")
	}

	servicePackage := nameSlice[len(nameSlice)-2]
	serviceName := nameSlice[len(nameSlice)-1]
	nextSlice := nameSlice[:len(nameSlice)-1]

	// 除了最后一个元素之外的其他元素
	nextPath := strings.Join(nextSlice, "/")

	handlerPath := fmt.Sprintf("app/handler/%s/%s_handler.go", strcase.ToSnakeWithIgnore(nextPath, "/"), strcase.ToSnake(serviceName))
	logicPath := fmt.Sprintf("app/logic/%s/%s_logic.go", strcase.ToSnakeWithIgnore(nextPath, "/"), strcase.ToSnake(serviceName))
	logicImportPackage := fmt.Sprintf("%s/app/logic/%s", g.GetModuleName(), strcase.ToSnakeWithIgnore(nextPath, "/"))
	typesPath := "app/types/types.go"

	return ConfigValue{
		ServicePackage:     servicePackage,
		ServiceName:        serviceName,
		HandlerPath:        handlerPath,
		LogicPath:          logicPath,
		LogicImportPackage: logicImportPackage,
		TypesPath:          typesPath,
	}
}

// 获取项目模块名称
func (g *GenCode) GetModuleName() string {
	filename := "go.mod"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("read file", "err", err)
	}

	mf, err := modfile.Parse(filename, data, nil)
	if err != nil {
		log.Fatal("modfile Parse", "err", err)
	}
	return mf.Module.Mod.Path
}

func (g *GenCode) CreateFile(path string, append bool) (*os.File, error) {
	// 获取文件所在的目录
	fileDir := filepath.Dir(path)

	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	if append {
		return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	}
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
}

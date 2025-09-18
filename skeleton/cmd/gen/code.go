package gen

import (
	"gin_app/app/core/tmpl"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type GenCode struct {
	GenEngine *tmpl.GenEngine
}

func NewGenCode() *GenCode {
	return &GenCode{
		GenEngine: tmpl.NewGenEngine().
			AddTemplate("handler", "cmd/template/handler.tpl").
			AddTemplate("logic", "cmd/template/logic.tpl"),
	}
}

type HandlerValue struct {
	ServiceName        string //        serviceName
	ServicePackage     string //     ServicePackage
	LogicImportPackage string // logicPath + "/" + logicPackage
}

func (g *GenCode) GenHandler(writer io.Writer, value HandlerValue) error {
	return g.GenEngine.GenCode(writer, "handler", value)
}

type LogicValue struct {
	ServiceName    string //        serviceName
	ServicePackage string //     handlerPackage
}

func (g *GenCode) GenLogic(writer io.Writer, value LogicValue) error {
	return g.GenEngine.GenCode(writer, "logic", value)
}

// 获取项目模块名称
func (g *GenCode) GetModuleName() string {
	filename := "go.mod"
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("read file", "err", err)
		panic(err)
	}

	mf, err := modfile.Parse(filename, data, nil)
	if err != nil {
		slog.Error("modfile Parse", "err", err)
		panic(err)
	}
	return mf.Module.Mod.Path
}

func (g *GenCode) CreateFile(path string) (*os.File, error) {
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

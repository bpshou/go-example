package gen

import (
	"gin_app/app/core"
	"gin_app/app/core/db"
	"log"

	"github.com/spf13/cobra"
)

func Index() {
	c := &cobra.Command{}

	cobra.OnInitialize(core.Init)
	c.AddCommand(GenApi())
	c.AddCommand(GenModel())
	// 执行
	c.Execute()
}

func GenApi() *cobra.Command {
	var name string

	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Api 生成工具",
		Long:  `生成 handler 和 logic， go run main.go api -n`,
		Run: func(cmd *cobra.Command, args []string) {
			genCode := NewGenCode()
			conf := genCode.GenConfig(name)

			handlerWriter, err := genCode.CreateFile(conf.HandlerPath, false)
			if err != nil {
				log.Fatal(err)
			}

			logicWriter, err := genCode.CreateFile(conf.LogicPath, false)
			if err != nil {
				log.Fatal(err)
			}

			typesWriter, err := genCode.CreateFile(conf.TypesPath, true)
			if err != nil {
				log.Fatal(err)
			}

			genCode.GenHandler(handlerWriter, HandlerValue{
				ServiceName:        conf.ServiceName,
				ServicePackage:     conf.ServicePackage,
				LogicImportPackage: conf.LogicImportPackage,
			})
			genCode.GenLogic(logicWriter, LogicValue{
				ServiceName:    conf.ServiceName,
				ServicePackage: conf.ServicePackage,
			})
			genCode.GenTypes(typesWriter, TypesValue{
				ServiceName: conf.ServiceName,
			})

			defer handlerWriter.Close()
			defer logicWriter.Close()
			defer typesWriter.Close()
		},
	}
	apiCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "服务名称必填，用于生成handler和logic的文件名")

	return apiCmd
}

func GenModel() *cobra.Command {
	var path string

	apiCmd := &cobra.Command{
		Use:   "model",
		Short: "Model 生成工具",
		Long:  `生成 model， go run main.go model -m`,
		Run: func(cmd *cobra.Command, args []string) {
			db.GromGen(db.Db, path)
		},
	}
	apiCmd.Flags().StringVarP(&path, "path", "p", "app/models", "model生成存放的路径")

	return apiCmd
}

package gen

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func Index() {
	c := &cobra.Command{}

	c.AddCommand(GenApi())
	// 执行
	c.Execute()
}

func GenApi() *cobra.Command {
	var servicePackage string
	var serviceName string

	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "Api 生成工具",
		Long:  `生成 handler 和 logic， go run main.go api -b -n`,
		Run: func(cmd *cobra.Command, args []string) {
			genCode := NewGenCode()
			basePath := genCode.GetModuleName()

			handlerPath := fmt.Sprintf("app/handler/%s/%s_handler.go", servicePackage, serviceName)
			logicPath := fmt.Sprintf("app/logic/%s/%s_logic.go", servicePackage, serviceName)

			handlerWriter, err := genCode.CreateFile(handlerPath)
			if err != nil {
				log.Fatal(err)
			}

			logicWriter, err := genCode.CreateFile(logicPath)
			if err != nil {
				log.Fatal(err)
			}

			genCode.GenHandler(handlerWriter, HandlerValue{
				ServiceName:        serviceName,
				ServicePackage:     servicePackage,
				LogicImportPackage: fmt.Sprintf("%s/app/logic/%s", basePath, servicePackage),
			})
			genCode.GenLogic(logicWriter, LogicValue{
				ServiceName:    serviceName,
				ServicePackage: servicePackage,
			})

			defer handlerWriter.Close()
			defer logicWriter.Close()
		},
	}
	apiCmd.PersistentFlags().StringVarP(&servicePackage, "servicePackage", "p", "", "服务名称必填，用于生成handler和logic的文件名")
	apiCmd.PersistentFlags().StringVarP(&serviceName, "serviceName", "n", "", "服务名称必填，用于生成handler和logic的文件名")

	return apiCmd
}

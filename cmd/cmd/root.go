package cmd

import (
	"os"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "component",
	Short: "A brief description of your application",
	Long: `hey littler stranger`,

	Run: func(cmd *cobra.Command, args []string) {
		
		fs,err := os.ReadDir(".")

		logger.Logger(args)

		if err != nil {
			logger.Logger(err)
		}	

		for _, dir := range fs {
			logger.Logger(dir.Name())

			var  fileName =  args[1]
			
			
			if dir.Name() != "components" || dir.Name() != "Components" {

				err:= os.Mkdir("components",0755)

				if err != nil {
					logger.Logger(err)
					return 
				}
			}
			os.Mkdir("components/"+fileName,0755)
			file,err := os.Create("components/"+fileName+"/index.tsx",)

			if err != nil {
				logger.Logger(err)
			}	
		
				file.WriteString(`
export function `+fileName+`(){
	return (
		<div>
			<span>Hey little strange</span>
		</div>
	)
}
`)

		}
		logger.Logger("🚀",cmd.Long,"command executed success")
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



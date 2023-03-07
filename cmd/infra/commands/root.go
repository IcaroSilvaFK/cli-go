package commands

import (
	"os"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
)

var componentCommand = &cobra.Command{
	Use:   "component",
	Short: "Create a new React Component",
	Long:  `hey little stranger your component created successfully!`,

	Run: func(cmd *cobra.Command, args []string) {

		var fileName = args[0]

		err := os.MkdirAll("./src/components/"+fileName, os.ModePerm)

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		file, err := os.Create("./src/components/" + fileName + "/index.tsx")

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		file.WriteString(`
		import {Container} from './style'

		export function ` + fileName + `(){
			return (
				<Container>
					<span>Hey little strange</span>
				</Container>
			)
		}
		`)
		stylesFile, err := os.Create("./src/components/" + fileName + "/styles.ts")

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		stylesFile.WriteString(`
		import styled from 'styled-components'

		export const Container = styled.div
					`)

		logger.Log("🚀", cmd.Long, "command executed success")
	},
}

func Execute() {

	rootCommand := &cobra.Command{
		Use: "rc",
	}

	rootCommand.AddCommand(PageCommand, componentCommand, RouteCommand, HttpClient)

	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	PageCommand.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

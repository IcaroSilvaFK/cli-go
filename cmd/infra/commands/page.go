package commands

import (
	"os"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
)

var PageCommand = &cobra.Command{
	Use:   "page",
	Short: "Create a new React Page",
	Long:  `hey little stranger your page created successfully!`,

	Run: func(cmd *cobra.Command, args []string) {

		fileName := args[0]

		err := os.MkdirAll("./src/pages/"+fileName, os.ModePerm)

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		file, err := os.Create("./src/pages/" + fileName + "/index.tsx")

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		file.WriteString(`
import {Container} from './styles'

export function ` + fileName + `(){
	return (
		<Container>
			<span>Hey little strange</span>
		</Container>
	)
}
`)
		stylesFile, err := os.Create("./src/pages/" + fileName + "/styles")

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

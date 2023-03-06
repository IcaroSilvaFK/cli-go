package commands

import (
	"os"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var componentCommand = &cobra.Command{
	Use:   "component",
	Short: "Create a new React Component",
	Long:  `hey little stranger your component created successfully!`,

	Run: func(cmd *cobra.Command, args []string) {

		logger.Log(args[0])

		fs, err := os.ReadDir(".")
		var fileName = args[0]
		var directories []string
		var path []string

		logger.Log(fileName)

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		for _, dir := range fs {
			directories = append(directories, dir.Name())
			if dir.Name() == "src" {
				path = append(path, dir.Name())
			}
		}

		logger.Log(directories)
		logger.Log(path)

		//TODO add src directory if not exists
		if !slices.Contains(directories, "src") {
			err = os.Mkdir("src", os.ModePerm)
			path = append(path, "src")
			if err != nil {
				logger.Log("[ERROR]", err)
				return
			}
		}

		// if !slices.Contains(directories, "components") {
		// 	// err = os.Mkdir("components", os.ModePerm)
		// 	path = append(path, "components")
		// 	// if err != nil {
		// 	// 	logger.Log("[ERROR]", err)
		// 	// 	return
		// 	// }
		// }

		var absolutePath string

		for _, d := range path {
			absolutePath = absolutePath + "/" + d
		}

		err = os.MkdirAll("./"+absolutePath+"/components/"+fileName, os.ModePerm)

		if err != nil {
			logger.Log("[ERROR]", err)
			return
		}

		file, err := os.Create("./" + absolutePath + "/components/" + fileName + "/index.tsx")

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
		stylesFile, err := os.Create("./" + absolutePath + "/components/" + fileName + "/styles.ts")

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

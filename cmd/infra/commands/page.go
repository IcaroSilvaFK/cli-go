package commands

import (
	"os"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var PageCommand = &cobra.Command{
	Use:"page",
	Short:"Create a new React Page",
	Long:`hey little stranger your page created successfully!`,

	Run: func(cmd *cobra.Command, args []string){


		fs,err := os.ReadDir(".")
		var directories []string;
		fileName := args[0]

		if err != nil {
			logger.Log("[ERROR]",err)
		}


		for _,dir := range fs {
			directories = append(directories, dir.Name())
		}
			//TODO add src direction of not exists
		// if !slices.Contains(directories,"src"){
		// 	err = os.Mkdir("src",os.ModePerm)
			
		// 	if err != nil {
		// 		logger.Log("[ERROR]",err)
		// 		return
		// 	}
		// }
		if !slices.Contains(directories,"pages"){
			err = os.Mkdir("pages",os.ModePerm)

			if err != nil {
				logger.Log("[ERROR]",err)
				return
			}
		}

		err = os.MkdirAll("pages/"+fileName, os.ModePerm)		

		if err != nil {
			logger.Log("[ERROR]",err)
			return
		}

		file, err := os.Create("pages/"+fileName+"/index.tsx")


		if err != nil {
			logger.Log("[ERROR]",err)
			return
		}

		file.WriteString(`
import {Container} from './styles'

export function `+fileName+`(){
	return (
		<Container>
			<span>Hey little strange</span>
		</Container>
	)
}
`)
stylesFile,err := os.Create("pages/"+fileName+"/styles.ts")

if err != nil {
	logger.Log("[ERROR]",err)
	return
}

stylesFile.WriteString(`
import styled from 'styled-components'

export const Container = styled.div
`)
	logger.Log("🚀",cmd.Long,"command executed success")
	}	,
}


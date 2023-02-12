package commands

import (
	"os"
	"strings"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
)

var RouteCommand = &cobra.Command{
	Use:"router-generate",
	Short:`Read your pages folder and generate Router component`,
	Long: `
		✨🎉Generate router component! Update your Main.tsx
	`,

	Run: func (command *cobra.Command, args []string){

		var allPages []string

		dirs,err := os.ReadDir("pages")

		if err != nil {
			logger.Log("[ERROR]",err)
			return 
		}

		for _,dir := range dirs{
			allPages = append(allPages, "\n<Route path='"+dir.Name()+"' component={<"+dir.Name()+"/>}/>")
		}


		os.MkdirAll("routes",os.ModePerm)

		file, err := os.Create("routes/index.tsx")


		if err != nil {
			logger.Log("[ERROR]",err)
			return 
		}

		_,err = file.WriteString(`
import {BrowserRouter,Routes as Router,Route} from 'react-router-dom'

export function Router(){
	return (
		<BrowserRouter>
			<Router>
					`+strings.Join(allPages, "\n")+`
			</Router>
		</BrowserRouter>
	)
}
`)

	if err != nil {
		logger.Log("[ERROR]",err)
		return 
	}

		logger.Log(command.Long)

	},	

}


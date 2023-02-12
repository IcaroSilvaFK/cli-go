package commands

import (
	"io/ioutil"
	"net/http"

	"github.com/IcaroSilvaFK/cli-go/cmd/infra/logger"
	"github.com/spf13/cobra"
)

var HttpClient = &cobra.Command{
	Use:"get",
	Long: `
		✨🎉Request made successfully
	`,
	Short:"⁉️Paste the endpoint to make the request",

	Run: func(command *cobra.Command, args []string){
		endpoint := args[0]

		if len(endpoint) == 0 {
			logger.Log("[ERROR]:provide an endpoint for the request🤔")
			return
		}

		resp,err := http.Get(endpoint)


		if err != nil {
			logger.Log("[ERROR]:request fail!😥")
			return
		}

		body, err := ioutil.ReadAll(resp.Body)		


		if err != nil {
			logger.Log("[ERROR]:body is invalid!😥")
			return
		}

		castBodyFromString := string(body)


		logger.JsonLogger(castBodyFromString)

		logger.Log(command.Long)
	},	

}
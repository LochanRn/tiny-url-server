package main

import (
	"log"
	"os"

	"github.com/LochanRn/tiny-url-server/config"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/LochanRn/tiny-url-server/gen/restapi"
	"github.com/LochanRn/tiny-url-server/gen/restapi/operations"
	"github.com/LochanRn/tiny-url-server/utils/logger"
)

func main() {

	logger.InitLogger(config.GetLogLevel())
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTinyURLServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "tiny-url-server"
	parser.LongDescription = "This is a tiny url server project"

	server.Host = "0.0.0.0"
	server.Port = 2112
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

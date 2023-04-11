package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-demo/pkg/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return terraform.Provider()
		},
	}

	if debugMode {
		// TODO: update this string with the full name of your provider as used in your configs
		userPluginDir := "/root/.terraform/plugins/hv/hitachi/2.0/linux_amd64/terraform-provider-demo"
		err := plugin.Debug(context.Background(), userPluginDir, opts)

		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}

package terraform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: nil, // schemaimpl.func
		ResourcesMap: map[string]*schema.Resource{
			"test": nil, // resourceimpl.func()
		},
		DataSourcesMap: map[string]*schema.Resource{
			"test2": nil,
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// TO DO
	return nil, nil
}

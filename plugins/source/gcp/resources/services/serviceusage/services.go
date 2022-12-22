// Code generated by codegen; DO NOT EDIT.

package serviceusage

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_serviceusage_services",
		Description: `https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplexEnabledServices("serviceusage.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "parent",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Parent"),
			},
			{
				Name:     "config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Config"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
		},
	}
}

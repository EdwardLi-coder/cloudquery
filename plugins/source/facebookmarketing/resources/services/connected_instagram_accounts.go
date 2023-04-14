package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ConnectedInstagramAccounts() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_connected_instagram_accounts",
		Resolver:  fetchConnectedInstagramAccounts,
		Transform: client.TransformWithStruct(&rest.ConnnectedInstagramAccount{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Resolver: client.ResolveAccountId,
				Type:     schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

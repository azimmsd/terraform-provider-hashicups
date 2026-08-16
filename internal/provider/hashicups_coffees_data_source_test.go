// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func TestHashicupsCoffeesDataSourceIsRegistered(t *testing.T) {
	p := New("test")()
	dataSources := p.DataSources(context.Background())

	if len(dataSources) == 0 {
		t.Fatal("expected at least one data source to be registered")
	}

	for _, factory := range dataSources {
		ds := factory()
		var resp datasource.MetadataResponse
		ds.Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "hashicups"}, &resp)

		if resp.TypeName == "hashicups_coffees" {
			return
		}
	}

	t.Fatal("hashicups_coffees data source was not registered")
}

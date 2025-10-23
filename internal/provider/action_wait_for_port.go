// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ action.Action = &WaitForPortAction{}

func NewWaitForPortAction() action.Action {
	return &WaitForPortAction{}
}

const interval = 500 * time.Millisecond

// WaitForPortAction defines tha action implementation.
type WaitForPortAction struct {
}

// WaitForPortActionModel describes the data source data model.
type WaitForPortActionModel struct {
	Host             types.String `tfsdk:"host"`
	Port             types.Int64  `tfsdk:"port"`
	TimeoutInSeconds types.Int64  `tfsdk:"timeout"`
}

func (d *WaitForPortAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_for_port"
}

func (d *WaitForPortAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "This action waits for the port to become available.",

		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "The host to try to connect to.",
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "The port to wait for",
				Required:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "The timeout in seconds to wait for the port to become available. Defaults to 120 seconds.",
				Optional:            true,
			},
		},
	}
}

func (d *WaitForPortAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var data WaitForPortActionModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	host := data.Host.ValueString()
	port := data.Port.ValueInt64()
	timeout := int(data.TimeoutInSeconds.ValueInt64())
	if timeout == 0 {
		timeout = 120
	}

	address := fmt.Sprintf("%s:%d", host, port)
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		select {
		case <-ctx.Done():
			resp.Diagnostics.AddError(
				"Context Canceled",
				fmt.Sprintf("The context was canceled while waiting for %q", address),
			)
			return
		default:
			conn, err := net.DialTimeout("tcp", address, time.Second)
			if err == nil {
				conn.Close()
				return // Port is available
			}

			if time.Now().After(deadline) {
				resp.Diagnostics.AddError(
					"Timeout Reached",
					fmt.Sprintf("Reached the timeout while waiting for %q", address),
				)
			}

			time.Sleep(interval) // Retry after a short delay
		}
	}
}

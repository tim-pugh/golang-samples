// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// findings contains example snippets for working with findings
// and there parent resource "sources".
package findings

// [START get_source]
import (
	"context"
	"fmt"
	"io"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"
)

// getSource retrieves a source by its resource name and print it to w.
// sourceName is the full resource name of the source to be updated.  Returns
// the Source.
func getSource(w io.Writer, sourceName string) (*securitycenterpb.Source, error) {
	// sourceName := "organizations/111122222444/sources/1234"
	// Instantiate a context and a security service client to make API calls.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error instantiating client %v\n", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	req := &securitycenterpb.GetSourceRequest{
		Name: sourceName,
	}
	source, err := client.GetSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Error getting source: %v", err)
	}
	fmt.Fprintf(w, "Source: %v", source)
	return source, nil
}

// [END get_source]
// Copyright 2019 Cruise LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    https://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package ingress

package main

import (
	"context"
	"flag"

	"github.com/cruise-automation/k-rail/server"
)

var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file: `<home>/.kube/config`")

func main() {
	flag.Parse()
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	server.Run(ctx)
}

//TODO: add kubeconfig flags (from unique_ingress_host)

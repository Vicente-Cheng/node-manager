/*
Copyright 2023 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/harvester/node-manager/pkg/generated/clientset/versioned/typed/node.harvesterhci.io/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNodeV1beta1 struct {
	*testing.Fake
}

func (c *FakeNodeV1beta1) Ksmtuneds() v1beta1.KsmtunedInterface {
	return &FakeKsmtuneds{c}
}

func (c *FakeNodeV1beta1) NodeConfigs(namespace string) v1beta1.NodeConfigInterface {
	return &FakeNodeConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNodeV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

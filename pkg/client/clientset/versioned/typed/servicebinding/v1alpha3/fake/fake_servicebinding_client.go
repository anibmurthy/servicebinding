/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/servicebinding/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicebindingV1alpha3 struct {
	*testing.Fake
}

func (c *FakeServicebindingV1alpha3) ServiceBindings(namespace string) v1alpha3.ServiceBindingInterface {
	return &FakeServiceBindings{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicebindingV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
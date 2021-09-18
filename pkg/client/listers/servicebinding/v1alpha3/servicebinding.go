/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/vmware-labs/service-bindings/pkg/apis/servicebinding/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceBindingLister helps list ServiceBindings.
// All objects returned here must be treated as read-only.
type ServiceBindingLister interface {
	// List lists all ServiceBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceBinding, err error)
	// ServiceBindings returns an object that can list and get ServiceBindings.
	ServiceBindings(namespace string) ServiceBindingNamespaceLister
	ServiceBindingListerExpansion
}

// serviceBindingLister implements the ServiceBindingLister interface.
type serviceBindingLister struct {
	indexer cache.Indexer
}

// NewServiceBindingLister returns a new ServiceBindingLister.
func NewServiceBindingLister(indexer cache.Indexer) ServiceBindingLister {
	return &serviceBindingLister{indexer: indexer}
}

// List lists all ServiceBindings in the indexer.
func (s *serviceBindingLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceBinding))
	})
	return ret, err
}

// ServiceBindings returns an object that can list and get ServiceBindings.
func (s *serviceBindingLister) ServiceBindings(namespace string) ServiceBindingNamespaceLister {
	return serviceBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceBindingNamespaceLister helps list and get ServiceBindings.
// All objects returned here must be treated as read-only.
type ServiceBindingNamespaceLister interface {
	// List lists all ServiceBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceBinding, err error)
	// Get retrieves the ServiceBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.ServiceBinding, error)
	ServiceBindingNamespaceListerExpansion
}

// serviceBindingNamespaceLister implements the ServiceBindingNamespaceLister
// interface.
type serviceBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceBindings in the indexer for a given namespace.
func (s serviceBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceBinding))
	})
	return ret, err
}

// Get retrieves the ServiceBinding from the indexer for a given namespace and name.
func (s serviceBindingNamespaceLister) Get(name string) (*v1alpha3.ServiceBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("servicebinding"), name)
	}
	return obj.(*v1alpha3.ServiceBinding), nil
}
/*
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSHostedZonePolicyLister helps list DNSHostedZonePolicies.
// All objects returned here must be treated as read-only.
type DNSHostedZonePolicyLister interface {
	// List lists all DNSHostedZonePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DNSHostedZonePolicy, err error)
	// DNSHostedZonePolicies returns an object that can list and get DNSHostedZonePolicies.
	DNSHostedZonePolicies(namespace string) DNSHostedZonePolicyNamespaceLister
	DNSHostedZonePolicyListerExpansion
}

// dNSHostedZonePolicyLister implements the DNSHostedZonePolicyLister interface.
type dNSHostedZonePolicyLister struct {
	indexer cache.Indexer
}

// NewDNSHostedZonePolicyLister returns a new DNSHostedZonePolicyLister.
func NewDNSHostedZonePolicyLister(indexer cache.Indexer) DNSHostedZonePolicyLister {
	return &dNSHostedZonePolicyLister{indexer: indexer}
}

// List lists all DNSHostedZonePolicies in the indexer.
func (s *dNSHostedZonePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.DNSHostedZonePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSHostedZonePolicy))
	})
	return ret, err
}

// DNSHostedZonePolicies returns an object that can list and get DNSHostedZonePolicies.
func (s *dNSHostedZonePolicyLister) DNSHostedZonePolicies(namespace string) DNSHostedZonePolicyNamespaceLister {
	return dNSHostedZonePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DNSHostedZonePolicyNamespaceLister helps list and get DNSHostedZonePolicies.
// All objects returned here must be treated as read-only.
type DNSHostedZonePolicyNamespaceLister interface {
	// List lists all DNSHostedZonePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DNSHostedZonePolicy, err error)
	// Get retrieves the DNSHostedZonePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DNSHostedZonePolicy, error)
	DNSHostedZonePolicyNamespaceListerExpansion
}

// dNSHostedZonePolicyNamespaceLister implements the DNSHostedZonePolicyNamespaceLister
// interface.
type dNSHostedZonePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DNSHostedZonePolicies in the indexer for a given namespace.
func (s dNSHostedZonePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DNSHostedZonePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSHostedZonePolicy))
	})
	return ret, err
}

// Get retrieves the DNSHostedZonePolicy from the indexer for a given namespace and name.
func (s dNSHostedZonePolicyNamespaceLister) Get(name string) (*v1alpha1.DNSHostedZonePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnshostedzonepolicy"), name)
	}
	return obj.(*v1alpha1.DNSHostedZonePolicy), nil
}

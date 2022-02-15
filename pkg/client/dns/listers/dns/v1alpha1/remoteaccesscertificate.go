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

// RemoteAccessCertificateLister helps list RemoteAccessCertificates.
// All objects returned here must be treated as read-only.
type RemoteAccessCertificateLister interface {
	// List lists all RemoteAccessCertificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RemoteAccessCertificate, err error)
	// RemoteAccessCertificates returns an object that can list and get RemoteAccessCertificates.
	RemoteAccessCertificates(namespace string) RemoteAccessCertificateNamespaceLister
	RemoteAccessCertificateListerExpansion
}

// remoteAccessCertificateLister implements the RemoteAccessCertificateLister interface.
type remoteAccessCertificateLister struct {
	indexer cache.Indexer
}

// NewRemoteAccessCertificateLister returns a new RemoteAccessCertificateLister.
func NewRemoteAccessCertificateLister(indexer cache.Indexer) RemoteAccessCertificateLister {
	return &remoteAccessCertificateLister{indexer: indexer}
}

// List lists all RemoteAccessCertificates in the indexer.
func (s *remoteAccessCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.RemoteAccessCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RemoteAccessCertificate))
	})
	return ret, err
}

// RemoteAccessCertificates returns an object that can list and get RemoteAccessCertificates.
func (s *remoteAccessCertificateLister) RemoteAccessCertificates(namespace string) RemoteAccessCertificateNamespaceLister {
	return remoteAccessCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RemoteAccessCertificateNamespaceLister helps list and get RemoteAccessCertificates.
// All objects returned here must be treated as read-only.
type RemoteAccessCertificateNamespaceLister interface {
	// List lists all RemoteAccessCertificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RemoteAccessCertificate, err error)
	// Get retrieves the RemoteAccessCertificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RemoteAccessCertificate, error)
	RemoteAccessCertificateNamespaceListerExpansion
}

// remoteAccessCertificateNamespaceLister implements the RemoteAccessCertificateNamespaceLister
// interface.
type remoteAccessCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RemoteAccessCertificates in the indexer for a given namespace.
func (s remoteAccessCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RemoteAccessCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RemoteAccessCertificate))
	})
	return ret, err
}

// Get retrieves the RemoteAccessCertificate from the indexer for a given namespace and name.
func (s remoteAccessCertificateNamespaceLister) Get(name string) (*v1alpha1.RemoteAccessCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("remoteaccesscertificate"), name)
	}
	return obj.(*v1alpha1.RemoteAccessCertificate), nil
}

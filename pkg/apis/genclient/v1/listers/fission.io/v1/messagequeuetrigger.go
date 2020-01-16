/*
Copyright The Fission Authors.

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

package v1

import (
	v1 "github.com/fission/fission/pkg/apis/fission.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MessageQueueTriggerLister helps list MessageQueueTriggers.
type MessageQueueTriggerLister interface {
	// List lists all MessageQueueTriggers in the indexer.
	List(selector labels.Selector) (ret []*v1.MessageQueueTrigger, err error)
	// MessageQueueTriggers returns an object that can list and get MessageQueueTriggers.
	MessageQueueTriggers(namespace string) MessageQueueTriggerNamespaceLister
	MessageQueueTriggerListerExpansion
}

// _messageQueueTriggerLister implements the MessageQueueTriggerLister interface.
type _messageQueueTriggerLister struct {
	indexer cache.Indexer
}

// NewMessageQueueTriggerLister returns a new MessageQueueTriggerLister.
func NewMessageQueueTriggerLister(indexer cache.Indexer) MessageQueueTriggerLister {
	return &_messageQueueTriggerLister{indexer: indexer}
}

// List lists all MessageQueueTriggers in the indexer.
func (s *_messageQueueTriggerLister) List(selector labels.Selector) (ret []*v1.MessageQueueTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessageQueueTrigger))
	})
	return ret, err
}

// MessageQueueTriggers returns an object that can list and get MessageQueueTriggers.
func (s *_messageQueueTriggerLister) MessageQueueTriggers(namespace string) MessageQueueTriggerNamespaceLister {
	return _messageQueueTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MessageQueueTriggerNamespaceLister helps list and get MessageQueueTriggers.
type MessageQueueTriggerNamespaceLister interface {
	// List lists all MessageQueueTriggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MessageQueueTrigger, err error)
	// Get retrieves the MessageQueueTrigger from the indexer for a given namespace and name.
	Get(name string) (*v1.MessageQueueTrigger, error)
	MessageQueueTriggerNamespaceListerExpansion
}

// _messageQueueTriggerNamespaceLister implements the MessageQueueTriggerNamespaceLister
// interface.
type _messageQueueTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MessageQueueTriggers in the indexer for a given namespace.
func (s _messageQueueTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1.MessageQueueTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MessageQueueTrigger))
	})
	return ret, err
}

// Get retrieves the MessageQueueTrigger from the indexer for a given namespace and name.
func (s _messageQueueTriggerNamespaceLister) Get(name string) (*v1.MessageQueueTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("messagequeuetrigger"), name)
	}
	return obj.(*v1.MessageQueueTrigger), nil
}

/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	networkv1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1"
	versioned "github.com/GoogleCloudPlatform/gke-networking-api/client/network/clientset/versioned"
	internalinterfaces "github.com/GoogleCloudPlatform/gke-networking-api/client/network/informers/externalversions/internalinterfaces"
	v1 "github.com/GoogleCloudPlatform/gke-networking-api/client/network/listers/network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SubnetworkInformer provides access to a shared informer and lister for
// Subnetworks.
type SubnetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SubnetworkLister
}

type subnetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSubnetworkInformer constructs a new informer for Subnetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSubnetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSubnetworkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSubnetworkInformer constructs a new informer for Subnetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSubnetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Subnetworks().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Subnetworks().Watch(context.TODO(), options)
			},
		},
		&networkv1.Subnetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *subnetworkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSubnetworkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *subnetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1.Subnetwork{}, f.defaultInformer)
}

func (f *subnetworkInformer) Lister() v1.SubnetworkLister {
	return v1.NewSubnetworkLister(f.Informer().GetIndexer())
}

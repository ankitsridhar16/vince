/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	vincev1alpha1 "github.com/vinceanalytics/vince/pkg/apis/vince/v1alpha1"
	versioned "github.com/vinceanalytics/vince/pkg/gen/client/vince/clientset/versioned"
	internalinterfaces "github.com/vinceanalytics/vince/pkg/gen/client/vince/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vinceanalytics/vince/pkg/gen/client/vince/listers/vince/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VinceInformer provides access to a shared informer and lister for
// Vinces.
type VinceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VinceLister
}

type vinceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVinceInformer constructs a new informer for Vince type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVinceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVinceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVinceInformer constructs a new informer for Vince type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVinceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StaplesV1alpha1().Vinces(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StaplesV1alpha1().Vinces(namespace).Watch(context.TODO(), options)
			},
		},
		&vincev1alpha1.Vince{},
		resyncPeriod,
		indexers,
	)
}

func (f *vinceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVinceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vinceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&vincev1alpha1.Vince{}, f.defaultInformer)
}

func (f *vinceInformer) Lister() v1alpha1.VinceLister {
	return v1alpha1.NewVinceLister(f.Informer().GetIndexer())
}

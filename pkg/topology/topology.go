package topology

import (
	"fmt"

	"github.com/gernest/vince/pkg/apis/vince/v1alpha1"
	vince_listers "github.com/gernest/vince/pkg/gen/client/vince/listers/vince/v1alpha1"
	"github.com/gernest/vince/pkg/k8s"
	apppsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	listers "k8s.io/client-go/listers/core/v1"
)

type Topology struct {
	serviceLister listers.ServiceLister
	vinceLister   vince_listers.VinceLister
	siteLister    vince_listers.SiteLister
	podLister     listers.PodLister
	secretsLister listers.SecretLister
	configLister  listers.ConfigMapLister
}

func (t *Topology) Build(filter *k8s.ResourceFilter) error {
	r, err := t.loadResources(filter)
	if err != nil {
		return err
	}
	r.Resolve()
	return nil
}

func (t *Topology) loadResources(filter *k8s.ResourceFilter) (*Resources, error) {
	r := &Resources{
		Services: make(map[string]*corev1.Service),
		Secrets:  make(map[string]*corev1.Secret),
		Configs:  make(map[string]*corev1.ConfigMap),
		Pods:     make(map[string]*corev1.Pod),
		Vinces:   make(map[string]*v1alpha1.Vince),
		Sites:    make(map[string]*v1alpha1.Site),
	}
	svc, err := t.serviceLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list services %v", err)
	}
	secrets, err := t.secretsLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list secrets %v", err)
	}
	config, err := t.configLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list config maps %v", err)
	}
	pods, err := t.podLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list pods maps %v", err)
	}
	vince, err := t.vinceLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list vinces maps %v", err)
	}
	site, err := t.siteLister.List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list vinces maps %v", err)
	}
	for _, o := range svc {
		if filter.IsIgnored(o) {
			continue
		}
		r.Services[key(o)] = o
	}
	for _, o := range secrets {
		if filter.IsIgnored(o) {
			continue
		}
		r.Secrets[key(o)] = o
	}
	for _, o := range config {
		if filter.IsIgnored(o) {
			continue
		}
		r.Configs[key(o)] = o
	}
	for _, o := range pods {
		if filter.IsIgnored(o) {
			continue
		}
		r.Pods[key(o)] = o
	}
	for _, o := range vince {
		if filter.IsIgnored(o) {
			continue
		}
		r.Vinces[key(o)] = o
	}
	for _, o := range site {
		if filter.IsIgnored(o) {
			continue
		}
		r.Sites[key(o)] = o
	}
	return r, nil
}

type Resources struct {
	Services map[string]*corev1.Service
	Secrets  map[string]*corev1.Secret
	Configs  map[string]*corev1.ConfigMap
	Pods     map[string]*corev1.Pod
	Vinces   map[string]*v1alpha1.Vince
	Sites    map[string]*v1alpha1.Site
}

func (r *Resources) Resolve() *ChangeSet {
	return nil
}

type ChangeSet struct {
	Secrets      []*corev1.Secret
	Configs      []*corev1.ConfigMap
	Services     []*corev1.Service
	VinceStatus  []v1alpha1.VinceStatus
	StatefulSets []*apppsv1.StatefulSet
}

func key(o metav1.Object) string {
	ts := types.NamespacedName{
		Namespace: o.GetNamespace(),
		Name:      o.GetName(),
	}
	return ts.String()
}

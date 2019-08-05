/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	settingsv1alpha1 "k8s.io/api/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/settings/v1alpha1"
)

var _ v1alpha1.PodPresetLister = &podPresetLister{}

var _ v1alpha1.PodPresetNamespaceLister = &podPresetNamespaceLister{}

// podPresetLister implements the PodPresetLister interface.
type podPresetLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPodPresetLister returns a new PodPresetLister.
func NewPodPresetLister(client kubernetes.Interface) v1alpha1.PodPresetLister {
	return NewFilteredPodPresetLister(client, nil)
}

func NewFilteredPodPresetLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1alpha1.PodPresetLister {
	return &podPresetLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all PodPresets in the indexer.
func (s *podPresetLister) List(selector labels.Selector) (ret []*settingsv1alpha1.PodPreset, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.SettingsV1alpha1().PodPresets(v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// PodPresets returns an object that can list and get PodPresets.
func (s *podPresetLister) PodPresets(namespace string) v1alpha1.PodPresetNamespaceLister {
	return podPresetNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// podPresetNamespaceLister implements the PodPresetNamespaceLister
// interface.
type podPresetNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all PodPresets in the indexer for a given namespace.
func (s podPresetNamespaceLister) List(selector labels.Selector) (ret []*settingsv1alpha1.PodPreset, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.SettingsV1alpha1().PodPresets(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the PodPreset from the indexer for a given namespace and name.
func (s podPresetNamespaceLister) Get(name string) (*settingsv1alpha1.PodPreset, error) {
	return s.client.SettingsV1alpha1().PodPresets(s.namespace).Get(name, v1.GetOptions{})
}
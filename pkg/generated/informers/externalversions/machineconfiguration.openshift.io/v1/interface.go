// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ControllerConfigs returns a ControllerConfigInformer.
	ControllerConfigs() ControllerConfigInformer
	// MachineConfigs returns a MachineConfigInformer.
	MachineConfigs() MachineConfigInformer
	// MachineConfigPools returns a MachineConfigPoolInformer.
	MachineConfigPools() MachineConfigPoolInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ControllerConfigs returns a ControllerConfigInformer.
func (v *version) ControllerConfigs() ControllerConfigInformer {
	return &controllerConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineConfigs returns a MachineConfigInformer.
func (v *version) MachineConfigs() MachineConfigInformer {
	return &machineConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MachineConfigPools returns a MachineConfigPoolInformer.
func (v *version) MachineConfigPools() MachineConfigPoolInformer {
	return &machineConfigPoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

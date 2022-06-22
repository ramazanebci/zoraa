// Package v1alpha1 contains API Schema definitions for the zora v1alpha1 API group
//+kubebuilder:object:generate=true
//+groupName=zora.undistro.io
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "zora.undistro.io", Version: "v1alpha1"}

	// SchemeGroupVersion is a copy of GroupVersion, used by client-gen
	SchemeGroupVersion = GroupVersion

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

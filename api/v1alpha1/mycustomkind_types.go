/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MyCustomKindSpec defines the desired state of MyCustomKind
type MyCustomKindSpec struct {
	// Foo represents the original key in v1alpha1.
	Foo string `json:"foo,omitempty"`
	// Zap represents a deprecated key of a non-string type in v1alpha1.
	Zap int32 `json:"zap,omitempty"`
}

// MyCustomKindStatus defines the observed state of MyCustomKind
type MyCustomKindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion

// MyCustomKind is the Schema for the mycustomkinds API
type MyCustomKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyCustomKindSpec   `json:"spec,omitempty"`
	Status MyCustomKindStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyCustomKindList contains a list of MyCustomKind
type MyCustomKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyCustomKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyCustomKind{}, &MyCustomKindList{})
}

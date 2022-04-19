/*
Copyright 2022.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HttpEchoSpec defines the desired state of HttpEcho
type HttpEchoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size allows to configure the amount of replicas for HttpEcho.
	Size int32 `json:"size"`
}

// HttpEchoStatus defines the observed state of HttpEcho
type HttpEchoStatus struct {
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HttpEcho is the Schema for the httpechoes API
type HttpEcho struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HttpEchoSpec   `json:"spec,omitempty"`
	Status HttpEchoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HttpEchoList contains a list of HttpEcho
type HttpEchoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HttpEcho `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HttpEcho{}, &HttpEchoList{})
}

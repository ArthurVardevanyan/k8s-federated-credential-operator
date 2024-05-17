/*
Copyright 2024.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SourceServiceAccount struct {
	// Name of Source Service Account
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type TargetServiceAccount struct {
	// Name of Target Service Account
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Namespace of Target Service Account
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace"`
	// The Endpoint used to exchange the token
	// +kubebuilder:validation:Required
	ClusterExchangeTokenEndpoint string `json:"clusterExchangeTokenEndpoint"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	// Contains the Fields Related to configuring the Source Service Account
	SourceServiceAccount SourceServiceAccount `json:"sourceServiceAccount"`
	// Contains the Fields Related to configuring the Source Service Account
	TargetServiceAccount TargetServiceAccount `json:"targetServiceAccount"`
	// Name of the Secret to Generate
	// +kubebuilder:validation:Required
	SecretName string `json:"secretName"`
}

// TokenStatus defines the observed state of Token
type TokenStatus struct {
	// When the Current Token Expires
	TokenExpiration string `json:"tokenExpiration,omitempty"`
	// Output of Any Errors
	Error string `json:"error,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Token is the Schema for the tokens API
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TokenSpec   `json:"spec,omitempty"`
	Status TokenStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TokenList contains a list of Token
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}

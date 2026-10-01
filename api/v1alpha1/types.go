/*
SPDX-FileCopyrightText: 2026 SAP SE or an SAP affiliate company and kafka-operator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// KafkaOperatorSpec defines the desired state of KafkaOperator.
type KafkaOperatorSpec struct {
	component.PlacementSpec `json:",inline"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	Image                          component.ImageSpec `json:"image"`
	component.KubernetesProperties `json:",inline"`
}

// KafkaOperatorStatus defines the observed state of KafkaOperator.
type KafkaOperatorStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=='Ready')].reason`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=`.metadata.creationTimestamp`
// +genclient

// KafkaOperator is the Schema for the kafkaoperators API.
type KafkaOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KafkaOperatorSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status KafkaOperatorStatus `json:"status,omitempty"`
}

var _ component.Component = &KafkaOperator{}

// +kubebuilder:object:root=true

// KafkaOperatorList contains a list of KafkaOperator.
type KafkaOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaOperator `json:"items"`
}

func (s *KafkaOperatorSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *KafkaOperator) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *KafkaOperator) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&KafkaOperator{}, &KafkaOperatorList{})
}

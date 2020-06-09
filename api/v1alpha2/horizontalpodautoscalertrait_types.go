/*


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

package v1alpha2

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HorizontalPodAutoscalerTraitSpec defines the desired state of HorizontalPodAutoscalerTrait
type HorizontalPodAutoscalerTraitSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	MaxReplicas int32 `json:"maxReplicas" protobuf:"varint,1,opt,name=maxReplicas"`

	Template autoscalingv1.HorizontalPodAutoscalerSpec `json:"template,omitempty"`

	// WorkloadReference to the workload this trait applies to.
	WorkloadReference runtimev1alpha1.TypedReference `json:"workloadRef"`
}

// HorizontalPodAutoscalerTraitStatus defines the observed state of HorizontalPodAutoscalerTrait
type HorizontalPodAutoscalerTraitStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	DesiredReplicas int32 `json:"desiredReplicas" protobuf:"varint,1,opt,name=desiredReplicas"`

	runtimev1alpha1.ConditionedStatus `json:",inline"`

	// Resources managed by this service trait
	Resources []runtimev1alpha1.TypedReference `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={crossplane,oam}

// HorizontalPodAutoscalerTrait is the Schema for the horizontalpodautoscalertraits API
type HorizontalPodAutoscalerTrait struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HorizontalPodAutoscalerTraitSpec   `json:"spec,omitempty"`
	Status HorizontalPodAutoscalerTraitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HorizontalPodAutoscalerTraitList contains a list of HorizontalPodAutoscalerTrait
type HorizontalPodAutoscalerTraitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HorizontalPodAutoscalerTrait `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HorizontalPodAutoscalerTrait{}, &HorizontalPodAutoscalerTraitList{})
}

// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerTrait) DeepCopyInto(out *HorizontalPodAutoscalerTrait) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerTrait.
func (in *HorizontalPodAutoscalerTrait) DeepCopy() *HorizontalPodAutoscalerTrait {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizontalPodAutoscalerTrait) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerTraitList) DeepCopyInto(out *HorizontalPodAutoscalerTraitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorizontalPodAutoscalerTrait, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerTraitList.
func (in *HorizontalPodAutoscalerTraitList) DeepCopy() *HorizontalPodAutoscalerTraitList {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerTraitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizontalPodAutoscalerTraitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerTraitSpec) DeepCopyInto(out *HorizontalPodAutoscalerTraitSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerTraitSpec.
func (in *HorizontalPodAutoscalerTraitSpec) DeepCopy() *HorizontalPodAutoscalerTraitSpec {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerTraitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerTraitStatus) DeepCopyInto(out *HorizontalPodAutoscalerTraitStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerTraitStatus.
func (in *HorizontalPodAutoscalerTraitStatus) DeepCopy() *HorizontalPodAutoscalerTraitStatus {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerTraitStatus)
	in.DeepCopyInto(out)
	return out
}

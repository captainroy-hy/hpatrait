package controllers

import (
	"context"
	"errors"
	"reflect"

	"github.com/crossplane/oam-kubernetes-runtime/pkg/oam"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	coreoamdevv1alpha2 "hpatrait/api/v1alpha2"
)

var (
	hpaKind         = reflect.TypeOf(autoscalingv1.HorizontalPodAutoscaler{}).Name()
	hpaGroupVersion = autoscalingv1.SchemeGroupVersion.String()
)

const LabelKey = "hpatrait.oam.crossplane.io"

func (r *HorizontalPodAutoscalerTraitReconciler) renderHPA(ctx context.Context, trait oam.Trait) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	t, ok := trait.(*coreoamdevv1alpha2.HorizontalPodAutoscalerTrait)
	if !ok {
		return nil, errors.New("not a hpa trait")
	}
	hpa := &autoscalingv1.HorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			Kind:       hpaKind,
			APIVersion: hpaGroupVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.GetName(),
			Namespace: t.GetNamespace(),
			Labels: map[string]string{
				LabelKey: string(t.GetUID()),
			},
		},
		Spec: autoscalingv1.HorizontalPodAutoscalerSpec{
			//TODO construct HPASpec according to workloadRef
			ScaleTargetRef: autoscalingv1.CrossVersionObjectReference{
				Kind:       "Deployment",
				Name:       "TestDeployment",
				APIVersion: "apps/v1",
			},
			MaxReplicas: t.Spec.MaxReplicas,
		},
	}
	if err := ctrl.SetControllerReference(trait, hpa, r.Scheme); err != nil {
		return nil, err
	}
	return hpa, nil
}

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	coreoamdevv1alpha2 "hpatrait/api/v1alpha2"
)

// HorizontalPodAutoscalerTraitReconciler reconciles a HorizontalPodAutoscalerTrait object
type HorizontalPodAutoscalerTraitReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.oam.dev.core.oam.dev,resources=horizontalpodautoscalertraits,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.oam.dev.core.oam.dev,resources=horizontalpodautoscalertraits/status,verbs=get;update;patch

func (r *HorizontalPodAutoscalerTraitReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("horizontalpodautoscalertrait", req.NamespacedName)

	log.Info("Reconcile HorizontalPodAutoscalerTrait")

	var hpatrait coreoamdevv1alpha2.HorizontalPodAutoscalerTrait
	if err := r.Get(ctx, req.NamespacedName, &hpatrait); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info("Get the HPA trait", "maxReplicas: ", hpatrait.Spec.MaxReplicas)

	// Fetch the workload this traiit is referring to
	// TODO: assume the target workload is a fixed deployment

	// TODO create hpa obj
	hpa, err := renderHPA(ctx, hpatrait)


	// TODO patch hpa obj

	return ctrl.Result{}, nil
}

func (r *HorizontalPodAutoscalerTraitReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&coreoamdevv1alpha2.HorizontalPodAutoscalerTrait{}).
		Complete(r)
}

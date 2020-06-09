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

	cpv1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/oam-controllers/pkg/oam/util"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	coreoamdevv1alpha2 "hpatrait/api/v1alpha2"
)

// Reconcile error strings.
const (
	errLocateWorkload    = "cannot find workload"
	errLocateResources   = "cannot find resources"
	errLocateStatefulSet = "cannot find statefulset"
	errRenderService     = "cannot render service"
	errApplyHPA          = "cannot apply the HPA"
	errGCHPA             = "cannot clean up HPA"
)

// HorizontalPodAutoscalerTraitReconciler reconciles a HorizontalPodAutoscalerTrait object
type HorizontalPodAutoscalerTraitReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.oam.dev.core.oam.dev,resources=horizontalpodautoscalertraits,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.oam.dev.core.oam.dev,resources=horizontalpodautoscalertraits/status,verbs=get;update;patch

// +kubebuilder:rbac:groups=core.oam.dev,resources=workloaddefinitions,verbs=get;list;watch

// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete

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
	hpa, err := r.renderHPA(ctx, &hpatrait)
	if err != nil {
		return ctrl.Result{}, err
	}

	// TODO patch hpa obj
	// server side apply the service, only the fields we set are touched
	applyOpts := []client.PatchOption{client.ForceOwnership, client.FieldOwner(hpa.Name)}
	if err := r.Patch(ctx, hpa, client.Apply, applyOpts...); err != nil {
		r.Log.Error(err, "Failed to apply a hpa")
		return util.ReconcileWaitResult,
			util.PatchCondition(ctx, r, &hpatrait, cpv1alpha1.ReconcileError(errors.Wrap(err, errApplyHPA)))
	}
	r.Log.Info("Successfully applied a HPA", "UID", hpa.UID)

	// TODO garbage collect

	// TODO record the new hpa
	hpatrait.Status.DesiredReplicas = hpatrait.Spec.MaxReplicas
	hpatrait.Status.Resources = nil
	hpatrait.Status.Resources = append(hpatrait.Status.Resources, cpv1alpha1.TypedReference{
		APIVersion: hpa.GetObjectKind().GroupVersionKind().GroupVersion().String(),
		Kind:       hpa.GetObjectKind().GroupVersionKind().Kind,
		Name:       hpa.GetName(),
		UID:        hpa.GetUID(),
	})
	if err := r.Status().Update(ctx, &hpatrait); err != nil {
		r.Log.Info("Failed update HPA_trait status", "err:", err)
		return util.ReconcileWaitResult, err
	}

	r.Log.Info("Successfully update HPA_trait status", "UID", hpatrait.GetUID())

	return ctrl.Result{}, util.PatchCondition(ctx, r, &hpatrait, cpv1alpha1.ReconcileSuccess())
}

func (r *HorizontalPodAutoscalerTraitReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&coreoamdevv1alpha2.HorizontalPodAutoscalerTrait{}).
		Complete(r)
}

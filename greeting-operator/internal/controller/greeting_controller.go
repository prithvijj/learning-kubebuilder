/*
Copyright 2026.

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

package controller

import (
	"context"
	"fmt"

	appsv1 "greeting-operator/api/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
)

// GreetingReconciler reconciles a Greeting object
type GreetingReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps.kubebuilder-lessons.dev,resources=greetings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.kubebuilder-lessons.dev,resources=greetings/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps.kubebuilder-lessons.dev,resources=greetings/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Greeting object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.24.1/pkg/reconcile
func (r *GreetingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var greeting appsv1.Greeting
	if err := r.Get(ctx, req.NamespacedName, &greeting); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil // object gone - nothing to do
		}
		return ctrl.Result{}, err
	}

	if greeting.Spec.Message == "" {
		meta.SetStatusCondition(&greeting.Status.Conditions, metav1.Condition{
			Type:    "Ready",
			Status:  metav1.ConditionFalse,
			Reason:  "EmptyMessage",
			Message: "spec.message must not be empty",
		})
		if err := r.Status().Update(ctx, &greeting); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
	}
	op, err := controllerutil.CreateOrUpdate(ctx, r.Client, cm, func() error {
		if cm.Data == nil {
			cm.Data = map[string]string{}
		}
		cm.Data["message"] = greeting.Spec.Message
		return controllerutil.SetControllerReference(&greeting, cm, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	meta.SetStatusCondition(&greeting.Status.Conditions, metav1.Condition{
		Type:    "Ready",
		Status:  metav1.ConditionTrue,
		Reason:  "ConfigMapSynced",
		Message: fmt.Sprintf("configmap %s is in sync", cm.Name),
	})
	if err := r.Status().Update(ctx, &greeting); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("reconciled greeting", "configmap-op", op)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GreetingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Greeting{}).
		Owns(&corev1.ConfigMap{}).
		Named("greeting").
		Complete(r)
}

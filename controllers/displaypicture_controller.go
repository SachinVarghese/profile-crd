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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	profilev1 "github.com/SachinVarghese/picture/api/v1"
)

// DisplayPictureReconciler reconciles a DisplayPicture object
type DisplayPictureReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=profile.sachinvarghese.github.io,resources=displaypictures,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=profile.sachinvarghese.github.io,resources=displaypictures/status,verbs=get;update;patch

func (r *DisplayPictureReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("displaypicture", req.NamespacedName)

	// your logic here
	log.Info("Reconciling", "displaypicture", req.Name)
	var dp profilev1.DisplayPicture
	if err := r.Get(ctx, req.NamespacedName, &dp); err != nil {
		// log.Error(err, "unable to fetch DisplayPicture")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		dpConfigmap := &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name + "-cm",
				Namespace: req.Namespace,
			},
		}
		log.Info("Deleteing cm for " + req.Name)
		err := r.Delete(ctx, dpConfigmap)
		if err != nil {
			log.Error(err, "unable to delete dp configmap")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	dpConfigmap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dp.ObjectMeta.Name + "-cm",
			Namespace: dp.ObjectMeta.Namespace,
		},
		Data: map[string]string{
			"imageSrc": dp.Spec.Src,
		},
	}
	log.Info("Creating cm for " + dp.ObjectMeta.Name)
	err := r.Create(ctx, dpConfigmap)
	if err != nil {
		log.Error(err, "unable to create dp configmap")
	}
	return ctrl.Result{}, nil
}

func (r *DisplayPictureReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&profilev1.DisplayPicture{}).
		Complete(r)
}

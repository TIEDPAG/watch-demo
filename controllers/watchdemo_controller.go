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
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/go-logr/logr"
	apiv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	demov1 "github.com/TIEDPAG/watch-demo/api/v1"
)

// WatchDemoReconciler reconciles a WatchDemo object
type WatchDemoReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=demo.zego.im,resources=watchdemoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=demo.zego.im,resources=watchdemoes/status,verbs=get;update;patch

func (r *WatchDemoReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("watchdemo", req.NamespacedName)

	// your logic here
	_log := &demov1.WatchDemo{}
	err := r.Get(ctx, req.NamespacedName, _log)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("watch log 没找到，忽略事件")
			return ctrl.Result{}, nil
		}
	}
	log.Info("集群发生事件，req:" + req.String())

	return ctrl.Result{}, nil
}

func (r *WatchDemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.WatchDemo{}).
		Owns(&apiv1.Deployment{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(event event.CreateEvent) bool {
				_log := r.Log.WithValues("create event")
				_log.Info("触发了事件")
				return true
			},
			DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
				_log := r.Log.WithValues("delete event")
				_log.Info("触发了事件")
				return true
			},
			UpdateFunc: func(updateEvent event.UpdateEvent) bool {
				_log := r.Log.WithValues("update event")
				_log.Info("触发了事件")
				return true
			},
			GenericFunc: func(genericEvent event.GenericEvent) bool {
				_log := r.Log.WithValues("generic event")
				_log.Info("触发了事件")
				return true
			},
		}).
		Complete(r)
}

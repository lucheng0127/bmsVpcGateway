package controller

import (
	"context"
	"fmt"

	"github.com/lucheng0127/bmsVpcGateway/pkg/apis/network/v1alpha1"
	"github.com/lucheng0127/bmsVpcGateway/pkg/client/clientset/versioned"
	"github.com/lucheng0127/bmsVpcGateway/pkg/client/informers/externalversions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type Controller struct {
	k8sClient   kubernetes.Interface
	crdClient   versioned.Interface
	crdSynced   cache.InformerSynced
	crdInformer externalversions.SharedInformerFactory
}

func NewController(
	k8sClient kubernetes.Interface,
	crdClient versioned.Interface,
	crdInformer externalversions.SharedInformerFactory,
) *Controller {
	controller := new(Controller)
	controller.k8sClient = k8sClient
	controller.crdClient = crdClient
	controller.crdSynced = crdInformer.Network().V1alpha1().VpcConnections().Informer().HasSynced
	controller.crdInformer = crdInformer

	// TODO(shawn): Add informer event handlers
	crdInformer.Network().V1alpha1().VpcConnections().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c := obj.(*v1alpha1.VpcConnection)
			klog.Infof("new vpc-conn %+v\n", c)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oc := oldObj.(*v1alpha1.VpcConnection)
			nc := newObj.(*v1alpha1.VpcConnection)
			klog.Infof("update vpc-conn %+v to %+v\n", oc, nc)
		},
		DeleteFunc: func(obj interface{}) {
			c := obj.(*v1alpha1.VpcConnection)
			klog.Infof("delete vpc-conn %s\n", c.Name)
		},
	})

	return controller
}

func (c *Controller) Run(ctx context.Context) error {
	klog.Info("Start controller")

	klog.Info("Waiting for inforer caches to sync")

	c.crdInformer.Start(ctx.Done())
	if ok := cache.WaitForCacheSync(ctx.Done(), c.crdSynced); !ok {
		return fmt.Errorf("failed to wait for cahces to sync")
	}

	klog.Info("Sync caches finished, start to work...")

	<-ctx.Done()
	klog.Info("Shutting down")
	return nil
}

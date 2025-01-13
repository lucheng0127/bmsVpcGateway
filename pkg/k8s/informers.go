package k8s

import (
	"time"

	"github.com/lucheng0127/bmsVpcGateway/pkg/client/clientset/versioned"
	"github.com/lucheng0127/bmsVpcGateway/pkg/client/informers/externalversions"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

const (
	defualtResync = 600 * time.Second
)

type InformerFactory interface {
	Kubernetes() informers.SharedInformerFactory
	CrdInformer() externalversions.SharedInformerFactory
}

type informerFactory struct {
	k8s informers.SharedInformerFactory
	crd externalversions.SharedInformerFactory
}

func NewInformerFactory(client kubernetes.Interface, crdClient versioned.Interface) InformerFactory {
	return &informerFactory{
		k8s: informers.NewSharedInformerFactory(client, defualtResync),
		crd: externalversions.NewSharedInformerFactory(crdClient, defualtResync),
	}
}

func (i *informerFactory) Kubernetes() informers.SharedInformerFactory {
	return i.k8s
}

func (i *informerFactory) CrdInformer() externalversions.SharedInformerFactory {
	return i.crd
}

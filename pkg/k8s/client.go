package k8s

import (
	"github.com/lucheng0127/bmsVpcGateway/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Interface interface {
	Kubernetes() *kubernetes.Clientset
	CrdClient() *versioned.Clientset
}

type Client struct {
	k8s *kubernetes.Clientset
	crd *versioned.Clientset
}

func NewForConfig(c *rest.Config) (Interface, error) {
	client := new(Client)

	kc, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	client.k8s = kc

	cc, err := versioned.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	client.crd = cc

	return client, nil
}

func (c *Client) Kubernetes() *kubernetes.Clientset {
	return c.k8s
}

func (c *Client) CrdClient() *versioned.Clientset {
	return c.crd
}

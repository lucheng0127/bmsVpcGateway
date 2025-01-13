package agent

import (
	"github.com/lucheng0127/bmsVpcGateway/pkg/controller"
	"github.com/lucheng0127/bmsVpcGateway/pkg/k8s"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Agent struct {
	Client          k8s.Interface
	InformerFactory k8s.InformerFactory
	Controller      controller.Controller
}

func NewAgent(kubeconfig string) (*Agent, error) {
	agent := new(Agent)

	var config *rest.Config
	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}

		config = cfg
	} else {
		cfg, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		config = cfg
	}

	client, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	agent.Client = client

	agent.InformerFactory = k8s.NewInformerFactory(agent.Client.Kubernetes(), agent.Client.CrdClient())

	agent.Controller = *controller.NewController(agent.Client.Kubernetes(), agent.Client.CrdClient(), agent.InformerFactory.CrdInformer())

	return agent, nil
}

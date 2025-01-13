package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PhaseIpset       = "ipset"
	PhaseIptables    = "iptables"
	PhaseIpRule      = "ipRule"
	PhasePolicyRoute = "policyRoute"
	PhaseDone        = "done"
)

type VpcConnectionSpec struct {
	VpcNat       string `json:"vpcNat"`
	VpcNatEip    string `json:"vpcNatEip"`
	VpcNatSubnet string `json:"vpcNatSubnet"`
}

type VpcConnectionStatus struct {
	Phase string `json:"phase"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster,shortName=vpc-conn
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"

type VpcConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcConnectionSpec   `json:"spec"`
	Status VpcConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VpcConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VpcConnection `json:"items"`
}

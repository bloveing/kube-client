package model

import (
	commonv1 "github.com/elastic/cloud-on-k8s/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Kibana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KibanaSpec   `json:"spec,omitempty"`
	Status KibanaStatus `json:"status,omitempty"`
	// assocConf holds the configuration for the Elasticsearch association
	assocConf *commonv1.AssociationConf `json:"-"`
	// entAssocConf holds the configuration for the Enterprise Search association
	entAssocConf *commonv1.AssociationConf `json:"-"`
	// monitoringAssocConf holds the configuration for the monitoring Elasticsearch clusters association
	monitoringAssocConfs map[types.NamespacedName]commonv1.AssociationConf `json:"-"`
}

// KibanaSpec holds the specification of a Kibana instance.
type KibanaSpec struct {
	// Version of Kibana.
	Version string `json:"version"`

	// Image is the Kibana Docker image to deploy.
	Image string `json:"image,omitempty"`

	// Count of Kibana instances to deploy.
	Count int32 `json:"count,omitempty"`

	// ElasticsearchRef is a reference to an Elasticsearch cluster running in the same Kubernetes cluster.
	ElasticsearchRef commonv1.ObjectSelector `json:"elasticsearchRef,omitempty"`

	// EnterpriseSearchRef is a reference to an EnterpriseSearch running in the same Kubernetes cluster.
	// Kibana provides the default Enterprise Search UI starting version 7.14.
	EnterpriseSearchRef commonv1.ObjectSelector `json:"enterpriseSearchRef,omitempty"`

	// Config holds the Kibana configuration. See: https://www.elastic.co/guide/en/kibana/current/settings.html
	// +kubebuilder:pruning:PreserveUnknownFields
	Config *commonv1.Config `json:"config,omitempty"`

	// HTTP holds the HTTP layer configuration for Kibana.
	HTTP commonv1.HTTPConfig `json:"http,omitempty"`

	// PodTemplate provides customisation options (labels, annotations, affinity rules, resource requests, and so on) for the Kibana pods
	// +kubebuilder:validation:Optional
	// +kubebuilder:pruning:PreserveUnknownFields
	PodTemplate corev1.PodTemplateSpec `json:"podTemplate,omitempty"`

	// SecureSettings is a list of references to Kubernetes secrets containing sensitive configuration options for Kibana.
	SecureSettings []commonv1.SecretSource `json:"secureSettings,omitempty"`

	// ServiceAccountName is used to check access from the current resource to a resource (eg. Elasticsearch) in a different namespace.
	// Can only be used if ECK is enforcing RBAC on references.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Monitoring enables you to collect and ship log and monitoring data of this Kibana.
	// See https://www.elastic.co/guide/en/kibana/current/xpack-monitoring.html.
	// Metricbeat and Filebeat are deployed in the same Pod as sidecars and each one sends data to one or two different
	// Elasticsearch monitoring clusters running in the same Kubernetes cluster.
	// +kubebuilder:validation:Optional
	Monitoring Monitoring `json:"monitoring,omitempty"`
}

// KibanaStatus defines the observed state of Kibana
type KibanaStatus struct {
	commonv1.DeploymentStatus `json:",inline"`
	// AssociationStatus is the status of any auto-linking to Elasticsearch clusters.
	// This field is deprecated and will be removed in a future release. Use ElasticsearchAssociationStatus instead.
	AssociationStatus commonv1.AssociationStatus `json:"associationStatus,omitempty"`
	// ElasticsearchAssociationStatus is the status of any auto-linking to Elasticsearch clusters.
	ElasticsearchAssociationStatus commonv1.AssociationStatus `json:"elasticsearchAssociationStatus,omitempty"`
	// EnterpriseSearchAssociationStatus is the status of any auto-linking to Enterprise Search.
	EnterpriseSearchAssociationStatus commonv1.AssociationStatus `json:"enterpriseSearchAssociationStatus,omitempty"`
	// MonitoringAssociationStatus is the status of any auto-linking to monitoring Elasticsearch clusters.
	MonitoringAssociationStatus commonv1.AssociationStatusMap `json:"monitoringAssociationStatus,omitempty"`
}

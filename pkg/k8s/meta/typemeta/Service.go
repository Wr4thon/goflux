package typemeta

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Version defines the versioned schema of this representation of an object.
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
type Version string

const (
	// VersionAppsV1 represents the Apps/V1 Version
	VersionAppsV1 Version = "apps/v1"
	// VersionV1 represents the V1 Version
	VersionV1 Version = "v1"
	// NetworkingV1Beta1 represents the networking.k8s.io/v1beta1 Version
	NetworkingV1Beta1 Version = "networking.k8s.io/v1beta1"
)

// Kind is a string value representing the REST resource this object represents.
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
type Kind string

const (
	// KindClusterIssuer represents the ClusterIssuer Kind
	KindClusterIssuer Kind = "ClusterIssuer"
	// KindClusterRole represents the ClusterRole Kind
	KindClusterRole Kind = "ClusterRole"
	// KindClusterRoleBinding represents the ClusterRoleBinding Kind
	KindClusterRoleBinding Kind = "ClusterRoleBinding"
	// KindConfigMap represents the ConfigMap Kind
	KindConfigMap Kind = "ConfigMap"
	// KindCronJob represents the CronJob Kind
	KindCronJob Kind = "CronJob"
	// KindCustomResourceDefinition represents the CustomResourceDefinition Kind
	KindCustomResourceDefinition Kind = "CustomResourceDefinition"
	// KindDeployment represents the Deployment Kind
	KindDeployment Kind = "Deployment"
	// KindHorizontalPodAutoscaler represents the HorizontalPodAutoscaler Kind
	KindHorizontalPodAutoscaler Kind = "HorizontalPodAutoscaler"
	// KindIngress represents the Ingress Kind
	KindIngress Kind = "Ingress"
	// KindKustomization represents the Kustomization Kind
	KindKustomization Kind = "Kustomization"
	// KindNamespace represents the Namespace Kind
	KindNamespace Kind = "Namespace"
	// KindPersistentVolume represents the PersistentVolume Kind
	KindPersistentVolume Kind = "PersistentVolume"
	// KindPersistentVolumeClaim represents the PersistentVolumeClaim Kind
	KindPersistentVolumeClaim Kind = "PersistentVolumeClaim"
	// KindPod represents the Pod Kind
	KindPod Kind = "Pod"
	// KindRole represents the Role Kind
	KindRole Kind = "Role"
	// KindRoleBinding represents the RoleBinding Kind
	KindRoleBinding Kind = "RoleBinding"
	// KindSealedSecret represents the SealedSecret Kind
	KindSealedSecret Kind = "SealedSecret"
	// KindSecret represents the Secret Kind
	KindSecret Kind = "Secret"
	// KindService represents the Service Kind
	KindService Kind = "Service"
	// KindServiceAccount represents the ServiceAccount Kind
	KindServiceAccount Kind = "ServiceAccount"
)

// Service provides functionality to create TypeMeta data
type Service interface {
	New(version Version, kind Kind) (metav1.TypeMeta, error)
}

type service struct {
}

// New creates a new instance of the Service
func New() Service {
	return &service{}
}

func (service *service) New(version Version, kind Kind) (metav1.TypeMeta, error) {
	return metav1.TypeMeta{
		APIVersion: string(version),
		Kind:       string(kind),
	}, nil
}

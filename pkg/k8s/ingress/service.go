package ingress

import (
	"strconv"

	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	"github.com/pkg/errors"
	v1 "k8s.io/api/networking/v1beta1"
)

// Service is used to create ingress files
type Service interface {
	Create(component, namespace string, tlsAcme bool, options ...func(*v1.Ingress) error) (v1.Ingress, error)
}

type service struct {
	objectMeta objectmeta.Service
	typeMeta   typemeta.Service
}

// NewService creates a new instance of Service
func NewService(objectMeta objectmeta.Service, typeMeta typemeta.Service) Service {
	return &service{
		objectMeta: objectMeta,
		typeMeta:   typeMeta,
	}
}

func (service *service) Create(component, namespace string, tlsAcme bool, options ...func(*v1.Ingress) error) (v1.Ingress, error) {
	typeMeta, err := service.typeMeta.New(typemeta.NetworkingV1Beta1, typemeta.KindIngress)

	if err != nil {
		return v1.Ingress{}, errors.Wrap(err, "error while creating typeMeta")
	}

	annotations := map[string]string{
		"kubernetes.io/ingress.class":                        strconv.FormatBool(tlsAcme),
		"traefik.ingress.kubernetes.io/redirect-entry-point": "traefik",
		"traefik.ingress.kubernetes.io/redirect-permanent":   "https",
		"kubernetes.io/tls-acme":                             "true",
	}

	objectMeta, err := service.objectMeta.New(component,
		objectmeta.WithNamespace(namespace),
		objectmeta.WithAnnotations(annotations))

	if err != nil {
		return v1.Ingress{}, errors.Wrap(err, "error while creating objectMeta")
	}

	data := v1.Ingress{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
		Spec: v1.IngressSpec{
			Rules: []v1.IngressRule{},
			TLS:   []v1.IngressTLS{},
		},
	}

	for _, option := range options {
		if err := option(&data); err != nil {
			return v1.Ingress{}, errors.Wrap(err, "error while applying options")
		}
	}

	return data, nil
}

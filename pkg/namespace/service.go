package namespace

import (
	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
)

// Service provides functionality to create namespace files
type Service interface {
	Create(name string) (v1.Namespace, error)
	setAnnotations(annotations map[string]string) error
}

type service struct {
	typeMeta    typemeta.Service
	objectMeta  objectmeta.Service
	annotations map[string]string
}

// New creates a new instance of service
func New(typeMeta typemeta.Service, objectMeta objectmeta.Service, options ...func(Service) error) (Service, error) {
	service := &service{
		typeMeta:   typeMeta,
		objectMeta: objectMeta,
		annotations: map[string]string{
			"linkerd.io/inject": "enabled",
		},
	}
	for _, option := range options {
		err := option(service)
		if err != nil {
			return nil, errors.Wrap(err, "error while applying option")
		}
	}

	return service, nil
}

// WithAnnotations can be used to set Annotations in the namespace
func WithAnnotations(annotations map[string]string) func(Service) error {
	return func(s Service) error {
		return s.setAnnotations(annotations)
	}
}

func (service *service) Create(name string) (v1.Namespace, error) {
	typeMeta, err := service.typeMeta.New(typemeta.VersionV1, typemeta.KindNamespace)

	if err != nil {
		return v1.Namespace{}, errors.Wrap(err, "error while creating typeMeta")
	}

	objectMeta, err := service.objectMeta.New(name,
		objectmeta.WithAnnotations(service.getAnnotations()))

	if err != nil {
		return v1.Namespace{}, errors.Wrap(err, "error while creating objectMeta")
	}

	data := v1.Namespace{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
	}

	return data, nil
}

func (service *service) getAnnotations() map[string]string {
	return mapCopy(service.annotations)
}

func (service *service) setAnnotations(annotations map[string]string) error {
	service.annotations = mapCopy(annotations)
	return nil
}

func mapCopy(in map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range in {
		result[k] = v
	}

	return result
}

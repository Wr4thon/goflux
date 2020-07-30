package service

import (
	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
)

// Service is a service service, lol
type Service interface {
	Create(name, namespace string, option ...func(*v1.Service) error) (v1.Service, error)
}

type service struct {
	typeMeta   typemeta.Service
	objectMeta objectmeta.Service
}

// New creates a new instance of Service
func New(typeMeta typemeta.Service, objectMeta objectmeta.Service) Service {
	return &service{
		typeMeta:   typeMeta,
		objectMeta: objectMeta,
	}
}

func (service *service) Create(name, namespace string, options ...func(*v1.Service) error) (v1.Service, error) {
	typeMeta, err := service.typeMeta.New(typemeta.VersionV1, typemeta.KindNamespace)

	if err != nil {
		return v1.Service{}, errors.Wrap(err, "error while creating typeMeta")
	}

	// TODO consts
	labels := map[string]string{
		"app":       namespace,
		"component": name,
	}

	objectMeta, err := service.objectMeta.New(name,
		objectmeta.WithNamespace(namespace),
		objectmeta.WithLabels(labels))

	if err != nil {
		return v1.Service{}, errors.Wrap(err, "error while creating objectMeta")
	}

	data := v1.Service{
		TypeMeta:   typeMeta,
		ObjectMeta: objectMeta,
		Spec: v1.ServiceSpec{
			Selector: labels,
		},
	}

	for _, option := range options {
		if err := option(&data); err != nil {
			return v1.Service{}, errors.Wrap(err, "error while applying options to service")
		}
	}

	return data, nil
}

package namespace

import (
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Service provides functionality to create namespace files
type Service interface {
	Create(name string) (v1.Namespace, error)
	setAnnotations(annotations map[string]string) error
}

type service struct {
	annotations map[string]string
}

// NewService creates a new instance of service
func NewService(options ...func(Service) error) (Service, error) {
	service := &service{
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
	data := v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Namespace",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Annotations: service.getAnnotations(),
		},
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

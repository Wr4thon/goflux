package objectmeta

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Service provides functionality to create TypeMeta data
type Service interface {
	New(name string, options ...func(*metav1.ObjectMeta) error) (metav1.ObjectMeta, error)
}

type service struct {
}

// New creates a new instance of the Service
func New() Service {
	return &service{}
}

func (service *service) New(name string, options ...func(*metav1.ObjectMeta) error) (metav1.ObjectMeta, error) {
	objMeta := metav1.ObjectMeta{
		Name: name,
	}

	for _, option := range options {
		if err := option(&objMeta); err != nil {
			return metav1.ObjectMeta{}, errors.New("error while applying options")
		}
	}

	return objMeta, nil
}

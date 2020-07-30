package namespace_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	"github.com/Nerzal/goflux/pkg/k8s/namespace"
)

const (
	namespaceName   string = "clarilab"
	annotationKey   string = "foo"
	annotationValue string = "bar"
)

func new(opts ...func(namespace.Service) error) (namespace.Service, error) {
	return namespace.New(typemeta.New(), objectmeta.New(), opts...)
}

func TestNamespace_Create(t *testing.T) {
	service, err := new()
	if err != nil {
		t.Error(err)
	}

	namespaceData, err := service.Create(namespaceName)

	if err != nil {
		t.Error(err)
	}

	if namespaceData.Name != namespaceName {
		t.Errorf("Name of namespace is incorrect. Actual: '%s', Expected: '%s'",
			namespaceData.Namespace, namespaceName)
	}
}

func TestNamespace_CreateWithAnnotations(t *testing.T) {
	expected := map[string]string{
		annotationKey: annotationValue,
	}

	service, err := new(namespace.WithAnnotations(expected))
	if err != nil {
		t.Error(err)
	}

	namespaceData, err := service.Create(namespaceName)

	if err != nil {
		t.Error(err)
	}

	if len(namespaceData.Annotations) != 1 ||
		namespaceData.Annotations[annotationKey] != annotationValue {

		t.Errorf("Annotations of namespace are incorrect. Actual: '%v', Expected: '%v'",
			namespaceData.Annotations,
			expected)
	}
}

func TestNamespace_CreateWithAnnotations_MutateAfterSetting(t *testing.T) {
	annotations := map[string]string{
		annotationKey: annotationValue,
	}

	service, err := new(namespace.WithAnnotations(annotations))
	if err != nil {
		t.Error(err)
	}

	annotations[annotationKey] = "foobar"

	namespaceData, err := service.Create(namespaceName)

	if err != nil {
		t.Error(err)
	}

	if len(namespaceData.Annotations) != 1 ||
		namespaceData.Annotations[annotationKey] != annotationValue {

		t.Errorf("Annotations of namespace are incorrect. Actual: '%v', Expected: '%v'",
			namespaceData.Annotations,
			map[string]string{
				annotationKey: annotationValue,
			})
	}
}

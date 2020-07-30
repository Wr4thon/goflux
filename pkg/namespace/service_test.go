package namespace_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/namespace"
)

const (
	namespaceName   string = "clarilab"
	annotationKey   string = "foo"
	annotationValue string = "bar"
)

func TestNamespace_Create(t *testing.T) {
	service, err := namespace.NewService()
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

	service, err := namespace.NewService(namespace.WithAnnotations(expected))
	if err != nil {
		t.Error(err)
	}

	namespaceData, err := service.Create(namespaceName)

	if err != nil {
		t.Error(err)
	}

	if len(namespaceData.Annotations) != 1 ||
		namespaceData.Annotations[annotationKey] != annotationValue {

		t.Errorf("Annotations of namespace are incorrect. Actual: '%s', Expected: '%s'",
			namespaceData.Annotations,
			expected)
	}
}

func TestNamespace_CreateWithAnnotations_MutateAfterSetting(t *testing.T) {
	annotations := map[string]string{
		annotationKey: annotationValue,
	}

	service, err := namespace.NewService(namespace.WithAnnotations(annotations))
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

		t.Errorf("Annotations of namespace are incorrect. Actual: '%s', Expected: '%s'",
			namespaceData.Annotations,
			map[string]string{
				annotationKey: annotationValue,
			})
	}
}

package namespace_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/namespace"
)

func TestNamespace_Create(t *testing.T) {
	service := namespace.NewService()
	namespaceName := "clarilab"
	namespaceData, err := service.Create(namespaceName)

	if err != nil {
		t.Error(err)
	}

	if namespaceData.Name != namespaceName {
		t.Errorf("Name of namespace is incorrect. Actual: '%s', Expected: '%s'", namespaceData.Namespace, namespaceName)
	}
}

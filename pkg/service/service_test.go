package service_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	"github.com/Nerzal/goflux/pkg/service"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	name      string = "kycnow-api"
	namespace string = "clarilab"
)

func new() service.Service {
	return service.New(typemeta.New(), objectmeta.New())
}

func TestService_Create(t *testing.T) {
	service := new()

	_, err := service.Create(name, namespace)
	if err != nil {
		t.Error(err)
	}
}

func TestService_Create_WithPort(t *testing.T) {
	serviceService := new()

	var port int32 = 80
	var targetPort int32 = 8080
	protocol := v1.ProtocolTCP

	foo, err := serviceService.Create(name, namespace, service.WithPort(port, targetPort, protocol))
	if err != nil {
		t.Error(err)
	}

	if len(foo.Spec.Ports) != 1 {
		t.FailNow()
	}

	specPort := foo.Spec.Ports[0]

	if specPort.Port != port ||
		specPort.TargetPort.IntVal != targetPort ||
		specPort.Protocol != protocol {

		t.Errorf("one or more portvalues was not applied. actual: '%v', expected: '%v'",
			specPort,
			v1.ServicePort{
				Protocol: protocol,
				Port:     port,
				TargetPort: intstr.IntOrString{
					IntVal: targetPort,
				},
			})
	}
}

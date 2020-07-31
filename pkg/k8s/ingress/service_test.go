package ingress_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/k8s/ingress"
	"github.com/Nerzal/goflux/pkg/k8s/meta/objectmeta"
	"github.com/Nerzal/goflux/pkg/k8s/meta/typemeta"
	v1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	name       = "kycnow-api"
	namespace  = "clarilab"
	host       = "app.foo.kycnow.de"
	tlsHost    = "*.foo.kycnow.de"
	secretName = "my-cert-de-crt"
	endpoint   = "/api"
	port       = 80
)

func new() ingress.Service {
	return ingress.NewService(objectmeta.New(), typemeta.New())
}

func TestIngress_Create(t *testing.T) {
	service := new()

	ingress, err := service.Create(name, namespace, true)
	if err != nil {
		t.Error(err)
	}

	if ingress.Namespace != namespace {
		t.Errorf("ingress namespace did not get set properly. Actual: '%s', Expected: '%s'", ingress.Namespace, namespace)
	}

	if ingress.Name != name {
		t.Errorf("ingress name did not get set properly. Actual: '%s', Expected: '%s'", ingress.Name, name)
	}
}

func TestIngress_Create_WithRule(t *testing.T) {
	service := new()

	ingress, err := service.Create(name,
		namespace,
		true,
		ingress.WithRule(host,
			ingress.WithBackend(endpoint, name, port)))
	if err != nil {
		t.Error(err)
	}

	if len(ingress.Spec.Rules) != 1 ||
		ingress.Spec.Rules[0].Host != host ||
		len(ingress.Spec.Rules[0].HTTP.Paths) != 1 ||
		ingress.Spec.Rules[0].HTTP.Paths[0].Path != endpoint ||
		ingress.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntVal != port ||
		ingress.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName != name {

		t.Errorf("ingress rule was set incorrectly. \nActual:   '%v'\nExpected: '%v'", ingress.Spec.Rules, []v1.IngressRule{
			{
				Host: host,
				IngressRuleValue: v1.IngressRuleValue{
					HTTP: &v1.HTTPIngressRuleValue{
						Paths: []v1.HTTPIngressPath{
							{
								Path: endpoint,
								Backend: v1.IngressBackend{
									ServiceName: name,
									ServicePort: intstr.FromInt(port),
								},
							},
						},
					},
				},
			},
		})
	}
}

func TestIngress_Create_WithTLS(t *testing.T) {
	service := new()

	ingress, err := service.Create(name,
		namespace,
		true,
		ingress.WithTLS(secretName, tlsHost))
	if err != nil {
		t.Error(err)
	}

	if len(ingress.Spec.TLS) != 1 ||
		ingress.Spec.TLS[0].SecretName != secretName ||
		len(ingress.Spec.TLS[0].Hosts) != 1 ||
		ingress.Spec.TLS[0].Hosts[0] != tlsHost {

		t.Errorf("ingress tls was set incorrectly. \nActual:   '%v'\nExpected: '%v'", ingress.Spec.TLS, []v1.IngressTLS{
			{
				SecretName: secretName,
				Hosts:      []string{tlsHost},
			},
		})
	}
}

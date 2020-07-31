package ingress

import (
	"github.com/pkg/errors"
	v1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// WithTLS can be used to add a tls configuration
func WithTLS(secretName, tlsHost string) func(*v1.Ingress) error {
	return func(ingress *v1.Ingress) error {
		tls := v1.IngressTLS{
			Hosts:      []string{tlsHost},
			SecretName: secretName,
		}

		ingress.Spec.TLS = append(ingress.Spec.TLS, tls)
		return nil
	}
}

// WithRule can be used to append a rule
func WithRule(host string, options ...func(*v1.IngressRule) error) func(*v1.Ingress) error {
	return func(ingress *v1.Ingress) error {
		rule := v1.IngressRule{
			Host: host,
			IngressRuleValue: v1.IngressRuleValue{
				HTTP: &v1.HTTPIngressRuleValue{
					Paths: []v1.HTTPIngressPath{},
				},
			},
		}

		for _, option := range options {
			if err := option(&rule); err != nil {
				return errors.Wrap(err, "error while applying options")
			}
		}

		ingress.Spec.Rules = append(ingress.Spec.Rules, rule)
		return nil
	}
}

// WithBackend can be used to append a new backend
func WithBackend(path, serviceName string, port int) func(*v1.IngressRule) error {
	return func(ruleValue *v1.IngressRule) error {
		ruleValue.IngressRuleValue.HTTP.Paths = append(ruleValue.IngressRuleValue.HTTP.Paths, v1.HTTPIngressPath{
			Path: path,
			Backend: v1.IngressBackend{
				ServiceName: serviceName,
				ServicePort: intstr.FromInt(port),
			},
		})
		return nil
	}
}

package service

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func WithPort(port int32, targetPort int, protocol v1.Protocol) func(*v1.Service) error {
	return WithNamedPort("", port, targetPort, protocol)
}

func WithNamedPort(name string, port int32, targetPort int, protocol v1.Protocol) func(*v1.Service) error {
	return func(s *v1.Service) error {
		port := v1.ServicePort{
			Protocol:   protocol,
			Name:       name,
			Port:       port,
			TargetPort: intstr.FromInt(targetPort),
		}

		s.Spec.Ports = append(s.Spec.Ports, port)
		return nil
	}
}

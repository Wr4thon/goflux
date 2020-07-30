package objectmeta

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func WithLabels(labels map[string]string) func(*metav1.ObjectMeta) error {
	return func(objMeta *metav1.ObjectMeta) error {
		objMeta.Labels = labels
		return nil
	}
}

func WithAnnotations(annotations map[string]string) func(*metav1.ObjectMeta) error {
	return func(objMeta *metav1.ObjectMeta) error {
		objMeta.Annotations = annotations
		return nil
	}
}

func WithNamespace(namespace string) func(*metav1.ObjectMeta) error {
	return func(objMeta *metav1.ObjectMeta) error {
		objMeta.Namespace = namespace
		return nil
	}
}

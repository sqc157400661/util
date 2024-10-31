package util

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetEnvVarFromSecret(sctName, name, key string, opt bool) corev1.EnvVar {
	return corev1.EnvVar{
		Name: name,
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: sctName,
				},
				Key:      key,
				Optional: &opt,
			},
		},
	}
}

func ContainsOnlyFinalizer(obj client.Object, finalizer string) bool {
	return len(obj.GetFinalizers()) == 1 && obj.GetFinalizers()[0] == finalizer
}

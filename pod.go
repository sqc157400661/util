package util

import (
	corev1 "k8s.io/api/core/v1"
)

func GetContainerFromPodSpec(podSpec *corev1.PodSpec, name string) *corev1.Container {
	if podSpec == nil || podSpec.Containers == nil {
		return nil
	}
	for i := range podSpec.Containers {
		c := &podSpec.Containers[i]
		if c.Name == name {
			return c
		}
	}
	return nil
}

func GetContainerFromPod(pod *corev1.Pod, name string) *corev1.Container {
	if pod == nil {
		return nil
	}
	return GetContainerFromPodSpec(&pod.Spec, name)
}

func IsPodRunning(pod *corev1.Pod) bool {
	return pod != nil && pod.Status.Phase == corev1.PodRunning
}

func IsPodReady(pod *corev1.Pod) bool {
	if pod == nil {
		return false
	}

	if pod.Status.Phase != corev1.PodRunning {
		return false
	}

	if pod.Status.ContainerStatuses == nil {
		return false
	}

	for _, containerStatus := range pod.Status.ContainerStatuses {
		if !containerStatus.Ready {
			return false
		}
	}

	return true
}

func IsAllPodsReady(pods []corev1.Pod) bool {
	for _, pod := range pods {
		if !IsPodReady(&pod) {
			return false
		}
	}
	return true
}

func IsPodDeletedOrFailed(po *corev1.Pod) bool {
	return IsPodDeleted(po) || IsPodFailed(po)
}

func IsPodDeleted(po *corev1.Pod) bool {
	return po != nil && !po.DeletionTimestamp.IsZero()
}

func IsPodFailed(po *corev1.Pod) bool {
	return po != nil && po.Status.Phase == corev1.PodFailed
}

func IsPodScheduled(po *corev1.Pod) bool {
	if po == nil {
		return false
	}
	for _, cond := range po.Status.Conditions {
		if cond.Type == corev1.PodScheduled &&
			cond.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

func FilterPodsBy(pods []corev1.Pod, filter func(pod *corev1.Pod) bool) []corev1.Pod {
	result := make([]corev1.Pod, 0)
	for _, pod := range pods {
		if filter(&pod) {
			result = append(result, pod)
		}
	}
	return result
}

func ArePodsAllReady(pods []corev1.Pod) bool {
	for _, pod := range pods {
		if !IsPodReady(&pod) {
			return false
		}
	}
	return true
}

func ArePodsAllScheduled(pods []corev1.Pod) bool {
	for _, pod := range pods {
		if !IsPodScheduled(&pod) {
			return false
		}
	}
	return true
}

func BuildPodMap(pods []corev1.Pod, keyFunc func(pod *corev1.Pod) string) map[string]*corev1.Pod {
	r := make(map[string]*corev1.Pod)
	for _, pod := range pods {
		r[keyFunc(&pod)] = &pod
	}
	return r
}

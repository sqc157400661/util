package util

import (
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
)

const (
	B   float64 = 1
	KiB         = 1024 * B
	MiB         = 1024 * KiB
	GiB         = 1024 * MiB
)

var convertMem = map[string]string{
	"":   "Gi",
	"gi": "Gi",
	"g":  "Gi",
	"gb": "Gi",
	"m":  "Mi",
	"mb": "Mi",
}

var convertCpu = map[string]string{
	"":  "",
	"c": "",
	"m": "m",
}

func GenerateResource(cpu float64, mem float64) (re map[corev1.ResourceName]resource.Quantity) {
	re = make(map[corev1.ResourceName]resource.Quantity)
	re[corev1.ResourceCPU], _ = ParseCPUWithUnit(cpu, "")
	re[corev1.ResourceMemory], _ = ParseMemoryWithUnit(mem, "gi")
	return
}

func GB(m *resource.Quantity) float64 {
	if m == nil {
		return 0
	}
	return float64(m.Value()) / GiB
}

func ParseCPUWithUnit(cpu float64, unit string) (re resource.Quantity, err error) {
	unit, err = getRealUnit(unit, convertCpu)
	if err != nil {
		return
	}
	cpuQuantity := fmt.Sprintf("%.2f%s", cpu, unit)
	return resource.ParseQuantity(cpuQuantity)
}

func ParseMemoryWithUnit(mem float64, unit string) (re resource.Quantity, err error) {
	unit, err = getRealUnit(unit, convertMem)
	if err != nil {
		return
	}
	cpuQuantity := fmt.Sprintf("%.2f%s", mem, unit)
	return resource.ParseQuantity(cpuQuantity)
}

func getRealUnit(u string, convert map[string]string) (unit string, err error) {
	if len(convert) == 0 {
		return u, nil
	}
	u = strings.ToLower(u)
	var has bool
	if unit, has = convert[u]; !has {
		err = errors.New("invalid unit")
		return
	}
	return
}

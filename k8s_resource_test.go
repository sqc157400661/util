package util

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

func TestA(t *testing.T) {
	cpuQuantity := fmt.Sprintf("%.2fGi", float64(0.25))
	re, err := resource.ParseQuantity(cpuQuantity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(GB(&re))
}

package util

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"hash/fnv"
	"k8s.io/apimachinery/pkg/util/rand"
)

func UnsafeMergeMap[K comparable, V any](tgt map[K]V, src map[K]V) map[K]V {
	if len(src) == 0 {
		return tgt
	}
	if tgt == nil {
		tgt = make(map[K]V)
	}
	for k, v := range src {
		if _, has := tgt[k]; !has {
			tgt[k] = v
		}
	}
	return tgt
}

func UnsafeMergeStringMap(tgt map[string]string, src map[string]string) map[string]string {
	if len(src) == 0 {
		return tgt
	}
	if tgt == nil {
		tgt = map[string]string{}
	}
	for k, v := range src {
		if _, has := tgt[k]; !has {
			tgt[k] = v
		}
	}
	return tgt
}

func ConvertToStruct(input map[string]interface{}, dst interface{}) error {
	if len(input) == 0 {
		return nil
		//return errors.New("not input")
	}
	if dst == nil {
		return errors.New("dst struct is nil")
	}
	arr, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(arr, &dst)
	if err != nil {
		return err
	}
	return nil
}

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	arr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	dst := map[string]interface{}{}
	err = json.Unmarshal(arr, &dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func InterfaceMapHash32(msg map[string]interface{}) (string, error) {
	hash := fnv.New32()
	_, err := fmt.Fprint(hash, msg)
	if err != nil {
		return "", err
	}
	return rand.SafeEncodeString(fmt.Sprint(hash.Sum32())), err
}

package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// SizeToBytes parses a string formatted by ByteSize as bytes.
func SizeToBytes(s string) (float64, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	idx := strings.IndexFunc(s, unicode.IsLetter)
	if idx == -1 {
		return strconv.ParseFloat(s, 64)
	}

	nums, err := strconv.ParseFloat(s[:idx], 64)
	if err != nil {
		return 0, err
	}

	switch s[idx:] {
	case "K":
		return nums * KiB, nil
	case "M":
		return nums * MiB, nil
	case "G":
		return nums * GiB, nil
	}
	return 0, fmt.Errorf("'%s' format error, must be a positive integer with a unit of measurement like K, M or G", s)
}

// Min returns the smallest int64 that was passed in the arguments.
func Min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// Max returns the largest int64 that was passed in the arguments.
func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// Bool returns a pointer to v.
func Bool(v bool) *bool { return &v }

// ByteMap initializes m when it points to nil.
func ByteMap(m *map[string][]byte) {
	if m != nil && *m == nil {
		*m = make(map[string][]byte)
	}
}

// Int32 returns a pointer to v.
func Int32(v int32) *int32 { return &v }

// Int64 returns a pointer to v.
func Int64(v int64) *int64 { return &v }

// String returns a pointer to v.
func String(v string) *string { return &v }

// StringMap initializes m when it points to nil.
func StringMap(m *map[string]string) {
	if m != nil && *m == nil {
		*m = make(map[string]string)
	}
}

func MD5Hash(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

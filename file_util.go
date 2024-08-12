package helper

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// reads a file and returns its contents as a byte slice
func SlurpAsBytes(filename string) ([]byte, error) {
	// #nosec G304
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading from file %s", filename)
	}
	return b, nil
}

// Writes a string slice into a file
// The file is created
func WriteStrings(lines []string, filename string, termination string) error {
	file, err := os.Create(filename) // #nosec G304
	if err != nil {
		return errors.Wrapf(err, "error creating file %s", filename)
	}
	defer file.Close() // #nosec G307

	w := bufio.NewWriter(file)
	for _, line := range lines {
		//fmt.Fprintln(w, line+termination)
		N, err := fmt.Fprintf(w, "%s", line+termination)
		if err != nil {
			return errors.Wrapf(err, "error writing to file %s ", filename)
		}
		if N == 0 {
			return nil
		}
	}
	return w.Flush()
}

// append a string into an existing file
func WriteString(line string, filename string) error {
	return WriteStrings([]string{line}, filename, "")
}

// returns true if a given file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// returns modify time of file
func FileModTime(filename string) int64 {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return fi.ModTime().Unix()
}

// returns true if a given directory exists
func DirExists(filename string) bool {
	f, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	filemode := f.Mode()
	return filemode.IsDir()
}

func GlobalTempDir() string {
	globalTmpDir := os.Getenv("TMPDIR")
	if globalTmpDir == "" {
		globalTmpDir = "/tmp"
	}
	return globalTmpDir
}

// Returns the base name of a file
func BaseName(filename string) string {
	return filepath.Base(filename)
}

// Returns the directory name of a file
func DirName(filename string) string {
	return filepath.Dir(filename)
}

// Returns the absolute path of a file
func AbsolutePath(value string) (string, error) {
	reTilde := regexp.MustCompile(`^~`)
	value = reTilde.ReplaceAllString(value, os.Getenv("HOME"))
	filename, err := filepath.Abs(value)
	if err != nil {
		return "", errors.Wrapf(err, "error getting absolute path for %s", value)
	}
	return filename, nil
}

func WriteRegularFile(fileName, text, directory string) (string, error) {
	fullName := path.Join(directory, fileName)
	err := WriteString(text, fullName)
	if err != nil {
		return "", err
	}
	return fullName, nil
}

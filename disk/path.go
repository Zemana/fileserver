package disk

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

/*
 * storageMap indicates the storage mapping for files using files first
 * character. It requires 8 disk to split files using hexadecimal format. Using
 * this method, every two pair will be in the same disk.
 */
var storageMap = map[uint8]string{
	'0': "G:",
	'1': "G:",
	'2': "E:",
	'3': "E:",
	'4': "F:",
	'5': "F:",
	'6': "I:",
	'7': "I:",
	'8': "J:",
	'9': "J:",
	'A': "K:",
	'B': "K:",
	'C': "L:",
	'D': "L:",
	'E': "H:",
	'F': "H:",
}

/*
 * ConvertToStoragePath function is an interface between program and disk. It
 * handles the disk logic inside.
 */
func ConvertToStoragePath(name string) (string, error) {
	if len(name) == 0 {
		return "", fmt.Errorf("Empty string")
	}

	dir, ok := storageMap[strings.ToUpper(name)[0]]
	if !ok {
		return "", fmt.Errorf("Could not convert to storage path %s", name)
	}

	return path.Join(dir, name), nil
}

/*
 * Exists function helps to check if file exists in the disk.
 */
func Exists(name string) (bool, error) {
	f, err := ConvertToStoragePath(name)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(f); err != nil {
		return false, err
	}
    
	return true, nil
}

/*
 * WriteToStorage helps to save files into disks.
 */
func WriteToStorage(input io.Reader, name string, size int64) error {
	f, err := ConvertToStoragePath(name)
	if err != nil {
		return err
	}

	output, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	s, err := io.Copy(output, input)
	if err != nil {
		return err
	}
	output.Close()

	if s != size {
		return io.ErrShortWrite
	}

	return nil
}

/*
 * Rename function helps to rename files in the disks
 */
func Rename(old string, new string) error {
	oldPath, err := ConvertToStoragePath(old)
	if err != nil {
		return err
	}

	newPath, err := ConvertToStoragePath(new)
	if err != nil {
		return err
	}

	return os.Rename(oldPath, newPath)
}

/*
 * GetRandomExt function helps to create random extensions
 */
func GetRandomExt() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return ".writing" + string(b)
}

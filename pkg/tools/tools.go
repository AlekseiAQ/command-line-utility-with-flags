package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func CalculateDirSize(directory string, isRecursive, isHuman bool) int64 {
	var size int64
	var wg sync.WaitGroup

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		} else if isRecursive && path != directory {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				s := CalculateDirSize(p, isRecursive, isHuman)
				fmt.Printf("%s: %s\n", p, FormatSize(s, isHuman))
			}(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", directory, err)
	}

	wg.Wait()

	return size
}

func FormatSize(size int64, isHuman bool) string {
	if !isHuman {
		return fmt.Sprintf("%d", size)
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

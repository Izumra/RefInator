package app

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	configparser "github.com/Izumra/RefInator/utils/config_parser"
	"github.com/brianvoe/gofakeit/v7"
)

type Changes struct {
	Classes map[string]string `yaml:"classes"`
	Funcs   map[string]string `yaml:"funcs"`
	Enums   map[string]string `yaml:"enums"`
	Structs map[string]string `yaml:"structs"`
}

type RefInator struct {
	excExts    map[string]bool
	excFiles   map[string]bool
	excFolders []string
	changes    Changes
}

func New(cfg configparser.Config) *RefInator {
	refInator := &RefInator{
		excExts:    make(map[string]bool),
		excFiles:   make(map[string]bool),
		excFolders: cfg.Exclusions.Folders,
	}

	for _, ext := range cfg.Exclusions.Extensions {
		refInator.excExts[ext] = true
	}

	for _, ext := range cfg.Exclusions.Files {
		refInator.excFiles[ext] = true
	}

	changes := Changes{
		Classes: make(map[string]string),
		Funcs:   make(map[string]string),
		Enums:   make(map[string]string),
		Structs: make(map[string]string),
	}

	for _, class := range cfg.Changes.Classes {
		changes.Classes[class] = gofakeit.Username()
	}

	for _, function := range cfg.Changes.Funcs {
		changes.Funcs[function] = gofakeit.Username()
	}

	for _, enum := range cfg.Changes.Enums {
		changes.Enums[enum] = gofakeit.Username()
	}

	for _, structure := range cfg.Changes.Structs {
		changes.Structs[structure] = gofakeit.Username()
	}

	refInator.changes = changes

	return refInator
}

func (r *RefInator) Refactor(folderPath string) error {
	return filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		for _, folder := range r.excFolders {
			if strings.HasPrefix(path, folder) {
				return nil
			}
		}

		if _, ok := r.excFiles[path]; ok {
			return nil
		}

		afterPath, _ := strings.CutPrefix(path, folderPath)

		if !d.IsDir() {

			ext := filepath.Ext(path)
			if _, ok := r.excExts[ext]; ok {
				return nil
			}

			fileReader, err := os.Open(path)
			if err != nil {
				log.Println(err)
				return nil
			}
			defer fileReader.Close()

			lines := []string{}
			scanner := bufio.NewScanner(fileReader)

			for scanner.Scan() {
				unchanged_line := scanner.Text()
				line := r.changeNamesWorker(unchanged_line) + "\n"

				lines = append(lines, line)
			}
			if err := scanner.Err(); err != nil {
				log.Printf("where is the error: %s,\n cause: %s", path, err)
				return nil
			}
			fileReader.Close()

			fileWriter, err := os.OpenFile(path, os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
				return nil
			}
			defer fileWriter.Close()

			writer := bufio.NewWriter(fileWriter)

			for i := range lines {
				if _, err := writer.WriteString(lines[i]); err != nil {
					log.Println(err)
					return nil
				}
			}
			writer.Flush()
			fileWriter.Close()

			newName := gofakeit.AppName() + ext
			pathWithoutFileName, _ := strings.CutSuffix(afterPath, "/"+d.Name())
			err = os.Rename(folderPath+afterPath, folderPath+pathWithoutFileName+"/"+newName)
			if err != nil {
				log.Println(err)
			}
		}

		return nil
	})
}

func (r *RefInator) changeNamesWorker(line string) string {
	line_older := line

	for class := range r.changes.Classes {
		if strings.Contains(line, class) {
			line = strings.ReplaceAll(line, class, r.changes.Classes[class])
		}
	}

	for function := range r.changes.Funcs {
		if strings.Contains(line, function) {
			line = strings.ReplaceAll(line, function, r.changes.Funcs[function])
		}
	}

	for enum := range r.changes.Enums {
		if strings.Contains(line, enum) {
			line = strings.ReplaceAll(line, enum, r.changes.Enums[enum])
		}
	}

	for structure := range r.changes.Structs {
		if strings.Contains(line, structure) {
			line = strings.ReplaceAll(line, structure, r.changes.Structs[structure])
		}
	}

	if line_older != line {
		log.Println(line_older, "\nChanged line: ", line)
	}

	return line
}

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Cell struct {
	CellType string   `json:"cell_type"`
	Source   []string `json:"source"`
}
type Notebook struct {
	Cells []Cell `json:"cells"`
}

func convert(filepath string) ([]string, error) {
	var out bytes.Buffer
	var catcmd *exec.Cmd
	if runtime.GOOS == "windows" {
		catcmd = exec.Command("cmd", "/c", "type", filepath)
	} else {
		catcmd = exec.Command("cat", filepath)
	}
	catcmd.Stdout = &out
	err := catcmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Error getting file content: %v", err)
	}
	var nb Notebook
	err = json.Unmarshal(out.Bytes(), &nb)
	if err != nil {
		return nil, fmt.Errorf("Error parsing file content: %v", err)
	}
	re := regexp.MustCompile(`(?m)^[#*_]+\s*(.*?)\s*[_*]*$`)
	var contarr []string
	for _, cell := range nb.Cells {
		sourceContent := ""
		cell.CellType = strings.ToLower(cell.CellType)
		if cell.CellType == "code" {
			for _, line := range cell.Source {
				sourceContent += line + "\n"
			}
		} else {
			for _, line := range cell.Source {
				line = re.ReplaceAllString(line, "$1")
				fmt.Println(line)
				sourceContent += "# " + line + "\n"
			}
		}
		contarr = append(contarr, sourceContent)
	}
	return contarr, nil
}

func main() {
	fileFlag := flag.String("file", "", "path to .ipynb file")
	flag.Parse()
	var files []string
	if *fileFlag != "" {
		files = append(files, *fileFlag)
	}
	files = append(files, flag.Args()...)
	if len(files) == 0 {
		fmt.Println("Error: No filepath found\nUsage: <script-name> [-file <file-path> / <file-path>]")
		return
	}
	for _, fpath := range files {
		extmtch, err := regexp.MatchString(`\.ipynb$`, fpath)
		if err != nil {
			fmt.Println("Error checking file:\n", err)
			return
		}
		if !extmtch {
			fmt.Println("Error: The file must have a .ipynb extension.")
			return
		}
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			fmt.Println("Error: No such file found")
			return
		}
		fname := regexp.MustCompile(`([^\\\/]+)\.[^\\\/]+$`).FindStringSubmatch(fpath)[1]
		data, err := convert(fpath)
		if err != nil {
			fmt.Println(err)
			return
		}
		pyfile, err := os.Create(fname + ".py")
		if err != nil {
			fmt.Println("Error creating file:\n", err)
			return
		}
		defer pyfile.Close()
		_, err = pyfile.WriteString(strings.Join(data, ""))
		if err != nil {
			fmt.Println("Error writing to file:\n", err)
			return
		}
		fmt.Printf("Python file '%s.py' created\n", fname)
	}
}

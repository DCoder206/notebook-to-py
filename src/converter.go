package main

import (
	"bytes"
	"encoding/json"
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

func convert(filepath string, contarr *[]string) {
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
		fmt.Println("Error getting file content:\n", err)
		return
	}
	var fdata interface{}
	err = json.Unmarshal(out.Bytes(), &fdata)
	if err != nil {
		fmt.Println("Error parsing file content:\n", err)
		return
	}
	var nb Notebook
	err = json.Unmarshal(out.Bytes(), &nb)
	if err != nil {
		fmt.Println("Error parsing file content:\n", err)
		return
	}
	re := regexp.MustCompile(`(?m)^[#*_]+\s*(.*?)\s*[_*]*$`)
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
		*contarr = append(*contarr, sourceContent)
	}
}
func main() {
	var fpath string
	fmt.Print("Enter file path >>> ")
	fmt.Scanln(&fpath)
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
	var data []string
	convert(fpath, &data)
	pyfile, err := os.Create(fname + ".py")
	if err != nil {
		fmt.Println("Error creating file:\n", err)
		return
	}
	defer pyfile.Close()
	_, err = pyfile.WriteString(strings.Join(data, ""))
	fmt.Println("Python file created successfully")
}

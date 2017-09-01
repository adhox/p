package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

// 1. Loading data from file or memory
// 2. Perform data cleaning and standardization
func Load(fname string, head bool) Panel {
	p := make(Panel)

	// check if filename is URL
	if u, err := url.Parse(fname); err == nil && fname[:4] == "http" {
		fmt.Printf("downloading (%s)\n", u.String())
		// if a URL, download file
		// and reassign fname
		fname = download(*u)
	}

	fmt.Printf("reading file (%s)\n", fname)

	switch path.Ext(fname) {
	case ".csv":
		return readCSV(fname, head).Clean()
	case ".xml", ".html":
		fmt.Printf("reading %s", fname)
		b, _ := ioutil.ReadFile(fname)
		fmt.Println(string(b))
		return nil

	case ".tsv":
		return nil

	case ".json":
		return nil

	case ".xlsx":
		// fmt.Println("reading EXCEL")
		return readXLSX(fname, head).Clean()

	default: // text file; treat like Hadoop
		// file, err := xlsx.OpenFile(fname)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	return p.Clean()
}

func Read(fname string, head bool) Panel {
	return Load(fname, head)
}

func download(u url.URL) string {

	res, _ := http.Get(u.String())
	defer res.Body.Close()

	fname := strings.ToLower(fmt.Sprintf("%s.%s", u.Host, strings.Split(path.Base(u.String()), "?")[0]))

	file, _ := os.Create(fname)
	body, _ := ioutil.ReadAll(res.Body)
	file.Write(body)

	return fname
}

func readCSV(fname string, head bool) Panel {
	p := make(Panel)

	h := []string{}
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}

	headless := false
	if !head {
		headless = true
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.FieldsPerRecord = -1
	for {
		if head {
			line, _ := r.Read()
			for _, val := range line {
				h = append(h, strings.ToLower(val))
			}
			head = false
		} else {
			line, err := r.Read()
			if err == io.EOF {
				break
			}
			for key, val := range line {
				if headless {
					k := strconv.Itoa(key)
					p[k] = append(p[k], val)
				} else {
					p[h[key]] = append(p[h[key]], val)
				}
			}
		}
	}
	return p
}

func readXLSX(fname string, head bool) Panel {
	p := make(Panel)
	h := []string{}

	file, err := xlsx.OpenFile(fname)
	if err != nil {
		fmt.Println(err)
	}

	headless := false
	if !head {
		headless = true
	}

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			if head {
				for _, cell := range row.Cells {
					val, _ := cell.String()
					h = append(h, strings.ToLower(val))
				}
				head = false
			} else {
				for cn, cell := range row.Cells {
					val, _ := cell.String()
					if headless {
						k := strconv.Itoa(cn)
						p[k] = append(p[k], val)
					} else {
						p[h[cn]] = append(p[h[cn]], val)
					}
				}
			}
		}
	}
	return p
}

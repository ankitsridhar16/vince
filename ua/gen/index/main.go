package index

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"

	"github.com/vinceanalytics/vince/tools"
)

func Make() {
	ls := []string{
		"mobile",
		"tablet",
		"laptop",
		"desktop",
	}
	seen := map[string]struct{}{
		"mobile":  {},
		"tablet":  {},
		"laptop":  {},
		"desktop": {},
	}
	dir, err := os.ReadDir(filepath.Join(tools.RootVince(), "ua"))
	if err != nil {
		tools.Exit(err.Error())
	}
	for _, v := range dir {
		if v.IsDir() {
			continue
		}
		if filepath.Ext(v.Name()) != ".json" {
			continue
		}
		b, err := os.ReadFile(v.Name())
		if err != nil {
			tools.Exit(err.Error())
		}
		var o []string
		err = json.Unmarshal(b, &o)
		if err != nil {
			tools.Exit(fmt.Sprintf("decoding %s %v", v.Name(), err))
		}
		for _, v := range o {
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			ls = append(ls, v)
		}
		tools.Remove(v.Name())
	}
	sort.Strings(ls)
	var b bytes.Buffer
	b.WriteString(`
	// DO NOT EDIT Code generated by ua/index/main.go"
	package ua 

	var commonIndex =map[string]uint16{
	`)
	for i, k := range ls {
		fmt.Fprintf(&b, "%q:%d,\n", k, i)
	}
	b.WriteString("}\n")
	b.WriteString(`
	var commonIndexReverse =map[uint16]string{
	`)
	for i, k := range ls {
		fmt.Fprintf(&b, "%d:%q,\n", i, k)
	}
	b.WriteString("}\n")
	x, err := format.Source(b.Bytes())
	if err != nil {
		tools.Exit("failed formatting go source ", err.Error())
	}
	tools.WriteFile("index.go", x)
}

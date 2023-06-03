package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"github.com/vinceanalytics/vince/tools"
	"gopkg.in/yaml.v2"
)

var index = map[string]bool{}
var refs = map[string]struct{}{}

type Domains struct {
	Type    string   `json:"-"`
	Name    string   `json:"-"`
	Index   int      `json:"-"`
	Domains []string `json:"domains"`
	Hosts   []string `json:"-"`
}

var domains []*Domains

type Name struct {
	name   string
	domain string
}

var names []*Name

const (
	repo = "git@github.com:snowplow-referer-parser/referer-parser.git"
	dir  = "referer-parser"
)

type refererData map[string]map[string]map[string][]string

func main() {
	root := tools.RootVince()
	tools.EnsureRepo(
		filepath.Join(root, "internal", "referrer"),
		repo, dir,
	)
	var data refererData
	bs := tools.ReadFile(filepath.Join(dir, "resources", "referers.yml"))
	err := yaml.NewDecoder(bytes.NewReader(bs)).Decode(&data)
	if err != nil {
		tools.Exit(err.Error())
	}

	var maxLen int
	var minLen = 6

	for metaType, o := range data {
		for metaName, xo := range o {
			xd := xo["domains"]
			x := Domains{
				Type:    metaType,
				Name:    metaName,
				Domains: xd,
			}
			names = append(names, &Name{
				name:   x.Name,
				domain: x.Domains[0],
			})
			for _, d := range x.Domains {
				refs[d] = struct{}{}
				u, _ := url.Parse("http://" + d)
				host := strings.TrimPrefix(u.Host, "www.")
				parts := strings.Split(host, ".")
				sort.Sort(sort.Reverse(StringSlice(parts)))
				if len(parts) > int(maxLen) {
					maxLen = len(parts)
				}
				host = strings.Join(parts, ".")
				if len(parts) < minLen {
					minLen = len(parts)
				}
				if index[host] {
					continue
				}
				x.Hosts = append(x.Hosts, host)
				index[host] = true
			}
			sort.Strings(x.Hosts)
			domains = append(domains, &x)
		}
	}

	var b bytes.Buffer

	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by referrer/make_referrer.go")
	fmt.Fprintln(&b, " package referrer")
	fmt.Fprintln(&b, " import \"sync\"")
	fmt.Fprintln(&b, " var Favicon =&sync.Map{}")
	fmt.Fprintln(&b, " func init() {")
	sort.SliceStable(names, func(i, j int) bool {
		return names[i].name < names[j].name
	})
	for _, v := range names {
		fmt.Fprintf(&b, " Favicon.Store(%q,%q)\n", v.name, v.domain)
	}
	fmt.Fprintln(&b, " }")
	fmt.Fprintf(&b, " const minReferrerSize=%d\n", minLen)
	fmt.Fprintf(&b, " const maxReferrerSize=%d\n", maxLen)
	fmt.Fprintln(&b, `
	type Medium struct {
		Type       string
		Name       string
	}
	`)
	fmt.Fprintln(&b, "var refList=map[string]*Medium{")
	sort.SliceStable(domains, func(i, j int) bool {
		return domains[i].Type < domains[j].Type &&
			domains[i].Name < domains[j].Name
	})
	for _, m := range domains {
		for _, h := range m.Hosts {
			fmt.Fprintf(&b, "%q:{Type:%q,Name:%q},\n", h, m.Type, m.Name)
		}
	}
	fmt.Fprintln(&b, "}")

	r, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	tools.WriteFile("referrer_data.go", r)

	b.Reset()
	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by referrer/make_referrer.go")
	fmt.Fprintln(&b, " package main")
	fmt.Fprintf(&b, " var domains=[%d]string{\n", len(refs))
	var ls []string
	for k := range refs {
		ls = append(ls, k)
	}
	sort.Strings(ls)
	for i, k := range ls {
		if i%4 == 0 {
			b.WriteByte('\n')
		}
		fmt.Fprintf(&b, "%q,", k)
	}
	fmt.Fprintln(&b, "\n}")

	r, err = format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	tools.WriteFile(
		filepath.Join(tools.RootVince(), "tools", "vince_load_gen", "domains.go"),
		r)
}

type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return i < j }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

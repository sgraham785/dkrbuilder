package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Cogility/dkr-img/internal/manifest/domain"
	"gopkg.in/yaml.v2"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var manifest string

	flag.StringVar(&manifest, "manifest", path+"/manifest.yaml", "Path to manifests file")
	flag.Parse()

	data, err := ioutil.ReadFile(manifest)
	if string(data) == "" {
		fmt.Println("No manifest found")
		os.Exit(1)
	}
	var e map[string]domain.Images
	if err := yaml.Unmarshal(data, &e); err != nil {
		fmt.Println(err.Error())
	}
	var img string
	var dkr string
	var versions []string
	var dockerfiles []string
	var c string

	for i := range e["images"].ImageNames {
		for v := range e["images"].ImageNames[i].Versions {
			if v != "dependents" {
				img = i + "-" + v
				fmt.Printf("%#v\n", i+"-"+v)
				versions = append(versions, img)
				if e["images"].ImageNames[i].Versions[v].Build.Dockerfile != "" {
					dockerfiles = append(dockerfiles, e["images"].ImageNames[i].Versions[v].Build.Dockerfile)
					// fmt.Println(string(e["images"].ImageNames[i].Versions[v].Dockerfile))
				} else {
					dkr = "./" + i + "/" + v + "/Dockerfile"
					dockerfiles = append(dockerfiles, dkr)
				}
			}
			for ch := range e["images"].ImageNames[i].Versions[v].Child {

				if v != "dependents" {
					c = ch
				}
				for cv := range e["images"].ImageNames[i].Versions[v].Child[c].Versions {
					if v != "dependents" {
						img = i + "-" + v + "-" + c + "-" + e["images"].ImageNames[i].Versions[v].Child[c].Versions[cv]
						// fmt.Printf("%#v\n", i+"-"+v+"-"+c+"-"+e["images"].ImageNames[i].Versions[v].Child[c].Versions[cv])
						versions = append(versions, img)

					}
				}
			}

		}
	}
	printImgSlice(versions)
	printDkrSlice(dockerfiles)
}

func printImgSlice(v []string) {
	fmt.Printf("versions=%s\n", v)
}

func printDkrSlice(v []string) {
	fmt.Printf("dockerfiles=%s\n", v)
}

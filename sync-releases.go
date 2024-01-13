package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Release struct {
	Tag    string `json:"tag"`
	Source Source `json:"source"`
}

type Source map[string]string

func main() {
	dir := "releases"
	files, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	var availableReleases []Release

	for _, file := range files {
		fileName := strings.Replace(file.Name(), ".zip", "", -1)
		fileName = strings.Replace(fileName, ".rar", "", -1)

		fmt.Println(fileName)

		parts := strings.Split(fileName, "-")

		version := parts[1]
		arch := parts[int(len((parts))-1)]

		alreadyIn := slices.IndexFunc(availableReleases, func(r Release) bool {
			return r.Tag == version
		})

		if alreadyIn == -1 {
			availableReleases = append(availableReleases, Release{
				Tag: version,
				Source: Source{
					arch: fileName,
				},
			})
		} else {
			availableReleases[alreadyIn].Source[arch] = fileName
		}
	}

	jsonData, err := json.MarshalIndent(availableReleases, "", "  ")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("releases.json", jsonData, os.ModePerm)

	if err != nil {
		panic(err)
	}
}

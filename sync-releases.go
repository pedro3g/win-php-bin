package main

import (
	"encoding/json"
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
		fileName := file.Name()

		splitDotLen := len(strings.Split(fileName, "."))
		extention := strings.Split(fileName, ".")[splitDotLen-1]
		fileName = strings.Replace(fileName, "."+extention, "", 1)
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
					arch: fileName + "." + extention,
				},
			})
		} else {
			availableReleases[alreadyIn].Source[arch] = fileName + "." + extention
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

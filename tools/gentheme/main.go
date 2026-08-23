// gentheme generates a go source file containing slices for all colors, icons and sizes.
package main

import (
	_ "embed"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"text/template"
)

const (
	baseURL      = "https://raw.githubusercontent.com/fyne-io/fyne/refs/tags/v2.8.0/theme/"
	colorFileURL = baseURL + "color.go"
	iconsFileURL = baseURL + "icons.go"
	sizeFileURL  = baseURL + "size.go"
)

//go:embed target.go.template
var tmplText string

var (
	reColors = regexp.MustCompile(`(ColorName\w+) fyne.ThemeColorName`)
	reIcons  = regexp.MustCompile(`(IconName\w+) fyne.ThemeIconName`)
	reSizes  = regexp.MustCompile(`(SizeName\w+) fyne.ThemeSizeName`)
)

var (
	packageFlag = flag.String("p", "main", "package name")
	output      = flag.String("out", "", "writes to given file when specified")
)

func main() {
	flag.Parse()

	colors, err := extract(reColors, colorFileURL)
	if err != nil {
		log.Fatal(err)
	}
	icons, err := extract(reIcons, iconsFileURL)
	if err != nil {
		log.Fatal(err)
	}
	sizes, err := extract(reSizes, sizeFileURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := generateFile(colors, icons, sizes); err != nil {
		log.Fatal(err)
	}
}

func extract(re *regexp.Regexp, url string) ([]string, error) {
	f, err := fetchFile(url)
	if err != nil {
		return nil, err
	}
	matches := re.FindAllStringSubmatch(f, -1)
	var result []string
	for _, match := range matches {
		result = append(result, match[1])
	}
	slices.Sort(result)
	return result, err
}

func fetchFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func generateFile(colors, icons, sizes []string) error {
	tmpl, err := template.New("").Parse(tmplText)
	if err != nil {
		return err
	}

	var out io.Writer
	if *output == "" {
		out = os.Stdout
	} else {
		f, err := os.Create(*output)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}
	err = tmpl.Execute(out, map[string]any{
		"Package": *packageFlag,
		"Colors":  colors,
		"Icons":   icons,
		"Sizes":   sizes,
	})
	if err != nil {
		return err
	}
	return nil
}

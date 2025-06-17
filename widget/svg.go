package widget

import (
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
)

func isFileSVG(path string) bool {
	return strings.EqualFold(filepath.Ext(path), ".svg")
}

// isResourceSVG checks if the resource is an SVG or not.
func isResourceSVG(res fyne.Resource) bool {
	if isFileSVG(res.Name()) {
		return true
	}

	if len(res.Content()) < 5 {
		return false
	}

	switch strings.ToLower(string(res.Content()[:5])) {
	case "<!doc", "<?xml", "<svg ":
		return true
	}
	return false
}

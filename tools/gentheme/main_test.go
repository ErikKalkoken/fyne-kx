package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchFile_ReturnsBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	defer srv.Close()

	got, err := fetchFile(srv.URL)

	require.NoError(t, err)
	assert.Equal(t, "hello world", got)
}

func TestFetchFile_ReturnsErrorForUnreachableURL(t *testing.T) {
	_, err := fetchFile("http://127.0.0.1:0")

	assert.Error(t, err)
}

func TestExtract_ParsesAndSortsMatchingNames(t *testing.T) {
	const source = `
package theme

const (
	ColorNameZebra fyne.ThemeColorName = "zebra"
	ColorNameApple fyne.ThemeColorName = "apple"
	IconNameHome   fyne.ThemeIconName  = "home"
)
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(source))
	}))
	defer srv.Close()

	got, err := extract(reColors, srv.URL)

	require.NoError(t, err)
	assert.Equal(t, []string{"ColorNameApple", "ColorNameZebra"}, got)
}

func TestExtract_ReturnsErrorForUnreachableURL(t *testing.T) {
	_, err := extract(reColors, "http://127.0.0.1:0")

	assert.Error(t, err)
}

func TestGenerateFile_WritesColorsIconsAndSizesToFile(t *testing.T) {
	oldPackage, oldOutput := *packageFlag, *output
	defer func() {
		*packageFlag, *output = oldPackage, oldOutput
	}()

	*packageFlag = "mytheme"
	*output = filepath.Join(t.TempDir(), "target.go")

	err := generateFile([]string{"ColorNameApple"}, []string{"IconNameHome"}, []string{"SizeNamePadding"})
	require.NoError(t, err)

	got, err := os.ReadFile(*output)
	require.NoError(t, err)

	content := string(got)
	assert.Contains(t, content, "package mytheme")
	assert.Contains(t, content, `{"ColorNameApple", theme.ColorNameApple}`)
	assert.Contains(t, content, `{"IconNameHome", theme.IconNameHome}`)
	assert.Contains(t, content, `{"SizeNamePadding", theme.SizeNamePadding}`)
}

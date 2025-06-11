# fyne-kx

A library with extensions and tools for the Fyne GUI toolkit.

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/fyne-kx)
[![Fyne](https://img.shields.io/badge/dynamic/regex?url=https%3A%2F%2Fgithub.com%2FErikKalkoken%2Ffyne-kx%2Fblob%2Fmain%2Fgo.mod&search=fyne%5C.io%5C%2Ffyne%5C%2Fv2%20(v%5Cd*%5C.%5Cd*%5C.%5Cd*)&replace=%241&label=Fyne&cacheSeconds=https%3A%2F%2Fgithub.com%2Ffyne-io%2Ffyne)](https://github.com/fyne-io/fyne)
[![build status](https://github.com/ErikKalkoken/fyne-kx/actions/workflows/go.yml/badge.svg)](https://github.com/ErikKalkoken/fyne-kx/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/ErikKalkoken/fyne-kx/graph/badge.svg?token=fDk5XvdhOQ)](https://codecov.io/gh/ErikKalkoken/fyne-kx)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-kx.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/fyne-kx)

> [!NOTE]
> The library is in active development. Any feedback or suggestions are welcome.

## Contents

- [Installation](#installation)
- [Extensions](#extensions)
  - [Dialogs](#dialogs)
  - [Layouts](#layouts)
  - [Modals](#modals)
  - [Themes](#modals)
  - [Widgets](#widgets)
- [Apps](#apps)
  - [demo](#demo)
  - [fynetheme](#fyne-theme)

## Installation

You can add this library to your current Fyne project with the following command:

```sh
go get github.com/ErikKalkoken/fyne-kx
```

## Extensions

This library contains several extensions for the [Fyne GUI toolkit](https://fyne.io/).

> [!TIP]
> For a live demo and example code please see the [demo app](#demo).

### Dialogs

The library provides helpers for building dialogs.

- [AddDialogKeyHandler](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/dialog#AddDialogKeyHandler) adds a key handler to a dialog. It enables the user to close the dialog by pressing the escape key.

### Layouts

- [Columns](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/layout#NewColumns) arranges all objects in a row, with each in their own column with a given minimum width.
It can be used to arrange subsequent rows of objects in columns.

- [RowWrap](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/layout#NewColumns) a layout that dynamically arranges objects of similar height in rows and wraps them dynamically.

### Modals

Modals are similar to Fyne dialogs, but do not require user interaction.
They are useful when you have a longer running process that the user needs to wait for before he can continue. e.g. opening a large file.

#### Progress

[Progress modals](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/modal#hdr-Progress_modals) are modals that show a progress indicator while an action function is running. The library provides several variants.

[Progress modal demo](https://github.com/user-attachments/assets/047c0464-0324-45c4-940e-f7d489b1ad11)

### Themes

Further, additional custom themes are provided:

- [DefaultWithFixedVariant](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/theme#DefaultWithFixedVariant) allows apps to set a permanent light or dark mode.

### Widgets

This library contains several Fyne widgets:

- [Badge](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Badge) is a variant of the Fyne label widget that renders a rounded box around the text.
- [FilterChipGroup](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#FilterChipGroup) allows the user to toggle multiple filters with filter chips.
- [FilterChipSelect](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#FilterChipSelect) is a filter chip that allows the user to select and de-select one option from a list of options.
- [Slider](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Slider) is a variation of the Slider widget that also displays the current value.
- [TappableIcon](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableIcon) is an icon widget which runs a function when tapped.
- [TappableImage](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableImage) is widget which shows an image and runs a function when tapped.
- [TappableLabel](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableLabel) is a variant of the Fyne Label which runs a function when tapped.
- [Switch](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Switch) is a widget implementing a digital switch with two mutually exclusive states: on/off.

The widgets can be used just like any other widget from the Fyne standard library. All widgets are themeable and unit tested.

Here is an example for the Switch widget:

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"

	kxwidget "github.com/ErikKalkoken/fyne-kx/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Example")
	s := kxwidget.NewSwitch(func(on bool) {
		fmt.Printf("Switch: %v\n", on)
	})
	w.SetContent(s)
	w.ShowAndRun()
}
```

## Apps

This library also contains two Fyne apps.

> [!TIP]
> To run any of the provided Fyne apps directly, you need to have Fyne installed and configured in your system.
> For more information on how to configure your system for Fyne please see: [Getting Started](https://docs.fyne.io/started/).

### Demo

Demo is a Fyne app for demonstrating the features provided by the fyne-kx library. It can also show you how each component of this library can be used in code.

To run the demo app directly use the following command:

```sh
go run github.com/ErikKalkoken/fyne-kx/cmd/demo@latest
```

![example](https://cdn.imgpile.com/f/UFSaUqd_xl.png)

### Fyne theme

Fynetheme is a Fyne app for showing details about the default Fyne theme like colors, icons and sizes and has a search functions to help find them more quickly. This app can be especially useful when creating your own widgets.

You can install this tool directly with the following command:

```sh
go install github.com/ErikKalkoken/fyne-kx/cmd/fynetheme@latest
```

Once installed it will be available with:

```sh
fynetheme
```

![Example](https://cdn.imgpile.com/f/vCHVA6I_xl.png)

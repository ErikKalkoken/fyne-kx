# fyne-kx

A library with extensions and tools for the Fyne GUI toolkit.

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/fyne-kx)
[![Fyne](https://img.shields.io/badge/dynamic/regex?url=https%3A%2F%2Fgithub.com%2FErikKalkoken%2Ffyne-kx%2Fblob%2Fmain%2Fgo.mod&search=fyne%5C.io%5C%2Ffyne%5C%2Fv2%20(v%5Cd*%5C.%5Cd*%5C.%5Cd*)&replace=%241&label=Fyne&cacheSeconds=https%3A%2F%2Fgithub.com%2Ffyne-io%2Ffyne)](https://github.com/fyne-io/fyne)
[![build status](https://github.com/ErikKalkoken/fyne-kx/actions/workflows/go.yml/badge.svg)](https://github.com/ErikKalkoken/fyne-kx/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/ErikKalkoken/fyne-kx/graph/badge.svg?token=fDk5XvdhOQ)](https://codecov.io/gh/ErikKalkoken/fyne-kx)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-kx.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/fyne-kx)

## Contents

- [Description](#description)
- [Installation](#installation)
- [Extensions](#extensions)
  - [Dialogs](#dialogs)
  - [Layouts](#layouts)
  - [Modals](#modals)
  - [Themes](#themes)
  - [Widgets](#widgets)
- [Demo](#demo)

## Description

fyne-kx extends the [Fyne GUI toolkit](https://fyne.io/) with additional widgets, layouts, modals, themes and dialog helpers. All extensions are unit tested and designed to fit seamlessly into Fyne apps.

Most extensions were originally developed for other Fyne apps, such as [EVE Buddy](https://github.com/ErikKalkoken/evebuddy) and [Janice](https://github.com/ErikKalkoken/janice/), and were later moved into this library to make them available to the Fyne community. Many are still used by those apps today.

> [!NOTE]
> This library is actively maintained. Feedback, bug reports and suggestions are always welcome, so please feel free to open an issue.

## Installation

You can add this library to your current Fyne project with the following command:

```sh
go get github.com/ErikKalkoken/fyne-kx
```

## Extensions

The following extensions are provided:

> [!TIP]
> For a live demo and example code please see the [demo app](#demo).

### Dialogs

The library provides helpers for building dialogs.

- [AddDialogKeyHandler](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/dialog#AddDialogKeyHandler) adds a key handler to a dialog. It enables the user to close the dialog by pressing the escape key.

### Layouts

- [Columns](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/layout#NewColumns) arranges all objects in a row, with each in their own column with a given minimum width.
It can be used to arrange subsequent rows of objects in columns.

### Modals

Modals are similar to Fyne dialogs, but do not require user interaction.
They are useful when you have a longer running process that the user needs to wait for before they can continue. e.g. opening a large file.

#### Progress

[Progress modals](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/modal#hdr-Progress_modals) are modals that show a progress indicator while an action function is running. The library provides several variants.

[Progress modal demo](https://github.com/user-attachments/assets/047c0464-0324-45c4-940e-f7d489b1ad11)

### Themes

Further, additional custom themes are provided:

- [DefaultWithFixedVariant](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/theme#DefaultWithFixedVariant) allows apps to set a permanent light or dark mode.

### Widgets

The widgets can be used just like any other widget from the Fyne standard library. All widgets are themeable. Please also see the included [demo](#demo) app that shows them all in action.

![widgets](https://github.com/user-attachments/assets/d4acd8d0-bd6d-4c19-a134-e54fd0717b31)

- [Badge](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Badge) is a variant of the Fyne label widget that renders a rounded box around the text.
- [FilterChip](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#FilterChip) an interactive chip for filtering content. It has a label and can be turned on or off.
- [FilterChipGroup](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#FilterChipGroup) allows the user to toggle multiple filters with filter chips.
- [FilterChipSelect](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#FilterChipSelect) is a filter chip that allows the user to select and de-select one option from a list of options.
- [IconButton](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#IconButton) is a widget which helps users take minor actions with one tap.
- [LoadingButton](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#LoadingButton) is a button that shows a loading indicator while its action is running.
- [Slider](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Slider) is a variation of the Slider widget that also displays the current value.
- [SortChip](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#SortChip) is a chip widget that shows current sorting (column & order) and allows the user to change it by selecting from a drop down menu.
- [Spinner](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Spinner) indicates that an operation is in progress with unknown duration. It shows an animated ring, similar to a Material Design circular progress indicator.
- [Switch](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Switch) is a widget implementing a digital switch with two mutually exclusive states: on/off.
- [TappableIcon](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableIcon) is an icon widget which runs a function when tapped.
- [TappableImage](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableImage) is widget which shows an image and runs a function when tapped.
- [TappableLabel](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableLabel) is a variant of the Fyne Label which runs a function when tapped.

### LoadingButton demo

The following is a video demonstration of different LoadingButtons in action.

[Loading button](https://github.com/user-attachments/assets/a1439d7f-dd9e-47b3-bee3-fdf040950a13)

### Example

Here is an example on how to use the Switch widget in code:

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

## Demo

Demo is a Fyne app that provides a live demonstration of the various fyne-kx extensions. Its code also serves as an example of how they can be used.

> [!IMPORTANT]
> To run any of the provided Fyne apps directly, you need to have Fyne installed and configured in your system.
> For more information on how to configure your system for Fyne please see: [Getting Started](https://docs.fyne.io/started/).

You can run the demo app directly from the repo with:

```sh
go run github.com/ErikKalkoken/fyne-kx/cmd/demo@latest
```

> [!NOTE]
> The previously included `fynetheme` app has moved to its own repo: [fyne-theme-explorer](https://github.com/ErikKalkoken/fyne-theme-explorer).

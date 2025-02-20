# fyne-kx

A library with extensions and tools for the Fyne GUI toolkit.

[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-kx.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx)

> [!NOTE]
> The library is still in active development and the API is not fully stable yet. Any feedback or suggestions are welcome.

## Contents

- [Installation](#installation)
- [Extensions](#extensions)
  - [Dialogs](#dialogs)
  - [Layouts](#layouts)
  - [Modals](#modals)
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

![Example](https://cdn.imgpile.com/f/0if8yhY_xl.png)

### Modals

Modals are similar to Fyne dialogs, but do not require user interaction.
They are useful when you have a longer running process that the user needs to wait for before he can continue. e.g. opening a large file.

#### Progress

[Progress modals](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/modal#hdr-Progress_modals) are modals that show a progress indicator while an action function is running. The library provides several variants.

[Progress modal demo](https://github.com/user-attachments/assets/047c0464-0324-45c4-940e-f7d489b1ad11)

### Widgets

This library contains several Fyne widgets:

- [Badge](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Badge) is a variant of the Fyne label widget that renders a rounded box around the text
- [Slider](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Slider) is a variation of the Slider widget that also displays the current value
- [TappableIcon](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableIcon) is an icon widget which runs a function when tapped
- [TappableImage](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableImage) is widget which shows an image and runs a function when tapped
- [TappableLabel](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableLabel) is a variant of the Fyne Label which runs a function when tapped
- [Switch](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Switch) is a widget implementing a digital switch with two mutually exclusive states: on/off

[Widget demo](https://github.com/user-attachments/assets/fb37a56a-dafa-49b5-92f2-e6c61457bdc4)

The widgets can be used just like any other widget from the Fyne standard library. Here is an example for the Switch widget:

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

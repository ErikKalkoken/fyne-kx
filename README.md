# fyne-kx

A library with extensions and tools for the Fyne GUI toolkit.

[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-kx.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx)

> [!NOTE]
> This is an early version of the library and the API may still change until version 1.0 is reached.

## Description

This library contains several extensions for the [Fyne GUI toolkit](https://fyne.io/):

- [Dialogs](#dialogs)
- [Layouts](#layouts)
- [Modals](#modals)
- [Widgets](#widgets)

In addition it contains two Fyne apps:

- [demo](#demo): Live demo of the features provided by this library
- [fynetheme](#fyne-theme): Shows insights about the default Fyne theme like colors and icons

For more details please see the [Go documentation](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx) for this package.

### Dialogs

The library provides helpers for building dialogs, e.g. a helper that enables closing any dialog with the escape key.

### Layouts

- [Columns](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/layout#NewColumns): Columns arranges all objects in a row, with each in their own column with a given minimum width.
It can be used to arrange subsequent rows of objects in columns.

![Example](https://cdn.imgpile.com/f/0if8yhY_xl.png)

### Modals

Modals are similar to Fyne dialogs, but do not require user interaction. They are useful when you have a longer running process that the user needs to wait for before he can continue. e.g. opening a large file.

#### Progress

[Progress modals](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/modal#hdr-Progress_modals) are modals that show a progress indicator while an action function is running. The library provides several variants.

![progress modal](https://cdn.imgpile.com/f/vZkxURa_xl.png)

### Widgets

This library contains several Fyne widgets:

- [Badge](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Badge) is a variant of the Fyne label widget that renders a rounded box around the text
- [Slider](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Slider) is a variation of the Slider widget that also displays the current value
- [TappableIcon](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableIcon) is an icon widget which runs a function when tapped
- [TappableImage](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableImage) is widget which shows an image and runs a function when tapped
- [TappableLabel](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#TappableLabel) is a variant of the Fyne Label which runs a function when tapped
- [Toggle](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx/widget#Toggle) is a widget implementing a digital switch with two mutually exclusive states: on/off

![example](https://cdn.imgpile.com/f/bpDBDRn_xl.png)

## Apps

> [!TIP]
> For more information on how to configure your system for Fyne please see: [Getting Started](https://docs.fyne.io/started/).

### Demo

Demo is a Fyne app for demonstrating the features provided by the fyne-kx library.

You can run it directly with the following command:

```sh
go run github.com/ErikKalkoken/fyne-kx/cmd/demo@latest
```

![example](https://cdn.imgpile.com/f/bpDBDRn_xl.png)

### Fyne theme

Fynetheme is a Fyne app for showing details about the default Fyne theme like colors, icons and sizes. You can run it directly with the following command:

```sh
go run github.com/ErikKalkoken/fyne-kx/cmd/themeinsight@latest
```

![Example](https://cdn.imgpile.com/f/vCHVA6I_xl.png)

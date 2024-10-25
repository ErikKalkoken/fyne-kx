# fyne-kx

A library with extensions for the Fyne GUI toolkit.

[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/fyne-kx.svg)](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx)

## Description

This library contains several extensions for the [Fyne GUI toolkit](https://fyne.io/):

- [Layouts](#layouts)
- [Modals](#modals)
- [Widgets](#widgets)

For more details please see the [Go documentation](https://pkg.go.dev/github.com/ErikKalkoken/fyne-kx) for this package.

For a live demo please see the [demo app](#demo).

### Layouts

- **Columns**: Columns arranges all objects in a row, with each in their own column with a given minimum width.
It can be used to arrange subsequent rows of objects in columns.

![Example](https://cdn.imgpile.com/f/maoyoP1_xl.png)

### Modals

Modals are similar to Fyne dialogs, but do not require user interaction. They are useful when you have a longer running process that the user needs to wait for before he can continue. e.g. opening a large file.

#### Progress

Progress modals are modals that show a progress indicator while an action function is running. The library provides several variants.

![progress modal](https://cdn.imgpile.com/f/p8NDn3O_xl.png)

### Widgets

This library contains several Fyne widgets:

- Badge is a variant of the Fyne label widget that renders a rounded box around the text
- SliderWithValue is a variation of the Slider widget that also displays the current value
- TappableIcon is an icon widget which runs a function when tapped
- TappableImage is widget which shows an image and runs a function when tapped
- TappableLabel is a variant of the Fyne Label which runs a function when tapped
- Toggle is a widget implementing a digital switch with two mutually exclusive states: on/off

![example](https://cdn.imgpile.com/f/swLOMWS_xl.png)

## Demo

For a live demo you can run the demo app with the following command:

```sh
go run github.com/ErikKalkoken/fyne-kx/cmd/demo@latest
```

> [!TIP]
> For more information on how to configure your system for Fyne please see: [Getting Started](https://docs.fyne.io/started/).

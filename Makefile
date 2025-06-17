bundle:
	fyne bundle resources/demo > cmd/demo/resource.go
	fyne bundle --package widget --prefix Icon resources/widget > widget/resource.go
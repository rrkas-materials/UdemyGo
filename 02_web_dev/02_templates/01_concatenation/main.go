package main

import "fmt"

func main() {
	name := "Rohnak Agarwal"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`

	fmt.Println(tpl)
}

/*
	$go run main.go > index.html

*/

package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maranqz/gopublicfield"
)

func main() {
	singlechecker.Main(gopublicfield.NewAnalyzer())
}

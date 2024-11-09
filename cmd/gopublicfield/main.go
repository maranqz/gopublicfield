package main

import (
	"github.com/maranqz/gopublicfield"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gopublicfield.NewAnalyzer())
}

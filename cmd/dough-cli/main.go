package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drewart/dough/data"
	"github.com/drewart/dough/util"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var action, filePath string

	flag.StringVar(&action, "action", "", "action [init,import,import-cats]")
	flag.StringVar(&filePath, "file", "", "import or db file file")

	flag.Parse()

	if action != "" {
		fmt.Println(action)
	}

	if action == "init" {
		data.InitSchema(&filePath)
	} else if action == "import" {
		f, err := os.Open(filePath)
		if err != nil {
			log.Printf("Unable to read input file "+filePath, err)
		}

		defer f.Close()
		entries, err := util.ImportCSVToAccount(f)
		if err != nil {
			log.Fatalf("import error: %s", err)
		}

		for _, entry := range entries {
			fmt.Printf("%v \n", entry)
		}

	} else if action == "import-cats" {
		f, err := os.Open(filePath)
		if err != nil {
			log.Printf("Unable to read input file "+filePath, err)
		}
		defer f.Close()

		util.ImportCatagories(f)
	}

}

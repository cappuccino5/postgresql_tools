package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"postgresql_tools/internal/parse"
)

func main() {
	flagSet := flag.NewFlagSet("postgres_seq", flag.ExitOnError)
	projectFileDirectoryPointer := flagSet.String("file", "./internal/example/id_seq.json", "path to project file")
	projectOutputFileDirectoryPointer := flagSet.String("target", "id_seq_sql.out", "output to target file")

	projectFileDirectory := *projectFileDirectoryPointer
	projectOutputFileDirectory := *projectOutputFileDirectoryPointer

	projectRootDirectory, err := filepath.Abs(projectFileDirectory)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Println("project file: " + projectRootDirectory)
	fmt.Println("target file: " + projectOutputFileDirectory)
	fileStat, err := os.Stat(projectRootDirectory)
	if os.IsNotExist(err) {
		fmt.Println("file IsNotExist:", projectRootDirectory)
		return
	}

	if path.Ext(fileStat.Name()) == ".json" {
		err = parse.PararWithIdSeq(projectRootDirectory, projectOutputFileDirectory)
		if err != nil {
			fmt.Errorf("parse error:%v", err)
		}
	}

}

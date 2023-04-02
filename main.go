package main

import (
    "flag"
    "fmt"
    "github.com/piglig/go-unitypackage"
)

func main() {
    // the unitypackage path
    packagePath := flag.String("p", "", "package path")
    // the output assets path
    assetsPath := flag.String("o", "~/Downloads/assets", "output assets path")
    flag.Parse()

    if *packagePath == "" {
        panic(fmt.Errorf("package path option is mandatory"))
    }
    // Unpackage command will extract content from unitypackage.
    err := unitypackage.UnPackage(*packagePath, *assetsPath)
    if err != nil {
        panic(err)
    }
}

package main
import (
    "log"
    "os"
    "path/filepath"
)

func main() {

  e := filepath.Walk("./deploy", func(path string, info os.FileInfo, err error) error {
      if info.IsDir() == false {
        println(path,info.Name())
      }
      return nil
  })
  if e != nil {
      log.Fatal(e)
  }
}

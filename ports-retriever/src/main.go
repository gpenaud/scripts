package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "path/filepath"
  "os"
  "strconv"

  "gopkg.in/yaml.v2"
)

type Ports struct {
  Backends map[string]string
  Ports    map[string]map[string]string
}

func main() {
  backend          := flag.String("backend", "", "The backend port to get (Required)")
  instance         := flag.String("instance", "", "The environment instance (Required)")
  environment_type := flag.String("type", "", "The environment type. (Required)")
  flag.Parse()

  if *environment_type == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  if *instance == "" {
    flag.PrintDefaults()
    os.Exit(2)
  }

  if *backend == "" {
    flag.PrintDefaults()
    os.Exit(3)
  }
  
  filename, _    := filepath.Abs("./ports.yml")
  yaml_file, err := ioutil.ReadFile(filename)

  if err != nil {
      panic(err)
  }

  var ports Ports

  err = yaml.Unmarshal(yaml_file, &ports)
  if err != nil {
      panic(err)
  }

  port_str      := ports.Ports[*environment_type][*instance]
  increment_str := ports.Backends[*backend]

  if port_str == "" {
    log.Fatal("Port range not found")
    os.Exit(4)
  }

  if increment_str == "" {
    log.Fatal("Backend not found")
    os.Exit(5)
  }

  port, _      := strconv.ParseInt(port_str, 10, 64)
  increment, _ := strconv.ParseInt(increment_str, 10, 64)

  fmt.Printf("Port: %#v\n", (port + increment))
}

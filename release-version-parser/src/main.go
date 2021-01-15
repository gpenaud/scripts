package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "os"
)

type Version struct {
  Services map[string][]map[string]string
  Database map[string]string
  Indexing map[string]string
  Messaging map[string]string
  Apigateway map[string]string
}

func get_service_version(service Version, name string) string {
  var version string

  for i, frontend := range service.Services["frontends"] {
    if frontend["name"] == name {
      version = service.Services["frontends"][i]["version"]
      break
    }
  }

  for i, backend := range service.Services["backends"] {
    if backend["name"] == name {
      version = service.Services["backends"][i]["version"]
      break
    }
  }

  return version
}

func main() {
  object    := flag.String("object", "", "The type of object to get (Possible choices: services, database, indexing, messaging, apigateway) (Required)")
  service   := flag.String("service", "", "The service from version to get (Required if entity \"services\" has been selected)")
  flag.Parse()

  if *object == "" {
    flag.PrintDefaults()
    log.Println("ERROR: Should take --object as first argument")
    os.Exit(1)
  }

  if *object == "services" && *service == "" {
    flag.PrintDefaults()
    log.Println("ERROR: Should take --service as second argument")
    os.Exit(2)
  }

  file_path      := os.Getenv("RELEASE_VERSION_FILE_PATH")
  json_file, err := ioutil.ReadFile(file_path)

  if err != nil {
      panic(err)
  }

  var versions Version

  err = json.Unmarshal(json_file, &versions)

  if err != nil {
      panic(err)
  }

  var version string

  switch object_type := *object; object_type {
	case "services":
		version = get_service_version(versions, *service)
	case "database":
		version = versions.Database["version"]
  case "indexing":
  	version = versions.Indexing["version"]
  case "messaging":
  	version = versions.Messaging["version"]
  case "apigateway":
  	version = versions.Apigateway["version"]
	default:
		log.Fatal("No version found for those parameters")
  }

  fmt.Print(version)
}

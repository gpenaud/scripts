package main

import (
  diskcache "github.com/gregjones/httpcache/diskcache"
  filepath  "path/filepath"
  gitlab    "github.com/xanzy/go-gitlab"
  json      "encoding/json"
  http      "net/http"
  log       "log"
  os        "os"
  template  "html/template"
	time      "time"
)

type GitlabProject struct {
  Branches []*gitlab.Branch
  Tags     []*gitlab.Tag
}

func get_gitlab_project_data(client *gitlab.Client, project *gitlab.Project, channel chan GitlabProject) <- chan GitlabProject {
  go func() {
    gitlab_project := GitlabProject{}

    branches, _, err := client.Branches.ListBranches(project.ID, nil, nil)
    if err != nil { log.Fatal(err); return }
    gitlab_project.Branches = branches

    tags, _, err := client.Tags.ListTags(project.ID, nil, nil)
    if err != nil { log.Fatal(err); return }
    gitlab_project.Tags = tags

    channel <- gitlab_project
  } ()

  return channel
}

func avs4_deployment_by_branches(w http.ResponseWriter, r *http.Request) {
  // Get path to template directory from env variable
  templates_directory := os.Getenv("TEMPLATES_DIR")

	templates_path := filepath.Join(templates_directory, "/avs4-deployment-by-branches.html")
	tmpl := template.Must(template.ParseFiles(templates_path))

  cache := diskcache.New("data/cache/project-data.json")

  data := make(map[string]GitlabProject)

  if value, found := cache.Get("project"); found {
		data_bytes := value
    json.Unmarshal(data_bytes, &data)
	} else {
    // local gitlab server
    git, err := gitlab.NewBasicAuthClient(nil, "http://gitlab", "root", "delivery")
    if err != nil { log.Fatal(err) }

    // opt := &ListProjectsOptions{Search: gitlab.String("root")}
    projects, _, err := git.Projects.ListProjects(nil)
    if err != nil { log.Fatal(err) }

    channel := make(chan GitlabProject)

    for _, project := range projects {
      data[project.Name] = <- get_gitlab_project_data(git, project, channel)
    }

    data_bytes, _ := json.Marshal(data)
    cache.Set("project", data_bytes)
    time.Sleep(10 * time.Millisecond)
  }

	tmpl.Execute(w, data)
}

func main() {
	// Create route to love-mountains web page
  http.HandleFunc("/", avs4_deployment_by_branches)
	// Launch web server on port 80
	log.Fatal(http.ListenAndServe(":8080", nil))
}

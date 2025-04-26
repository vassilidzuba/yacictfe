package yacicclient

import (
		"log"
	//	"strconv"
		"net/http"
		"io/ioutil"
		"encoding/json"
)

type Branch struct {
	Branch  string   `json:"branchId"`
	Dir     string   `json:"branchDir"`
}

type Project struct {
	ProjectId   string   `json:"projectId"`
	Repo        string   `json:"repo"`
	Branches    []Branch `json:"branches"`
}


type Build struct {
	ProjectId  string `json:"projectId"`
	BranchId   string `json:"branchId"`
	Timestamp  string `json:"timestamp"`
	Status     string `json:"status"`
	Duration   int    `json:"duration"`
}

type Step struct {
	ProjectId string `json:"projectId"`
	BranchId  string `json:"branchId"`
	Timestamp string `json:"timestamp"`
	StepId    string `json:"stepId"`
	Seq       int    `json:"seq"`
	Status    string `json:"status"`
	Duration  int    `json:"duration"`
}

var Projects []Project
var Builds   []Build
var Steps    []Step

func InitProjectList() {
	host := "localhost:8080"
	username := "vassili"
	password := "sekret"
	
	url := "http://" + host + "/yacic/project/list"

	log.Println("url:", url)
		
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		//panic("bad")
	}

	// print the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &Projects)
}

func InitBuildList(projectId string, branchId string) {
	host := "localhost:8080"
	username := "vassili"
	password := "sekret"

	url := "http://" + host + "/yacic/build/list?project=" + projectId
	if branchId != "" {
		url = url + "&branch=" + branchId
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		//panic("bad")
	}

	// print the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &Builds)
}


func InitStepList(projectId string, branchId string, timestamp string) {
	host := "localhost:8080"
	username := "vassili"
	password := "sekret"

		url := "http://" + host + "/yacic/step/list?project=" + projectId
	if branchId != "" {
		url = url + "&branch=" + branchId
	}
	url = url + "&timestamp=" + timestamp

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if resp.StatusCode != http.StatusOK {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		//panic("bad")
	}

	// print the response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &Steps)
}

func GetBranches(projectId string) []Branch {
	for _, p := range Projects {
		if (projectId == p.ProjectId) {
			return p.Branches			
		}
	}
	return nil
}

package tools

import (
	"encoding/json"
	"net/http"
)

type GitHubProfile struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}

func GetGitHubProfile(username string) (GitHubProfile, error) {
	resp, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		return GitHubProfile{}, err
	}
	defer resp.Body.Close()

	var profile GitHubProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return GitHubProfile{}, err
	}

	return profile, nil
}

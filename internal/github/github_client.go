package github

import (
	"context"
	"os"
	"time"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

// GitHubClient 包装 go-github 库客户端
type GitHubClient struct {
	client *github.Client
}

// NewGitHubClient 初始化 GitHub 客户端
func NewGitHubClient() *GitHubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &GitHubClient{client: client}
}

// GetIssues 获取项目的 issues
func (c *GitHubClient) GetIssues(owner, repo string, state string, since time.Time) ([]*github.Issue, error) {
	options := &github.IssueListByRepoOptions{
		State:     state,
		Since:     since,
		Sort:      "updated",
		Direction: "desc",
	}
	issues, _, err := c.client.Issues.ListByRepo(context.Background(), owner, repo, options)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

// GetPullRequests 获取项目的 pull requests
func (c *GitHubClient) GetPullRequests(owner, repo string, state string, since time.Time) ([]*github.PullRequest, error) {
	options := &github.PullRequestListOptions{
		State:     state,
		Sort:      "updated",
		Direction: "desc",
		Base:      "main",
	}
	pullRequests, _, err := c.client.PullRequests.List(context.Background(), owner, repo, options)
	if err != nil {
		return nil, err
	}
	return pullRequests, nil
}

// GetLatestRelease 获取项目的最新发布信息
func (c *GitHubClient) GetLatestRelease(owner, repo string) (*github.RepositoryRelease, error) {
	release, _, err := c.client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		return nil, err
	}
	return release, nil
}

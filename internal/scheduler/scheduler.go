package scheduler

import (
	"context"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/zysuper/github-sentinel-go/internal/repository"
	"github.com/zysuper/github-sentinel-go/pkg/logger"
	"golang.org/x/oauth2"
)

// ManualUpdate 手动触发更新
func ManualUpdate() {
	repos := repository.ListSubscriptions()

	// 获取 GitHub 客户端
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, repo := range repos {
		// 使用 go-github 库获取仓库事件
		ownerRepo := splitRepo(repo)
		events, _, err := client.Activity.ListRepositoryEvents(ctx, ownerRepo[0], ownerRepo[1], nil)
		if err != nil {
			logger.Error("Failed to fetch updates for repo:", repo, err)
			continue
		}

		logger.Info("Fetched updates for", repo, ":", events)
	}
}

func splitRepo(repo string) []string {
	return strings.Split(repo, "/")
}

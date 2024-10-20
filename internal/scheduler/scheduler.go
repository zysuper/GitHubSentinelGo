package scheduler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/zysuper/github-sentinel-go/internal/repository"
	"github.com/zysuper/github-sentinel-go/pkg/logger"
	"golang.org/x/oauth2"
)

// ManualUpdate 负责获取所有订阅仓库的最新发布信息
func ManualUpdate() {
	// 获取订阅的仓库列表
	repos := repository.ListSubscriptions()

	// 遍历每个仓库并获取其最新发布信息
	for _, repo := range repos {
		fetchLatestRelease(repo)
	}
}

// fetchLatestRelease 获取仓库的最新发布版本信息
func fetchLatestRelease(repo string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ownerRepo := splitRepo(repo)

	release, _, err := client.Repositories.GetLatestRelease(ctx, ownerRepo[0], ownerRepo[1])
	if err != nil {
		logger.Error("Failed to fetch latest release for repo:", repo, err)
		return
	}

	// 汇总发布版本信息
	report := fmt.Sprintf(
		"Repository: %s\nLatest Release: %s\nTag: %s\nPublished At: %s\nURL: %s\n",
		repo,
		release.GetName(),
		release.GetTagName(),
		release.GetPublishedAt(),
		release.GetHTMLURL(),
	)

	// 输出到终端或保存到文件
	logger.Info(report)
}

// splitRepo 将 owner/repo 拆分为 owner 和 repo
func splitRepo(repo string) []string {
	return strings.Split(repo, "/")
}

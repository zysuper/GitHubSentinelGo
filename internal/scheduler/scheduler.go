package scheduler

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zysuper/github-sentinel-go/internal/github"
	"github.com/zysuper/github-sentinel-go/internal/llm"
	"github.com/zysuper/github-sentinel-go/internal/repository"
	"github.com/zysuper/github-sentinel-go/pkg/logger"
)

// Scheduler 负责管理项目的每日进展、版本更新和报告生成
type Scheduler struct {
	githubClient *github.GitHubClient
	llmClient    *llm.LLMClient
}

// NewScheduler 创建一个新的 Scheduler 实例
func NewScheduler() *Scheduler {
	return &Scheduler{
		githubClient: github.NewGitHubClient(),
		llmClient:    llm.NewLLMClient(),
	}
}

// DailyProgress 获取项目 issues、pull requests 和最新发布信息，并导出成 Markdown 文件
func (s *Scheduler) DailyProgress() {
	repos := repository.ListSubscriptions()
	date := time.Now().Format("2006-01-02")

	for _, repo := range repos {
		ownerRepo := splitRepo(repo)
		issues, _ := s.githubClient.GetIssues(ownerRepo[0], ownerRepo[1])
		pullRequests, _ := s.githubClient.GetPullRequests(ownerRepo[0], ownerRepo[1])

		filename := fmt.Sprintf("%s_%s.md", repo, date)
		file, err := os.Create(filename)
		if err != nil {
			logger.Error("Failed to create report file:", err)
			continue
		}
		defer file.Close()

		// 写入 Issues 和 Pull Requests
		file.WriteString(fmt.Sprintf("# Daily Progress for %s\n\n## Issues\n", repo))
		for _, issue := range issues {
			file.WriteString(fmt.Sprintf("- %s (URL: %s)\n", issue.GetTitle(), issue.GetHTMLURL()))
		}
		file.WriteString("\n## Pull Requests\n")
		for _, pr := range pullRequests {
			file.WriteString(fmt.Sprintf("- %s (URL: %s)\n", pr.GetTitle(), pr.GetHTMLURL()))
		}

		// 获取并写入最新发布信息
		releaseInfo := s.fetchLatestRelease(repo)
		file.WriteString("\n## Latest Release\n")
		file.WriteString(releaseInfo)

		logger.Info(fmt.Sprintf("Progress report for %s saved to %s", repo, filename))
	}
}

// fetchLatestRelease 使用 GitHubClient 获取仓库的最新发布信息并返回格式化字符串
func (s *Scheduler) fetchLatestRelease(repo string) string {
	ownerRepo := splitRepo(repo)
	release, err := s.githubClient.GetLatestRelease(ownerRepo[0], ownerRepo[1])
	if err != nil {
		logger.Error("Failed to fetch latest release for repo:", repo, err)
		return "No release information available."
	}

	return fmt.Sprintf(
		"- Name: %s\n- Tag: %s\n- Published At: %s\n- URL: %s\n",
		release.GetName(),
		release.GetTagName(),
		release.GetPublishedAt(),
		release.GetHTMLURL(),
	)
}

// ManualUpdate 先更新 DailyProgress，然后生成并输出正式报告
func (s *Scheduler) ManualUpdate() {
	// 先生成最新的每日进展
	s.DailyProgress()

	// 然后生成正式报告
	repos := repository.ListSubscriptions()
	for _, repo := range repos {
		s.GenerateDailyReport(repo)
	}
}

// GenerateDailyReport 读取每日进展文件并使用 LLM 生成正式日报告
func (s *Scheduler) GenerateDailyReport(repo string) {
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s_%s.md", repo, date)
	content, err := os.ReadFile(filename)
	if err != nil {
		logger.Error("Failed to read progress file:", err)
		return
	}

	report, err := s.llmClient.GenerateReport(string(content))
	if err != nil {
		logger.Error("Failed to generate report:", err)
		return
	}

	// 输出到终端
	fmt.Printf("=== Formal Report for %s ===\n%s\n", repo, report)
}

// splitRepo 将 owner/repo 拆分为 owner 和 repo
func splitRepo(repo string) []string {
	return strings.Split(repo, "/")
}

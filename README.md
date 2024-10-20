# GitHub Sentinel Go

GitHub Sentinel Go 是一个开源的 CLI 工具，专为开发者和项目管理人员设计。它能够自动获取并汇总订阅的 GitHub 仓库的最新动态。通过这个工具，您可以轻松跟踪多个项目的进展，及时响应仓库更新事件，提升团队协作效率。

## 特性

- **订阅管理**：可以轻松地订阅和取消订阅 GitHub 仓库。
- **自动更新**：定期获取订阅仓库的最新动态（Pull Requests、Issues、Commits 等）。
- **事件查看**：通过 CLI 立即查看订阅仓库的最新活动。
- **使用 go-github 库**：依赖 GitHub 官方 API 客户端库 `go-github`。

## 安装

### 克隆项目
首先，克隆项目代码：
```bash
git clone https://github.com/zysuper/github-sentinel-go.git
cd github-sentinel-go


```

确保已经安装 Go 1.20 及以上版本，然后运行以下命令来安装依赖：

```bash
go mod tidy
```

### 编译项目

使用以下命令编译 CLI 工具：

```bash
go build -o github-sentinel cmd/main.go
```

### 运行 CLI

编译完成后，您可以通过以下命令来运行工具并进入交互式 REPL 环境：

```bash
./github-sentinel
```

在 REPL 模式下，您可以执行以下命令：

- `subscribe owner/repo`：订阅 GitHub 仓库。
- `unsubscribe owner/repo`：取消订阅 GitHub 仓库。
- `list`：列出所有已订阅的仓库。
- `update`：手动获取所有订阅仓库的更新。
- `help`：显示可用命令列表。
- `exit`：退出 REPL 环境。

#### 示例

```
> subscribe octocat/Hello-World
Successfully subscribed to: octocat/Hello-World

> list
Subscribed repositories:
- octocat/Hello-World

> update
Fetching updates...
```

## 如何配置项目



### 配置环境变量

在开始使用 GitHub Sentinel Go 之前，请确保已设置 GitHub 访问令牌。您可以通过 `GITHUB_TOKEN` 环境变量指定 GitHub API 访问令牌。

```
export GITHUB_TOKEN=your_github_token_here
```

您可以在 GitHub [Personal Access Tokens](https://github.com/settings/tokens) 页面创建一个新的令牌，并将其存储为环境变量。令牌需要具备读取仓库的权限。

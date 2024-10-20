package repository

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

const filePath = "subscriptions.json"

// Subscribe 订阅新的 GitHub 仓库
func Subscribe(repo string) error {
    repos, err := loadSubscriptions()
    if err != nil {
        return err
    }
    
    repos = append(repos, repo)
    return saveSubscriptions(repos)
}

// Unsubscribe 取消订阅
func Unsubscribe(repo string) error {
    repos, err := loadSubscriptions()
    if err != nil {
        return err
    }
    
    for i, r := range repos {
        if r == repo {
            repos = append(repos[:i], repos[i+1:]...)
            break
        }
    }
    
    return saveSubscriptions(repos)
}

// ListSubscriptions 列出所有订阅的仓库
func ListSubscriptions() []string {
    repos, err := loadSubscriptions()
    if err != nil {
        return []string{}
    }
    return repos
}

// 加载订阅信息
func loadSubscriptions() ([]string, error) {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return []string{}, nil
    }

    file, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var repos []string
    if err := json.Unmarshal(file, &repos); err != nil {
        return nil, err
    }

    return repos, nil
}

// 保存订阅信息
func saveSubscriptions(repos []string) error {
    data, err := json.Marshal(repos)
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filePath, data, 0644)
}

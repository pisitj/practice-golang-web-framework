# วิธีติดตั้งและใช้งาน Go Modules (Go 1.11+)
https://igokuz.com/%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5%E0%B8%95%E0%B8%B4%E0%B8%94%E0%B8%95%E0%B8%B1%E0%B9%89%E0%B8%87%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%87%E0%B8%B2%E0%B8%99-go-modules-go-1-11-62c36a0c5fd9

### go mod
- download    download modules to local cache
- edit        edit go.mod from tools or scripts
- graph       print module requirement graph
- init        initialize new module in current directory
- tidy        add missing and remove unused modules
- vendor      make vendored copy of dependencies
- verify      verify dependencies have expected content
- why         explain why packages or modules are needed

### initialize
- go mod init package-name
- go mod init github.com/{your-username}/{repo-name}

### update dependencies by creating .exe
​- go build

### remove unnecessary dependencies
- go mod tidy


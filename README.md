<h1 align="center">MinecraftCacher 📂</h1>

<p align="center">
  Caching Mojang's API requests to prevent Ratelimits issues
</p>

## 🔎 Requirements
- Redis database
- Go (to compile)


## 🏗️ How to run

* Clone this repo.
* Build code using `go build -o app.exe .`, to build for different platforms use
  * Powershell: ` $env:GOOS="linux";$env:GOARCH="amd64"; go build -o app .`
  * Bash: `GOOS="linux" GOARCH="amd64" go build -o app .`
  * Check available platforms using: `go tool dist list`
* Run executable. 

> **Note**: By default, the app will run on port `:8080`, you can change it by setting the `PORT` environment variable.

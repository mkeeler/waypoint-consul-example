project = "example-go"

app "example-go" {
  labels = {
    "service" = "example-go",
    "env"     = "dev"
  }

  build {
    use "pack" {}
  }

  config {
    env = {
      SERVE_ENV_DATA = dynamic("consul", {
        key  = "env-data"
      })
    }
  }

  deploy {
    use "docker" {}
  }
}

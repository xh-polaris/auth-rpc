## Badges

[![CI](https://github.com/xh-polaris/account-svc/actions/workflows/static-analysis.yml/badge.svg)](https://github.com/xh-polaris/account-svc/actions/workflows/static-analysis.yml)
[![Build](https://github.com/xh-polaris/account-svc/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/xh-polaris/account-svc/actions/workflows/docker-publish.yml)

## Get started

**Start services**

```bash
go run rpc/account.go -f rpc/etc/account.yaml
```

```bash
go run api/account.go -f api/etc/account.yaml
```

Before starting the server, please replace the default config file in `etc` directory.

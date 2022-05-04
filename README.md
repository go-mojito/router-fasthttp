<h1 align="center"><strong>fasthttp for Mojito</strong></h1>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/go-mojito/router-fasthttp" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/go-mojito/router-fasthttp" /></a>
	<a href="https://github.com/go-mojito/router-fasthttp" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/go-mojito/router-fasthttp.svg" /></a>
	<a href="https://godoc.org/github.com/go-mojito/router-fasthttp" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/go-mojito/router-fasthttp/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>
<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>
		
</p>
<p align="center">
fasthttp for Mojito provides a fasthttp Router implementation that was designed specifically for Mojito with full compatibility.
</p>

<p align="center"><strong>SonarCloud Report</strong></p>
<p align="center">
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=alert_status" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=sqale_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=reliability_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=security_rating" /></a>
	<br>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=vulnerabilities" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=code_smells" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_router-fasthttp" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_router-fasthttp&metric=bugs" /></a>
</p>

<p align="center"><strong>How to register</strong></p>
<p align="center">To register the fasthttp router as default, do the following in your main file</p>

```go
func init() {
    fasthttp.AsDefault()
}
```
<p align="center">To register fasthttp as a named router, do the following in your main file</p>

```go
func init() {
    fasthttp.As("myRouter")
}
```

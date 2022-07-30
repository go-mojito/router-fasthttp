<p align="center">
    <img src="/.github/assets/gopher.png"
        height="300">
</p>

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

<p align="center"><strong>Quickstart</strong></p>
<p align="center">To start using fasthttp as your router just enable it as the default router like this:</p>

```go
func init() {
    fasthttp.AsDefault()
}
```

<p align="center">Or to use fasthttp as a secondary router you can register it as a named router like this:</p>

```go
func init() {
    fasthttp.As("myRouter")
}
```

<p align="center"><strong>Documentation</strong></p>
<p align="center">
	Read our
	<a href="https://go-mojito.infinytum.co/docs">Documentation</a>, check out the 
	<a href="https://go-mojito.infinytum.co/">Project Website</a>.
</p>

<p align="center"><sub>Icon made with <a href="https://gopherize.me">Gopherize</a> and <a href="https://www.flaticon.com/free-icon/mojito_920710">flaticon</a>.</sub></p>

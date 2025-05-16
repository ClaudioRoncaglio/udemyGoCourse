# udemyGoCourse

Repo con gli esercizi del corso **Learn How To Code: Google's Go (golang) Programming Language** su [Udemy](https://www.udemy.com).

## Cross Compiling
Per compilare i programmi per altri sistemi operativi, basta anteporre al comando `go build` l'inizializzazione della variabile d'ambiente GOOS. 

Per vedere le variabili d'ambiente a disposizione ed i loro valori, basta lancire `go env`.

Comandi per compilare un programma per *MacOs, Linux  e Windows*:
- GOOS=darwin go build
- GOOS=linux go build
- GOOS=windows go build

## Principali fonti per la documentazione
- [Go User Manual](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Standard Go Library](https://pkg.go.dev/std)
- [Release History](https://go.dev/doc/devel/release)
- [Third Party Packages](https://pkg.go.dev/)

## Letture interessanti
- [Go Proverbs](https://go-proverbs.github.io/) (by Rob Pike)
- [Our Software Dependency Problem](https://research.switch.com/deps)(by Russ Cox)
- [Go 1 and the Future of Go Programs]( htsps://go.dev/doc/go1compat )
- [Go Blog](https://go.dev/blog/protobuf-opaque)

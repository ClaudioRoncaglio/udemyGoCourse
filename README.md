# udemyGoCourse

Repo con gli esercizi del corso **Learn How To Code: Google's Go (golang) Programming Language** su [Udemy](https://www.udemy.com).

## Cross Compiling
Per compilare i programmi per altri sistemi operativi, basta anteporre al comando `go build` l'inizializzazione della variabile d'ambiente GOOS. 

Per vedere le variabili d'ambiente a disposizione ed i loro valori, basta lancire `go env`.

Comandi per compilare un programma per *MacOs, Linux  e Windows*:
- GOOS=darwin go build
- GOOS=linux go build
- GOOS=windows go build

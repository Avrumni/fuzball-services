# Fuzball services

## Get "GO"ing

Download GO from <https://golang.org/doc/install?download=go1.10.3.darwin-amd64.pkg> 

Make go workspace

```
➜  mkdir -p ~/go/src/motome.com.au
➜  cd ~/go/src/motome.com.au
➜  go get -u github.com/kardianos/govendor
➜  
```

Starting application

```
➜  git clone _____
➜  cd _____
➜  ~/go/bin/govendor fetch +out
➜  go install
➜  ~/go/bin/fuzball-services
➜  go install && DATABASE_URL=postgres://localhost:5432/fuzball?sslmode=disable ~/go/bin/fuzball-services
➜  
```
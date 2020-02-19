# cursogolang
Curso da Udemy Golang

Apresento abaixo alguns dos principais comandos da linguagem Golang, com eles é possível verificar erros, biuldar os fonte, consultar as rotas do gopath goroot e outros.


Comand godoc
Ter acesso na sua maquina local ao go document mesmo ofline:
godoc -http=:[num_port] ex: godoc -http=:6060

Comand go env
tras toda a estrutura que o go criou na instalação tipo to GOPATH GOROOT e outras informaçes de instalação

Go build - cria arquivo de build
ex:	go build file.go

Go dev - executando dev checar erro no codigo
	go dev file.go

Go run - executa e compila arquivo
	go run file.go
    go fun *.go // varios arquivos
    
Go clean - limpa file build
	go clean file.go 

Comand go doc cmd/vet
O Vet examina o código-fonte Go e relata construções suspeitas, como
Chamadas Printf cujos argumentos não estão alinhados com a sequência de formato. Vet usa
heurísticas que não garantem todos os relatórios são problemas genuínos, mas
pode encontrar erros não capturados pelos compiladores.

Vet is normally invoked using the go command by running "go vet":

    go vet

vets the package in the current directory.

    go vet package/path/name

vets the package whose path is provided.

Use "go help packages" to see other ways of specifying which packages to vet.

ex:  go vet primeiro.go 

instalação de pacotes no go







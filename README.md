# cursogolang
Curso da Udemy Golang

Comand godoc
Ter acesso na sua maquina local ao go document mesmo ofline:
godoc -http=:[num_port] ex: godoc -http=:6060

Comand go env
tras toda a estrutura que o go criou na instalação tipo to GOPATH GOROOT e outras informaçes de instalação

Comand go doc cmd/vet
O Vet examina o codigo Go e reporta construções suspeitas como procurar Printf 
Vet examines Go source code and reports suspicious constructs, such as
Printf calls whose arguments do not align with the format string. Vet uses
heuristics that do not guarantee all reports are genuine problems, but it
can find errors not caught by the compilers.

Vet is normally invoked using the go command by running "go vet":

    go vet

vets the package in the current directory.

    go vet package/path/name

vets the package whose path is provided.

Use "go help packages" to see other ways of specifying which packages to
vet.

ex:  go vet primeiro.go 

comando para compilar 
go build file.go

comando para compilar e executar
go run file.go
go fun *.go // varios arquivos

instalação de pacotes no go







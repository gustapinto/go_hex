# Go HEX

Um projeto simples para aperfeiçoar arquitetura hexagonal, baseado no artigo [Ready for changes with Hexagonal Architecture
](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749), do time de engenharia de software da Netflix.

Esse projeto também visa se tornar um "*template*" para meus futuros desenvolvimentos de microsserviços em Go, usando novas funcionalidades da versão 1.22 da linguagem, podendo futuramente ser expandido com novos *cmds* ou *datasources*.

## Arquitetura hexagonal

A arquitetura hexagonal é uma arquitetura de *software* que prega o desacoplamento entre os componentes de um sistema, de forma que o desenvolvimento do mesmo seja altamente flexível, permitindo que as implementações de cada um de seus módulos sejam facilmente trocadas no futuro, com as dependências entre os mesmos se dando a partir do diagrama abaixo: 

<p align="center">
    <img src="https://miro.medium.com/v2/resize:fill:1200:632/g:fp:0.49:0.49/1*NfFzI7Z-E3ypn8ahESbDzw.png">
</p>

## Instalando o projeto

### Pré-requisitos

1. Go 1.22 ou mais recente
2. Compilador para [*protocol buffers*](https://grpc.io/docs/protoc-installation/) e das dependências do mesmo para Go e gRPC

### Executando a API REST (cmds/rest)

1. Clone esse repositório
2. Navegue até o diretório `go_hex/cmds/rest`
3. Em um terminal execute o comando `go run .`
4. Pronto! :rocket:

### Executando o servidor gRPC (cmds/grpc)

1. Clone esse repositório
2. Navegue até o diretório `go_hex/cmds/grpc`
3. (Opcional) Gere novo código para o projeto usando o script `./dev/gen/sh`
4. Em um terminal execute o comando `go run .`
5. Pronto! :rocket:

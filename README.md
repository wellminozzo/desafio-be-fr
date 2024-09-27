# Documentação da API em Go
## Introdução

Esta documentação fornece instruções sobre como configurar e rodar a API desenvolvida em Go. A API é uma aplicação simples que se conecta a um banco de dados MariaDB e oferece endpoints para manipulação de dados.
Pré-requisitos

Antes de iniciar, você precisa ter os seguintes pré-requisitos instalados em seu ambiente:

    Go: Certifique-se de ter o Go instalado. Você pode baixar a versão mais recente aqui.
    Docker: Para rodar o banco de dados MariaDB em um container, você precisará do Docker. Acesse Docker para instruções de instalação.
    Docker Compose: Para facilitar a orquestração de múltiplos containers, é recomendável usar o Docker Compose.
    Postman ou outra ferramenta similar pra fazer os requests obrigatório!

Configuração do Ambiente
1. Clone o Repositório

Clone o repositório da API para sua máquina local:


git clone git@github.com:wellminozzo/desafio-be-fr.git
cd seu-repositorio

feito isso execute o comando docker-compose build
depois que for feito a build você pode rodar o comando docker-compose ps
ele vai listar os container e se tudo estiver certo vai está rodando nosso container

você pode verificar no arquivo docker-compose.yml a porta do servidor

feito isso vai ter um arquivo chamado "a.json"
neste arquivo contem o json de entrada pra fazer a requisição post

no Insomnia fazer começar um novo HTTP request e escolher o POST
em seguida colocar a url http://localhost:9090/quote e no body o json do arquivo a.json

feito isso vc vai receber a reposta da requisição
com essa reposta ja podemos fazer um GET

http://localhost:9090/metrics?last_quotes=10

http://localhost:9090/metrics/carrierprice

http://localhost:9090/metrics/cheaper

http://localhost:9090/metrics/expensive

seguinda essas rotas obtemos oque foi proposta na ROTA2 do desafio


## Documentação da API em Go
### Introdução

Esta documentação fornece instruções para configurar e rodar a API desenvolvida em Go. A API é uma aplicação simples que se conecta a um banco de dados MariaDB e oferece endpoints para manipulação de dados.
Pré-requisitos

Antes de iniciar, você precisará ter os seguintes componentes instalados em seu ambiente:

    Go: Certifique-se de ter a versão mais recente instalada.
    Docker: Necessário para rodar o banco de dados MariaDB em um container.
    Docker Compose: Recomendado para facilitar a orquestração de múltiplos containers.
    Postman ou outra ferramenta similar: Para realizar requisições HTTP.

Configuração do Ambiente
1. Clone o Repositório

Clone o repositório da API para sua máquina local:

bash

git clone git@github.com:wellminozzo/desafio-be-fr.git
cd desafio-be-fr

2. Build e Execução do Docker

Execute o comando para construir os containers:

bash

docker-compose build

Após a build, verifique se os containers estão rodando:

bash

docker-compose ps

3. Verificação da Porta do Servidor

Você pode verificar a porta do servidor no arquivo docker-compose.yml.
4. Configuração da Requisição POST

    Localize o arquivo a.json, que contém o JSON de entrada para a requisição POST.
    Abra o Insomnia ou sua ferramenta de requisições preferida.
    Crie um novo HTTP request, selecionando o método POST.
    Insira a URL: http://localhost:9090/quote.
    No corpo da requisição, insira o JSON do arquivo a.json.

5. Realizando Requisições GET

Após receber a resposta da requisição POST, você pode realizar as seguintes requisições GET:

    Últimas Cotações:
    http://localhost:9090/metrics?last_quotes=10

    Preço do Transportador:
    http://localhost:9090/metrics/carrierprice

    Cotações Mais Baratas:
    http://localhost:9090/metrics/cheaper

    Cotações Mais Caras:
    http://localhost:9090/metrics/expensive

Seguindo essas rotas, você obterá os resultados propostos na ROTA2 do desafio.
Conclusão

Siga essas instruções para configurar e rodar sua API em Go. Para mais detalhes sobre a implementação, consulte o código-fonte no repositório.

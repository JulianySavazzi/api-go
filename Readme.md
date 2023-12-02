# API REST EM GO - Live Full Cycle

## Este projeto foi criado com base na live esquenta Full Cycle - APIs e Mensageria com Golang, meu objetivo com ele era estudar sobre APIS REST e para conhecer novas tecnologias - Docker, Golang e Kafka.

### Resumo sobre a live da Full Cycle e anotações sobre o projeto:

A Golang é uma linguagem hexagonal, o objetivo desse projeto era criar uma aplicação que consegue se conectar com adaptadores, sem modificar o seu núcleo (evitar acoplamento);

Tecnologias utilizadas: Golang, BD MySQL, Docker, Mensageria com Apache Kafka;

#### Estrutura do projeto
Pastas: internal (tudo que acontece dentro da aplicação, e somente ela tem acesso), pkg (pacotes que podem ser utilizados por outras aplicações). 
Internal: entity (dominio e regras de dominio, entidades), infra(repository, conexão com banco de dados), usecase(ações do sistema, ações do usuario, orquestraçao, camada de aplicação);

cmd -> app -> onde vamos rodar o executável da aplicação;





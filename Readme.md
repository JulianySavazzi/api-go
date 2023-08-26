API REST EM GO

criando aplicacao que consegue se conectar com adaptadores, sem modificar o nucleo da aplicacao;

estrutura do projeto -> pastas: internal (tudo que acontece dentro da aplicacao, somente ela tem acesso), pkg (pacotes que podem ser utilizados por outras aplicacoes);
 internal -> entity (dominio e regras de dominio, entidades), infra(repository, conexao com banco de dados), usecase(acoes do sistema, acoes do usuario, orquestracao, camada de aplicacao);
cmd -> app -> onde vamos rodar o executavel da aplicacao;



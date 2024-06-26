# OCI Novo Backend

Nova versão da backend da Olímpiada Cearence de Informática.

## Ambiente

Para fazer o setup ambiente de desenvolvimento primeiramente deve ser instalado as seguintes dependências:

- [Docker](https://www.docker.com/) e Docker Compose (normalmente vem junto com o pacote do docker)

Após a instalação, você deve criar um arquivo `.env` na raiz do projeto com os
seguintes atributos:

| Nome               | Descrição                                   |
| ------------------ | ------------------------------------------- |
| `DB_ADMIN_USER`    | Nome do administrador do banco              |
| `DB_ADMIN_PWD`     | Senha do administrador do banco             |
| `PETIANO_USER_PWD` | Senha do usuario de cargo "petiano"         |
| `ESCOLA_USER_PWD`  | Senha do usuario de cargo "escola"          |
| `ALUNO_USER_PWD`   | Senha do usuario de cargo "aluno"           |
| `API_USER_PWD`     | Senha do usuario usado para funções da API |

Aqui está um exemplo deste arquivo `.env`:

```.env
DB_ADMIN_USER="db_admin"
DB_ADMIN_PWD="senha_exemplo"
PETIANO_USER_PWD="senha_exemplo"
ESCOLA_USER_PWD="senha_exemplo"
ALUNO_USER_PWD="senha_exemplo"
API_USER_PWD="senha_exemplo"
```

## Rodar

Para iniciar a aplicação em, execute o comando:

```bash
$ docker compose up
```

ou, caso queira rodar em background:

```bash
$ docker compose up -d
```

No momento apenas uma menssagem de _"hello world"_ é dada como resposta na porta
`8080`. Assim, para verificar se a aplicação está funcionando rode:

```bash
$ curl localhost:8080
```

O banco será disponibilizado apenas dentro do container docker `db`, desse
modo se necessário entrar diretamente pelo psql o seguinte comando deve ser
executado:

```bash
$ docker exec -it oci-novo-back-db-1 psql -U <usuario> -h localhost
```

O armazenamento do banco é armazenado em um volume docker de forma persistente
e, portanto, para que os arquivos de inicilização do banco em `db/setup` sejam
executados você deve excluir o volume atual do banco:

```bash
$ docker volume rm oci-novo-back_pgdata
```

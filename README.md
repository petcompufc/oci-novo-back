# OCI Novo Backend

Nova versão da backend da Olímpiada Cearence de Informática.

## Rodar

Primeiramente deve ser instalado as seguintes dependências:

- [Docker](https://www.docker.com/) e Docker Compose (normalmente vem junto com o pacote do docker)


Para iniciar a aplicação em, execute o comando:

```
$ docker compose up
```

ou, caso queira rodar em background:

```
$ docker compose up -d
```

No momento apenas uma menssagem de _"hello world"_ é dada como resposta na porta
`8080`. Assim, para verificar se a aplicação está funcionando rode:

```
$ curl localhost:8080
```

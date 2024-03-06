-- create databases
CREATE DATABASE oci_dados;
\c oci_dados;

-- tipos de dado
CREATE TYPE cargo AS ENUM ('petiano', 'aluno', 'escola');

CREATE TYPE serie AS ENUM ('6f', '7f', '8f', '9f', '1m', '2m', '3m');

CREATE TYPE genero AS ENUM (
  'masculino',
  'feminino',
  'nao-binario',
  'genero-fluido',
  'agenero',
  'bigenero',
  'transgenero',
  'outro'
);

CREATE TYPE perfil_acess AS ENUM (
  'fisica',
  'sensorial',
  'cognitiva/intelectual',
  'psiquiatrica',
  'invisivel'
);

CREATE TYPE rede_escola AS ENUM ('particular', 'publica');

CREATE TYPE oci_modalidade AS ENUM ('iniciacao A', 'iniciacao B', 'programacao');

CREATE TYPE oci_fase AS ENUM ('fase 1', 'fase 2', 'fase 3');

-- tabela dos usuarios
CREATE TABLE IF NOT EXISTS usuario (
  id_usuario INT GENERATED ALWAYS AS IDENTITY,
  hash_senha BYTEA NOT NULL,
  cargo cargo,
  ultimo_login TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY (id_usuario)
);

CREATE TABLE IF NOT EXISTS petiano (
  id_petiano INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  login VARCHAR(32) NOT NULL,
  PRIMARY KEY (id_petiano),
  CONSTRAINT fk_petiano_global FOREIGN KEY (id_global) REFERENCES usuario (id_usuario) ON DELETE CASCADE,
  UNIQUE (id_global),
  UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS endereco (
  id_endereco INT GENERATED ALWAYS AS IDENTITY,
  cep CHAR(9) NOT NULL,
  bairro VARCHAR(46) NOT NULL,
  cidade VARCHAR(46) NOT NULL,
  estado CHAR(2) NOT NULL,
  rua VARCHAR(46) NOT NULL,
  numero SMALLINT NOT NULL,
  complemento VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY(id_endereco)
);

CREATE TABLE IF NOT EXISTS escola (
  id_escola INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  cod_inep CHAR(12) NOT NULL,
  nome VARCHAR(63) NOT NULL,
  id_endereco INT NOT NULL,
  email VARCHAR(100) NOT NULL,
  telefone VARCHAR(13) NOT NULL,
  nome_coordenador VARCHAR(63) NOT NULL,
  email_coordenador VARCHAR(100) NOT NULL,
  telefone_coordenador VARCHAR(13),
  tipo_escola rede_escola NOT NULL,
  PRIMARY KEY (id_escola),
  CONSTRAINT fk_ecola_global FOREIGN KEY (id_global) REFERENCES usuario (id_usuario) ON DELETE CASCADE,
  CONSTRAINT fk_ecola_endereco FOREIGN KEY (id_endereco) REFERENCES endereco (id_endereco) ON DELETE CASCADE,
  UNIQUE (id_global),
  UNIQUE(cod_inep)
);

CREATE TABLE IF NOT EXISTS aluno (
  id_aluno INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  id_escola INT NOT NULL, 
  cpf CHAR(14) NOT NULL,
  nome VARCHAR(63),
  serie_atual serie NOT NULL,
  genero genero NOT NULL,
  data_nasc DATE NOT NULL,
  perfis_acess perfil_acess[] NOT NULL DEFAULT ARRAY[]::perfil_acess[],
  PRIMARY KEY (id_aluno),
  CONSTRAINT fk_aluno_global FOREIGN KEY (id_global) REFERENCES usuario (id_usuario) ON DELETE CASCADE,
  CONSTRAINT fk_aluno_escola FOREIGN KEY (id_escola) REFERENCES escola (id_escola),
  UNIQUE (id_global),
  UNIQUE (cpf)
);

-- tabelas da olimpiada
CREATE TABLE IF NOT EXISTS edicao (
  id_edicao INT GENERATED ALWAYS AS IDENTITY,
  edicao VARCHAR(7),
  abertura_das_inscricoes DATE NOT NULL,
  meta_arrecadacao BIGINT NOT NULL DEFAULT 0,
  arrecadacao BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (id_edicao),
  UNIQUE(edicao)
);

CREATE TABLE IF NOT EXISTS inscricao_aluno (
  id_aluno INT NOT NULL,
  id_edicao INT NOT NULL,
  modalidade oci_modalidade NOT NULL,
  CONSTRAINT fk_aluno_inscrito FOREIGN KEY (id_aluno) REFERENCES aluno (id_aluno),
  CONSTRAINT fk_inscricao_edicao FOREIGN KEY (id_edicao) REFERENCES edicao (id_edicao),
  PRIMARY KEY (id_aluno, id_edicao)
);

CREATE TABLE IF NOT EXISTS aluno_fase (
  id_edicao INT NOT NULL,
  id_aluno INT NOT NULL,
  fase oci_fase NOT NULL,
  acertos BOOLEAN[] DEFAULT NULL,
  CONSTRAINT fk_aluno_fase FOREIGN KEY (id_aluno) REFERENCES aluno (id_aluno),
  CONSTRAINT fk_edicao_fase FOREIGN KEY (id_edicao) REFERENCES edicao (id_edicao),
  PRIMARY KEY(id_edicao, id_aluno)
);

CREATE TABLE IF NOT EXISTS inscricao_escola (
  id_edicao INT NOT NULL,
  id_escola INT NOT NULL,
  valor_pago BIGINT NOT NULL DEFAULT 0, 
  CONSTRAINT fk_escola_pagante FOREIGN KEY (id_escola) REFERENCES escola (id_escola),
  CONSTRAINT fk_edicao_escola FOREIGN KEY (id_edicao) REFERENCES edicao (id_edicao),
  PRIMARY KEY (id_escola, id_edicao)
);

-- oci_dados roles
CREATE ROLE ver_aluno;
GRANT SELECT ON aluno TO ver_aluno;

CREATE ROLE ver_escola;
GRANT SELECT ON escola TO ver_escola;
GRANT SELECT ON endereco TO ver_escola;

CREATE ROLE ver_petiano;
GRANT SELECT ON petiano TO ver_petiano;

CREATE ROLE ver_usuario;
GRANT SELECT ON usuario TO ver_usuario;
GRANT ver_aluno TO ver_usuario;
GRANT ver_escola TO ver_usuario;
GRANT ver_petiano TO ver_usuario;

CREATE ROLE ver_metadados;
GRANT SELECT ON edicao TO ver_metadados;
GRANT SELECT ON inscricao_aluno TO ver_metadados;
GRANT SELECT ON aluno_fase TO ver_metadados;
GRANT SELECT ON inscricao_escola TO ver_metadados;

CREATE ROLE gerenciar_aluno;
GRANT UPDATE ON aluno TO gerenciar_aluno;

CREATE ROLE gerenciar_escola;
GRANT UPDATE ON escola TO gerenciar_escola;
GRANT UPDATE ON endereco TO gerenciar_escola;

CREATE ROLE gerenciar_metadados;
GRANT UPDATE ON edicao TO gerenciar_metadados;
GRANT UPDATE ON inscricao_aluno TO gerenciar_metadados;
GRANT UPDATE ON inscricao_escola TO gerenciar_metadados;
GRANT UPDATE ON aluno_fase TO gerenciar_metadados;

CREATE ROLE inscrever_aluno;
GRANT INSERT ON inscricao_aluno TO inscrever_aluno;

CREATE ROLE inscrever_escola;
GRANT INSERT ON inscricao_escola TO inscrever_escola;

CREATE ROLE adicionar_usuario;
GRANT INSERT ON usuario TO adicionar_usuario;
GRANT INSERT ON petiano TO adicionar_usuario;
GRANT INSERT ON aluno TO adicionar_usuario;
GRANT INSERT ON escola TO adicionar_usuario;

CREATE ROLE inserir_metadados;
GRANT INSERT ON edicao TO ver_metadados;
GRANT INSERT ON aluno_fase TO ver_metadados;

-- atribuir para os usuários
GRANT ver_usuario TO api_user;
GRANT adicionar_usuario TO api_user;

GRANT ver_aluno TO aluno_user;
GRANT gerenciar_aluno TO aluno_user;
GRANT inscrever_aluno TO aluno_user;

GRANT ver_escola TO escola_user;
GRANT gerenciar_escola TO escola_user;
GRANT ver_aluno TO escola_user;
GRANT inscrever_escola TO escola_user;

GRANT ver_usuario TO petiano_user;
GRANT ver_metadados TO petiano_user;
GRANT inserir_metadados TO petiano_user;

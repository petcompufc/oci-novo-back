-- users
CREATE USER petiano_user;

CREATE USER aluno_user;

CREATE USER escola_user;

CREATE USER anon_user;

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
  cargo cargo,
  ultimo_login TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY (id_usuario),
);

CREATE TABLE IF NOT EXISTS petiano (
  id_petiano INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  login VARCHAR(32) NOT NULL,
  hash_senha BYTEA NOT NULL,
  PRIMARY KEY (id_petiano),
  CONSTRAINT fk_petiano_global FOREIGN KEY (id_global) REFERENCES usuario (id_usuario) ON DELETE CASCADE,
  UNIQUE (id_global),
  UNIQUE (login),
);

CREATE TABLE IF NOT EXISTS uluno (
  id_aluno INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  id_escola INT NOT NULL, 
  cpf CHAR(14) NOT NULL,
  nome VARCHAR(63),
  hash_senha BYTEA NOT NULL,
  serie_atual serie NOT NULL,
  genero genero NOT NULL,
  data_nasc DATE NOT NULL,
  perfis_acess perfil_acess[] NOT NULL DEFAULT array[],
  PRIMARY KEY (id_aluno),
  CONSTRAINT fk_aluno_global FOREIGN KEY (id_global) REFERENCES usuario (id_usuario) ON DELETE CASCADE,
  CONSTRAINT fk_aluno_escola FOREIGN KEY (id_escola) REFERENCES escola (id_escola),
  UNIQUE (id_global),
  UNIQUE (cpf),
);

CREATE TABLE IF NOT EXISTS escola (
  id_escola INT GENERATED ALWAYS AS IDENTITY,
  id_global INT NOT NULL,
  hash_senha BYTEA NOT NULL,
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
  UNIQUE(cod_inep),
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
    PRIMARY KEY(id_endereco),
);

-- dados da edicao atual
CREATE TABLE IF NOT EXISTS alunos_inscritos (
    id_aluno INT,
    modalidade oci_modalidade NOT NULL,
    CONSTRAINT fk_aluno_inscrito FOREIGN KEY (id_aluno) REFERENCES aluno (id_aluno),
    PRIMARY KEY (id_aluno)
);

CREATE TABLE IF NOT EXISTS escolas_inscritas (
    id_escola INT,
    valor_pago BIGINT NOT NULL DEFAULT 0, 
    CONSTRAINT fk_escola_pagante FOREIGN KEY (id_escola) REFERENCES escola (id_escola),
    PRIMARY KEY (id_escola),
);

CREATE TABLE IF NOT EXISTS resultados_alunos (
    id_aluno INT NOT NULL,
    fase oci_fase NOT NULL,
    acertos BOOLEAN[],
    CONSTRAINT fk_resultado_aluno FOREIGN KEY (id_aluno) REFERENCES aluno (id_aluno),
    PRIMARY KEY (id_aluno, fase),
);

-- dados persistentes
CREATE TABLE IF NOT EXISTS edicao (
    id_edicao INT GENERATED ALWAYS AS IDENTITY,
    edicao VARCHAR(7),
    abertura_das_inscricoes DATE NOT NULL,
    meta_arrecadacao BIGINT NOT NULL DEFAULT 0,
    arrecadacao BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id_edicao),
    UNIQUE(edicao)
);

CREATE TABLE IF NOT EXISTS fase (
    id_edicao INT NOT NULL,
    fase oci_fase NOT NULL,
    qtd_pub_iniA INT NOT NULL DEFAULT 0,
    qtd_par_iniA INT NOT NULL DEFAULT 0,
    qtd_pub_iniB INT NOT NULL DEFAULT 0,
    qtd_par_iniB INT NOT NULL DEFAULT 0,
    qtd_pub_prog INT NOT NULL DEFAULT 0,
    qtd_par_prog INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id_edicao, fase),
    CONSTRAINT fk_edicao_fase FOREIGN KEY (id_edicao) REFERENCES aluno (id_edicao),
);

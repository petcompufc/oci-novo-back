\c oci_dados;

--inserindo valores na tabela usuário
INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
  	(decode('Senha_Petiano', 'escape'), 'petiano', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Petiano', 'escape'), 'petiano', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Petiano', 'escape'), 'petiano', '2024-03-08 00:00:00+00:00');
  
INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
	(decode('Senha_Aluno', 'escape'), 'aluno', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Aluno', 'escape'), 'aluno', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Aluno', 'escape'), 'aluno', '2024-03-08 00:00:00+00:00');

INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
	(decode('Senha_Escola', 'escape'), 'escola', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Escola', 'escape'), 'escola', '2024-03-08 00:00:00+00:00'),
  	(decode('Senha_Escola', 'escape'), 'escola', '2024-03-08 00:00:00+00:00');
	
--inserindo valores na tabela petiano
INSERT INTO petiano (id_global, login)
VALUES
	(1, 'login_petiano1'),
	(2, 'login_petiano2'),
	(3, 'login_petiano3');
	
--inserindo valores na tabela endereco
INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento)
VALUES
	('00000-000', 'bairro1', 'cidade1', 'CE', 'rua1', 123, 'apto'),
  	('11111-111', 'bairro2', 'cidade2', 'CE', 'rua2', 456, NULL),
  	('22222-222', 'bairro3', 'cidade3', 'CE', 'rua3', 789, 'casa');

--inserindo valores na tabela escola
INSERT INTO escola (id_global, cod_inep, nome, id_endereco, email, telefone, nome_coordenador, email_coordenador, telefone_coordenador, tipo_escola)
VALUES
	(7, '123456789012', 'Escola1', 1, 'escola1@example.com', '(11)1234-5678', 'coordenador1', 'coordenador1@example.com', '(11)9876-5432', 'publica'),
  	(8, '987654321098', 'Escola2', 2, 'escola2@example.com', '(21)9876-5432', 'coordenador2', 'coordenador2@example.com', '(21)1234-5678', 'particular'),
  	(9, '543210987654', 'Escola3', 3, 'escola3@example.com', '(41)1234-5678', 'coordenador3', 'coordenador3@example.com', '(41)9876-5432', 'publica');

--inserindo valores na tabela aluno
INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess)
VALUES
  	(4, 1, '12345678901', 'Aluno1', '6f', 'masculino', '2010-05-15', ARRAY['fisica']::perfil_acess[]),
  	(5, 2, '98765432109', 'Aluno2', '9f', 'feminino', '2009-08-20', ARRAY['sensorial']::perfil_acess[]),
  	(6, 3, '45678901234', 'Aluno3', '3m', 'masculino', '2007-03-10', ARRAY['psiquiatrica']::perfil_acess[]);


--inserindo valores nas tabelas da olimpíada
INSERT INTO edicao (edicao, abertura_das_inscricoes, meta_arrecadacao)
VALUES
	('2023', '2023-04-12', 16000),
	('2022', '2022-06-10', 16000),
	('2020', '2020-06-11', 16000);

INSERT INTO inscricao_aluno (id_aluno, id_edicao, modalidade)
VALUES
	(3, 1, 'programacao'),
	(3, 2, 'programacao'),
	(3, 3, 'iniciacao B'),
	(1, 1, 'iniciacao A'),
	(2, 1, 'iniciacao B'),
	(2, 2, 'iniciacao B'),
	(2, 3, 'iniciacao A');

INSERT INTO aluno_fase (id_edicao, id_aluno, fase, acertos)
VALUES 

	--ERRO: não tá dando de inserir linhas de cada fase de um mesmo aluno
	-- aluno 3 - programacao
	-- edição 2023
	(1, 3, 'fase 1', ARRAY[true, false, true]),
	(1, 3, 'fase 2', ARRAY[true, true, true]),
	(1, 3, 'fase 3', ARRAY[true, true, false]),

	-- edição 2022
	(2, 3, 'fase 1', ARRAY[true, false, true]),
	(2, 3, 'fase 2', ARRAY[true, true, true]),
	(2, 3, 'fase 3', ARRAY[true, true, false]),

	-- edição 2020
	(3, 3, 'fase 1', ARRAY[true, false, true]),
	(3, 3, 'fase 2', ARRAY[true, true, true]),
	
	-- aluno 1
	(1, 1, 'fase 1', ARRAY[true, true, true]),
	(1, 1, 'fase 2', ARRAY[false, true, true]),
	
	-- aluno 2 - iniciação - teste (não participam fase 3)
	(1, 2, 'fase 1', ARRAY[true, true, true]),
	(1, 2, 'fase 2', ARRAY[false, true, true]),
	(1, 2, 'fase 3', ARRAY[false, false, false]);


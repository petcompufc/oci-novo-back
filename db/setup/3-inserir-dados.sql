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
	(21, 'login_petiano1'),
	(22, 'login_petiano2'),
	(23, 'login_petiano3');

--inserindo valores na tabela escola
INSERT INTO escola (id_global, cod_inep, nome, id_endereco, email, telefone, nome_coordenador, email_coordenador, telefone_coordenador, tipo_escola)
VALUES
	(27, '123456789012', 'Escola1', 1, 'escola1@example.com', '(11)1234-5678', 'coordenador1', 'coordenador1@example.com', '(11)9876-5432', 'publica'),
  	(28, '987654321098', 'Escola2', 2, 'escola2@example.com', '(21)9876-5432', 'coordenador2', 'coordenador2@example.com', '(21)1234-5678', 'particular'),
  	(29, '543210987654', 'Escola3', 3, 'escola3@example.com', '(41)1234-5678', 'coordenador3', 'coordenador3@example.com', '(41)9876-5432', 'publica');
	
--inserindo valores na tabela endereco
INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento)
VALUES
	('00000-000', 'bairro1', 'cidade1', 'CE', 'rua1', 123, 'apto'),
  	('11111-111', 'bairro2', 'cidade2', 'CE', 'rua2', 456, NULL),
  	('22222-222', 'bairro3', 'cidade3', 'CE', 'rua3', 789, 'casa');

--inserindo valores na tabela aluno
INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess)
VALUES
  	(24, 1, '12345678901', 'Aluno1', '6f', 'masculino', '2010-05-15', ARRAY['fisica']::perfil_acess[]),
  	(25, 2, '98765432109', 'Aluno2', '9f', 'feminino', '2009-08-20', ARRAY['sensorial']::perfil_acess[]),
  	(26, 3, '45678901234', 'Aluno3', '3m', 'masculino', '2007-03-10', ARRAY['psiquiatrica']::perfil_acess[]);
	

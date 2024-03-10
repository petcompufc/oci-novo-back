\c oci_dados;

--inserindo valores na tabela usuário
INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
  	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-08 00:00:00+00:00'),-- id: 1
  	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-08 00:00:00+00:00'),-- id: 2
	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-09 00:00:00+00:00'),-- id: 3
	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-09 00:00:00+00:00'),-- id: 4 
	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-09 00:00:00+00:00'),-- id: 5
	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-10 00:00:00+00:00'),-- id: 6
	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-10 00:00:00+00:00'),-- id: 7
  	(decode('senha_petiano', 'escape'), 'petiano', '2024-03-10 00:00:00+00:00');-- id: 8
  
INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-08 00:00:00+00:00'),-- id: 9
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-08 00:00:00+00:00'),-- id: 10
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-09 00:00:00+00:00'),-- id: 11
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-09 00:00:00+00:00'),-- id: 12
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-09 00:00:00+00:00'),-- id: 13
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 14
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 15
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 16
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 17
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 18
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 19
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 20
	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 21
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00'),-- id: 22
  	(decode('senha_aluno', 'escape'), 'aluno', '2024-03-10 00:00:00+00:00');-- id: 23

INSERT INTO usuario (hash_senha, cargo, ultimo_login)
VALUES
	(decode('senha_escola', 'escape'), 'escola', '2024-03-08 00:00:00+00:00'),-- id: 24
  	(decode('senha_escola', 'escape'), 'escola', '2024-03-08 00:00:00+00:00'),-- id: 25
	(decode('senha_escola', 'escape'), 'escola', '2024-03-09 00:00:00+00:00'),-- id: 26
  	(decode('senha_escola', 'escape'), 'escola', '2024-03-09 00:00:00+00:00'),-- id: 27
  	(decode('senha_escola', 'escape'), 'escola', '2024-03-09 00:00:00+00:00');-- id: 28
	
--inserindo valores na tabela petiano
INSERT INTO petiano (id_global, login)
VALUES
	(1, 'login_petiano1'),
	(2, 'login_petiano2'),
	(3, 'login_petiano3'),
	(4, 'login_petiano4'),
	(5, 'login_petiano5'),
	(6, 'login_petiano6'),
	(7, 'login_petiano7'),
	(8, 'login_petiano8');
	
--inserindo valores na tabela endereco
INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento)
VALUES
	('00000-000', 'bairro1', 'cidade1', 'CE', 'rua1', 123, 'próximo ao parque eolico'),
  	('11111-111', 'bairro2', 'cidade2', 'CE', 'rua2', 456, 'ao lado da fabrica'),
	('22222-222', 'bairro3', 'cidade3', 'CE', 'rua3', 789, 'proximo a subestação de energia eletrica'),
  	('33333-333', 'bairro4', 'cidade4', 'CE', 'rua4', 012, 'proximo ao presidio'),
  	('44444-444', 'bairro5', 'cidade5', 'CE', 'rua5', 345, 'prédio branco');

--inserindo valores na tabela escola
INSERT INTO escola (id_global, cod_inep, nome, id_endereco, email, telefone, nome_coordenador, email_coordenador, telefone_coordenador, tipo_escola)
VALUES
	(24, '123456789012', 'Escola1', 1, 'escola1@example.com', '85 91234-5678', 'coordenador1', 'coordenador1@example.com', '85 99876-5432', 'publica'),
  	(25, '987654321098', 'Escola2', 2, 'escola2@example.com', '85 98197-6619', 'coordenador2', 'coordenador2@example.com', '85 91234-5678', 'particular'),
	(26, '345678901234', 'Escola2', 3, 'escola3@example.com', '85 99667-8431', 'coordenador3', 'coordenador3@example.com', '85 95678-9012', 'particular'),
	(27, '456789012345', 'Escola2', 4, 'escola4@example.com', '85 99876-5432', 'coordenador4', 'coordenador4@example.com', '85 98901-2345', 'particular'),
  	(28, '543210987654', 'Escola3', 5, 'escola5@example.com', '85 91234-5678', 'coordenador5', 'coordenador5@example.com', '88 99876-5432', 'publica');

--inserindo valores na tabela aluno
INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess)
VALUES
  	(1, 1, '12345678901', 'Aluno1', '6f', 'masculino', '2013-05-15', ARRAY['fisica']::perfil_acess[]),
  	(2, 2, '98765432109', 'Aluno2', '9f', 'feminino', '2009-08-20', ARRAY['sensorial']::perfil_acess[]),
  	(3, 3, '45678901234', 'Aluno3', '3m', 'masculino', '2006-03-10', ARRAY['psiquiatrica']::perfil_acess[]),
	(4, 4, '01234567890', 'Aluno4', '6f', 'masculino', '2013-11-13', ARRAY['cognitiva/intelectual']::perfil_acess[]),
  	(5, 5, '34567890123', 'Aluno5', '9f', 'feminino', '2010-02-21', ARRAY['invisivel']::perfil_acess[]),
  	(6, 1, '78901234567', 'Aluno6', '3m', 'masculino', '2007-01-14', ARRAY['psiquiatrica']::perfil_acess[]),
	(7, 2, '09876543211', 'Aluno7', '8f', 'masculino', '2011-05-15', ARRAY['invisivel']::perfil_acess[]),
  	(8, 3, '13573486534', 'Aluno8', '9f', 'feminino', '2009-04-03', ARRAY['sensorial']::perfil_acess[]),
  	(9, 4, '34215687091', 'Aluno9', '3m', 'masculino', '2007-03-07', ARRAY['psiquiatrica']::perfil_acess[]),
	(10, 5, '55566677788', 'Aluno10', '7f', 'masculino', '2012-12-25', ARRAY['fisica']::perfil_acess[]),
  	(11, 1, '12309876543', 'Aluno11', '9f', 'feminino', '2009-08-20', ARRAY['fisica']::perfil_acess[]),
  	(12, 2, '12563490789', 'Aluno12', '3m', 'masculino', '2007-03-11', ARRAY['fisica']::perfil_acess[]),
	(13, 3, '00100200300', 'Aluno13', '6f', 'masculino', '2013-05-15', ARRAY['fisica']::perfil_acess[]),
  	(14, 4, '10210310411', 'Aluno14', '8f', 'feminino', '2011-06-27', ARRAY['sensorial']::perfil_acess[]),
  	(15, 5, '02112345690', 'Aluno15', '3m', 'masculino', '2007-11-10', ARRAY['fisica']::perfil_acess[]);


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

INSERT INTO inscricao_escola (id_edicao, id_escola, valor_pago)
VALUES
	(1, 2, 2000),
	(1, 1, 0),
	(1, 3, 0),

	(2, 2, 3500),
	(2, 1, 0),
	(2, 3, 0),

	(3, 2, 500),
	(3, 1, 0),
	(3, 3, 0);

INSERT INTO escola_para_analisar (id_escola)
VALUES 
	(1),
	(2),
	(3);

CREATE ROLE petiano_user LOGIN;

CREATE ROLE aluno_user LOGIN;

CREATE ROLE escola_user LOGIN;

CREATE ROLE api_user LOGIN;

-- oci_dados roles
CREATE ROLE ver_aluno;
GRANT SELECT ON oci_dados.aluno TO ver_aluno;

CREATE ROLE ver_escola;
GRANT SELECT ON oci_dados.escola TO ver_escola;
GRANT SELECT ON oci_dados.endereco TO ver_escola;

CREATE ROLE ver_petiano;
GRANT SELECT ON oci_dados.petiano TO ver_petiano;

CREATE ROLE ver_usuario;
GRANT SELECT ON oci_dados.usuario TO ver_usuario;
GRANT ver_aluno TO ver_usuario;
GRANT ver_escola TO ver_usuario;
GRANT ver_petiano TO ver_usuario;

CREATE ROLE ver_metadados;
GRANT SELECT ON oci_dados.edicao TO ver_metadados;
GRANT SELECT ON oci_dados.fase TO ver_metadados;
GRANT SELECT ON oci_dados.alunos_inscritos TO ver_metadados;
GRANT SELECT ON oci_dados.escolas_inscritas TO ver_metadados;
GRANT SELECT ON oci_dados.resultados_alunos TO ver_metadados;

CREATE ROLE gerenciar_aluno;
GRANT UPDATE ON oci_dados.aluno TO gerenciar_aluno;

CREATE ROLE gerenciar_escola;
GRANT UPDATE ON oci_dados.escola TO gerenciar_escola;
GRANT UPDATE ON oci_dados.endereco TO gerenciar_escola;

CREATE ROLE gerenciar_metadados;
GRANT UPDATE ON oci_dados.edicao TO gerenciar_metadados;
GRANT UPDATE ON oci_dados.alunos_inscritos TO gerenciar_metadados;
GRANT UPDATE ON oci_dados.escolas_inscritas TO gerenciar_metadados;
GRANT UPDATE ON oci_dados.resultados_alunos TO gerenciar_metadados;

CREATE ROLE inscrever_aluno;
GRANT INSERT ON oci_dados.alunos_inscritos TO inscrever_aluno;

CREATE ROLE inscrever_escola;
GRANT INSERT ON oci_dados.escolas_inscritas TO inscrever_escola;

CREATE ROLE adicionar_usuario;
GRANT INSERT ON oci_dados.usuario TO adicionar_usuario;
GRANT INSERT ON oci_dados.petiano TO adicionar_usuario;
GRANT INSERT ON oci_dados.aluno TO adicionar_usuario;
GRANT INSERT ON oci_dados.escola TO adicionar_usuario;

CREATE ROLE inserir_resultados;
GRANT INSERT ON oci_dados.resultados_alunos TO inserir_resultados;
-- TODO: oci_interface roles

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
GRANT gerenciar_metadados TO petiano_user;
GRANT inserir_resultados TO petiano_user;

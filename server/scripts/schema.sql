-- Tabela de usuários
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    perfil TEXT NOT NULL CHECK (perfil IN ('liberal', 'clinica_admin', 'supervisor', 'terapeuta')),
    clinica_id INTEGER,
    supervisor_id INTEGER,
    ativo BOOLEAN NOT NULL DEFAULT TRUE,
    data_inicio_inatividade TIMESTAMP,
    data_fim_inatividade TIMESTAMP,
    motivo_inatividade TEXT,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela de clínicas
CREATE TABLE clinicas (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    cnpj TEXT UNIQUE,
    endereco TEXT,
    telefone TEXT,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Relacionamentos
ALTER TABLE usuarios
    ADD CONSTRAINT fk_clinica FOREIGN KEY (clinica_id) REFERENCES clinicas(id) ON DELETE SET NULL;

ALTER TABLE usuarios
    ADD CONSTRAINT fk_supervisor FOREIGN KEY (supervisor_id) REFERENCES usuarios(id) ON DELETE SET NULL;
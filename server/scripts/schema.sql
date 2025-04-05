-- Tabela de usuários
CREATE TABLE IF NOT EXISTS usuarios (
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
CREATE TABLE IF NOT EXISTS clinicas (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    cnpj TEXT UNIQUE,
    endereco TEXT,
    telefone TEXT,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Relacionamentos
ALTER TABLE IF EXISTS usuarios
    ADD CONSTRAINT IF NOT EXISTS fk_clinica FOREIGN KEY (clinica_id) REFERENCES clinicas(id) ON DELETE SET NULL;

ALTER TABLE IF EXISTS usuarios
    ADD CONSTRAINT IF NOT EXISTS fk_supervisor FOREIGN KEY (supervisor_id) REFERENCES usuarios(id) ON DELETE SET NULL;

-- Usuários de teste (evita duplicação)
INSERT INTO clinicas (id, nome, cnpj, endereco, telefone)
VALUES (1, 'Clínica Modelo', '12345678000100', 'Rua das Terapias, 100', '(11) 99999-0000')
ON CONFLICT (id) DO NOTHING;

INSERT INTO usuarios (nome, email, senha, perfil, clinica_id, ativo)
VALUES ('Ana Supervisor', 'ana@exemplo.com',
        '$2a$10$yL6wspZAR7K96oFcvQv3Se0vLybYdYobWnHPYkkyVgJ0GMUgXEsr2',
        'supervisor', 1, TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO usuarios (nome, email, senha, perfil, clinica_id, supervisor_id, ativo)
VALUES ('Carlos Terapeuta', 'carlos@exemplo.com',
        '$2a$10$yL6wspZAR7K96oFcvQv3Se0vLybYdYobWnHPYkkyVgJ0GMUgXEsr2',
        'terapeuta', 1, 1, TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO usuarios (nome, email, senha, perfil, clinica_id, ativo)
VALUES ('Fernanda Admin', 'admin@exemplo.com',
        '$2a$10$yL6wspZAR7K96oFcvQv3Se0vLybYdYobWnHPYkkyVgJ0GMUgXEsr2',
        'clinica_admin', 1, TRUE)
ON CONFLICT (email) DO NOTHING;
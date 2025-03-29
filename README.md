# MSD - Sistema de Supervisão Terapêutica para TEA

Este repositório contém o sistema completo voltado para acompanhamento e supervisão de terapias aplicadas a pessoas com Transtorno do Espectro Autista (TEA), com suporte para psicólogos, supervisores e clínicas.

---

## 📁 Estrutura do Projeto

```
MSD/
├── server/              # Backend em Go (API REST + JWT + PostgreSQL)
├── msd-tea-terapy/      # Frontend em React (Dashboard e interface de uso)
```

---

## 🔧 Backend (Go) – `server/`

### ✔ Requisitos

- Go 1.21+
- Docker (opcional, para subir banco PostgreSQL)
- PostgreSQL 15+

### ▶ Como rodar localmente

```bash
cd server/
go mod tidy
go build -o msd-therapy-api.exe
./msd-therapy-api.exe
```

### 🔐 Variáveis de ambiente (.env)

Exemplo de `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=msdtea
JWT_SECRET=chave_super_secreta
APP_PORT=9000
```

### 📦 Banco de Dados

Estrutura em PostgreSQL disponível em `/server/scripts/schema.sql` (em construção).

---

## 💻 Frontend (React) – `msd-tea-terapy/`

### ✔ Requisitos

- Node.js 18+
- Yarn ou npm

### ▶ Como rodar

```bash
cd msd-tea-terapy/
npm install
npm start
```

### 🌐 URL padrão
```bash
http://localhost:3000
```

---

## 📌 Funcionalidades Planejadas

- Cadastro de clínica ou profissional liberal
- Cadastro hierárquico (admin → supervisor → terapeuta AT)
- Login com autenticação JWT
- Dashboard de acompanhamento
- Controle de permissões
- Registro de rotinas terapêuticas
- Inativação temporária de usuários

---

## 📘 Licença

Projeto privado e em desenvolvimento por [@ailtonbrc](https://github.com/ailtonbrc). Direitos reservados.
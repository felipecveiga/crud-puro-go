# 📘 CRUD de Usuários em Golang Puro com MongoDB

Este projeto é um **CRUD (Create, Read, Update, Delete)** de usuários desenvolvido em **Golang puro**, com conexão ao banco de dados **MongoDB**, sem o uso de frameworks.

O objetivo principal é aprofundar o conhecimento na linguagem Go, explorando diretamente sua sintaxe, boas práticas e comportamento sem abstrações externas. Ideal como projeto de **estudo**.

## 🚀 Tecnologias Utilizadas

- [Go (Golang)](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

## 🧠 Objetivos do Projeto

- Aprender e praticar Go puro sem o uso de frameworks.
- Compreender melhor como a linguagem se comporta em aplicações reais.
- Manipular um banco NoSQL com Go, utilizando o driver oficial do MongoDB.
- Criar um CRUD completo e funcional.
- Servir como base de estudo.

## 📌 Funcionalidades Implementadas

✅ Criar usuário (Create)

✅ Listar usuário (Read)

✅ Listar usuários (Read) — em desenvolvimento

❌ Atualizar dados do usuário (Update) — em desenvolvimento

✅ Remover usuário (Delete) — em desenvolvimento

A cada funcionalidade finalizada, este README será atualizado.

### 📂 Estrutura de Pastas

```bash
.
├── config/
│   └── db/
│       └── db.go           # Configuração do cliente MongoDB
│   └── errs/
│       └── error.go        # Define os erros
├── handler/
│   └── user.go             # Manipula as requisições (entrada/saída)
├── model/
│   └── user.go             # Define os dados e structs (User, Residence)
├── repository/
│   └── user.go             # Interações com o banco de dados
├── service/
│   └── user.go             # Regras de negócio (validações e lógica)
├── main.go                 # Ponto de entrada da aplicação
├── go.mod                  # Dependências do projeto
├── go.sum                  # Sums de verificação
├── Makefile                # Comando make run para rodar o projeto
└── README.md               # Documentação
```

## 📝 Observações
 - Este projeto não usa frameworks para fins didáticos, permitindo o máximo de contato com o funcionamento nativo da linguagem.

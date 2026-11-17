# Task API
### API de tarefas com Go, Gin e SQLite

API REST para criação e gerenciamento de tarefas, organizada em camadas de controller, service e repository.

## Tecnologias

**Go 1.25 · Gin · GORM · SQLite**

## Recursos

- Criar, listar, consultar e excluir tarefas.
- Editar tarefas e atualizar seu status.
- Persistir dados em SQLite.
- Criar a estrutura da tabela a partir do modelo com GORM.

## Executar localmente

Pré-requisitos: Go 1.25 e ambiente compatível com CGO para o driver SQLite, incluindo compilador C.

```bash
git clone https://github.com/JoaoMFserafim/Go-api.git
cd Go-api/backend
go mod download
go run .
```

O servidor usa `http://localhost:8888` e o arquivo `banco.db` no diretório de execução.

## Endpoints

| Método | Rota | Operação |
| --- | --- | --- |
| POST | `/task` | Criar tarefa |
| GET | `/task` | Listar tarefas |
| GET | `/task/:id` | Consultar tarefa |
| PUT | `/task/:id` | Editar tarefa |
| PUT | `/task/:id/status` | Atualizar status |
| DELETE | `/task/:id` | Excluir tarefa |

Para os formatos de requisição usados no projeto, consulte [test.http](backend/test.http).

Exemplo de consulta:

```bash
curl http://localhost:8888/task
```

## Estrutura

```text
backend/
├── controller/  # Entrada e resposta HTTP
├── models/      # Modelo de dados
├── repository/  # Acesso ao banco
├── service/     # Camada de serviço
└── main.go      # Banco, dependências e rotas
```

## Estado atual

Projeto de estudo de back-end. A configuração atual usa CORS permissivo e não configura autenticação no ponto de entrada; a publicação para uso real exige revisar esses controles.

---

Desenvolvido por [JoaoMFserafim](https://github.com/JoaoMFserafim).

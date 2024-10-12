# Projeto Teste

Este projeto é uma API escrita em Go que faz uso do banco de dados SQLite. O projeto foi estruturado com repositórios genéricos, autenticação JWT, e uso do framework Echo.

## Requisitos

Antes de começar, você precisará instalar as seguintes dependências:

- [Docker](https://docs.docker.com/get-docker/) (recomendado para rodar o projeto em container)
- [Docker Compose](https://docs.docker.com/compose/install/) (opcional, se preferir rodar com `docker-compose`)
- Go 1.23 ou superior (caso queira rodar o projeto localmente fora do Docker)
- GCC (caso esteja rodando localmente, o Go com CGO habilitado necessita do GCC)

### Dependências Go

Se você estiver rodando o projeto localmente (fora de containers Docker), você precisará instalar as dependências do Go:

```bash
go mod tidy
```

Este comando vai baixar todas as dependências e garantir que o ambiente esteja configurado corretamente.

## Configuração do Projeto

### 1. Estrutura do Projeto

Aqui está um breve resumo da estrutura do projeto:

```bash
.
├── cmd/                   # Código fonte da aplicação
├── internal/              # Códigos internos (repositórios, serviços, etc)
│   ├── adapters/          # Adaptadores de banco de dados e HTTP
│   ├── domain/            # Lógica de negócios e serviços de domínio
│   ├── ports/             # Interfaces de repositórios e serviços
├── Dockerfile             # Arquivo Docker para build e execução do projeto
├── Makefile               # Arquivo Make para facilitar a execução de tarefas
└── go.mod                 # Gerenciador de dependências do Go
```

### 2. Ambiente de Desenvolvimento Local

Se você deseja rodar o projeto localmente, siga estas instruções:

#### Instalar Dependências Localmente

1. Instale o Go (1.23 ou superior).
2. Instale o GCC para compilar o SQLite com CGO:

   - Para sistemas baseados em Debian (como Ubuntu):

     ```bash
     sudo apt-get update
     sudo apt-get install build-essential
     ```

   - Para sistemas baseados em Arch Linux:

     ```bash
     sudo pacman -S base-devel
     ```

3. Baixe as dependências Go:

   ```bash
   go mod tidy
   ```

#### Executar o Projeto Localmente

Para rodar o projeto localmente, após instalar as dependências, execute:

```bash
go run cmd/main.go
```

O projeto rodará no endereço: `http://localhost:8080`

### 3. Executar o Projeto com Docker

#### Build da Imagem Docker

Para compilar a imagem Docker, execute o seguinte comando no diretório raiz do projeto:

```bash
docker build -t teste-app .
```

Este comando irá gerar uma imagem Docker chamada `teste-app`.

#### Rodar o Container

Depois de gerar a imagem, você pode rodar o container:

```bash
docker run -p 8080:8080 teste-app
```

O projeto estará disponível em: `http://localhost:8080`

### 4. Rodar com Docker Compose (opcional)

Se você preferir usar Docker Compose, crie um arquivo `docker-compose.yml` semelhante ao exemplo abaixo:

```yaml
version: '3'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APPNAME=teste
```

Depois, você pode iniciar o projeto com:

```bash
docker-compose up
```

### 5. Uso do Makefile

Um `Makefile` foi incluído para facilitar a execução de algumas tarefas comuns, como o build e a execução do projeto.

#### Comandos Disponíveis:

- **Build**: Para compilar a aplicação:

  ```bash
  make build
  ```

- **Run**: Para rodar o projeto localmente:

  ```bash
  make run
  ```

- **Clean**: Para remover os arquivos de build:

  ```bash
  make clean
  ```

### 6. Banco de Dados

O projeto usa SQLite como banco de dados. O banco de dados será automaticamente criado e inicializado ao rodar o projeto pela primeira vez. Você não precisa se preocupar com a criação manual de tabelas.

### 7. Rotas Disponíveis

Abaixo estão as rotas configuradas no projeto:

- **GET** `/ping`: Verifica o status da API.
- **POST** `/users`: Registra um novo usuário.
- **POST** `/login`: Autentica um usuário com email e senha e retorna um token JWT.
- **POST** `/recover`: Autentica um usuário com email e retorna um token JWT de recuperação.
- **GET** `/users`: Lista todos os usuários (requer autenticação JWT).

### 8. Autenticação JWT

O projeto utiliza JWT (JSON Web Tokens) para autenticação. A rota `/login` gera um token JWT, e para acessar as rotas protegidas, o token deve ser enviado no cabeçalho `Authorization` com o formato `Bearer <token>`.

### 9. Erros Comuns

- **CGO Error**: Se você ver erros relacionados ao CGO, certifique-se de que o GCC esteja instalado e de que a variável de ambiente `CGO_ENABLED=1` está configurada corretamente ao compilar o projeto.
- **Banco de Dados Não Encontrado**: Se o SQLite não puder encontrar o banco de dados, verifique se o caminho está correto e se a tabela foi criada.

---
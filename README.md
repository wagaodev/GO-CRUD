# CRUD Tweet API with Go

### 🇺🇸 The README of the project is translated into two languages, I will keep this custom to explain in both languages the purpose of the project and the explanation about the implementation, I hope you like it 🥹

### 🇧🇷 O README do projeto está traduzido em dois idiomas, irei manter este costume para explicar em ambos idiomas o propósito do projeto e a explicação sobre a implementação, espero que gostem 🥹

🇺🇸 This is a simple API that allows you to create, fetch, and delete tweets. It's built using the Go programming language and the Gin framework.

🇧🇷 Esta é uma API simples que permite criar, buscar e deletar tweets. Ela foi construída utilizando a linguagem de programação Go e o framework Gin.

## 🚀 Project Structure

🇺🇸 The project is divided into three main parts:

1. `controllers`: 🇺🇸 This is where the main logic of the API resides. We have a controller for handling the CRUD operations of tweets.
   / 🇧🇷 Aqui é onde a lógica principal da API reside. Nós temos um controller para lidar com as operações de CRUD dos tweets.
2. `entities`: 🇺🇸 This is where the data models are defined. In this case, we have the Tweet model.
   / 🇧🇷 Aqui é onde os modelos de dados são definidos. Neste caso, temos o modelo de Tweet.
3. `routes`: 🇺🇸 This is where the API routes are defined. We are using Gin's router to handle the routes.
   / 🇧🇷 Aqui é onde as rotas da API são definidas. Estamos utilizando o router do Gin para lidar com as rotas.

### 🔧 Installation

🇺🇸 First, clone this repository to your local machine:
🇧🇷 Primeiro, clone este repositório para a sua máquina local:

```bash
git clone https://github.com/wagaodev/GO-CRUD
```

🇺🇸 Then, navigate into the project directory:
🇧🇷 Em seguida, navegue até o diretório do projeto:

```bash
cd GO-CRUD
```

🇺🇸 To install all necessary dependencies, you can use the `go get` command:
🇧🇷 Para instalar todas as dependências necessárias, você pode usar o comando `go get`:

```bash
go get -d ./...
```

## ⚙️ Running

🇺🇸 To run the application, use the `go run` command:
🇧🇷 Para executar a aplicação, use o comando go run:

```bash
go run main.go
```

### 📋 Endpoints

🇺🇸 The API has the following endpoints:
🇧🇷 A API tem os seguintes endpoints:

- `GET /v1/tweets`: fetches all tweets.
- `POST /v1/tweet`: creates a new tweet. The request body should be a JSON that follows this format:

```json
{
  "name": "User's Name",
  "age": "User's Age",
  "profession": "User's Profession",
  "description": "Tweet Content"
}
```

- `DELETE /v1/tweet/:id`: deletes the tweet with the specified ID.:

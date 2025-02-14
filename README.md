# Eventos calendario google GO

Este é um projeto simples de CLI (Interface de Linha de Comando) para gerenciamento de calendários, desenvolvido em Go. Ele permite, visualizar e gerenciar eventos em um calendário do Google de forma fácil e prática.

## Pré-requisitos

- [Go](https://golang.org/doc/install) instalado na máquina.

## Como Executar

1. **Clonar o repositório:**

   ```bash
   git clone https://github.com/cleberbonifacio/go-cli-calendar.git
   cd go-cli-calendar

2. **Instalar dependências:**

   ```bash
   cd go-cli-calendar
   go mod tidy

3. **Criar conta de serviço no google cloud**
4. **Criar chaves para a conta de serviço**
5. **Vincular chave a alguma agenda do google calendar**
6. **Salvar credenciais no arquivo credentials.json**    

7. **Compilar o projeto::**

   ```bash
   go build -o calendar-cli main.go

## Como Usar

  ```bash
  ./calendar-cli events week
  ./calendar-cli events today

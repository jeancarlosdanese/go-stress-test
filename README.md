# Go Stress Test

Uma ferramenta de linha de comando em Go para realizar **testes de carga (load testing)** em qualquer serviço web.

---

## Funcionalidades

- Simula múltiplas requisições HTTP de forma concorrente
- Gera relatório ao final com:
  - Tempo total da execução
  - Total de requisições
  - Total de respostas 200 (OK)
  - Distribuição de códigos HTTP

---

## Como usar com Docker

### 1. Build da imagem:

```bash
docker build -t go-stress-test .
```

### 2. Executar o teste de carga:

```bash
docker run --rm go-stress-test exec   --url=http://example.com   --requests=100   --concurrency=10
```

Ou utilizando as **short options**:

```bash
docker run --rm go-stress-test exec   -u http://example.com   -r 100   -c 10
```

> A URL precisa estar acessível a partir do container Docker.

---

## Parâmetros

| Parâmetro              | Descrição                                     | Obrigatório    |
| ---------------------- | --------------------------------------------- | -------------- |
| `--url` / `-u`         | URL do serviço a ser testado                  | sim            |
| `--requests` / `-r`    | Total de requisições a serem executadas       | default: `100` |
| `--concurrency` / `-c` | Número de chamadas simultâneas (concorrência) | default: `10`  |

---

## Exemplo de Saída

```bash
Iniciando teste de carga:
- URL: http://example.com
- Requisições: 100
- Concorrência: 10

====== Load Test Report ======
Tempo total: 4.045s
Total de requisições: 100
Respostas 200 OK: 100

Distribuição de status HTTP:
  200: 100
================================
```

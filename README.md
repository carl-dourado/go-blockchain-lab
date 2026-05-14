# Go Blockchain Lab

Blockchain local e didatica em Go. A ideia e entender blocos, hash, proof-of-work simples, validacao e uma API pequena sem fingir que isso e uma rede blockchain real.

## O que tem

- Blocos com `index`, `timestamp`, registros, `previousHash`, `hash`, `nonce` e `difficulty`.
- Registros pendentes que viram bloco quando o comando `mine` roda.
- Proof-of-work simples com prefixo de zeros no hash SHA-256.
- Validacao da cadeia inteira.
- Persistencia local em JSON.
- CLI e API REST local.

## O que nao tem

- Rede P2P.
- Carteiras.
- Assinaturas digitais.
- Token ou moeda.
- Consenso distribuido.
- Segurança de producao.

## Rodar

```bash
go test ./...
go run ./cmd/go-blockchain-lab init
go run ./cmd/go-blockchain-lab add --author carlos --data "first local blockchain record"
go run ./cmd/go-blockchain-lab mine
go run ./cmd/go-blockchain-lab validate
go run ./cmd/go-blockchain-lab print
```

Os arquivos locais ficam em `data/chain.json` e `data/pending.json`.

## API local

```bash
go run ./cmd/go-blockchain-lab serve --addr 127.0.0.1:8789
```

Em outro terminal:

```bash
curl http://127.0.0.1:8789/health
curl http://127.0.0.1:8789/chain
curl -X POST http://127.0.0.1:8789/records \
  -H 'Content-Type: application/json' \
  -d '{"author":"carlos","data":"api record"}'
curl -X POST http://127.0.0.1:8789/mine
curl http://127.0.0.1:8789/validate
```

## Nivel do projeto

Esse projeto e propositalmente pequeno. Ele serve para praticar Go, JSON, HTTP, testes e conceitos basicos de blockchain. Um proximo passo honesto seria registrar eventos simulados de mainframe, como submit de JCL ou copia de dataset, mas o core atual fica como blockchain pura.

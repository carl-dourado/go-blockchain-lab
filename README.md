# Go Blockchain Lab

Blockchain local e didatica em Go. A ideia e entender blocos, hash, proof-of-work simples, validacao e uma API pequena sem fingir que isso e uma rede blockchain real.

## O que tem

- Blocos com `index`, `timestamp`, registros, `previousHash`, `hash`, `nonce` e `difficulty`.
- Registros pendentes que viram bloco quando o comando `mine` roda.
- Proof-of-work simples com prefixo de zeros no hash SHA-256.
- Validacao da cadeia inteira.
- Persistencia local em JSON.
- CLI e API REST local.
- Explorer web local servido pelo proprio binario.

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

## Explorer local

```bash
go run ./cmd/go-blockchain-lab serve --addr 127.0.0.1:8789
```

Depois abra:

```text
http://127.0.0.1:8789
```

O explorer mostra:

- altura da cadeia;
- quantidade de registros pendentes;
- status da validacao;
- ultimo hash;
- blocos minerados;
- registros pendentes;
- formulario para adicionar record;
- botao para minerar pendentes.

Ele e uma tela de estudo em cima da API local. Nao e um explorer de rede publica.

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

## Exemplos de records

```bash
go run ./cmd/go-blockchain-lab add --author carlos --data "transfer 10 from alice to bob"
go run ./cmd/go-blockchain-lab add --author carlos --data "jcl submit JOB001 status=ok"
go run ./cmd/go-blockchain-lab add --author carlos --data "dataset COPY from INPUT.PDS to OUTPUT.PDS"
go run ./cmd/go-blockchain-lab mine
go run ./cmd/go-blockchain-lab validate
```

Esses textos sao apenas dados de exemplo. O projeto ainda nao implementa saldo, assinatura, carteira, UTXO ou regra financeira.

## Nivel do projeto

Esse projeto e propositalmente pequeno. Ele serve para praticar Go, JSON, HTTP, testes e conceitos basicos de blockchain.

O foco atual e:

- entender como o hash muda quando o bloco muda;
- entender por que `previousHash` encadeia os blocos;
- ver registros saindo de `pending` e entrando em um bloco;
- validar a cadeia depois de minerar;
- ter uma tela local simples para enxergar o estado.

Proximos passos honestos, sem pular nivel:

- adicionar assinatura Ed25519 em records;
- criar uma regra simples de "transfer" apenas para estudo;
- registrar eventos simulados de mainframe, como submit de JCL ou copia de dataset;
- mostrar no explorer quando um bloco foi adulterado.

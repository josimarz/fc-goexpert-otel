# Full Cycle - Go Expert - Laboratório | Tracing distribuído com OpenTelemetry

Este projeto faz parte de um laboratório do curso de pós graduação em Go - Go Expert - da Full Cycle.

## Introdução

Este projeto implementa um servidor distribuído que permite consultar a temperatura de uma localidade em tempo real com base no CEP.

## Como executar

Para executar este projeto, através do terminal acesse o diretório raiz deste repositório e execute o seguinte comando:

```sh
docker compose up -d
```

Utilize o arquivo arquivo `api.http`, localizado na raiz deste repositório para executar as requisições.

## Como observar o tracing

Para visulizar o resultado do *tracing*, através de um navegador acesse o Zipkin (http://localhost:9411).
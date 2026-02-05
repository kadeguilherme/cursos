# Chapter 1 - A Linguagem de Programação Go

Este diretório contém os exercícios do Capítulo 1 do livro "A Linguagem de Programação Go".

## Estrutura

```
chapter-1/
├── README.md              # Este arquivo
└── echo/                  # Exercícios relacionados ao comando echo
    ├── ex1_1/             # Exercício 1.1 - Exibir nome do programa
    │   └── main.go
    ├── ex1_2/             # Exercício 1.2 - Exibir argumentos com índice
    │   └── main.go
    └── ex1_3/             # Exercício 1.3 - Comparar performance
        ├── main.go
        └── echo_test.go   # Testes unitários e benchmarks
```

## Executando os Exercícios

Para executar cada exercício:

```bash
# Exercício 1.1
cd echo/ex1_1 && go run main.go

# Exercício 1.2
cd echo/ex1_2 && go run main.go arg1 arg2 arg3

# Exercício 1.3
cd echo/ex1_3 && go run main.go arg1 arg2 arg3
```

## Descrição dos Exercícios

### Exercício 1.1
Modifique o programa echo para imprimir apenas o primeiro argumento (nome do programa).

### Exercício 1.2
Modifique o programa echo para imprimir o índice e o valor de cada argumento.

### Exercício 1.3
Compare a performance entre concatenação de strings e o uso de `strings.Join`.

**Nota**: Os testes comparativos de desempenho foram implementados, mas a avaliação sistemática não foi realizada porque ainda não cheguei na seção 11.4.

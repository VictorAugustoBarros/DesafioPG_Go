package model

import "errors"

// ErrLerArquivoYaml -> Erro aprensentado quando há problemas ao abrir arquivo pra leitura
var ErrLerArquivoYaml = errors.New(("Erro na leitura do arquivo .yml"))

// ErrConverterArquivoYaml -> Erro aprensentado quando há problemas na conversão dos dados para JSON
var ErrConverterArquivoYaml = errors.New(("Erro ao converter dados yaml para JSON"))

// ErrLerArquivoCsv -> Erro aprensentado quando há problemas ao abrir arquivo pra leitura do arquivo
var ErrLerArquivoCsv = errors.New(("Erro na leitura do arquivo .csv"))

// ErrNenhumArquivoEncontrado -> Erro aprensentado quando não foram encontrados arquivos no diretório
var ErrNenhumArquivoEncontrado = errors.New(("Nenhum arquivo encontrado"))

package controller

import (
	"desafioPG/model"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// pathArquivos -> Diretório dos arquivos
var pathArquivos = "files/"

// DadosTemp ->
type DadosTemp struct {
	Nome   string   `json:"nome"`
	Linhas []string `json:"linhas"`
}

// DadosArquivos -> variável para salvar todos os registros dos arquivos
var DadosArquivos []struct {
	DadosTemp
}

// ReadFiles -> Função responsável por retornar os dados dos arquivos no diretório "files"
func ReadFiles() (dados []string, err error) {
	err = nil

	files, err := ioutil.ReadDir(pathArquivos)
	if err != nil {
		err = errors.New(model.ErrNenhumArquivoEncontrado.Error() + " -> [" + err.Error() + "]")
	}

	for _, f := range files {
		arquivo, errFile := os.Open(pathArquivos + f.Name())
		if errFile != nil {
			err = errors.New(model.ErrLerArquivoCsv.Error() + " -> [" + errFile.Error() + "]")
			return
		}

		leitorCsv := csv.NewReader(arquivo)
		conteudo, errFile := leitorCsv.ReadAll()
		if errFile != nil {
			err = errors.New(model.ErrLerArquivoCsv.Error() + " -> [" + errFile.Error() + "]")
			return
		}

		var linhasArquivo []string
		for _, linha := range conteudo {
			linhasArquivo = append(linhasArquivo, linha[0])
		}

		dataTmp := DadosTemp{f.Name(), linhasArquivo}
		fmt.Println(dataTmp)

		DadosArquivos = append(DadosArquivos)
	}

	return DadosArquivos, err
}

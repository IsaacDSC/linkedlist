package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run profile_analyzer.go <arquivo.prof>")
		return
	}

	profileFile := os.Args[1]
	// Verificar se o arquivo existe
	if _, err := os.Stat(profileFile); os.IsNotExist(err) {
		fmt.Printf("Arquivo de perfil não encontrado: %s\n", profileFile)
		return
	}

	fmt.Printf("Analisando perfil: %s\n", profileFile)
	fmt.Println("Executando 'go tool pprof' para análise...")

	// Executar go tool pprof com o comando top10
	cmd := exec.Command("go", "tool", "pprof", "-top", profileFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao executar go tool pprof: %v\n", err)
		return
	}

	fmt.Println("\nPara uma análise mais detalhada, execute:")
	fmt.Printf("go tool pprof %s\n", profileFile)
	fmt.Println("Ou para visualização gráfica (requer Graphviz):")
	fmt.Printf("go tool pprof -web %s\n", profileFile)
}

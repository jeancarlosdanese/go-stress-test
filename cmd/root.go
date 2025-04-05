// cmd/root.go

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando raiz
var rootCmd = &cobra.Command{
	Use:   "go-stress-test",
	Short: "Um sistema CLI para teste de carga HTTP",
	Long:  "Uma ferramenta simples para executar testes de carga HTTP com relatórios ao final.",
}

// Execute executa o comando raiz
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}

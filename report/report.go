// report/report.go

package report

import (
	"fmt"
	"time"

	"github.com/jeancarlosdanese/go-stress-test/tester"
)

// Generate imprime um relatório final do teste de carga
func Generate(results []tester.Result, duration time.Duration) {
	total := len(results)
	statusCounts := make(map[int]int)
	success := 0

	// Conta os status HTTP
	for _, r := range results {
		statusCounts[r.StatusCode]++
		if r.StatusCode == 200 {
			success++
		}
	}

	// Relatório
	fmt.Println("\n====== Relatório do Teste de Carga ======")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requisições: %d\n", total)
	fmt.Printf("Respostas 200 OK: %d\n", success)
	fmt.Println("\nDistribuição de status HTTP:")

	for code, count := range statusCounts {
		label := fmt.Sprintf("%d", code)
		if code == 0 {
			label = "Erro de conexão"
		}
		fmt.Printf("  %s: %d\n", label, count)
	}
	fmt.Println("=========================================")
}

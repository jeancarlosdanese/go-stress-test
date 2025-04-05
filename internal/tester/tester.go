// tester/tester.go

package tester

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
}

// RunTest executa os testes de carga com concorrência definida
func RunTest(url string, totalRequests, concurrency int) []Result {
	results := make([]Result, 0, totalRequests)
	resultsCh := make(chan Result, totalRequests)
	jobs := make(chan struct{}, totalRequests)
	var wg sync.WaitGroup

	// Preenche o canal de jobs
	for i := 0; i < totalRequests; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	// Cria os workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := &http.Client{
				Timeout: 5 * time.Second,
			}

			for range jobs {
				status := doRequest(client, url)
				resultsCh <- Result{StatusCode: status}
			}
		}()
	}

	// Aguarda os workers finalizarem
	wg.Wait()
	close(resultsCh)

	// Coleta os resultados
	for r := range resultsCh {
		results = append(results, r)
	}

	return results
}

// doRequest realiza uma requisição HTTP e retorna o status
func doRequest(client *http.Client, url string) int {
	resp, err := client.Get(url)
	if err != nil {
		return 0 // erro de rede ou timeout
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

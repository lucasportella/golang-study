package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int)
	quit := make(chan int)
	// a ordem das chamadas das funções importa. Primeiro tem que ser a go routine (receive) pois caso a sendToChannel seja chamada antes, como ela é síncrona, vai terminar sem enviar nada para o canal e a função receive não vai conseguir receber os vaores, resultando em deadlock
	go receive(chan1, quit)
	sendToChannel(chan1, quit)
}

// função receive espera receber valor de um canal para printar, como está em go routine ela é executada de forma assíncrona, podendo a função sendToChannel enviar valores para ela de forma assíncrona.
func receive(chan1 chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-chan1)
	}
	quit <- 0
}

func sendToChannel(chan1 chan int, quit chan int) {
	contador := 1
	for i := 0; i < 50; i++ {
		select {
		// envia valor de contador para o canal 1 que o despejará na função receive
		case chan1 <- contador:
			contador++
		case <-quit: // quando terminar o loop da função receive o canal quit vai receber um valor, cair nesse case e retornar a função sendToChannel, encerrando o programa
			return
		}
	}
}

// por que esse programa não precisa de sync.WaitGroup?
// porque a função receive fica como uma go routine sendo executada depois, assim executa-se a função sendToChannel que é um loop infinito, e vai enviando o contador para o channel 1. Como a função receive está limitada a 50 iterações, ela vai receber no máximo 50 valores do channel 1 e depois vai enviar um valor (0) para o channel quit.
// como no select temos o case do channel quit receber um valor, nesse caso a função sendToChannel retorna undefined e termina sua execução, terminando o programa sem nenhum problema de sincronicidade.

// caso o loop da função receive fosse infinito, aí sim a função sendToChannel nunca iria parar de enviar valores para chan1
// caso o loop da função sendToChannel NÃO fosse infinito e maior que 50, após a 50ª iteração não haveria mais a go routine da função receive para receber valores, não sendo possível imprimir mais valores, pois a go routine foi fechada com o comando quit <- 0

// importante notar que o comando quit <- 0 não termina a go routine, a go routine termina pq ela terminou de executar a última linha. O que o comando quit <- 0 faz é sinalizar para o select statement que a função sendToChannel deve retornar. Como a função sendToChannel retornou e ela era a última linha de código da função main, logo o comando quit <- 0, indiretamente, encerrou o programa inteiro.

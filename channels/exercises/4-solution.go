package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)
	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // all values sent to the "c" channel. (the select statement will not know the "c" channel was close, hence we must use the command below)
		q <- 0 // send message to "q" channel so the select loop ends
	}()
	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

/*
Esse exercício tá uma bagunça, até a Ellen se perdeu numa hora. O pessoal que usou comma ok conseguiu fazer rodar, mas não consumiu todos os statements do select. Tem alguns conceitos importantes que estão passando despercebidos. 
Vamos analisar primeiramente o exercício: A proposta é mostrar os valores enviados pro canal "c" com um select statement. Além disso existe um canal "q" que é para terminar o for infinito e encerrar o programa. 

Para encerrar corretamente o programa com o select statement do canal “q” é preciso enviar um valor para esse canal. E aí está o grande problema. Uma galera não percebeu o que ela falou e fez em 4:33. Ela trocou a direção do canal no parâmetro da função gen, permitindo que o canal recebesse informações, ao invés de enviá-las(como o canal “q” não recebe nenhum valor, fica impossível o select statement ler um valor do canal “q”, impossibilitando a função receive encerrar (e o envio de valor ao canal “q” tem que ser necessariamente feito dentro da função gen, pois a função receive é síncrona, não dá para colocar na main senão a mensagem para o canal “q” vai chegar antes da função receive começar ou nunca vai chegar pois a função receive vai travar o código).

Teve o pessoal que usou comma ok, tiveram que fazer isso porque no código deles dá pra ver que não trocaram a direção do canal no parâmetro da função gen, logo a solução deles foi usar o próprio statement do canal “c” com comma ok para encerrar a função receive. Só que isso torna inútil o statement deles do “case <-q”. Podem testar, se removerem o segundo case do código deles, o resultado é o mesmo. Nesse caso, não precisa de um select, daria pra fazer com um simples loop direto no canal: https://go.dev/play/p/gCJNhyt36pw A vantagem desse caso é que o loop direto no canal detecta automaticamente quando o canal é fechado e assim para de ler as informações do canal. Isso não acontece com um select statement.

O problema é que o exercício pediu para usar select statement.

Se não trocarmos a direção do canal na função gen, o que acontecerá é que o loop da função gen vai iterar 100 vezes e mostrar os valores até 99. Depois disso o canal é fechado. Como o select não sabe que o canal fechou, ele fica lendo infinitamente o canal. Como o canal está fechado, o select case do canal “c” só vai conseguir ler o valor padrão de um canal de inteiros, o valor 0. Aí para indicar para o select que o canal fechou, usa-se comma ok que verifica se o valor 0 realmente foi enviado ou se é o valor padrão não inicializado. Se é valor não inicializado retornam a função.

Quem usou comma ok não consegue encerrar o programa pelo case do "q", pois nunca vai cair nele, já que não enviaram nenhum valor para "q", logo o programa deles vai dar um loop até o valor 99, o loop vai se encerrar.

A solução mais correta para esse exercício seria o que a Ellen fez, trocar a direção do canal e enviar uma mensagem qualquer para o canal “q”, encerrando a função receive logo após o envio do valor 99 e o fechamento do canal.

*/

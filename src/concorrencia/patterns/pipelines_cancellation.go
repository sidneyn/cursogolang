package main

import (
	"fmt"
	"sync"
	//"sync"
)

/************************** PIPELINE **************************

	Informalmente, um pipeline é uma série de estágios conectados por canais,
	em que cada estágio é um grupo de goroutines executando a mesma função.
	Em cada etapa, as goroutines
	- receber valores do upstream (sob o fluxo com o dados) via canais de entrada
	- executar alguma função nesses dados, geralmente produzindo novos valores
	- enviar valores a downstream (desce o fluxo com o dados) através de canais de saída.

****************************************************************/
/**
A função interna make aloca e inicializa um objeto do tipo
fatia, mapa ou chan (apenas). Como novo, o primeiro argumento é um tipo, não um
valor. Diferente do novo, o tipo de retorno do make é igual ao tipo do seu
argumento //, não um ponteiro para ele. A especificação do resultado depende do
tipo:
Fatia: o tamanho especifica o comprimento. A capacidade da fatia é
 igual ao seu comprimento. Um segundo argumento inteiro pode ser fornecido para
 especificar uma capacidade diferente; não deve ser menor que o
 comprimento. Por exemplo, make ([] int, 0, 10) aloca uma matriz subjacente
 do tamanho 10 e retorna uma fatia do comprimento 0 e da capacidade 10 que é
 apoiado por esta matriz subjacente.
 Mapa: um mapa vazio é alocado com espaço suficiente para armazenar o
 número especificado de elementos. O tamanho pode ser omitido; nesse caso,
 um pequeno tamanho inicial é alocado.
 Canal: o buffer do canal é inicializado com a capacidade especificada do buffer. Se zero, ou o tamanho for omitido, o canal será // sem buffer.
*/

/* Função gen CANAL DE ENTRADA PARA SERVIDOR
@range nums
@return chan int
  1. cria e inicializa a variavel out  em seguida converte os numeros inteiros em um canal para serem emitidos na lista.
  2. inicia goroutine para enviar os números inteiros ao canal apos enviar todos ele é fechado.
*/

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums)) // Canal: o buffer do canal é inicializado com a capacidade especificada do buffer. Se zero, ou o tamanho for omitido, o canal será sem buffer.
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

/**
Função sq CANAL DE SAIDA PARA O CLIENTE
	recebe o canal de saida para o cliente dos numeros inteiros:
	1. cria e inicializa a variavel out em seguida converte os numeros inteiros em um canal para serem emitidos na lista.
	2. inicia uma sub função goroutine que calcula o quadrado da lista de numeros inteiros, e dar saida para o canal em seguida o canal é fechado.
*/
func sq(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}

		}
		close(out)
	}()
	return out
}

func main() {

	// Set up the pipeline.
	/*
		c := gen(2, 3, 4 )  // NA
		out := sq(c)

		// consume the output
		fmt.Println(<-out) // 4
		fmt.Println(<-out) // 9
	*/

	// Set up the pipeline and consume the output
	/*
		for n := range sq(sq(gen(2, 3, 4))) {
			fmt.Println(n) // 16 then 81
		}
	*/

	/*
		in := gen(2, 3, 4)

		c1 := sq(in)
		c2 := sq(in)
		// Consume a saída mesclada de c1 e c2,
		for n := range merge(c1) {
			fmt.Println(n) // 4 then 9, or 9 then 4
		}
	*/
	/**************************************
		in := gen(2,3)
		// distribuindo o trabalho sql em duas gourotines que são lidas em
		c1 := sq(in)
		c2 := sq(in)

		// Consumo o primeiro valor da saída.
		done := make (chan struct{}, 2)
		out := merge(done, c1, c2)
		fmt.Println(<- out)  // 4 or 9
		fmt.Println(<- out)  // 4 or 9
		//fmt.Println(<- out)  // 4 or 9

		// Diga aos remetenntes restantes que estamos saindo.
	//	done <- struct{}{}
		//done <- struct{}{}
	******************************************/
	// configure um canal pronto que seja compartilhado por todo pipeline,
	// e feche esse canal quando esse pipelione sair, como um sinal
	// para todas as goroutines que começamos a sair.
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)
	// distribuindo o trabalho sql em duas gourotines que são lidas em
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consome o primeiro valor da saída
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 ou 9

	// concluiído será fechado pela chamada adiada.

}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup // organiza a sincronização
	out := make(chan int)

	// Start an output goroutin for each input channel in cs. output
	// copies values from c to out until c is closed, then calls wg.Done.

	output := func(c <-chan int) {
		defer wg.Done() // defer garanti fechar a sincronização no final
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}

	}
	wg.Add(len(cs)) // total de (canais) serão sincronizados
	// percorre a lista de canais para ler na função output
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

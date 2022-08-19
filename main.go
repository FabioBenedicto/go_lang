package main

import "fmt"
//0import "encoding/json"

/*type Aluno struct{
	RA string `json: "ra"`
	Nome string `json: "nome"`
}

func (aluno *Aluno) getNome() string {
	return aluno.Nome
}
*/

func main(){
	/*var variavel1 string
	variavel2 := "ola"
	var(
		variavel3 string
		variavel4 int
	)
	const variavel5 = 5

	fmt.Println(variavel1, " ", variavel2, " ", variavel3, " ", variavel4, " ", variavel5, " ")
	*/
	funcao := func(){
		fmt.Println("Hello World")
	}

	funcao()

	resultado, ok := Divisao(3, 0)

	if ok {
		fmt.Println(resultado)
	} else {
		fmt.Println("Impossível dividir por 0")
	}

	//fmt.Println(Somar(3, 4))

	fmt.Println(Somar(1,2,3,4,5))
	fmt.Println(Somar(1,2,3))

	

	fmt.Println( " ", func(a int, b int) int {
		return a * b
	}(3, 4))

	//sophia := Aluno{RA: "202501", Nome: "Sophia"}
	//fmt.Println(sophia.RA, " ", sophia.Nome)

	/*var sophia Aluno
	sophia.RA = "202501"
	sophia.Nome = "Sophia"
	fmt.Println(sophia.RA, " ", sophia.Nome)

	aluno := Aluno{RA: "200154", Nome: "Victor"}
	fmt.Println(aluno.getNome())*/

	/*alunoJson := Aluno{RA: "200154", Nome: "Victor"}
	json, erro := json.Marshal(aluno)
	if erro != nil {
		fmt.Println("Erro: ", erro.Error())
		return
	}

	var alunoNovo Aluno
	erro := json.Unmarshal(alunoJson, &alunoNovo)
	if erro != nil {
		fmt.Println("Erro: ", erro.Error())
		return
	}
	
	stringJson := string(json)
	fmt.Println(stringJson)
	*/

	array := [5] int {1, 2, 3, 4, 5}
	array[0] = 1
	slice1 := array[2:3]
	slice1 := array[2:]
	slice1 := array[:4]
	fmt.Println(slice1)

	dicionario := map[string] string{}
	dicionario["oii"] = "Um"
	dicionario["dsd"] = "Zebra"
	dicionario["dso"] = "Girafa"
	fmt.Println(dicionario)

	dicionario := map[string] string{
		"1": "um",
		"2": "dois",
	}

	fmt.Println(dicionario["1"])
	delete(dicionario, "1")

	if 1 == 1 {
		panic("1 é um igual a 1")
	}

	//panic, ,defer print hello word com recover

}

/*func Somar(a int, b int) (int c){
	c = a + b
	return
}*/

func Divisao(a int, b int) (c int, ok bool){
	if b == 0{
		ok = false
		c = -1
	} else {
		c = a/b
		ok = true
	}
	return
}

func Somar(numeros ...int) int{
	soma := 0
	for _, valor := range numeros {
		soma+=valor
	}
	return soma
}

/*func Fibonacci(number int) int{
	if number == 0 {
		return 0
	} 
	if number == 1{
		return 1
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
}*/

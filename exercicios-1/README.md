# Revisão 1

## Questão 1

Em um dia quente de verão, Pete e seu amigo Billy decidiram comprar uma melancia. Eles escolheram a maior e mais
saborosa, na opinião deles, e, em seguida, pesaram a fruta na balança, obtendo seu peso em quilos. Morrendo de sede,
correram para casa com a melancia e decidiram dividi-la. No entanto, enfrentaram um problema difícil.

Como grandes fãs de números pares, Pete e Billy querem dividir a melancia de tal forma que cada uma das duas partes pese
um número par de quilos, sem que as partes necessariamente tenham pesos iguais. Mas os meninos estão extremamente
cansados e querem começar a refeição o mais rápido possível. Portanto, precisam de ajuda para descobrir se é possível
dividir a melancia da maneira que desejam. É importante destacar que cada um deles deve receber uma parte de peso
positivo.

A função deve retornar um valor booleano, indicando se é possível ou não dividir a melancia da forma desejada. Se o peso
da melancia for menor ou igual a 0, a função deve retornar um erro.

### Exemplo de entrada

```
Peso da melancia: 8
```

### Exemplo de saída:

```
É possível: true
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

Nesse caso, é possível dividir a melancia em duas partes de peso par: 4kg e 4kg. Portanto, a função retorna True.

### Exemplo de entrada

```
Peso da melancia: 6
```

### Exemplo de saída:

```
É possível: true
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

Nesse caso, é possível dividir a melancia em duas partes de peso par: 4kg e 2kg. Portanto, a função retorna True.

### Exemplo de entrada

```
Peso da melancia: 5
```

### Exemplo de saída:

```
É possível: false
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

Nesse caso, não é possível dividir a melancia em duas partes de peso par, já que 5 é ímpar. Portanto, a função retorna
False.

### Exemplo de entrada

```
Peso da melancia: 0
```

### Exemplo de saída:

```
É possível: false
Erro: "peso inválido"
```

### Explicação:

Nesse caso, o peso da melancia é 0, o que não é um peso válido. Portanto, a função retorna um erro indicando que o peso
é inválido.

## Questão 2

Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de
programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os
amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso
contrário, os amigos não escreveriam a solução do problema.

Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você
receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e
Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false"
indica o contrário.

Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.

### Exemplo de entrada

```
[[true, true, false],
[true, false, true],
[true, true, true],
[false, false, true]]
```

### Exemplo de saída:

```
3
```

### Explicação:

Neste exemplo, há 4 problemas oferecidos no concurso. A primeira linha da matriz indica que Pedro e Vanessa têm certeza
da solução para o primeiro problema, mas Tônia não. A segunda linha indica que apenas Pedro tem certeza da solução para
o segundo problema, e assim por diante. Os amigos só vão escrever uma solução para um problema se pelo menos dois deles
tiverem certeza da solução. Portanto, eles escreverão soluções para os problemas 1, 2 e 3 (porque nesses casos, pelo
menos dois amigos têm certeza da solução), mas não para o problema 4.

## Questão 3

Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
tabuleiro, seguindo as seguintes condições:

1. Cada peça de dominó cobre completamente dois quadrados.
2. Nenhuma duas peças de dominó se sobrepõem.
3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.

Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.

Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

### Exemplo de entrada

```
M = 3
N = 3
```

### Exemplo de saída:

```
4
```

### Explicação:

Um dos possíveis tabuleiros é mostrado abaixo. As peças de dominó são mostradas em barras.

```
--|
--|
--.
```

### Exemplo de entrada

```
M = 2
N = 3
```

### Exemplo de saída:

```
3
```

### Explicação:

Um dos possíveis tabuleiros é mostrado abaixo. As peças de dominó são mostradas em barras.

```
--|
--|
```

## Questão 4

Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.

Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
Caso a lista possua apenas um elemento, a função deve retornar 3.

### Exemplo de entrada

```
[1, 2, 3, 4, 5]
```

### Exemplo de saída:

```
1
```

### Explicação:

Nesse caso, a lista está em ordem crescente, então a função retorna 1.

### Exemplo de entrada

```
[5, 4, 3, 2, 1]
```

### Exemplo de saída:

```
2
```

### Explicação:

Nesse caso, a lista está em ordem decrescente, então a função retorna 2.

### Exemplo de entrada

```
[]
```

### Exemplo de saída:

```
Erro: "lista vazia"
```

### Explicação:

Nesse caso, a lista está vazia, então a função retorna um erro.

## Questão 5

Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
minúsculas, ele:

* deleta todas as vogais;
* insere um caractere "." antes de cada consoante;
* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.

As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
programa na sequência de caracteres inicial.

Ajude Pedro a lidar com esta tarefa fácil.

### Exemplo de entrada

```
Tour
```

### Exemplo de saída:

```
.t.r
```

### Explicação:

Neste caso, o programa deve deletar as vogais "o" e "u" e inserir um ponto antes de "t" e "r". Além disso, ele
substitui "T" por "t", resultando na string ".t.r".

### Exemplo de entrada

```
CEUB
```

### Exemplo de saída:

```
.c.b
```

### Explicação:

Neste caso, o programa deve deletar as vogais "E" e "U" e inserir um ponto antes de "C" e "B". Além disso, ele
substitui "C" por "c" e "B" por "b", resultando na string ".c.b".

## Questão Bônus

Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.

Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

### Exemplo de entrada

```
[1, 2, 3]
```

### Exemplo de saída:

```
1 altura máxima
3 torres
```

### Explicação:

Nesse caso, Joãozinho pode construir 3 torres utilizando as 3 barras de madeira.

```
x | xx | xxx
```

### Exemplo de entrada

```
[6, 5, 6, 7]
```

### Exemplo de saída:

```
2 altura máxima
3 torres
```

### Explicação:

Nesse caso, Joãozinho pode construir três torres utilizando as 4 barras de madeira.

```
xxxxxx
xxxxxx | xxxxx | xxxxxxx
```

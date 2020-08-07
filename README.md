# Introdução

O teor deste desafio é bastante voltado a alguns problemas que resolvemos com frequência, e vai nos ajudar a descobrir como você raciocina e quais são suas habilidades.

Como constantemente precisamos pensar em formas diferentes de resolver os desafios que enfrentamos, acreditamos que uma base teórica sólida é mais importante que apenas ser bom em uma linguagem ou framework. Além disso queremos ver o seu melhor, e por isso não é preciso ficar limitado às tecnologias exigidas na descrição da vaga.

# O desafio

O desafio consiste em 2 problemas técnicos parecidos com os que enfrentamos na vida real.
A resolução deles não necessariamente requer a implementação de um mega-projeto, porém ambos precisam de um certo conhecimento para serem resolvidos.

Leia atentamente os enunciados dos problemas, pois apesar de não haver nenhuma pegadinha, os detalhes importam.

Outros requisitos do projeto incluem:
- Deve funcionar em um ambiente Linux
- Deve ter testes automatizados
- Deve ter um README explicando como instalar as dependências, executar as soluções e os testes.

Também sugerimos que o projeto seja organizado de uma forma parecida com esta:

```
part-1/
    src/
    test/
part-2/
    src/
    test/
README.md
```

## Parte 1 - API de produtos

Precisamos de uma API para receber a atualização de dados cadastrais de produtos. Ela deve receber um corpo no formato JSON, onde o tamanho varia desde alguns poucos Kb até alguns Gb.
Experiências anteriores mostram que alguns clientes costumam enviar o mesmo corpo repetidas vezes ao longo de um curto espaço de tempo.
Isso nos causou alguns problemas, como o fato de ter que escalar nossos bancos de dados muito além do necessário afim de aguentar a carga extra desnecessária.

## Parte 2 - Agregador de URLs

Recebemos um dump com lista de URLs de imagens de produtos que vamos utilizar para manter nossa base de dados atualizada.
As URLs pertencem a uma empresa terceirizada que hospeda a maioria destas imagens, e ela nos cobra um valor fixo por cada request.
Já sabemos que o dump de origem não tem uma boa confiabilidade, pois encontramos várias imagens repetidas e boa parte delas também retornam status 404.
Como não é interessante atualizar nossa base com dados ruins, filtramos apenas as URLs que retornam status 200.

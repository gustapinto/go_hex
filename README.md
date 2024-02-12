# Go HEX

Um projeto simples para aperfeiçoar arquitetura hexagonal, baseado no artigo [Ready for changes with Hexagonal Architecture
](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749), do time de engenharia de software da Netflix. Também usando novas funcionalidades da linguagem Go 1.22.

## O que é arquitetura hexagonal?

Arquitetura hexagonal é uma arquitetura de *software* que prega o desacoplamento entre os componentes de um sistema, de forma que o desenvolvimento do mesmo seja altamente flexível, permitindo que as implementações de cada um de seus módulos sejam facilmente trocadas no futuro, com as dependências entre os mesmos se dando a partir do diagrama abaixo: 

<p align="center">
    <img src="https://miro.medium.com/v2/resize:fill:1200:632/g:fp:0.49:0.49/1*NfFzI7Z-E3ypn8ahESbDzw.png">
</p>

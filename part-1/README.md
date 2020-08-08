<h1 align="center">
  <img style="background-color: #312e38; border-radius: 10px;" alt="gobarber-logo" src="https://www.linx.com.br/app/themes/linx/crystals/dist/assets/static/logo.png" />
  <p align="center">
    <a href="https://nodejs.org/en/">
      <img src="https://img.shields.io/badge/-NodeJS-006400?style=flat&logo=Node.js&logoColor=#339933" />
    <a href="https://www.typescriptlang.org/">
      <img src="https://img.shields.io/badge/-TypeScript-007ACC?style=flat&logo=TypeScript&logoColor=#007ACC" />
    </a>
    <a href="https://jestjs.io/">
      <img src="https://img.shields.io/badge/-Jest-C21325?style=flat&logo=Jest&logoColor=FFFFF" />
    <a href="https://www.mongodb.com/">
      <img src="https://img.shields.io/badge/-MongoDB-47A248?style=flat&logo=MongoDB&logoColor=006400" />
    </a>
  </p>
</h1>

## 🔖 Sobre o projeto

O projeto **API de Produtos**, desenvolvido para teste de programador na [Linx](https://www.linx.com.br/ 'Linx'). Tendo como objetivo, a atualização de dados cadastrais de produtos.

## 💻 Tecnologias

- <img width="20px" src="https://img.icons8.com/color/2x/nodejs.png" /> [NodeJS](https://nodejs.org/en/ 'NodeJS')
- <img width="20px" src="https://img.icons8.com/color/2x/typescript.png" /> [TypeScript](https://www.typescriptlang.org/ 'TypeScript')
- <img width="20px" src="https://res.cloudinary.com/practicaldev/image/fetch/s--00h6CjGb--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://www.maxrooted.com/panduan-membangun-rest-api-expressjs-mysql/cover.png" /> [Express](https://expressjs.com/ 'Express')
- <img width="20px" src="https://img.icons8.com/color/2x/mongodb.png"/> [MongoDB](https://www.mongodb.com/ 'MongoDB')
- <img width="20px" src="https://avatars2.githubusercontent.com/u/20165699?s=400&v=4" /> [TypeORM](https://typeorm.io/#/ 'TypeORM')
- <img width="20px" src="https://simpleicons.org/icons/jest.svg" /> [Jest](https://jestjs.io/ 'Jest')
- <img width="20px" src="https://img.icons8.com/dusk/2x/docker.png" /> [Docker](https://www.docker.com/ 'Docker')

## ▶️ Getting Started

- **Passo 1️⃣** : git clone do projeto [Linx](https://github.com/rafaelsanzio/linx 'Linx')
- **Passo 2️⃣** : executar a instalação de: [Node](https://nodejs.org/en/ 'Node') e [Docker](https://www.docker.com/ 'Docker')
- **Passo 3️⃣** : rodando a aplicação executando os seguintes comandos:

```bash
   # Navegando até a pasta do projeto
   $ cd linx/part-1

   # Instalando todas as depêndencias necessárias
   $ npm install ou yarn install

   # Criando cluster no banco de dados mongoDB usando o docker
   $ docker run --name mongodb-linxapi -p 27017:27017 -d -t mongo

   # Iniciando o banco de dados
   $ docker start mongodb-linxapi

   # Starting o backend da aplicação
   $ npm dev:server ou yarn dev:server

   # Rodando os testes
   $ npm test ou yarn test
```

## 📸 Imagens

- Imagem de exemplo de request e response: 

 <p align="center">
 	<img src="https://user-images.githubusercontent.com/18368947/89713900-39a70380-d971-11ea-910e-7bf010baa909.png" />
 </p>

## ㊗ ️ Considerações

- Projeto desenvolvido by:

  - <a href="https://github.com/rafaelsanzio">
    <img src="https://img.shields.io/badge/-Rafael%20Sanzio-000000?style=flat&logo=GitHub&logoColor=#000000" />
  </a>

  - <a href="https://www.linkedin.com/in/rafael-sanzio-012778143/">
    <img src="https://img.shields.io/badge/-Rafael%20Sanzio-0077B5?style=flat&logo=LinkedIN&logoColor=#000000" />
  </a>

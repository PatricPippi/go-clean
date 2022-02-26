## SERVICE 
Core da aplicação, não depende de outras camadas
- ### Domain *Tipos de dados e regras internas*
- ### UseCase *Lógicas de negócio da aplicação*
- ### Infrastructure *implementa requisitos do usecase* 
  - Repository *implementa drivers de banco para o usecase*

## ITERFACE ADAPTERS
Adaptam dados recebidos de fora da aplicação para o core da aplicação e vice-versa, interagem com o mundo externo

- ### Api *Cria servidor http para o mundo externo*
  - Controller *Logica de negócio para ser utilizado por rotas, utiliza use cases* 
  - Middleware *Middlewares*
  - Routes *Rotas*
  - Presenter *Dados utilizados pelo mundo externo*
- ### gRPC *Cria servidor grpc com rpcs para o mundo externo*
        
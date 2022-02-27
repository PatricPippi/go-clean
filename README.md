
## Domain Layer
Core da aplicação
  - ### Domain *Tipo de dados, regras internas e requisitos da aplicação*
  - ### UseCase *Lógicas de negócio da aplicação*

## Data layer
Serve e adapta dados ao domain layer
  - ## Infra
    - Repository *implementa drivers de banco para o usecase*
    - Gateway *implementa dados externos, exemplo: microserviços, apis*

## Presentation layer
Expõe a aplicação para o mundo externo e adapta dados de request e response para domain layer
  - ## Delivery
    - Api Server
    - Grpc Server

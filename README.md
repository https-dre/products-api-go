# Products Api Com Go!

Um crud básico para gerenciar produtos de um sistema através de uma API construída com Go-Lang.
O serviço usa Gin-Gonic em conjunto com um Banco De Dados Postgre-Sql em um container Docker.

## Rotas

Todas as rotas da aplicação com seus UseCases

#### Listar todos os produtos cadastrados
```http
GET http://host/products HTTP/1.1
```

#### Cadastrar um produto
```http
POST http://host/create-products HTTP/1.1
Content-Type: application/json; charset=utf-8

{
    "product_name": "Bala Fini Menta",
    "product_price": 12.00
}
```

#### Encontrar um produto por Id
```http
GET http://localhost:8000/product-by-id/5 HTTP/1.1
```
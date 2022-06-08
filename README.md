# API - MercadoFresco - Grupo: A Sociedade do GO

## Sellers

## Warehouses
### Cadastrar Warehouse
- uri:  `localhost:8080/api/v1/warehouses`
- método: `POST`
- body: 
  ```
  {
    "warehouse_code": string, unique
    "address": string
    "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
    "minimum_capacity": number, integer, greater than 0
    "minimum_temperature": float
  }
  ```
- responses em caso de sucesso: 
    - status: 201
      - body:
        ```
        "data": {
          "id": number
          "warehouse_code": string, unique
          "address": string
          "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
          "minimum_capacity": number, integer, greater than 0
          "minimum_temperature": float
        }
        ```
- responses em caso de falha: 
    - status: 400
    - status: 422
    - status: 500
      - body: comum para todas as requisições com falha
        ```
        {
          "error": string
        }
        ```

### Listar todas as Warehouses
- uri:  `localhost:8080/api/v1/warehouses`
- método: `GET`

- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": [
          {
            "id": number
            "warehouse_code": string, unique
            "address": string
            "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
            "minimum_capacity": number, integer, greater than 0
            "minimum_temperature": float
          },
          ...
        ]
        ```
- responses em caso de falha: 
    - status: 500
      - body:
        ```
        {
          "error": string
        }
        ```

### Listar Warehouse por Id
- uri:  `localhost:8080/api/v1/warehouses/:id`
- método: `GET`
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": {
          "id": number
          "warehouse_code": string, unique
          "address": string
          "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
          "minimum_capacity": number, integer, greater than 0
          "minimum_temperature": float
        }
        ```
- responses em caso de falha: 
    - status: 400
    - status: 404
    - status: 500
      - body: comum para todas as requisições com falha
        ```
        {
          "error": string
        }
        ```

### Atualizar Warehouse
- uri:  `localhost:8080/api/v1/warehouses/:id`
- método: `PATCH`
- body: 
  ```
  {
    "warehouse_code": string, unique
    "address": string
    "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
    "minimum_capacity": number, integer, greater than 0
    "minimum_temperature": float
  }
  ```
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": {
          "id": number
          "warehouse_code": string, unique
          "address": string
          "telephone": string, format: (xx) xxxxx-xxxx or (xx) xxxx-xxxx
          "minimum_capacity": number, integer, greater than 0
          "minimum_temperature": float
        }
        ```
- responses em caso de falha: 
    - status: 400
    - status: 404
    - status: 422
    - status: 500
      - body: comum para todas as requisições com falha
        ```
        {
          "error": string
        }
        ```

### Deletar Warehouse
- uri:  `localhost:8080/api/v1/warehouses/:id`
- método: `DELETE`
- responses em caso de sucesso: 
    - status: 204
- responses em caso de falha: 
    - status: 400
    - status: 404
    - status: 500
      - body: comum para todas as requisições com falha
        ```
        {
          "error": string
        }
        ```


## Sections

## Products

## Employees

## Buyers

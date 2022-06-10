# API - MercadoFresco - Grupo: A Sociedade do GO

## Sellers
### Cadastrar Seller
- uri:  `localhost:8080/api/v1/seller`
- método: `POST`
- body: 
  ```
  {
    "Cid": string, unique
    "CompanyName": string
    "Address": string
    "Telephone": string

  }
  ```
- responses em caso de sucesso: 
    - status: 201
      - body:
        ```
        "data": {
          "id": number
          "Cid": string, unique
          "CompanyName": string
          "Address": string
          "Telephone": string
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

### Listar todas os sellers
- uri:  `localhost:8080/api/v1/seller`
- método: `GET`

- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": [
          {
            "id": number
           "Cid": string, unique
           "CompanyName": string
           "Address": string
           "Telephone": string
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

### Listar seller por Id
- uri:  `localhost:8080/api/v1/seller/:id`
- método: `GET`
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": {
          "id": number
          "Cid": string, unique
          "CompanyName": string
          "Address": string
          "Telephone": string
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

### Atualizar seller
- uri:  `localhost:8080/api/v1/seller/:id`
- método: `PATCH`
- body: 
  ```
  {
   
          "Cid": string, unique
          "CompanyName": string
          "Address": string
          "Telephone": string
  }
  ```
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        "data": {
          "Cid": string, unique
          "CompanyName": string
          "Address": string
          "Telephone": string
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

### Deletar seller
- uri:  `localhost:8080/api/v1/seller/:id`
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

## Employees
### Cadastrar Employee
- url:  `localhost:8080/api/v1/employees`
- método: `POST`
- body: 
  ```
  {
    "card_number_id": INT, UNIQ,
    "first_name": STRING,
    "last_name": STRING,
    "warehouse_id": INT (Precisa ter um wareHouse ID Válido)
  }
  ```
- responses em caso de sucesso: 
    - status: 201
      - body:
        ```
        {
          "data": {
            "id": INT,
            "card_number_id": INT,
            "first_name": "STRING",
            "Last_name": "STRING",
            "warehouse_id": INT
          }
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

### Listar todos os Employees
- url:  `localhost:8080/api/v1/employees`
- método: `GET`

- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        {
          "data": [
            {
              "id": INT,
              "card_number_id": INT,
              "first_name": STRING,
              "Last_name": STRING,
              "warehouse_id": INT
            }
          ]
        }
        ```
- responses em caso de falha: 
    - status: 500
      - body:
        ```
        {
          "error": string
        }
        ```

### Listar Employee por Id
- url:  `localhost:8080/api/v1/employees/id`
- método: `GET`
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
      {
        "data": {
          "id": INT,
          "card_number_id": INT,
          "first_name": STRING,
          "Last_name": STRING,
          "warehouse_id": INT
        }
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

### Atualizar Employee
- url:  `localhost:8080/api/v1/employees/id`
- método: `PATCH`
- body: 
  ```
      {
        "id": INT UNIQ,
        "card_number_id": INT,
        "first_name": STRING,
        "Last_name": STRING,
        "warehouse_id": INT
      }
  ```
- responses em caso de sucesso: 
    - status: 200
      - body:
        ```
        {
          "data": {
          "id": INT,
          "card_number_id": INT,
          "first_name": STRING,
          "Last_name": STRING,
          "warehouse_id": INT
          }
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

### Deletar Employee
- url:  `localhost:8080/api/v1/employees/id`
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

## Buyers

## Sections

## Products

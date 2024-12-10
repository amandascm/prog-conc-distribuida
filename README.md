# Programação concorrente e distribuída

### Descrição

Nesta implementação, o recurso a ser invocado a partir de uma pool de instâncias (padrão pooling) representa uma simples calculadora. O servidor fornecerá as instâncias de tal forma que elas possam ser usadas concorrentemente a partir do seu lifecycle manager (responsável por ativar, desativar, obter e liberar objetos de acordo com a strategy em uso, que neste caso é uma pool de `Calculator`).

Observe a partir dos prints durante a execução que diferentes IDs de instâncias da pool são usadas (não sequencialmente), como no exemplo abaixo:

```
Got object available at channel with ID: 4
Returned to pool object with ID: 3
Got object available at channel with ID: 3
Got object available at channel with ID: 2
Returned to pool object with ID: 7
Got object available at channel with ID: 7
```

### Execução

Para executar o cliente e servidor, siga os passos abaixo:

1. Execute o naming service
    ```
    go run project/services/naming/namingserver/namingserver.go
    ```

2. Execute o servidor
    ```
    go run project/app/server/servidor.go
    ```

3. Execute o cliente (um ou mais)
    ```
    go run project/app/client/cliente.go
    ```

# Programação concorrente e distribuída

## Atividade 1

### Descrição

Nesta implementação, o recurso a ser invocado a partir de uma pool de instâncias (padrão pooling) representa um simples coletor de logs ou métricas. O servidor fornecerá as instâncias de tal forma que elas possam ser usadas concorrentemente. 

Observe a partir dos prints durante a execução que diferentes IDs de instâncias da pool são usadas (não sequencialmente), como no exemplo abaixo:

```
Got object available at channel with ID: 6
    Received log with message: this is a log message 2245
Returned to pool object with ID: 9
Got object available at channel with ID: 9
    Received log with message: this is a log message 2473
Returned to pool object with ID: 4
```

### Execução

Para executar o cliente e servidor, siga os passos abaixo:

1. Execute o naming service
    ```
    go run atv1/services/naming/namingserver/namingserver.go
    ```

2. Execute o servidor
    ```
    go run atv1/app/server/servidor.go
    ```

3. Execute o cliente (um ou mais)
    ```
    go run atv1/app/client/cliente.go
    ```

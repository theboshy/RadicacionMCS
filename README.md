# RadicacionMCS
implementacion de micorservicio con radicacion / grpc y rest end point (mux (github.com/gorilla/mux))

Ejecucion de nuevo (o cambios agregados) agregados al protobuf

 `#app/mcs/grpc/proto/radicacionmcs.proto build`

```console
~$ protoc  ./app/mcs/grpc/proto/radicacionmcs.proto --go_out=plugins=grpc:./
```

## Librerias utilizadas 
```console
~$ go get -d .
```

```console

-- proto 3 
~$ go get golang.org/x/net/context
~$ go get google.golang.org/grpc
~$ go get google.golang.org/grpc/reflection
~$ go get github.com/gin-gonic/gin
~$ go get github.com/gin-contrib/cors
~$ go get github.com/gorilla/mux
~$ go get github.com/gorilla/handlers
~$ go get github.com/go-xorm/core
~$ go get github.com/go-xorm/xorm
~$ go get github.com/go-sql-driver/mysql
```


## Configuracion DB

> {raiz}/configuration.json

### Test conexion
> {raiz}/main.go

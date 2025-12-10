# Go web Gorilla

Ejercicio de uso de Gorilla web toolkit en Go

https://gorilla.github.io/

# Dependencias

Paquete gorilla mux, en la terminal ubicados en el directorio del proyecto:

go get github.com/gorilla/mux


# Ejecutar proyecto

Iniciar servicio -> En terminal 1:

go run main.go

Terminar servicio -> En terminal 1:

Ctrl + C

# Prueba incial 

- Verificar endpoint usuarios

Metodo GET En terminal 2:

curl http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

Consultar todo usuario

Metodo POST En terminal 2:

curl -X POST http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

Crear usuario en terminal 2:

curl -X POST -d @crearjs.json http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

{Juan Perez jp@test.com 33114400}


Metodo PATCH En terminal 2:

curl -X PATCH http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

Actualizar usuario

Metodo DELETE En terminal 2:

curl -X DELETE http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

Eliminar usuario

- Verificar endpint cursos -> En terminal 2:

curl http://localhost:3333/cursos

Se verifica en terminal 1 mientras se ejecuta el programa:

got /cursos




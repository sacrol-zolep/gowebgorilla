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

- Verificar endpint usuarios -> En terminal 2:

curl http://localhost:3333/usuarios

Se verifica en terminal 1 mientras se ejecuta el programa:

got /usuarios


- Verificar endpint cursos -> En terminal 2:

curl http://localhost:3333/cursos

Se verifica en terminal 1 mientras se ejecuta el programa:

got /cursos




# go4it

Es la base para crear diferentes aplicaciones que utilicen bases de datos y algunas cosita más.

Es mi SDK base para poder crear herramientas que no requieran tener una web online sino que se manejen por commandos.

## IMPORTANTE

go4it utiliza algunas librerías como DBmanager que se utilizan en gasonline. Tengo que encontrar la forma de poder organizar esos mismos proyectos como una base que pueda ser compartida y utilizada por estos proyectos. Por el momento será copiar y pegar.

# Dev mode

Por defecto core/dbmanager.go incluye base de datos mysql,postgress y sqlite. En caso de no querer alguna deberán borrarse para obtener un ejecutable más ligero.

# filepaths

Carpetas dentro de los archivos

./core: hay código común utilizado para la apps por ejemplo dbmanager para conectarse a las bases de datos o leer archivos json.
./src: archivos específicos del programa.
./cmd: comandos para enviar y ejecutar en la app. Las funciones dentro de src deberán ser llamadas por aquí o en el main directamente.



# light build

al momento de compilar agregar estos dos parámetros para quitar accesorios de debug y que el programa sea más liviano:

    go build -ldflags "-w -s"

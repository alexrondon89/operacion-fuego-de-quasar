# Operacion-fuego-de-quasar
prueba tecnica

##URL en donde este hosteado el servicio
El servicio donde se ejecuta el codigo es AWS y la url publica es: https://r5ovozjea2.execute-api.us-east-1.amazonaws.com/dev

##Como ejecutar el servicio.
Existen dos opciones para ejecutar el servicio.

1) Desde el swaggerhub a traves de la url: 
   https://app.swaggerhub.com/apis-docs/alexrondon1/Ejercicio-meli/1.3
    
2) Desde la terminal, ubicado en la carpeta del ejercicio en cuestion, se puede ingresar el comando: local="true" go run main.go . Donde local="true" es un flag para que el acceso a los servicios de AWS sea a traves de las credenciales locales guardadas en la carpeta ./aws/credentials con el profile [default] 
    
##Documentacion del proyecto.

Se desarrollo una libreria commons donde se encuentran todos los recursos compartidos entre los folders ejercicio_2 y ejercicio_3; estos servicios son:

    -Interfaces, necesarias para definir la implementacion
    -Infraestructura para obtener session en aws y cliente para el recurso DynamoDB en AWS
    -Builders, para la construccion de input a DynamoDB
    -Dto, para definir los requests y outputs
    -Services, donde estan los servicios que implementan las interfaces
    -Models, donde se definen los modelos involucrados
    -Utils, donde se almacenan metodos reutilizables
    -Tests, contienen los mocks de las interfaces aplicadas, necesarios para los tests unitarios

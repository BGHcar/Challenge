# Indexador de Correos Electrónicos
Este proyecto es una solución para el desafío técnico de indexar una base de datos de correos electrónicos y proporcionar una interfaz para buscar información dentro de ella.

# Descripción del Proyecto
El objetivo de este proyecto es crear una aplicación que indexe el contenido de una base de datos de correos electrónicos utilizando ZincSearch y proporcione una interfaz para buscar y visualizar la información de manera eficiente.

# Estructura del Proyecto
-backend: Contiene el código del backend de la aplicación.
-database: Contiene el código para indexar la base de datos de correos electrónicos.
-frontend: Contiene el código del frontend de la aplicación.

# Tecnologías Utilizadas
-Lenguaje Backend: Go
-Base de Datos: ZincSearch
-API Router: chi
-Interfaz de Usuario: Vue 3
-CSS: Tailwind

# Instrucciones de Uso

Instalación

Backend
Instalar dependencias del backend
```
cd backend
```
```
go mod tidy
```
Frontend
Instalar dependencias del frontend
```
cd frontend
```
```
npm install
```

Ejecutar Backend
```
cd backend
```
```
go run main.go
```

Ejecutar Frontend
```
cd frontend
```
```
npm run serve
```
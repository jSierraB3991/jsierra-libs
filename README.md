# jsierra-libs

Este repositorio contiene una colección de funciones útiles para el desarrollo rápido en Golang. Diseñado para facilitar tareas comunes y mejorar la productividad, `jsierra-libs` ofrece una serie de herramientas fáciles de comprender y usar.

## Características

- Funciones útiles para tareas comunes en Golang
- Manipulación de fechas y tiempos
- Manejo de tokens y claims
- Utilidades para paginación
- Conversiones de tipos de datos
- Manejo de archivos y rutas
- Utilidades para HTTP y headers
- Y mucho más...

## Instalación

Para utilizar estas librerías en tu proyecto, puedes clonar el repositorio:

```
git clone https://github.com/jSierraB3991/jsierra-libs.git
```

O importar directamente en tu código Go:

```go
import "github.com/jSierraB3991/jsierra-libs"
```

## Funciones principales

Aquí hay una lista de las funciones principales disponibles en `jsierra-libs`:

### Manipulación de datos y conversiones
- `GenerateCombinations(A []int, K int) [][]int`
- `GetFloatStringToInt(data int) string`
- `GetFloatStringToUInt(data uint) string`
- `GetFloatToInt(data int) float64`
- `GetFloatToUInt(data uint) float64`
- `GetNumberForString(number string) int`
- `GetStringNumberFor(number int) string`
- `GetStringUNumber64For(number uint64) string`
- `GetStringUNumberFor(number uint) string`

### Manejo de fechas y tiempos
- `GetDateByString(date string) (time.Time, error)`
- `GetDateTimeByString(date string) (time.Time, error)`
- `GetStringToTimeLongFormat(date time.Time) string`
- `GetStringToTimeShortFormat(date time.Time) string`
- `ScheduleTask(targetTime time.Time, task func()) error`

### Utilidades HTTP y headers
- `GetHeader(r *http.Request, header string) []string`
- `GetHeaderMail(requet *http.Request) (string, error)`
- `GetLanguageHeader(r *http.Request) string`
- `PublicMiddleWare(route string, method string) bool`

### Manejo de tokens y claims
- `GetClaimByToken(tokenString, claim string) (interface{}, error)`

### Utilidades para paginación
- `GetFirstForPagination(limit, page int) int`
- `GetMaxForPagination(limit, page int) int`

### Manejo de archivos y rutas
- `IsFileExist(dir string) bool`
- `PathsEqual(pattern, path string) bool`
- `ReadFile(templateName string) string`
- `ReplaceTextInFile(templateName string, MapForReplace map[string]string) string`

### Variables de entorno
- `GetDataOfEnviroment(enviromentVariable string) string`
- `GetDataOfEnviromentRequired(enviromentVariable string) (string, error)`

### Manejo de mensajes
- `GetMessage(code string) string`
- `GetMessageParam(code string, accept string, name string) string`

### Tipos de error personalizados
- `InvalidDurationOfTask`
- `StringToDateParseError`

## Uso

Aquí hay algunos ejemplos de cómo usar algunas de las funciones de `jsierra-libs`:

```go
import "github.com/jSierraB3991/jsierra-libs"

// Ejemplo de uso de GetDateByString
date, err := jsierra.GetDateByString("2024-10-07")
if err != nil {
    log.Fatal(err)
}
fmt.Println(date)

// Ejemplo de uso de GetFloatToInt
floatValue := jsierra.GetFloatToInt(42)
fmt.Println(floatValue)

// Ejemplo de uso de IsFileExist
exists := jsierra.IsFileExist("/path/to/file")
fmt.Println("File exists:", exists)
```

## Contribución

Las contribuciones son bienvenidas. Por favor, abre un issue para discutir cambios mayores antes de crear un pull request.


## Contacto
Para cualquier pregunta, sugerencia o contribución, puedes contactar al autor:

GitHub: jSierraB3991
Email: eliotandelon@gmail.com, judas3991@gmail.com

No dudes en abrir un issue en el repositorio si encuentras algún problema o tienes alguna sugerencia de mejora.

## Licencia

Este proyecto está licenciado bajo los términos de la [Licencia Pública General GNU, versión 3 (GPLv3)](https://www.gnu.org/licenses/gpl-3.0.html).

### Resumen de la licencia:

- Puedes copiar, distribuir y modificar este proyecto siempre que mantengas la misma licencia.
- Cualquier modificación realizada y distribuida debe estar bajo los mismos términos.
- No hay garantías sobre el uso del software.

Puedes encontrar una copia completa de la licencia en el archivo [LICENSE](./LICENSE) en el repositorio de este proyecto.
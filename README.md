# Sistema de Monitoreo de Recursos del Sistema

Este proyecto en Go permite monitorear el uso de recursos del sistema (CPU, RAM y Disco) y ofrece dos funcionalidades principales: un servidor web para visualizar las estadísticas en tiempo real y la generación de un reporte en formato JSON con los datos obtenidos.

## Características

- Monitoreo en tiempo real de los recursos del sistema (CPU, RAM, Disco).
- Servidor web que muestra las estadísticas del sistema en un formato HTML, con auto-refresco cada 5 segundos.
- Generación de reporte en formato JSON con las estadísticas del sistema.

## Requisitos

- **Go 1.18 o superior**: Asegúrate de tener Go instalado en tu máquina.
- **Dependencias**: Para ejecutar el proyecto, se utilizan las siguientes librerías:
  - `gopsutil`: Para obtener las estadísticas del sistema (CPU, RAM, Disco).
  - `fatih/color`: Para manejo de colores en consola.

## Uso

### Ejecutar el Servidor Web

Al ejecutar el comando con el flag `--web`, se inicia un servidor web que muestra las estadísticas del sistema en tiempo real. El servidor corre en `http://localhost:8000`, y la página se actualiza automáticamente cada 5 segundos.

### Generar Reporte en JSON

Si prefieres guardar las estadísticas en un archivo JSON, puedes ejecutar el proyecto con el flag `--report`. Esto generará un archivo llamado `system_report.json` con las estadísticas del sistema en formato estructurado.

## Comandos

- `--web`: Inicia el servidor web para visualizar las estadísticas del sistema en tiempo real.
- `--report`: Genera un reporte en formato JSON con las estadísticas actuales del sistema.

## Ejemplo de Reporte JSON

El archivo JSON generado con el flag `--report` tendrá la siguiente estructura:

- `cpu_usage`: El porcentaje de uso de la CPU.
- `ram_usage`: El porcentaje de uso de la RAM.
- `disk_usage`: El porcentaje de uso del disco.
- `time_stamp`: La fecha y hora en la que se generaron las estadísticas.

## Estructura del Código

### SystemStats

Se define una estructura llamada `SystemStats` que contiene los siguientes campos:

- `CPUUsage`: Porcentaje de uso de la CPU.
- `RAMUsage`: Porcentaje de uso de la RAM.
- `DISKUsage`: Porcentaje de uso del disco.
- `TIMEStamp`: Fecha y hora en la que se obtuvieron las estadísticas.

### getSystemStats

Esta función recopila las estadísticas del sistema utilizando la librería `gopsutil`. Obtiene el uso de la CPU, la memoria y el disco, y genera un timestamp con la fecha y hora actual.

### Servidor Web

Cuando se ejecuta con el flag `--web`, el proyecto arranca un servidor web que se ejecuta en `http://localhost:8000`. En esta página, los usuarios pueden ver las estadísticas del sistema en tiempo real, las cuales se actualizan cada 5 segundos.

### Generación del Reporte

Si se utiliza el flag `--report`, el proyecto genera un archivo JSON con las estadísticas del sistema. Este archivo se guarda con el nombre `system_report.json` en el directorio actual.

### Manejo de Errores

El código maneja posibles errores al obtener las estadísticas del sistema o generar el reporte. Si ocurre un error, se muestra un mensaje en la consola para informar al usuario sobre el problema.

## Licencia

Este proyecto está licenciado bajo la licencia MIT. Puedes ver los detalles completos de la licencia en el archivo `LICENSE`.

# Task-CLI

`Task-CLI` es una herramienta de línea de comandos (CLI) para gestionar tareas de forma simple y eficiente. Puedes agregar, listar, actualizar y cambiar el estado de tus tareas directamente desde la terminal.

## Características

- **Agregar tareas**: Crea una nueva tarea con facilidad.
- **Listar tareas**: Muestra todas las tareas con sus respectivos estados.
- **Actualizar tareas**: Cambia el nombre de una tarea existente.
- **Cambiar estado**:
  - Marcar como "completada".
  - Marcar como "en progreso".
- **Persistencia**: Las tareas se guardan automáticamente para su uso futuro.

## Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/tuusuario/task-cli.git
   ```

2. Navega al directorio del proyecto:
   ```bash
   cd task-cli
   ```

3. Compila el binario:
   ```bash
   go build -o task-cli
   ```

4. Ejecuta el binario desde tu terminal:
   ```bash
   ./task-cli
   ```

## Uso

### Comandos disponibles

- **Agregar una tarea**:
  ```bash
  ./task-cli add "Nombre de la tarea"
  ```

- **Listar todas las tareas**:
  ```bash
  ./task-cli list
  ```

- **Actualizar el nombre de una tarea**:
  ```bash
  ./task-cli update <id_tarea> "Nuevo nombre"
  ```

- **Marcar una tarea como completada**:
  ```bash
  ./task-cli mark-done <id_tarea>
  ```

- **Marcar una tarea como en progreso**:
  ```bash
  ./task-cli mark-in-progress <id_tarea>
  ```

### Ejemplo

```bash
$ ./task-cli add "Terminar proyecto CLI"
$ ./task-cli list
1 - Terminar proyecto CLI
$ ./task-cli mark-done 1
$ ./task-cli list
1 - Terminar proyecto CLI ✅
```

## Desarrollo

Si deseas contribuir o mejorar este proyecto:

1. Asegúrate de tener instalado [Go](https://golang.org/doc/install).
2. Haz un fork de este repositorio y clónalo.
3. Crea una nueva rama para tus cambios:
   ```bash
   git checkout -b nombre-de-tu-rama
   ```
4. Realiza los cambios necesarios y haz un commit:
   ```bash
   git commit -m "Descripción de los cambios"
   ```
5. Envía tus cambios a tu repositorio fork y abre un Pull Request.

## Licencia

Este proyecto está licenciado bajo la [MIT License](LICENSE).

---

¡Gracias por usar Task-CLI! Si tienes preguntas o sugerencias, no dudes en abrir un issue.

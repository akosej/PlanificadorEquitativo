# Planificador de Guardias

Este repositorio alberga un programa en Go diseñado para planificar las guardias de dos grupos de personas de manera equitativa y considerando restricciones específicas. La herramienta toma en cuenta la fecha de inicio, la cantidad deseada de guardias por persona y las fechas en las que cada persona no puede tomar guardia debido a afectaciones personales.

## Características Clave

- Planificación equitativa de guardias entre dos grupos de personas.
- Consideración de restricciones, como la cantidad deseada de guardias por persona.
- Gestión de fechas en las que las personas no pueden tomar guardia debido a afectaciones personales.
- Generación de informes detallados de la planificación, incluyendo la fecha y las parejas de guardia asignadas.
- Resumen de guardias asignadas a cada persona en cada grupo.

## Instrucciones de Uso

1. Clone este repositorio en su máquina local.
2. Configure la configuración en un archivo JSON, especificando la fecha de inicio, los grupos de personas, las afectaciones personales y la cantidad deseada de guardias por persona.
3. Ejecute el programa en Go para planificar las guardias.
4. Obtenga los resultados en formato JSON, incluyendo un resumen detallado de las guardias asignadas.

Este proyecto es útil para cualquier organización que requiera una planificación eficiente de guardias, como hospitales, equipos de soporte técnico o centros de atención de emergencia.

## Archivo `config.json`

El archivo `config.json` es un archivo de configuración esencial para ejecutar el Planificador de Guardias. Permite especificar los parámetros clave para la planificación de guardias. A continuación, se describe la estructura y los valores admitidos en el archivo `config.json`:

### Ejemplo de `config.json`

```json
    {
        "fechaInicial": "01/11/2023",
        "grupo1": [
            "G1_Persona1",
            "G1_Persona2",
            "G1_Persona3",
            "G1_Persona4",
            "G1_Persona5"
        ],
        "grupo2": [
            "G2_Persona1",
            "G2_Persona2",
            "G2_Persona3",
            "G2_Persona4",
            "G2_Persona5"
        ],
        "guardiasDeseadas": 6,
        "afectacionesPorPersona": {
            "G1_Persona1": [
                "03/11/2023",
                "11/11/2023"
            ],
            "G2_Persona2": [
                "05/11/2023"
            ]
        }
    }
```

## Descripción de los Campos

- **fechaInicial**: La fecha en la que comenzará la planificación de guardias en formato "dd/mm/aaaa".

- **grupo1**: Un arreglo de las personas en el Grupo 1 que participarán en la planificación.

- **grupo2**: Un arreglo de las personas en el Grupo 2 que participarán en la planificación.

- **guardiasDeseadas**: El número deseado de guardias que cada persona debe realizar durante el período especificado.

- **afectacionesPorPersona**: Un mapa que contiene afectaciones por persona. Cada persona puede especificar las fechas en las que no está disponible para tomar guardias. Las fechas de afectación se indican en formato "dd/mm/aaaa".

Asegúrate de que el archivo `config.json` esté correctamente configurado antes de ejecutar el Planificador de Guardias. Puedes personalizar los valores según las necesidades de tu organización y los requisitos específicos de planificación de guardias.


## Contribuciones

Siéntase libre de contribuir a este proyecto mediante solicitudes de extracción. Sus contribuciones son bienvenidas para mejorar la funcionalidad y la usabilidad de esta herramienta.

## Licencia

Este proyecto se distribuye bajo la Licencia MIT. Consulte el archivo [LICENSE.md](LICENSE.md) para obtener más detalles.

¡Gracias por usar el Planificador de Guardias!

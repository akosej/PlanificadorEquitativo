package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	FechaInicial           string              `json:"fechaInicial"`
	Grupo1                 []string            `json:"grupo1"`
	Grupo2                 []string            `json:"grupo2"`
	GuardiasDeseadas       int                 `json:"guardiasDeseadas"`
	AfectacionesPorPersona map[string][]string `json:"afectacionesPorPersona"`
}

type ResultadoGuardias struct {
	Fecha         string
	DiaDeLaSemana string
	PersonaGrupo1 string
	PersonaGrupo2 string
}

type ResumenGuardias struct {
	Grupo1 map[string]int
	Grupo2 map[string]int
}

func main() {
	// Cargar la configuración desde un archivo JSON.
	config, err := cargarConfiguracion("config.json")
	if err != nil {
		fmt.Println("Error al cargar la configuración:", err)
		return
	}

	// Parsea la fecha de inicio.
	fecha, err := time.Parse("02/01/2006", config.FechaInicial)
	if err != nil {
		fmt.Println("Error al parsear la fecha:", err)
		return
	}

	// Map para realizar un seguimiento de las guardias asignadas a cada persona.
	guardiasPorPersona := make(map[string]int)
	ultimaGuardia := make(map[string]time.Time)

	// Cantidad de guardias que deben hacer cada persona por igual.
	guardiasDeseadas := config.GuardiasDeseadas

	// Almacena los resultados en una lista.
	var resultados []ResultadoGuardias

	// Map para realizar un seguimiento de las guardias asignadas a cada persona.
	resumenGuardias := ResumenGuardias{
		Grupo1: make(map[string]int),
		Grupo2: make(map[string]int),
	}

	// Imprime los 30 días siguientes con el día de la semana correspondiente y parejas de guardia equitativas.
	for i := 0; i < 31; i++ {
		diaDeLaSemana := fecha.Weekday()
		personaGrupo1, personaGrupo2 := generarParejasEquitativas(config.Grupo1, config.Grupo2, guardiasPorPersona, guardiasDeseadas, ultimaGuardia, fecha, config.AfectacionesPorPersona)

		// Calcula el conteo de guardias por persona y actualiza el resumen.
		contadorGrupo1 := guardiasPorPersona[personaGrupo1]
		contadorGrupo2 := guardiasPorPersona[personaGrupo2]
		resumenGuardias.Grupo1[personaGrupo1] = contadorGrupo1
		resumenGuardias.Grupo2[personaGrupo2] = contadorGrupo2

		resultado := ResultadoGuardias{
			Fecha:         fecha.Format("02/01/2006"),
			DiaDeLaSemana: diaDeLaSemana.String(),
			PersonaGrupo1: personaGrupo1,
			PersonaGrupo2: personaGrupo2,
		}
		resultados = append(resultados, resultado)

		// Avanza al siguiente día.
		fecha = fecha.AddDate(0, 0, 1)
	}

	// Imprime la lista de personas divididas por grupos con el número de guardias planificadas.
	fmt.Println("Grupo 1:")
	for _, persona := range config.Grupo1 {
		fmt.Printf("%s: %d guardias\n", persona, guardiasPorPersona[persona])
	}

	fmt.Println("Grupo 2:")
	for _, persona := range config.Grupo2 {
		fmt.Printf("%s: %d guardias\n", persona, guardiasPorPersona[persona])
	}

	// Guarda los resultados en un archivo JSON, incluyendo el resumen de guardias.
	resultadosJSON, err := json.Marshal(resultados)
	if err != nil {
		fmt.Println("Error al serializar resultados a JSON:", err)
		return
	}

	resumenJSON, err := json.Marshal(resumenGuardias)
	if err != nil {
		fmt.Println("Error al serializar resumen de guardias a JSON:", err)
		return
	}

	archivoResultados, err := os.Create("resultados.json")
	if err != nil {
		fmt.Println("Error al crear el archivo de resultados:", err)
		return
	}
	defer archivoResultados.Close()

	archivoResultados.Write(resultadosJSON)

	archivoResumen, err := os.Create("resumen_guardias.json")
	if err != nil {
		fmt.Println("Error al crear el archivo de resumen de guardias:", err)
		return
	}
	defer archivoResumen.Close()

	archivoResumen.Write(resumenJSON)

	fmt.Println("Resultados guardados en resultados.json")
	fmt.Println("Resumen de guardias guardado en resumen_guardias.json")
}

func cargarConfiguracion(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func generarParejasEquitativas(grupo1, grupo2 []string, guardiasPorPersona map[string]int, guardiasDeseadas int, ultimaGuardia map[string]time.Time, fecha time.Time, afectacionesPorPersona map[string][]string) (string, string) {
	// Baraja aleatoriamente los grupos para formar las parejas.
	rand.Shuffle(len(grupo1), func(i, j int) {
		grupo1[i], grupo1[j] = grupo1[j], grupo1[i]
	})
	rand.Shuffle(len(grupo2), func(i, j int) {
		grupo2[i], grupo2[j] = grupo2[j], grupo2[i]
	})

	personaGrupo1 := ""
	personaGrupo2 := ""

	// Función para verificar si una persona puede tomar guardia en una fecha específica.
	puedeTomarGuardia := func(persona string, fecha time.Time) bool {
		ultimaGuardia, ok := ultimaGuardia[persona]
		if !ok {
			return true
		}
		return fecha.Sub(ultimaGuardia).Hours() >= 24*4
	}

	// Función para verificar si una persona tiene una afectación en la fecha dada.
	tieneAfectacion := func(persona string, fecha time.Time) bool {
		afectaciones, ok := afectacionesPorPersona[persona]
		if !ok {
			return false
		}
		for _, afectacion := range afectaciones {
			afectacionFecha, _ := time.Parse("02/01/2006", afectacion)
			if fecha.Equal(afectacionFecha) {
				return true
			}
		}
		return false
	}

	// Encuentra las parejas de guardia equitativas, evitando las fechas afectadas.
	for i := 0; i < len(grupo1); i++ {
		persona1 := grupo1[i]
		if guardiasPorPersona[persona1] < guardiasDeseadas && puedeTomarGuardia(persona1, fecha) && !tieneAfectacion(persona1, fecha) {
			for j := 0; j < len(grupo2); j++ {
				persona2 := grupo2[j]
				if guardiasPorPersona[persona2] < guardiasDeseadas && puedeTomarGuardia(persona2, fecha) && !tieneAfectacion(persona2, fecha) {
					personaGrupo1 = persona1
					personaGrupo2 = persona2
					guardiasPorPersona[persona1]++
					guardiasPorPersona[persona2]++
					ultimaGuardia[persona1] = fecha
					ultimaGuardia[persona2] = fecha
					break
				}
			}
		}
		if personaGrupo1 != "" && personaGrupo2 != "" {
			break
		}
	}

	return personaGrupo1, personaGrupo2
}

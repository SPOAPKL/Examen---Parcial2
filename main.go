package main

import (
	"errors"
	"fmt"
	"time"
)

type Medicamento struct {
	nombre           string
	fabricante       string
	fechaFabricacion time.Time
	vidaUtilMeses    int
}

func NuevoMedicamento(nombre, fabricante string, fechaFab time.Time, vidaUtil int) (*Medicamento, error) {
	if nombre == "" || fabricante == "" {
		return nil, errors.New("El nombre del producto y el fabricante no pueden estar vacios")
	}
	if vidaUtil <= 0 {
		return nil, errors.New("La vida util tiene que ser un numero positivo")
	}
	return &Medicamento{nombre, fabricante, fechaFab, vidaUtil}, nil
}

func (m *Medicamento) GetNombre() string { return m.nombre }
func (m *Medicamento) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("Nombre no valido")
	}
	m.nombre = nombre
	return nil
}
func (m *Medicamento) GetFabricante() string { return m.fabricante }

func (m *Medicamento) CalcularFechaCaducidad() time.Time {
	return m.fechaFabricacion.AddDate(0, m.vidaUtilMeses, 0)
}

type Tableta struct {
	Medicamento
	dosisPorTableta float64
	requiereReceta  bool
}

func NuevaTableta(nombre, fabricante string, fechaFab time.Time, vidaUtil int, dosis float64, receta bool) (*Tableta, error) {
	base, err := NuevoMedicamento(nombre, fabricante, fechaFab, vidaUtil)
	if err != nil {
		return nil, err
	}
	if dosis <= 0 {
		return nil, errors.New("La dosis tiene que ser mayor a cero")
	}
	return &Tableta{*base, dosis, receta}, nil
}

type Jarabe struct {
	Medicamento
	volumen float64
	sabor   string
}

func NuevoJarabe(nombre, fabricante string, fechaFab time.Time, vidaUtil int, vol float64, sabor string) (*Jarabe, error) {
	base, err := NuevoMedicamento(nombre, fabricante, fechaFab, vidaUtil)
	if err != nil {
		return nil, err
	}
	if vol <= 0 {
		return nil, errors.New("El volumen tiene que ser mayor a cero")
	}
	return &Jarabe{*base, vol, sabor}, nil
}

func (t *Tableta) MostrarDetalles() {
	fmt.Printf("TABLETA Nombre: %-15s - Lab: %-10s - Dosis: %6.2fmg - Receta: %-5v - Caducidad: %s\n",
		t.nombre, t.fabricante, t.dosisPorTableta, t.requiereReceta, t.CalcularFechaCaducidad().Format("28/06/2006"))
}
func (j *Jarabe) MostrarDetalles() {
	fmt.Printf("JARABE Nombre: %-15s - Lab: %-10s - Volumen: %8.2fml - Sabor: %-5s - Caducidad: %s\n",
		j.nombre, j.fabricante, j.volumen, j.sabor, j.CalcularFechaCaducidad().Format("28/06/2028"))
}

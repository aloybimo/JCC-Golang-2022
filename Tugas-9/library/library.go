package library

import (
	"fmt"
	"math"
	"strconv"
	//"strings"
)

type SegitigaSamaSisi struct{
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas*s.Tinggi/2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3*s.Alas
}
  
type PersegiPanjang struct{
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang*p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2*(p.Panjang+p.Lebar)
}

type Tabung struct{
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi*t.JariJari*t.JariJari*t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	var a = 2*math.Pi*t.JariJari*t.JariJari
	var b = 2*math.Pi*t.JariJari*t.Tinggi
	return a+b
}

type Balok struct{
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang*b.Lebar*b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	var a = b.Panjang*b.Tinggi
	var c = b.Lebar*b.Tinggi
	var d = b.Panjang*b.Lebar
	return float64(2*(a+c+d))
}

type HitungbangunDatar interface{
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface{
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct{
	Name, Brand string
	Year int
	Colors []string
}

func (p Phone) TampilData(nama string) string{
	a := ""
	if nama == "name" {
		a += "name : " + p.Name 
	} else if  nama == "brand" {
		a += "brand : "+p.Brand 
	} else if  nama == "year" {
		a += "year : "+ strconv.Itoa(p.Year)
	} else if  nama == "colors" {
		a += "colors : "
		for _, item := range p.Colors{
			a += item + ", "
		}
	}
	fmt.Println(a)
	return a
}

type Data interface {
	TampilData(string) string
}

func LuasPersegi(sisi int, kondisi bool) interface{} {
	if sisi != 0 && kondisi == true {
		var a = sisi*sisi
		return "luas persegi dengan sisi " + strconv.Itoa(sisi) + " adalah " + strconv.Itoa(a) + " cm"
	} else if sisi != 0 && kondisi == false {
		var a = sisi*sisi
		return a
	} else if sisi == 0 && kondisi == true {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && kondisi == false {
		return nil
	}
	return nil
}
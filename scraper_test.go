package silscraper

import (
	"os"
	"testing"
)

//declaring test files
var doc1 = LegislatorProfile{
	Nombre:              "Estefan Garfias, José Antonio",
	UrlReference:        "http://sil.gobernacion.gob.mx/Librerias/pp_PerfilLegislador.php?SID=&Referencia=1",
	UrlFoto:             "http://sil.gobernacion.gob.mx/Archivos/Fotos/1.jpg",
	Cargo:               "Diputado",
	PorlaLegis:          "LVII Legislatura",
	Estatus:             "BAJA",
	Partido:             "PRI",
	NacimientoFecha:     "16/12/1954",
	NacimientoEstado:    "OAXACA",
	NacimientoCiudad:    "Tehuantepec",
	PrinEleccion:        "Mayoría Relativa",
	ZonaEntidad:         "Oaxaca",
	ZonaDistrito:        "5 (Santo Domingo Tehuantepec)",
	ZonaCircunscripcion: "",
	FechaProtesta:       "01/09/1997",
	Suplente:            "Acevedo Gutiérrez, Agustina",
	UltEstudios:         "Licenciatura",
	PrepAcademic:        "Económico-financiera",
}

var doc2 = LegislatorProfile{
	Nombre:              "Castillo Juárez, Laura Itzel",
	UrlReference:        "http://sil.gobernacion.gob.mx/Librerias/pp_PerfilLegislador.php?SID=&Referencia=2",
	UrlFoto:             "http://sil.gobernacion.gob.mx/Archivos/Fotos/2.jpg",
	Cargo:               "Diputada",
	PorlaLegis:          "LVII Legislatura",
	Estatus:             "BAJA",
	Partido:             "PRD",
	NacimientoFecha:     "16/11/1957",
	NacimientoEstado:    "Distrito Federal",
	NacimientoCiudad:    "México",
	PrinEleccion:        "Representación Proporcional",
	ZonaEntidad:         "",
	ZonaDistrito:        "",
	ZonaCircunscripcion: "",
	FechaProtesta:       "01/09/1997",
	Suplente:            "Alemán García, José Antonio",
	UltEstudios:         "",
	PrepAcademic:        "",
}

func TestScraper(t *testing.T) {
	file, err := os.Open("./datatest/test.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	_, err = docQueryFromFile(file)
	if err != nil {
		t.Error(err)
	}
	t.Error("Nothing has been wrote yet")

}

func TestScrapProfile(t *testing.T) {
	file, err := os.Open("./datatest/test.html")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	html, err := docQueryFromFile(file)
	if err != nil {
		t.Error(err)
	}

	var profile LegislatorProfile
	scrapProfile(&profile, html)
	if profile.Nombre != doc1.Nombre {
		t.Error("Expected " + doc1.Nombre + " got " + profile.Nombre)
	}
}

func TestScrapTrayectorias(t *testing.T) {
	t.Error("Nothing has been wrote yet")
}

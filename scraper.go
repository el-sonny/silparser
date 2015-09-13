package silscraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func ScrapLegislatorById(id int) (profile LegislatorProfile, err error) {
	url := fmt.Sprintf(urlBase, id)
	urlfoto := fmt.Sprintf(urlFoto, id)
	profile.UrlReference = url
	profile.UrlFoto = urlfoto
	html, err := docQueryFromUrl(url)
	if err != nil {
		return
	}
	scrapProfile(&profile, html)
	scrapTrayectorias(&profile, html)
	return
}
func scrapTrayectorias(profile *LegislatorProfile, html *goquery.Document) {
	var activated bool
	var titulo string
	const trayectoria = "trayectoria"
	html.Find("table").Each(func(i int, table *goquery.Selection) {
		if activated {
			activated = false
			var trayectoria Trayectoria
			table.Find(".tddatosazul").Each(func(j int, td *goquery.Selection) {
				if j%3 == 0 {
					trayectoria.From = td.Text()
				} else if j%3 == 1 {
					trayectoria.To = td.Text()
				} else {
					trayectoria.Action = td.Text()
					trayectoria.Tipo = titulo
					profile.Trayectorias = append(profile.Trayectorias, trayectoria)
				}
			})
		} else {
			// si trayectoria en el body se rompe
			titulo = normalize(table.Text())
			if strings.Contains(titulo, trayectoria) {
				activated = true
			}
		}
	})
	return

}

func scrapProfile(profile *LegislatorProfile, html *goquery.Document) {
	html.Find("tr").Each(func(i int, tr *goquery.Selection) {
		key := tr.Find(".tdcriterio").First().Text()
		val := tr.Find(".tddatosazul").First().Text()
		key = normalize(key)
		switch key {
		case "nombre":
			p, n, pl := scrapNombre(val)
			profile.Cargo = p
			profile.Nombre = n
			profile.PorlaLegis = pl
			break
		case "estatus":
			profile.Estatus = trimSpace(val)
			break
		case "partido":
			profile.Partido = trimSpace(val)
			break
		case "nacimiento":
			f, e, c := scrapNacimiento(val)
			profile.NacimientoFecha = f
			profile.NacimientoEstado = e
			profile.NacimientoCiudad = c
			break

		case "principio de eleccion":
			profile.PrinEleccion = trimSpace(val)
			break
		case "zona":
			e, d, c := scrapZona(val)
			profile.ZonaEntidad = e
			profile.ZonaDistrito = d
			profile.ZonaCircunscripcion = c
			break
		case "toma de protesta":
			profile.FechaProtesta = trimSpace(val)
			break
		case "suplente":
			profile.Suplente = trimSpace(val)
			break
		case "ultimo grado de estudios":
			profile.UltEstudios = trimSpace(val)
			break
		case "preparacion academica":
			profile.PrepAcademic = trimSpace(val)
			break
		default:
			break
		}
	})
}

func scrapNombre(val string) (cargo, nombre, porla string) {
	vals := strings.Split(val, ":")
	cargo = trimSpace(vals[0])
	oterVals := strings.Split(vals[1], "por la")
	nombre = trimSpace(oterVals[0])
	porla = trimSpace(oterVals[1])
	return
}

func scrapNacimiento(nacimiento string) (fecha, estado, ciudad string) {
	nacimiento = normalize(nacimiento)
	reFecha := regexp.MustCompile("Fecha:\\s+([0-9]+/[0-9]+/[0-9]+)")
	reEstado := regexp.MustCompile("Entidad:\\s+([a-z|A-Z|\\s]+)Ciudad")
	reCiudad := regexp.MustCompile("Ciudad:\\s+([a-z|A-Z]+)")
	f := reFecha.FindStringSubmatch(nacimiento)
	e := reEstado.FindStringSubmatch(nacimiento)
	c := reCiudad.FindStringSubmatch(nacimiento)
	if len(f) >= 2 {
		fecha = f[1]
	}
	if len(e) >= 2 {
		estado = e[1]
	}
	if len(c) >= 2 {
		ciudad = c[1]
	}
	return
}

func scrapZona(zona string) (entidad, distrito, circunscripcion string) {
	zona = normalize(zona)
	reEntidad := regexp.MustCompile("entidad([a-z|A-Z|\\s]+)distrito")
	reDistrito := regexp.MustCompile("distrito([0-9|a-z|A-Z|\\s]+)")
	e := reEntidad.FindStringSubmatch(zona)
	d := reDistrito.FindStringSubmatch(zona)
	if len(e) >= 2 {
		entidad = e[1]
	}
	if len(d) >= 2 {
		distrito = d[1]
	}
	return
}

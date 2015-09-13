# silparser
silparser genera un json con la info del profile 

tomado de http://sil.gobernacion.gob.mx/Librerias/pp_PerfilLegislador.php?SID=\u0026Referencia=6911048

Ejemplo :
```js
[
	{
		"Nombre": "Borge Angulo, Roberto",
		"Cargo": "Diputado Propietario",
		"PorlaLegis": "LXI Legislatura",
		"Estatus": "BAJA",
		"Partido": "PRI",
		"NacimientoFecha": "",
		"NacimientoEstado": "",
		"NacimientoCiudad": "",
		"PrinEleccion": "Mayoría Relativa",
		"ZonaEntidad": " quintana roo",
		"ZonaDistrito": " 1 playa del carmen",
		"ZonaCircunscripcion": "",
		"FechaProtesta": "29/08/2009",
		"Suplente": "Hurtado Vallejo, Susana",
		"UltEstudios": "Licenciatura",
		"PrepAcademic": "Administración y contaduría",
		"Trayectorias": [
			{
				"From": "2002",
				"To": "2003",
				"Action": "Jefe de relaciones públicas de la Secretaría de Desarrollo Económico de Quintana Roo. ",
				"Tipo": "trayectoria administrativa"
			},
			{
				"From": "2005",
				"To": "2005",
				"Action": "Secretario particular del gobernador de Quintana Roo.",
				"Tipo": "trayectoria administrativa"
			},
			{
				"From": "2005",
				"To": "2006",
				"Action": "Tesorero del gobierno de Quintana Roo.",
				"Tipo": "trayectoria administrativa"
			},
			{
				"From": "2006",
				"To": "2008",
				"Action": "Oficial mayor del gobierno de Quintana Roo.",
				"Tipo": "trayectoria administrativa"
			},
			{
				"From": "1997",
				"To": "",
				"Action": "Integrante del Frente Juvenil Revolucionario (FJR) del PRI. ",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "1998",
				"To": "1998",
				"Action": "Coordinador de eventos especiales de la Asociación de Estudiantes del ITESM en Quintana Roo. ",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "1999",
				"To": "1999",
				"Action": "Secretario de la Asociación de Estudiantes del ITESM en Quintana Roo.",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2001",
				"To": "",
				"Action": "Integrante de la Confederación Nacional de Organizaciones Populares (CNOP).",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2001",
				"To": "2001",
				"Action": "Vicepresidente de la Asociación de Estudiantes del ITESM en Quintana Roo.",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2003",
				"To": "",
				"Action": "Colaborador en campañas para la presidencia municipal de Cozumel, diputación federal y gubernatura de Quintana Roo.",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2003",
				"To": "2004",
				"Action": "Secretario privado de diputado federal de Quintana Roo en la LIX legislatura.",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2008",
				"To": "2009",
				"Action": "Presidente del  CDE  del PRI en Quintana Roo.",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "2010",
				"To": "2010",
				"Action": "Candidato al gobierno de Quintana Roo. ",
				"Tipo": "trayectoria politica"
			},
			{
				"From": "1997",
				"To": "2001",
				"Action": "Licenciatura en Administración de Empresas  por el Instituto Tecnológico de Estudios Superiores de Monterrey (ITESM), Campus Monterrey.",
				"Tipo": "trayectoria academica"
			},
			{
				"From": "2000",
				"To": "2000",
				"Action": "Curso de Valores en la Administración en la Universidad Pontificia de Comillas, Madrid, España.",
				"Tipo": "trayectoria academica"
			},
			{
				"From": "2001",
				"To": "2001",
				"Action": "Programa de Liderazgo Empresarial (PLEI), misión Euro Asia.",
				"Tipo": "trayectoria academica"
			},
			{
				"From": "",
				"To": "",
				"Action": "Asistente de administración en Viajes Caribe Tours, en Cancún, Quintana Roo.",
				"Tipo": "trayectoria empresarial/iniciativa privada"
			}
		],
		"UrlReference": "http://sil.gobernacion.gob.mx/Librerias/pp_PerfilLegislador.php?SID=\u0026Referencia=6911048",
		"UrlFoto": "http://sil.gobernacion.gob.mx/Archivos/Fotos/6911048.jpg"
	}
]
```
#Ejemplo de uso en GO
```go
package main

import (
  "encoding/json"
  "fmt"
  "github.com/tugorez/silscraper"
  "io/ioutil"
  "log"
)

func main() {
  var docs []silscraper.LegislatorProfile
 
  //seleciona un rango de ids para buscar y descargar perfiles
  idfrom:=6911048
  idto:=6911149
  for i := idfrom; i < idto; i++ {
    profile, err := silscraper.ScrapLegislatorById(i)
    if err != nil {
      log.Fatal(fmt.Sprintf("getting the doc %d", i))
    }
    if profile.Nombre == "" {
      continue
    }
    docs = append(docs, profile)
  }
  //
  file, err := json.MarshalIndent(docs, "", "\t")
  if err != nil {
    log.Fatal(err)
  }
  ioutil.WriteFile("testall.json", file, 0644)

}
```


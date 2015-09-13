package silscraper

type LegislatorProfile struct {
	Nombre              string
	Cargo               string
	PorlaLegis          string
	Estatus             string
	Partido             string
	NacimientoFecha     string
	NacimientoEstado    string
	NacimientoCiudad    string
	PrinEleccion        string
	ZonaEntidad         string
	ZonaDistrito        string
	ZonaCircunscripcion string
	FechaProtesta       string
	Suplente            string
	UltEstudios         string
	PrepAcademic        string
	Trayectorias        []Trayectoria
	UrlReference        string
	UrlFoto             string
}
type Trayectoria struct {
	From   string
	To     string
	Action string
	Tipo   string
}

const urlBase = "http://sil.gobernacion.gob.mx/Librerias/pp_PerfilLegislador.php?SID=&Referencia=%d"
const urlFoto = "http://sil.gobernacion.gob.mx/Archivos/Fotos/%d.jpg"

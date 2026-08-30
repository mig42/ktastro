package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type ImportCmd struct {
	Force bool `help:"Vaciar base de datos antes de importar." short:"f"`

	DirCatastro   string `required:"" name:"dir-catastro" help:"Directorio de ficheros de catastro." type:"existingdir" short:"c"`
	DirShapefiles string `required:"" name:"dir-shapefiles" help:"Directorio de ficheros shapefile." type:"existingdir" short:"s"`
	Out           string `required:"" name:"output" help:"Ruta de destino para la importación." type:"path" short:"o"`
}

func (i *ImportCmd) Run(ctx *Context) error {
	fmt.Printf("Importando entradas de catastro desde '%s' y '%s' en '%s'\n", i.DirCatastro, i.DirShapefiles, i.Out)
	return nil
}

type CLI struct {
	Debug bool `help:"Habilitar modo de depuración." short:"d"`

	Import ImportCmd `cmd:"" default:"true" help:"Importar entradas de catastro."`
}

func Run() {
	var cli CLI
	ctx := kong.Parse(
		&cli,
		kong.Name("ktastro"),
		kong.Description("Herramienta CLI para importar archivos CAT y Shapefiles a un archivo SQLite."),
		kong.UsageOnError(),
	)

	_ = ctx.Run(&Context{Debug: cli.Debug})
}

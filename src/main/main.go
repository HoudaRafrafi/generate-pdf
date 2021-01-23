package main

import (
	"log"								 
	"gopdf-master"
	"io/ioutil"
)

func main() {

	var err error

	files, err := ioutil.ReadDir(".")							//liste tout les fichiers du répertoire courant 
	if err != nil {
		log.Fatal(err)
	}

	pdf := gopdf.GoPdf{}									//Création d'un fichier pdf
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 }) 

	for _, f := range files {								// parcourt les fichiers du répertoire courant
		if f.Name() != "main.go" {	
																
			pdf.AddPage()								// insertion des images dans chaque page pdf 
			pdf.Image(f.Name(), 0, 50, nil) 				
		}
	}

	pdf.WritePdf("image.pdf")								//l'écriture dans le fichier pdf 
}

package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 5 || ocena > 10 {
		fmt.Println("Neveljavna ocena!")
		return
	}

	student, ok := studenti[vpisnaStevilka]
	if ok {
		student.ocene = append(student.ocene, ocena)
		//JF PAZI, student je v tem primeru kopija, in ko jo spremenis se to ne piše v redovalnico! Zato jo samodejno vpiši!
		studenti[vpisnaStevilka] = student
		fmt.Printf("Dodal oceno %d k študentu: %s!\n", ocena, vpisnaStevilka)
	} else {
		fmt.Println("Student z to vpisno številko ne obstaja!")
	}

}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if ok {
		if len(student.ocene) < 6 {
			return 0
		}
		var vsota float64 = 0
		for _, v := range student.ocene {
			vsota += float64(v)
		}

		return vsota / float64(len(student.ocene))
	} else {
		return -1
	}
}

func IzpisRedovalnice(studenti map[string]Student) {
	for k, v := range studenti {
		fmt.Print(k, " - ")
		fmt.Println(v.ime, v.priimek, ": ", v.ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for k, v := range studenti {
		fmt.Print(k, " -  ")
		avg := povprecje(studenti, k)
		fmt.Print(v.ime, v.priimek, " povprečna ocena: ", avg, " -> ")

		switch {
		case avg >= 9:
			fmt.Println("Odličen študent!")
		case avg < 9 && avg >= 6:
			fmt.Println("Povprečen študent!")
		case avg < 6:
			fmt.Println("Neuspešen študent!")
		}
	}
}

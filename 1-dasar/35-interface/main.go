package main

type Hewan interface {
	suara() string
	jenis() string
}

type Kucing struct{}

func (k Kucing) suara() string {
	return "Meong"
}

func (k Kucing) jenis() string {
	return "Mamalia"
}

func cetakSuara(h Hewan) {
	println("Suara:", h.suara())
	println("Jenis:", h.jenis())
}

func main() {
	k := Kucing{}
	cetakSuara(k)
}

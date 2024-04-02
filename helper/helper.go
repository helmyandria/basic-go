package helper

var version = "1.0.0"      //huruf awal kecil tidak bisa diakses dari luar class
var Application = "golang" //huruf awal besar bisa diakses dari luar class

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

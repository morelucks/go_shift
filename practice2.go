package main
import ("fmt")


const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAmerica
	canSeeAsia
)
func main(){
	var roles byte= isAdmin|isHeadquarters|canSeeAsia
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin?: %v\n", isAdmin & roles == isAdmin )
	fmt.Printf("Is Asia?: %v\n", canSeeAsia & roles == canSeeAsia)
	fmt.Printf("Is America?: %v\n", canSeeAmerica & roles == canSeeAmerica )

}
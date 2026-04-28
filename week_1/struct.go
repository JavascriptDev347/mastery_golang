package main


type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func (g gasEngine) mailLeft() uint8 {
	return g.gallons * g.mpg
}

// func main() {
// 	var myEngine gasEngine = gasEngine{
// 		mpg:     2,
// 		gallons: 4,
// 		owner: owner{
// 			name: "Russel",
// 		},
// 	}

// 	fmt.Println(myEngine.mailLeft())
// }

// TypeScript version
//// 1. owner interfeysi (Go'dagi owner struct)
//interface Owner {
//name: string;
//}
//
//// 2. GasEngine klassi (Go'dagi gasEngine struct + method)
//class GasEngine {
//mpg: number;
//gallons: number;
//owner: Owner;   // embedded struct → property sifatida
//
//constructor(mpg: number, gallons: number, owner: Owner) {
//this.mpg     = mpg;
//this.gallons = gallons;
//this.owner   = owner;
//}
//
//// Go'dagi (g gasEngine) milesLeft() uint8 → class method
//milesLeft(): number {
//return this.gallons * this.mpg;
//}
//}
//
//// 3. main() funksiyasi
//function main(): void {
//const myEngine = new GasEngine(
//2,           // mpg
//4,           // gallons
//{ name: "Russel" }
//);
//
//console.log(myEngine.milesLeft());  // → 8
//console.log(myEngine.owner.name);   // → "Russel"
//}
//
//main();

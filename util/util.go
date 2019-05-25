package util

import(
    "math/rand"
    "time"
)
//DiceInt creates random int from 1 to 6
func DiceInt() int{

	rand.Seed(time.Now().Unix())
    return rand.Intn(7 - 1) + 1

}
//CoinFlipInt creates Int for Coinflips
func CoinFlipInt() int{
    rand.Seed(time.Now().Unix())
    return rand.Intn(3 - 1) + 1
}
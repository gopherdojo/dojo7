package omikuji

import "time"

var (
	probTable   map[Omikuji]int
	omikujiList []Omikuji
)

type Omikuji struct {
	Result string `json:"result"`
}

func init() {
	probTable = map[Omikuji]int{
		Omikuji{"dai-kichi"}: 1,
		Omikuji{"chu-kichi"}: 3,
		Omikuji{"sho-kichi"}: 4,
		Omikuji{"kyo"}:       2,
	}

	for k, v := range probTable {
		for i := 0; i < v; i++ {
			omikujiList = append(omikujiList, k)
		}
	}
}

func DrawOmikuji(dt time.Time, seed int) Omikuji {
	if int(dt.Month()) == 1 {
		if dt.Day() == 1 || dt.Day() == 2 || dt.Day() == 3 {
			return Omikuji{"dai-kichi"}
		}
	}

	return omikujiList[seed%len(omikujiList)]
}

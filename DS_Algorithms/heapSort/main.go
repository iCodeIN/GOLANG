package main

import (
	"log"
	"time"

	"github.com/khvr1993/GOLANG/DS_Algorithms/heapSort/heapSort"
)

func main() {
	start := time.Now()
	a := []int{917,
		249,
		638,
		752,
		910,
		299,
		404,
		320,
		375,
		853,
		936,
		440,
		648,
		845,
		376,
		809,
		193,
		207,
		264,
		389,
		932,
		397,
		978,
		753,
		699,
		724,
		862,
		85,
		247,
		7,
		714,
		806,
		415,
		768,
		446,
		467,
		821,
		349,
		223,
		620,
		711,
		595,
		338,
		536,
		999,
		776,
		957,
		979,
		820,
		570,
		566,
		424,
		14,
		817,
		948,
		257,
		225,
		667,
		793,
		509,
		361,
		202,
		442,
		605,
		338,
		427,
		765,
		886,
		823,
		374,
		38,
		209,
		555,
		740,
		927,
		776,
		934,
		846,
		680,
		988,
		998,
		700,
		806,
		622,
		924,
		820,
		257,
		953,
		730,
		230,
		669,
		612,
		78}
	heapSort.HeapSort(&a)
	elapsed := time.Since(start)
	log.Printf("Heap Sort took %s", elapsed)
}
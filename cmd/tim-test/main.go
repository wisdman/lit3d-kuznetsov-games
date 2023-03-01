package main

import "fmt"

const ONE_PERIOD = 65536
const HALF_PERIOD = 32768
const ENC_DEL = 1000
const ENC_MAX = 200

var encoder_prev int32 = 0
var encoder_value uint16 = 0

func unwrap_encoder_diff(in uint16, prev *int32) int32 {
	var c32 int32 = int32(in) - HALF_PERIOD
	var diff int32 = (c32 - *prev)

	var mod_diff int32 = ((diff + HALF_PERIOD) % ONE_PERIOD) - HALF_PERIOD
	if diff < -HALF_PERIOD {
		mod_diff += ONE_PERIOD
	}
	
	*prev = *prev + mod_diff
	return mod_diff
}

func set_range(diff int32) uint16 {
	var newValue int32 = int32(encoder_value) + ( diff / ENC_DEL)
	if newValue <= 0 {
		encoder_value = 0
		return encoder_value
	}

	if newValue >= ENC_MAX {
		encoder_value = ENC_MAX
		return encoder_value
	}

	encoder_value = uint16(newValue)
	return encoder_value
}

func main() {
	var cnt uint16 = 0
	for i := 0; i < 100; i++ {
		cnt += 10000;
		diff := unwrap_encoder_diff(cnt, &encoder_prev)
		out := set_range(diff)
		fmt.Printf("CNT: %d; DIFF: %d OUT: %d\n", cnt, diff, out)
	}



}
package etl

import "strings"

//Transform converts legacy scoring system data to new data format
func Transform(in map[int][]string) map[string]int {
  out := map[string]int{}
  for k, v:= range in {
    for _, l := range v {
      out[strings.ToLower(l)] = k
    }
  }
  return out
}

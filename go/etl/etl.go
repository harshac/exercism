package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
  out := map[string]int{}
  for k, v:= range in {
    for _, l := range v {
      out[strings.ToLower(l)] = k
    }
  }
  return out
}

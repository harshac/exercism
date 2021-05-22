package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(predicate func(int) bool) Ints {
  if(i == nil){
    return nil
  }
  out:= Ints{}
  for _, el := range i{
    if(predicate(el)){
      out = append(out, el)
    }
  }
  return out
}

func (i Ints) Discard(predicate func(int) bool) Ints {
  if(i == nil){
    return nil
  }
  out:= Ints{}
  for _, el := range i{
    if(!predicate(el)){
      out = append(out, el)
    }
  }
  return out

}

func (l Lists) Keep(predicate func([]int) bool) Lists {
  if(l == nil){
    return nil
  }
  out:= Lists{}
  for _, el := range l{
    if(predicate(el)){
      out = append(out, el)
    }
  }
  return out
}

func (s Strings) Keep(predicate func(string) bool) Strings {
  if(s == nil){
    return nil
  }
  out:= Strings{}
  for _, el := range s{
    if(predicate(el)){
      out = append(out, el)
    }
  }
  return out

}

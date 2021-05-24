package strain

//Ints is a collection of integers
type Ints []int
//Lists is a collection of collection of integers
type Lists [][]int
//Strings is a collection of strings
type Strings []string

//Keep returns a new Ints containing those elements where the predicate is true
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

//Discard returns a new Ints containing those elements where the predicate is false
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

//Keep returns a new Lists containing those elements where the predicate is true
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

//Keep returns a new Strings containing those elements where the predicate is true
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

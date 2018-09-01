package main

type PoS struct {
	block *Block
}

func CreatePoS(block *Block) PoS {
	pos := &PoS{block}

	return pos
}

func (pos *PoS) prepareValidators(validators *Validators) (int map[*Validator]int) {
  validatorPool := map[*Validators]int

  mutex.lock()
  setValidators := validators
  mutex.Unlock()

  var prefixSum int
  for _, validator := range setValidators {
    prefixSum += validator.Balance
    validatorPool[validator] = prefixSum
  }


  return validatorPool, prefixSum
}

func (pos *PoS) pickValidator(validators *Validators) Validator {
  sum, validatorPool := pos.prepareValidators(validators)

  rand.Seed(time.Now().UnixNano())
  r := rand.Intn(sum)

  left := 0
  right := sum
  for left < right {
    mid := left + (right - left)/2
    if
  }
}

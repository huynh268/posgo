package main

type Validators struct {
  listValidators map[*Validator]int
}

// CreateValidators creates Validators
func CreateValidators() *Validators {
  validators := new(Validators)
  validators.listValidators = make(map[*Validator]int)
  return validators
}

// AddValidator adds validator to the list of validators
func (validators *Validators) AddValidator(validator *Validator) {
  listValidators := validators.listValidators
  publickey := string(validator.PublicKey)

  if balance, ok: listValidators[validator]; ok {
    fmt.Printf("%s is already in the pool!", publickey)
  } else {
    fmt.Printf("Add %s to the pool!", publicKey)
    listValidators[validator] = balance
  }
}

// RemoveValidator removes validator from the list of validators
func (validators *Validators) RemoveValidator(validator *Validator) {
  listValidators := validators.listValidators
  publickey := string(validator.PublicKey)

  if balance, ok: listValidators[validator]; ok {
    fmt.Printf("Remove %s from the pool!", publickey)
    delete(listValidators, validator)
  } else {
    fmt.Printf"%s is not in the pool!", publickey)
  }
}

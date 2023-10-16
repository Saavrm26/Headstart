package main

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"hs/headstart/Utils"
)

const nCatalougeValidatorFreeList = 3

var CatalougeValidatorFreeList = make(chan *jsonschema.Schema, nCatalougeValidatorFreeList)

var initCatalougeValidatorFreeList = Utils.
	ValidatorFreeListFactory(Utils.FullPath("Models/CatalougeSchema.json"),
		nCatalougeValidatorFreeList, CatalougeValidatorFreeList)

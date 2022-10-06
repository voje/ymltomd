package testdocs

var Doc1 = []byte(`
dict_struct: 
  subfield_one: "sub val one"
  subfield_two: "sub val two"
field_one: "some val"
field_two: "second val"
`)

var Doc2 = []byte(`---
## Test yaml document
field_one: "some val"
field_two: "second val"

# Let's try a dictionary
dict_struct:
  subfield_one: "sub val one"
  # Second subfield
  subfield_two: "sub val two"
`)
## Interfaces 

All types are concrete types, except the interface type. The interface type is an abstract type and unlike with a concrete type, we cannot directly use an interface type to create a value.

Interfaces can be usefull in case a method does something that can be re-used for a variety of objects. Interfaces in Go are oftentimes explained as something that is used to define a behavior, like printing, writing etc.

You can create a variable that is an interface type and that has the methods that belong to the interface. This makes the method a custom type as well as a collection of methods.

Interfaces are implicit. We do not manually define the correlation between the interface custom type and the other methods and functions.

Interfaces can be seen as a contract to help manage types. The Go compiler will check all the values and the returns involved, but (obviously) will not check if the logic that is used is sound.




Fun facts:
- if the struct passed as a receiver is not used by the method, we can just pass the type
- structs do not need to have any fields defined. A field-less struct can be used as placeholder for methods and interfaces

It is possible to attach a method to almost any type (even functions!):
- can use pointer and receiver values:
  - int
  - array
  - string
  - struct
  - float64
- use receiver values:
  - slice
  - map
  - func
  - chan


'The bigger the interface, the weaker the abstraction' - Rob Pike.
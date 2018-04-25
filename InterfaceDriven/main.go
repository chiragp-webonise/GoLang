package main

import "github.com/InterfaceDriven/pets"

// This interface represents any type
// that has both Walk and Sit methods.
type FourLegged interface {
    Walk()
    Sit()
}

// We can replace DemoDog and DemoCat
// with this single function.
func Demo(animal FourLegged) {
    animal.Walk()
    animal.Sit()
}

func main() {
    dog := pets.Dog{"Fido", "Terrier"}
    cat := pets.Cat{"Fluffy", "Siamese"}
    Demo(dog)
    // The above call (again) outputs:
    // Fido walks across the room
    // Fido sits down
    Demo(cat)
    // The above call (again) outputs:
    // Fluffy walks across the room
    // Fluffy sits down
}
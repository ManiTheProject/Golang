package main
import "fmt"

/*
For the purpose of this challenge, the 12 musical notes in the chromatic scale are named:

C  C#  D  D#  E  F  F#  G  G#  A  A#  B
The interval between each pair of notes is called a semitone, and the sequence wraps around. So for instance, E is 1 semitone above D#,
C is 1 semitone above B, F# is 4 semitones above D, and C# is 10 semitones above D#. (This also means that every note is 12 semitones 
above itself.)

A major scale comprises 7 out of the 12 notes in the chromatic scale. There are 12 different major scales, one for each note.
For instance, the D major scale comprises these 7 notes:

D  E  F#  G  A  B  C#
The notes in a major scale are the notes that are 0, 2, 4, 5, 7, 9, and 11 semitones above the note that the 
scale is named after. In the movable do solfège system, these are referred to by the names Do, Re, Mi, Fa, So, La, and Ti,
respectively. So for instance, Mi in the D major scale is F#, because F# is 4 semitones above D.

(In general, a note can have more than one name. For instance A# is also known as Bb. 
Depending on the context, one or the other name is more appropriate. You'd never hear it referred to as the A# major 
scale in real music. Instead it would be called Bb major. Don't worry about that for this challenge. Just always 
use the names of the notes given above.)

Challenge
Write a function that takes the name of a major scale and the solfège name of a note, and returns the corresponding note in that scale.

note("C", "Do") -> "C"
note("C", "Re") -> "D"
note("C", "Mi") -> "E"
note("D", "Mi") -> "F#"
note("A#", "Fa") -> "D#"
*/

var solfege = map[string]int{"Do": 0, "Re": 2, "Mi": 4, "Fa": 5, "So": 7, "La": 9, "Ti": 11}
var notes = [...]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

func main() {
    note("C", "Do")
    note("C", "Re")
    note("C", "Mi")
    note("D", "Mi")
    note("A#", "Fa")
}

func note(scale, move string) {
    start := 0
    // find index
    for i := 0; i < len(notes); i++ {
        if notes[i] == scale {
            start = i
            break
        }
    }
    fmt.Println(notes[(start+solfege[move])%12])
}

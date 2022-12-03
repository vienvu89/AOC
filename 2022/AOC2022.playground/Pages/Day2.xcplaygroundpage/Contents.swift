//: [Previous](@previous)

import Foundation

//
let a = readfile()!

let points = ["X": 1,
              "Y": 2,
              "Z": 3]

func calculatePoints(elf: String, human: String) -> Int {
    switch (elf, human) {
    case ("A", "X"),
        ("B", "Y"),
        ("C", "Z"):
        return 3
    case ("A", "Y"),
        ("B", "Z"),
        ("C", "X"):
        return 6
    default:
        return 0
    }
}

///Part 1
let total1 = a.split(separator: "\n")
    .map { game in
        let plays = game.components(separatedBy: " ")
        
        return calculatePoints(elf: plays[0], human: plays[1]) + points[plays[1]]!
    }
    .reduce(0, { $0 + $1 })

print(total1)


//Part 2
let maps = ["A": 1,
            "B": 2,
            "C": 3]

let results = ["X": 0,
               "Y": 3,
               "Z": 6]

func calculatePlay(elf: String, result: String) -> String? {
    switch (elf, result) {
    case (_, "Y"):
        return elf
    case ("A", "X"),
        ("B", "Z"):
        return "C"
    case ("A", "Z"),
        ("C", "X"):
        return "B"
    case ("B", "X"),
        ("C", "Z"):
        return "A"
    default:
        return nil
    }
}

let total2 = a.split(separator: "\n")
    .map { game in
        let plays = game.components(separatedBy: " ")
        var value = 0
        if let a =  calculatePlay(elf: plays[0], result: plays[1]) {
            value = maps[a]!
        }
        
        return value + results[plays[1]]!
    }
    .reduce(0, { $0 + $1 })

print(total2)

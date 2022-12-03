//: [Previous](@previous)

import Foundation

//
if let a = readfile() {
    let elves = a.split(separator: "\n\n")
    //Part 1
    var sum = 0
    for elf in elves {
        var total = elf.split(separator: "\n").map{ Int($0)! }.reduce(0, { $0 + $1 })
        if total > sum { sum = total }
    }
    print(sum)
    
    //Part 2
    let a = elves.map { $0.split(separator: "\n").map{ Int($0)! }.reduce(0, { $0 + $1 })}.sorted { $0 > $1 }
    print(a[0] + a[1] + a[2])
}

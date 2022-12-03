//: [Previous](@previous)

import Foundation

if let a = readfile() {
    let elves = a.split(separator: "\n\n")
    var sum = 0
    for elf in elves {
        var total = elf.split(separator: "\n").map{ Int($0)! }.reduce(0, { $0 + $1 })
        if total > sum { sum = total }
    }
    print(sum)
}

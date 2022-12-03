import Foundation

public func readfile(_ name: String = "input") -> String? {
    // get the file path for the file "test.json" in the playground bundle
    let filePath = Bundle.main.path(forResource:name, ofType: "txt")

    // get the contentData
    let contentData = FileManager.default.contents(atPath: filePath!)

    // get the string
    let content = String(data:contentData!, encoding:String.Encoding.utf8)

    return content
}

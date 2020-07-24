List = {}
function List.new()
    List.at = -1
end
function List.push(value)
    List[value] = value
    List.at = value
end
function List.printvalue()
    print(List[List.at])
end
List.new()
List.push(1)
List.push(2)
List.printvalue()
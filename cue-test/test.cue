one: 1
two: 1

"float": 2.5

// Lis
list: [
    1,
    2,
    3,
]

// Object can contain other object
a: {
    content: list
}

// Define large capital
largeCapital: {
    name: string
    population: >5
    capital: true
}

// Merge element
largeCapital: {
    test: true
}

// Constraint
dataType: {
    name: string
    content: {}
    test: "Other props"
}

data: dataType
data: {
    name: "MyData"
    content: largeCapital
}
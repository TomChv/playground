#Foo: {
    foo: string
}

#Data: {
    name: string
    content: #Foo
    ...
}

foo: #Foo & {
    foo: "bar"
}

data: #Data & {
    name: "Here"
    content: foo
    kind: "bloc"
}
import envbyjson from "envbyjson";

// it loads env.json file in the root directory
envbyjson.load()
console.log(process.env.A)

// it loads env2.json file in the testdata directory
envbyjson.load("testdata/env2.json")
console.log(process.env.B)

// it loads env3.json and env4.json file in the testdata directory
envbyjson.loadProp("testdata/env3.json", "testdata/env4.json", "C") 
console.log(process.env.D)
console.log(process.env.E)

try { 
    // it throws an error because one of the prop is not string
    envbyjson.load("testdata/env_err.json")
} catch(e) {
    console.log(e.message)
}

try {
    // it throws an error because one of the prop is not string
    envbyjson.loadProp("testdata/env_err2.json", "C")
} catch(e) {
    console.log(e.message)
}
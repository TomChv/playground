package main

import (
	da "dagger.io/dagger"
	b "universe.dagger.io/bash"
)

da.#Plan & {
	actions: {
		sayHello: b.#Run & {
			script: contents: "echo hello world"
		}
	}
}

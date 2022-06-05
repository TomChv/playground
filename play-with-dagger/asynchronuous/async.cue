package main

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"
)

dagger.#Plan & {
	actions: {
		image: core.#Pull & {
			source: "postgres:14.3-alpine"
		}

		test: {
			start: core.#Start & {
				input: image.output
				args: ["docker-entrypoint.sh", "postgres"]
				env: {
					POSTGRES_PASSWORD: "password"
					POSTGRES_USER:     "foo"
					POSTGRES_DB:       "test"
					image.config.env
				}
			}
		}
	}
}

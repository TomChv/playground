"""Execute a command."""

import sys

import anyio

import dagger


async def test():
    async with dagger.Connection(dagger.Config()) as conn:
        await conn.container().from_("alpine:edge").with_exec(["mkdir", "-p", "/test/abc"]).with_exec(["/test/abc"]).stdout()


anyio.run(test)
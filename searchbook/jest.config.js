module.exports = {
	roots: ['<rootDir>'],
	transform: {
		'^.+\\.tsx?$': 'ts-jest',
	},
	testMatch: ['**/tests/index.ts'],
	moduleFileExtensions: ['ts', 'js'],
};

import { IResolvers } from 'apollo-server';

const mutations: IResolvers = {
	Mutation: {
		add: (_: unknown, args: { number1: number; number2: number }) => {
			const { number1, number2 } = args;
			return number1 + number2;
		},
	},
};

export default mutations;

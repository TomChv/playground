import { IResolvers } from 'apollo-server';

const queries: IResolvers = {
	Query: {
		hello: (): string => {
			return 'Hello world';
		},

		personalHello: (_: unknown, args: { name: string }) => {
			return `Hello ${args.name}`;
		},
	},
};

export default queries;

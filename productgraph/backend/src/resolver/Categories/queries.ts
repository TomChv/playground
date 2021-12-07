import { IResolvers } from 'apollo-server';
import { Category } from '@prisma/client';
import { Context } from '../../context';

const queries: IResolvers = {
	Query: {
		categories: async (_: unknown, args: void, ctx: Context): Promise<Category[]> => {
			return ctx.db.category.findMany({
				include: {
					products: true,
				},
			});
		},

		category: async (_: unknown, args: { id: string }, ctx: Context): Promise<Category | null> => {
			return ctx.db.category.findUnique({
				where: args,
				include: {
					products: true,
				},
			});
		},
	},
};

export default queries;

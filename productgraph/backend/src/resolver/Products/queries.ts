import { IResolvers } from 'apollo-server';
import { Product } from '@prisma/client';
import { Context } from '../../context';

const queries: IResolvers = {
	Query: {
		products: async (_: unknown, args: void, ctx: Context): Promise<Array<Product>> => {
			return ctx.db.product.findMany({
				include: {
					categories: true,
				},
			});
		},

		product: async (_: unknown, args: { id: string }, ctx: Context): Promise<Product | null> => {
			return ctx.db.product.findUnique({
				where: args,
				include: {
					categories: true,
				},
			});
		},
	},
};

export default queries;

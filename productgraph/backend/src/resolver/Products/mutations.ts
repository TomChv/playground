import { IResolvers, UserInputError } from 'apollo-server';

import { Product } from '@prisma/client';
import { Context } from '../../context';
import { NotFoundError } from '../../helpers/Errors';

interface ProductInput {
	name: string;
	desc: string;
	price: number;
	categories: Array<string>;
}

const mutations: IResolvers = {
	Mutation: {
		createProduct: async (_: unknown, args: { input: ProductInput }, ctx: Context): Promise<Product> => {
			if (Object.keys(args.input).length < 3) {
				throw new UserInputError('Miss arguments');
			}

			const { name, desc, price, categories } = args.input;

			if (price <= 0) {
				throw new NotFoundError('Price must be positive');
			}

			let connect: Array<{ id: string }> = [];
			if (categories !== undefined) {
				connect = categories.map((e) => ({ id: e }));
			}

			return ctx.db.product.create({
				data: {
					name,
					desc,
					price,
					categories: { connect },
				},
				include: {
					categories: true,
				},
			});
		},

		updateProduct: async (
			_: unknown,
			args: { id: string; input: ProductInput },
			ctx: Context,
		): Promise<Product | null> => {
			const { id } = args;

			const product = await ctx.db.product.findUnique({ where: { id } });
			if (product === null) {
				throw new NotFoundError('Product not found');
			}

			const { name, desc, price, categories } = args.input;
			if (price <= 0) {
				throw new NotFoundError('Price must be positive');
			}

			let connect: Array<{ id: string }> = [];
			if (categories !== undefined) {
				connect = categories.map((e) => ({ id: e }));
			}

			return ctx.db.product.update({
				where: { id },
				data: {
					name,
					desc,
					price,
					categories: { connect },
				},
				include: {
					categories: true,
				},
			});
		},

		deleteProduct: async (_: unknown, args: { id: string }, ctx: Context): Promise<Product> => {
			const { id } = args;

			const product = await ctx.db.product.findUnique({ where: { id } });
			if (product === null) {
				throw new NotFoundError('Product not found');
			}

			return ctx.db.product.delete({
				where: { id },
				include: {
					categories: true,
				},
			});
		},
	},
};

export default mutations;

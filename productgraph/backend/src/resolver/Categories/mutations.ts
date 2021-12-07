import { IResolvers, UserInputError } from 'apollo-server';

import { Category } from '@prisma/client';
import { Context } from '../../context';
import { NotFoundError } from '../../helpers/Errors';

interface CategoryInput {
	name: string;
	products: Array<string>;
}

const mutations: IResolvers = {
	Mutation: {
		createCategory: async (_: unknown, args: { input: CategoryInput }, ctx: Context): Promise<Category> => {
			const { name, products } = args.input;

			let connect: Array<{ id: string }> = [];
			if (products !== undefined) {
				connect = products.map((e) => ({ id: e }));
			}

			return ctx.db.category.create({
				data: {
					name,
					products: { connect },
				},
				include: {
					products: true,
				},
			});
		},

		updateCategory: async (
			_: unknown,
			args: { id: string; input: CategoryInput },
			ctx: Context,
		): Promise<Category | null> => {
			const { id } = args;

			const product = await ctx.db.category.findUnique({ where: { id } });
			if (product === null) {
				throw new NotFoundError('Category not found');
			}

			const { name, products } = args.input;
			let connect: Array<{ id: string }> = [];
			if (products !== undefined) {
				connect = products.map((e) => ({ id: e }));
			}

			return ctx.db.category.update({
				where: { id },
				data: {
					name,
					products: { connect },
				},
				include: {
					products: true,
				},
			});
		},

		deleteCategory: async (_: unknown, args: { id: string }, ctx: Context): Promise<Category> => {
			const { id } = args;

			const product = await ctx.db.category.findUnique({ where: { id } });
			if (product === null) {
				throw new NotFoundError('Category not found');
			}

			return ctx.db.category.delete({
				where: { id },
				include: {
					products: true,
				},
			});
		},
	},
};

export default mutations;

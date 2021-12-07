import React from 'react';

import { Product } from '../Api/Entities';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import { v4 as uuidv4 } from 'uuid';

import ProductTab from './ProductTab';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			margin: 10,
			width: 'calc(100vw - 35vw)',
			height: 'calc(100vh - 15vh)',
			flexDirection: 'row'
		},
	})
);

interface ProductsGridsProps {
	value: number;
	index: number
	products: Product[];
}

export default function ProductsGrid(props: ProductsGridsProps) {
	const {products, value, index: i} = props;
	const classes = useStyles();

	if (value === i) {
		return (
			<Grid container className={classes.root} spacing={3}>
				{products.map((product) =>
					(<ProductTab index={i} value={value} key={uuidv4().toString()} product={product}/>))
				}
			</Grid>
		)
	} else {
		return (<div/>)
	}
}
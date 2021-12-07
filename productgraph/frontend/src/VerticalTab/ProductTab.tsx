import React from 'react';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';

import { Product } from '../Api/Entities';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';

const useStyles = makeStyles((theme: Theme) =>
	createStyles({
		box: {
			height: 180,
			width: 180,
			background: '#FBFBFB',
			boxShadow: '0px 4px 3px rgba(0, 0, 0, 0.25)',
			padding: 10,
			margin: 10,
			display: 'flex',
			flexDirection: 'column',
			flexWrap: 'wrap',
			alignContent: 'space-between',
		},
		title: {
			marginBottom: 5,
		},
		desc: {
		},
		price: {
			fontWeight: 'bold',
			alignSelf: 'left',
		}
	}),
)

interface ProductTabProps {
	index: number;
	value: number;
	product: Product;
}

export default function ProductTab(props: ProductTabProps) {
	const {product, value, index, ...other} = props;
	const classes = useStyles();

	return (
		<div
			role="tabpanel"
			hidden={value !== index}
			id={`vertical-tabpanel-${index}`}
			aria-labelledby={`vertical-tab-${index}`}
			{...other}
		>
			{value === index && (
				<Box p={3} className={classes.box}>
					<div>
						<Typography className={classes.title} variant={'h5'}>{product.name}</Typography>
					</div>
					<div>
						<Typography className={classes.desc}>{product.desc}</Typography>
					</div>
					<div>
						<Typography className={classes.price}>{product.price}€</Typography>
					</div>

				</Box>
			)}
		</div>
	);
}

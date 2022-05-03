import AddIcon from '@mui/icons-material/Add';
import { IconButton, Theme } from '@mui/material';

export default function Add() {
	return (
		<IconButton aria-label="add" sx={{
			backgroundColor: "lightblue",
			position: "absolute",
			right: "5vw",
			bottom: "10vh",
			":hover": {
				backgroundColor: (theme: Theme) => theme.palette.primary.light
			}
		}}>
			<AddIcon sx={ {
				width: 40,
				height: 40,
				color: "black"
			} }/>
		</IconButton>
	);
};
import { Card, Typography } from '@mui/material';

export default function Task() {
	return (
		<Card sx={{
			minHeight: "20vh",
			minWidth: "20vw"
		}}>
			<Typography variant="h5">
				Title
			</Typography>
			<Typography variant="body1">
				Lorem ipsum...
			</Typography>
		</Card>
	)
};
import { Box, Grid } from '@mui/material';

import Task from './Task'

export default function TasksList() {
	return (
		<Box sx={{ flexGrow: 1 }}>
			<Grid container spacing={{ xs: 2, md: 4 }} columns={{ xs: 4, sm: 8, md: 12 }}>
				{Array.from(Array(6)).map((_, index) => (
					<Grid item xs={2} sm={4} md={4} key={index}>
						<Task/>
					</Grid>
				))}
			</Grid>
		</Box>
	)
}
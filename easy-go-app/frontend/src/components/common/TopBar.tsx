import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { Container } from '@mui/material';

export default function TopBar() {
	return (
		<Box sx={{ flexGrow: 1 }}>
			<AppBar position="static" color="primary">
				<Toolbar variant="dense">
					<Container sx={{
						textAlign: "center"
					}}>
						<Typography variant="h6" color="inherit" component="div">
							Easy go
						</Typography>
					</Container>
				</Toolbar>
			</AppBar>
		</Box>
	);
}

import { Box, Container, Typography } from '@mui/material';

export default function Footer() {
	return (
		<Box
			component="footer"
			sx={ {
				py: 3,
				px: 2,
				mt: 'auto',
				backgroundColor: (theme) => theme.palette.primary.light
			} }
		>
			<Container maxWidth="sm">
				<Typography variant="body1" textAlign="center" fontWeight="bold" color="white">
					Made with love by Vasek
				</Typography>
			</Container>
		</Box>
	);
}
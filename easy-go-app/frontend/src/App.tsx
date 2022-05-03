import React from 'react';

import { Footer } from './components/common';
import { Box, Container } from '@mui/material';
import { AddButton } from './components/buttons';
import TopBar from './components/common/TopBar';
import TasksList from './components/task/TasksList';

function App() {
	return (
		<Box sx={ { display: 'flex', flexDirection: 'column', minHeight: '100vh' } }>
			<TopBar/>
			<Container
				component="main"
				sx={ {
					mt: '20px',
					mb: '20px',
					height: '80vh',
					width: '100vw',
					overflow: 'auto'
				} }>
				<TasksList/>
				<AddButton/>
			</Container>
			<Footer/>
		</Box>
	);
}

export default App;

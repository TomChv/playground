import { NestFactory } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';

import { apiConfig } from './appConfig';
import AppModule from './appModule';

const { host, port } = apiConfig;

async function bootstrap() {
	const app = await NestFactory.create(AppModule);

	app.useGlobalPipes(new ValidationPipe());

	const options = new DocumentBuilder()
		.setTitle('Note API')
		.setDescription('Note API description')
		.setVersion('1.0')
		.addTag('PoC')
		.build();
	const documentation = SwaggerModule.createDocument(app, options);
	SwaggerModule.setup('doc', app, documentation);

	app.enableCors();
	await app.listen(port);

	console.log(`Server listen on http://${host}:${port}`);
}
bootstrap();

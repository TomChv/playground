import { Injectable, NestMiddleware, Logger } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';

@Injectable()
export default class AppLogger implements NestMiddleware {
	private logger = new Logger('HTTP');

	use(request: Request, response: Response, next: NextFunction): void {
		const { ip, method, path: url } = request;

		response.on('close', () => {
			const { statusCode } = response;

			this.logger.log(
				`[${method} ${url}] from ${ip} response with ${statusCode}`,
			);
		});

		next();
	}
}

import { ENV } from '../config/env';

export class ApiError extends Error {
	constructor(public status: number, message: string) {
		super(message);
		this.name = 'ApiError';
	}
}

export async function http(path: string, options: RequestInit = {}) {
	const response = await fetch(`${ENV.API_URL}${path}`, {
		...options,
		headers: {
			'Content-Type': 'application/json',
			...options.headers
		}
	});

	if (!response.ok) {
		throw new ApiError(response.status, `HTTP Error: ${response.status}`);
	}

	return response.json();
}

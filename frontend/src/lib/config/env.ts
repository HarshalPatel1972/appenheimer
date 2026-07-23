let apiUrl = import.meta.env.VITE_API_URL || '';

if (typeof apiUrl === 'string') {
	apiUrl = apiUrl.trim();
}

if (!apiUrl || apiUrl.includes('vercel.app')) {
	console.error('CRITICAL: VITE_API_URL is missing or incorrectly set to Vercel. Falling back to Render backend.');
	apiUrl = 'https://appenheimer.onrender.com';
}

export const ENV = {
	API_URL: apiUrl,
};
